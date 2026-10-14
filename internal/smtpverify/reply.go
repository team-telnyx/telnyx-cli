package smtpverify

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"
	"unicode/utf8"
)

// Reply is one SMTP server reply, possibly multi-line, kept verbatim.
type Reply struct {
	Code     int
	Enhanced string   // RFC 3463 enhanced status code ("5.7.8"), "" if absent
	Lines    []string // raw reply lines, CRLF stripped
}

// Text renders the reply on a single line for verdicts: the lines joined
// with " | " (a multi-line failure reply is rare, but must not be truncated),
// with control bytes escaped (see printable).
func (r Reply) Text() string {
	lines := make([]string, len(r.Lines))
	for i, l := range r.Lines {
		lines[i] = printable(l)
	}
	return strings.Join(lines, " | ")
}

// printable escapes C0/C1 control characters, DEL and invalid UTF-8 as
// \xNN / \u00NN. Replies before STARTTLS are attacker-controllable, and
// raw escape sequences could rewrite what a terminal shows (for example
// paint a fake VERIFIED line). Parsing always uses the raw bytes.
func printable(s string) string {
	var b strings.Builder
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		switch {
		case r == utf8.RuneError && size == 1:
			fmt.Fprintf(&b, "\\x%02x", s[i])
		case r == '\t':
			b.WriteRune(r)
		case r < 0x20 || r == 0x7f:
			fmt.Fprintf(&b, "\\x%02x", r)
		case r >= 0x80 && r <= 0x9f:
			fmt.Fprintf(&b, "\\u%04x", r)
		default:
			b.WriteString(s[i : i+size])
		}
		i += size
	}
	return b.String()
}

// Positive reports a 2xx or 3xx reply.
func (r Reply) Positive() bool {
	return r.Code >= 200 && r.Code < 400
}

const (
	maxReplyLineBytes = 2048
	maxReplyLines     = 64
)

var enhancedCodeRE = regexp.MustCompile(`^([245])\.(\d{1,3})\.(\d{1,3})(?:\s|$)`)

// errMalformedReply marks a reply that is not valid SMTP (RFC 5321 §4.2).
var errMalformedReply = errors.New("malformed reply")

// readReply reads one complete (possibly multi-line) reply.
func readReply(r *bufio.Reader) (Reply, error) {
	var reply Reply
	for {
		line, err := readLine(r)
		if err != nil {
			if errors.Is(err, io.EOF) && len(reply.Lines) > 0 {
				err = io.ErrUnexpectedEOF
			}
			return reply, err
		}
		code, sep, err := parseReplyLine(line)
		if err != nil {
			return reply, err
		}
		if len(reply.Lines) == 0 {
			reply.Code = code
			reply.Enhanced = parseEnhancedCode(code, line)
		} else if code != reply.Code {
			return reply, fmt.Errorf("%w: code changed mid-reply in %q", errMalformedReply, line)
		}
		reply.Lines = append(reply.Lines, line)
		if sep != '-' {
			return reply, nil
		}
		if len(reply.Lines) >= maxReplyLines {
			return reply, fmt.Errorf("%w: more than %d continuation lines", errMalformedReply, maxReplyLines)
		}
	}
}

// readLine reads one CRLF- (or bare LF-) terminated line, bounded in length.
func readLine(r *bufio.Reader) (string, error) {
	var b strings.Builder
	for {
		chunk, isPrefix, err := r.ReadLine()
		if err != nil {
			return "", err
		}
		b.Write(chunk)
		if b.Len() > maxReplyLineBytes {
			return "", fmt.Errorf("%w: line longer than %d bytes", errMalformedReply, maxReplyLineBytes)
		}
		if !isPrefix {
			return b.String(), nil
		}
	}
}

// parseReplyLine validates "NNN", "NNN text" or "NNN-text" and returns the
// code and separator (' ' for a final line, '-' for a continuation).
func parseReplyLine(line string) (int, byte, error) {
	if len(line) < 3 {
		return 0, 0, fmt.Errorf("%w: %q", errMalformedReply, line)
	}
	code := 0
	for i := 0; i < 3; i++ {
		c := line[i]
		if c < '0' || c > '9' {
			return 0, 0, fmt.Errorf("%w: %q", errMalformedReply, line)
		}
		code = code*10 + int(c-'0')
	}
	if code < 200 || code > 599 {
		return 0, 0, fmt.Errorf("%w: reply code %d out of range in %q", errMalformedReply, code, line)
	}
	if len(line) == 3 {
		return code, ' ', nil
	}
	if sep := line[3]; sep == ' ' || sep == '-' {
		return code, sep, nil
	}
	return 0, 0, fmt.Errorf("%w: %q", errMalformedReply, line)
}

// parseEnhancedCode extracts an RFC 3463 code whose class matches the reply
// code's first digit ("535 5.7.8 ..." -> "5.7.8").
func parseEnhancedCode(code int, line string) string {
	if len(line) < 5 {
		return ""
	}
	m := enhancedCodeRE.FindStringSubmatch(line[4:])
	if m == nil || int(m[1][0]-'0') != code/100 {
		return ""
	}
	return m[1] + "." + m[2] + "." + m[3]
}

const uuidPattern = `[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`

var (
	queuedUUIDRE = regexp.MustCompile(`(?i)queued as (` + uuidPattern + `)\b`)
	uuidRE       = regexp.MustCompile(`\b` + uuidPattern + `\b`)
)

// messageUUID returns the message UUID carried by an accepting DATA reply
// (Telnyx: "250 2.0.0 Ok: queued as <uuid>"), or "" if there is none. The
// "queued as" form wins; otherwise the first UUID in the reply.
func messageUUID(r Reply) string {
	for _, line := range r.Lines {
		if m := queuedUUIDRE.FindStringSubmatch(line); m != nil {
			return strings.ToLower(m[1])
		}
	}
	for _, line := range r.Lines {
		if m := uuidRE.FindString(line); m != "" {
			return strings.ToLower(m)
		}
	}
	return ""
}

// capabilities parses an EHLO reply into upper-cased keyword -> parameters.
// The first line is the server greeting, not a capability.
func capabilities(r Reply) map[string]string {
	caps := map[string]string{}
	for _, line := range r.Lines[1:] {
		if len(line) < 5 {
			continue
		}
		keyword, params, _ := strings.Cut(line[4:], " ")
		caps[strings.ToUpper(keyword)] = params
	}
	return caps
}

// hasAuthMechanism reports whether an EHLO "AUTH" capability lists mech.
func hasAuthMechanism(caps map[string]string, mech string) bool {
	params, ok := caps["AUTH"]
	if !ok {
		return false
	}
	for _, m := range strings.Fields(params) {
		if strings.EqualFold(m, mech) {
			return true
		}
	}
	return false
}
