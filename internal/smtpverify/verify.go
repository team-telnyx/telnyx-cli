// Package smtpverify runs one real SMTP submission session against a Telnyx
// SMTP submission host — STARTTLS, AUTH PLAIN with the literal username
// "apikey" and the caller's API key, MAIL FROM, RCPT TO, DATA, QUIT — and
// prints a stage-by-stage transcript of every server reply, verbatim
// (credentials redacted), then a hint or portal note, and last a single
// verdict line.
package smtpverify

import (
	"bufio"
	"context"
	"crypto/rand"
	"crypto/tls"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net"
	"net/mail"
	"net/textproto"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)

// Stage names a step of the submission session; FAILED verdicts name the
// stage whose reply (or local check) ended the session.
type Stage string

const (
	StageConnect  Stage = "CONNECT"
	StageEHLO     Stage = "EHLO"
	StageSTARTTLS Stage = "STARTTLS"
	StageAuth     Stage = "AUTH"
	StageMailFrom Stage = "MAIL FROM"
	StageRcptTo   Stage = "RCPT TO"
	StageData     Stage = "DATA"
	StageQuit     Stage = "QUIT"
)

// AuthUsername is the fixed SMTP username; the API key is the password.
const AuthUsername = "apikey"

// Defaults mirror the Telnyx SMTP submission contract: the server allows
// 30 s for AUTH-stage commands and 300 s for DATA.
const (
	DefaultPort           = 587
	DefaultCommandTimeout = 30 * time.Second
	DefaultDataTimeout    = 300 * time.Second
	DefaultEHLOName       = "localhost"
)

// quitTimeout bounds the courtesy QUIT after a failure.
const quitTimeout = 5 * time.Second

// Config describes one verification session.
type Config struct {
	Host   string
	Port   int
	APIKey string
	From   string // envelope sender and From: header
	To     string // recipient; defaults to From

	EHLOName       string
	CommandTimeout time.Duration // per reply, every stage but DATA content
	DataTimeout    time.Duration // message content through the final DATA reply

	// TLSConfig is cloned for STARTTLS; ServerName defaults to Host.
	TLSConfig *tls.Config
	// Dial defaults to a net.Dialer bounded by CommandTimeout.
	Dial func(ctx context.Context, network, address string) (net.Conn, error)
	// Now defaults to time.Now (used for the Date header and subject).
	Now func() time.Time
}

// Result is the outcome of a session.
type Result struct {
	Verified    bool
	MessageUUID string
	Stage       Stage  // stage that ended the session (StageQuit when verified)
	Reply       *Reply // terminal server reply; nil for local failures
	Detail      string // local failure description when Reply is nil
	Hint        string // user guidance for a failure
}

// Verdict is the single summary line: "VERIFIED: message UUID <uuid>" or
// "FAILED at <stage>: <exact reply or local failure>".
func (r Result) Verdict() string {
	switch {
	case r.Verified:
		return "VERIFIED: message UUID " + r.MessageUUID
	case r.Reply != nil:
		return fmt.Sprintf("FAILED at %s: %s", r.Stage, r.Reply.Text())
	default:
		// Detail can embed peer-controlled text (e.g. certificate names in
		// an x509 error), so it is escaped like server replies.
		return fmt.Sprintf("FAILED at %s: %s", r.Stage, printable(r.Detail))
	}
}

// Run validates cfg, performs the session, and writes the transcript, a
// hint or portal note, and the verdict to out; the verdict is always the
// last line. The error is non-nil only for an invalid Config; every
// session outcome, including failures, is reported in the Result. Neither
// the output nor the Result carries the API key: server text that echoes
// it, or the AUTH PLAIN payload, is redacted. Nothing is written until the
// session ends, as all of its server text is redacted together.
func Run(ctx context.Context, cfg Config, out io.Writer) (Result, error) {
	cfg, err := cfg.normalize()
	if err != nil {
		return Result{}, err
	}
	s := &session{cfg: cfg, t: &transcript{w: out, redact: newRedactor(cfg.APIKey)}}
	res := s.t.flush(s.run(ctx))
	if res.Verified {
		fmt.Fprintf(out, "\nLook the message up by the UUID below in the Mission Control portal message log.\n%s\n", res.Verdict())
	} else {
		fmt.Fprintf(out, "\nhint: %s\n%s\n", res.Hint, res.Verdict())
	}
	return res, nil
}

func (c Config) normalize() (Config, error) {
	c.Host = strings.TrimSpace(c.Host)
	if c.Host == "" {
		return c, errors.New("host is required")
	}
	if strings.ContainsAny(c.Host, " \t\r\n/@[]") {
		return c, fmt.Errorf("invalid host %q", c.Host)
	}
	if strings.Contains(c.Host, ":") && net.ParseIP(c.Host) == nil {
		return c, fmt.Errorf("invalid host %q: pass the port separately with --port", c.Host)
	}
	if c.Port == 0 {
		c.Port = DefaultPort
	}
	if c.Port < 1 || c.Port > 65535 {
		return c, fmt.Errorf("invalid port %d", c.Port)
	}
	if c.APIKey == "" {
		return c, errors.New("an API key is required")
	}
	if strings.ContainsAny(c.APIKey, "\x00\r\n") {
		return c, errors.New("the API key contains invalid characters")
	}
	from, err := parseAddress("from", c.From)
	if err != nil {
		return c, err
	}
	c.From = from
	if strings.TrimSpace(c.To) == "" {
		c.To = c.From
	} else if c.To, err = parseAddress("to", c.To); err != nil {
		return c, err
	}
	if c.EHLOName == "" {
		c.EHLOName = DefaultEHLOName
	}
	if c.CommandTimeout <= 0 {
		c.CommandTimeout = DefaultCommandTimeout
	}
	if c.DataTimeout <= 0 {
		c.DataTimeout = DefaultDataTimeout
	}
	if c.Now == nil {
		c.Now = time.Now
	}
	if c.Dial == nil {
		d := &net.Dialer{Timeout: c.CommandTimeout}
		c.Dial = d.DialContext
	}
	return c, nil
}

// parseAddress accepts "user@domain" or "Name <user@domain>" and returns the
// bare address. CR/LF are rejected up front: the address lands in SMTP
// commands and message headers.
func parseAddress(field, s string) (string, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return "", fmt.Errorf("a %s address is required", field)
	}
	if strings.ContainsAny(s, "\r\n") {
		return "", fmt.Errorf("invalid %s address %q", field, s)
	}
	a, err := mail.ParseAddress(s)
	if err != nil || !plainMailbox(a.Address) {
		return "", fmt.Errorf("invalid %s address %q: use a plain ASCII address such as name@example.com", field, s)
	}
	return a.Address, nil
}

// plainMailbox accepts dot-atom@domain in ASCII only. net/mail unquotes
// quoted local parts, so without this check `"a@b> X"@example.com` would
// inject into the MAIL FROM line; non-ASCII would need SMTPUTF8.
func plainMailbox(addr string) bool {
	at := strings.LastIndexByte(addr, '@')
	if at <= 0 || at == len(addr)-1 {
		return false
	}
	local, domain := addr[:at], addr[at+1:]
	for _, atom := range strings.Split(local, ".") {
		if atom == "" {
			return false
		}
		for i := 0; i < len(atom); i++ {
			c := atom[i]
			if !isAlnum(c) && !strings.ContainsRune("!#$%&'*+-/=?^_`{|}~", rune(c)) {
				return false
			}
		}
	}
	for _, label := range strings.Split(domain, ".") {
		if label == "" || label[0] == '-' || label[len(label)-1] == '-' {
			return false
		}
		for i := 0; i < len(label); i++ {
			if c := label[i]; !isAlnum(c) && c != '-' {
				return false
			}
		}
	}
	return true
}

func isAlnum(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9'
}

// authPlainPayload is the AUTH PLAIN initial response: authzid empty,
// authcid "apikey", password = API key.
func authPlainPayload(apiKey string) string {
	return base64.StdEncoding.EncodeToString([]byte("\x00" + AuthUsername + "\x00" + apiKey))
}

// redactor masks the API key and the AUTH PLAIN payload that carries it. A
// server or proxy may echo the AUTH command back, in a reply, in a
// malformed line quoted by an error, or in a cut-off fragment, so every
// transcript line and the Result are passed through it. The key is matched
// raw and as printable and %q render it. A reply may also wrap an echo
// across its lines, so a reply's lines are first masked together (see
// reply), along with the bytes that ended the read if it failed (see
// replyErr). A secret can also be split across replies, or left in pieces
// that each look harmless, so the transcript masks all of a session's
// server text together before writing any of it (see session).
type redactor struct {
	r       *strings.Replacer
	wrapped []wrappedSecret
}

// wrappedSecret matches one secret form wrapped across the lines of a reply
// or not, and the mask it gets: whole, or any start of it that runs to the
// end of the text, where a cut-off reply may stop inside it.
type wrappedSecret struct {
	secret       string
	whole, start *regexp.Regexp
	mask         string
}

// replyLineBreak is what may sit between two pieces of a secret a reply
// wraps: the line break, with any blanks before it, then the next line's
// reply code, separator and enhanced status code. Each part of that framing
// is optional, as the next line may be one readReply rejects for lacking it.
const replyLineBreak = `(?:[\t\r ]*\n(?:\d{3}[ -]?)?[\t ]*(?:\d\.\d{1,3}\.\d{1,3}[\t ]+)?)?`

const (
	redactedKey     = "<redacted: API key>"
	redactedPayload = "<redacted: AUTH PLAIN credentials>"
)

func newRedactor(apiKey string) redactor {
	pairs := []string{authPlainPayload(apiKey), redactedPayload} // first: wins over a key inside it
	quoted := strconv.Quote(apiKey)
	for _, k := range []string{apiKey, printable(apiKey), quoted[1 : len(quoted)-1]} {
		if !slices.Contains(pairs, k) {
			pairs = append(pairs, k, redactedKey)
		}
	}
	rd := redactor{r: strings.NewReplacer(pairs...)}
	for i := 0; i < len(pairs); i += 2 {
		var whole, start strings.Builder
		for j, c := range pairs[i] {
			char := regexp.QuoteMeta(string(c))
			if j == 0 {
				whole.WriteString(char)
				start.WriteString(char)
				continue
			}
			whole.WriteString(replyLineBreak + char)
			start.WriteString(replyLineBreak + `(?:` + char + `|\z)`) // past a cut, the rest never arrived
		}
		start.WriteString(`\z`)
		rd.wrapped = append(rd.wrapped, wrappedSecret{
			secret: pairs[i],
			whole:  regexp.MustCompile(whole.String()),
			start:  regexp.MustCompile(start.String()),
			mask:   pairs[i+1],
		})
	}
	return rd
}

func (rd redactor) mask(s string) string { return rd.r.Replace(s) }

// reply masks a secret wrapped across a reply's lines, which mask misses
// because no line holds it whole. The masked span takes in the line breaks
// inside the secret, so the lines it spans become one and no piece is left
// to rejoin. Each line still goes through mask afterwards.
func (rd redactor) reply(lines []string) []string {
	if len(lines) == 0 {
		return lines
	}
	text := strings.Join(lines, "\n") // reply lines never hold a LF
	for _, w := range rd.wrapped {
		text = w.whole.ReplaceAllLiteralString(text, w.mask)
	}
	return strings.Split(text, "\n")
}

// replyErr is reply for a reply that readReply gave up on: lines are the
// lines it read, and err may quote the bytes that ended the read, a line that
// is not valid SMTP or one cut off before its CRLF. A secret wrapped across
// the lines can run on into those bytes, so they are masked together with
// the lines, as one more line; lines that a secret spans fold into the
// quoted bytes. A cut can also stop inside a secret, so the start of one that
// runs from the lines into the cut is masked too. err comes back quoting the
// masked bytes, its cause kept; any other err comes back as is.
func (rd redactor) replyErr(lines []string, err error) ([]string, error) {
	switch e := err.(type) {
	case *badLineError:
		lines = rd.reply(append(slices.Clone(lines), e.line))
		return lines[:len(lines)-1], &badLineError{why: e.why, line: lines[len(lines)-1]}
	case *incompleteLineError:
		lines = rd.cutOff(rd.reply(append(slices.Clone(lines), e.partial)))
		return lines[:len(lines)-1], &incompleteLineError{partial: lines[len(lines)-1], err: e.err}
	}
	return rd.reply(lines), err
}

// cutOff masks the start of a secret left at the end of lines, the last of
// which was cut off before its CRLF, if it runs on from an earlier line. A
// start within the last line alone may be chance: a cut-off upper-case UUID
// ending in "A" would begin the AUTH PLAIN payload.
func (rd redactor) cutOff(lines []string) []string {
	text := strings.Join(lines, "\n")
	for _, w := range rd.wrapped {
		if loc := w.start.FindStringIndex(text); loc != nil && strings.Contains(text[loc[0]:], "\n") {
			text = text[:loc[0]] + w.mask
		}
	}
	return strings.Split(text, "\n")
}

// minFragment is how many characters of a secret in a row are masked
// wherever they are, on their own too: a line cut off before the "=="
// padding of the AUTH PLAIN payload decodes, unpadded, to the whole key.
// Text that is not a piece of this session's own secrets does not hold 16
// of their characters in a row by chance.
const minFragment = 16

// heardText is a string of server bytes that the output shows: a reply
// line, or the bytes a read error quotes.
type heardText struct {
	text   string
	quoted bool // quoted by an error, as %q renders it
	cut    bool // cut off before its CRLF: whatever followed never arrived
}

// session masks all of a session's server text together, so that no split
// of a secret leaves pieces to rejoin: across the lines of a reply, across
// replies, between a reply and the bytes an error quotes, or between the
// transcript and the Result. heard is that text in the order the output
// shows it, the Result's repeat of it last. Joined line by line, it is
// searched for
//   - each run of minFragment or more characters of a secret, or all of a
//     shorter secret;
//   - each start of a secret that runs to the end of a cut-off text, the
//     rest never having arrived, and crosses a line break on the way. Blanks
//     or a CR may follow it before the cut. A start within the cut-off line
//     alone may be chance: a cut-off upper-case UUID ending in "A" would
//     begin the AUTH PLAIN payload.
//
// Gaps may split a run (see gapEnds). A run's characters are masked, and the
// gaps between them within a line; the framing that opens a line is kept. A
// text shown more than once is masked the same way each time, by the runs
// found in any of its showings: one showing left unmasked would give away
// what another masked. session returns what may be shown of each text.
func (rd redactor) session(heard []heardText) map[string]string {
	var b strings.Builder
	var cuts []int
	for _, h := range heard {
		b.WriteString(h.text)
		if h.cut {
			cuts = append(cuts, b.Len())
		}
		b.WriteByte('\n')
	}
	text := b.String()
	masks := make([]int, len(text))             // 1 + the index in rd.wrapped of the secret masked at each byte, or 0
	for i := len(rd.wrapped) - 1; i >= 0; i-- { // the first secret wins, as in mask
		for p, in := range runs(text, rd.wrapped[i].secret, cuts) {
			if in {
				masks[p] = i + 1
			}
		}
	}
	union := map[string][]int{}
	at := 0
	for _, h := range heard {
		m := masks[at : at+len(h.text)]
		at += len(h.text) + 1
		u, ok := union[h.text]
		if !ok {
			union[h.text] = slices.Clone(m)
			continue
		}
		for i, v := range m {
			if v != 0 && (u[i] == 0 || v < u[i]) {
				u[i] = v
			}
		}
	}
	shown := make(map[string]string, len(union))
	for s, m := range union {
		var out strings.Builder
		for i := 0; i < len(s); {
			j := i + 1
			for j < len(s) && m[j] == m[i] {
				j++
			}
			if m[i] == 0 {
				out.WriteString(s[i:j])
			} else {
				out.WriteString(rd.wrapped[m[i]-1].mask)
			}
			i = j
		}
		shown[s] = out.String()
	}
	return shown
}

// runs reports, for each byte of text, whether session masks it for secret;
// cuts are the ends of cut-off texts. A run is a path through states, each a
// byte of text read as a character of secret, from one state to the state
// reading the next character at the next byte or past a gap. The longest
// run through a state is the longest path into it joined to the longest
// path out of it.
func runs(text, secret string, cuts []int) []bool {
	n := len(text)
	var at [256][]int // the indexes in secret of each byte value
	for j := 0; j < len(secret); j++ {
		at[secret[j]] = append(at[secret[j]], j)
	}
	// The states of text[p] are numbered from first[p], one per index in
	// at[text[p]]; pos and char map a state back to its byte and index.
	first := make([]int, n+1)
	for p := 0; p < n; p++ {
		first[p+1] = first[p] + len(at[text[p]])
	}
	pos, char := make([]int, first[n]), make([]int, first[n])
	for p := 0; p < n; p++ {
		for i, j := range at[text[p]] {
			pos[first[p]+i], char[first[p]+i] = p, j
		}
	}
	gaps := map[int][]int{}
	ends := func(q int) []int {
		e, ok := gaps[q]
		if !ok {
			e = gapEnds(text, q)
			gaps[q] = e
		}
		return e
	}
	// next calls f with each state that can follow state s, the byte it
	// reads and whether a gap comes between.
	next := func(s int, f func(s2, r int, gap bool)) {
		p, j := pos[s], char[s]
		if j+1 == len(secret) {
			return
		}
		follow := func(r int, gap bool) {
			if i := slices.Index(at[text[r]], j+1); i >= 0 {
				f(first[r]+i, r, gap)
			}
		}
		if p+1 < n {
			follow(p+1, false)
		}
		for _, r := range ends(p + 1) {
			if r < n {
				follow(r, true)
			}
		}
	}
	into := make([]int, first[n]) // longest path into each state, the state included
	// For each state a path from secret[0] leads into: whether one crosses
	// a gap, and the states before it on such paths.
	crossed := make([]bool, first[n])
	before := map[int][]int{}
	for s := range into {
		into[s] = max(into[s], 1)
		next(s, func(s2, _ int, gap bool) {
			into[s2] = max(into[s2], into[s]+1)
			if into[s] == char[s]+1 {
				crossed[s2] = crossed[s2] || crossed[s] || gap
				before[s2] = append(before[s2], s)
			}
		})
	}
	out := make([]int, first[n]) // longest path out of each state, the state included
	for s := len(out) - 1; s >= 0; s-- {
		out[s] = 1
		next(s, func(s2, _ int, _ bool) { out[s] = max(out[s], out[s2]+1) })
	}

	in := make([]bool, n)
	// gap marks the bytes between two characters of a run read at p and r,
	// unless they end a line: the framing of the next line is kept.
	gap := func(p, r int) {
		if !strings.Contains(text[p+1:r], "\n") {
			for q := p + 1; q < r; q++ {
				in[q] = true
			}
		}
	}
	// Where bytes read as a run in more than one way, as when the framing
	// that opens a line could also be read as characters of the secret, the
	// longest reading wins: a state is masked only on a longest path, one of
	// tight edges from a state nothing leads into (head) to a state that
	// leads nowhere (tail).
	tight := func(s, s2 int) bool { return into[s2] == into[s]+1 && out[s] == out[s2]+1 }
	head, tail := make([]bool, first[n]), make([]bool, first[n])
	for s := range head {
		head[s] = head[s] || into[s] == 1
		next(s, func(s2, _ int, _ bool) {
			if head[s] && tight(s, s2) {
				head[s2] = true
			}
		})
	}
	for s := len(tail) - 1; s >= 0; s-- {
		tail[s] = out[s] == 1
		next(s, func(s2, _ int, _ bool) {
			if tail[s2] && tight(s, s2) {
				tail[s] = true
			}
		})
	}
	long := min(minFragment, len(secret))
	for s := range into {
		if into[s]+out[s]-1 < long || !head[s] || !tail[s] {
			continue
		}
		in[pos[s]] = true
		next(s, func(s2, r int, isGap bool) {
			if isGap && tight(s, s2) && tail[s2] {
				gap(pos[s], r)
			}
		})
	}
	// blankTo reports whether only blanks, after a gap or not, lie between q
	// and the cut at e. A CR at the cut ends a line, so is a gap.
	blankTo := func(q, e int) bool {
		for _, g := range append([]int{q}, ends(q)...) {
			if g <= e && strings.Trim(text[g:e], " \t") == "" {
				return true
			}
		}
		return false
	}
	walked := map[int]bool{}
	for _, e := range cuts {
		for s := range into {
			p := pos[s]
			if p >= e || into[s] != char[s]+1 || !blankTo(p+1, e) {
				continue
			}
			if !crossed[s] && !strings.ContainsAny(strings.TrimRight(text[p+1:e], " \t\r"), "\r\n") {
				continue
			}
			for todo := []int{s}; len(todo) > 0; {
				s := todo[len(todo)-1]
				todo = todo[:len(todo)-1]
				if walked[s] {
					continue
				}
				walked[s] = true
				in[pos[s]] = true
				for _, s0 := range before[s] {
					gap(pos[s0], pos[s])
					todo = append(todo, s0)
				}
			}
		}
	}
	return in
}

// gapEnds returns each i such that text[q:i] is a gap: what may sit between
// two characters of a secret that a server splits. That is one or more line
// breaks, each with any blanks around it and the framing that may open the
// next line: reply code, separator and enhanced status code, each optional,
// as in replyLineBreak. A break is a LF, or a CR: a server may end a line
// with a bare CR, which readLine keeps in the line it reads.
func gapEnds(text string, q int) []int {
	if q >= len(text) || strings.IndexByte(" \t\r\n", text[q]) < 0 {
		return nil
	}
	var ends []int
	seen := map[int]bool{}
	end := func(i int) {
		if !seen[i] {
			seen[i] = true
			ends = append(ends, i)
		}
	}
	blanks := func(i int) int { // a gap may end after each blank from i
		for i < len(text) && (text[i] == ' ' || text[i] == '\t') {
			i++
			end(i)
		}
		return i
	}
	for todo, done := []int{q}, map[int]bool{}; len(todo) > 0; {
		i := todo[len(todo)-1]
		todo = todo[:len(todo)-1]
		if done[i] {
			continue
		}
		done[i] = true
		for i < len(text) && (text[i] == ' ' || text[i] == '\t') {
			i++
		}
		if i == len(text) || text[i] != '\r' && text[i] != '\n' {
			continue
		}
		for i < len(text) && strings.IndexByte(" \t\r\n", text[i]) >= 0 {
			i++
			end(i)
		}
		k := i
		if k+3 <= len(text) && isDigit(text[k]) && isDigit(text[k+1]) && isDigit(text[k+2]) {
			k += 3
			end(k)
			if k < len(text) && text[k] == '-' {
				k++
				end(k)
			}
			k = blanks(k)
		}
		for _, j := range []int{i, k} {
			if m := enhancedFraming.FindStringIndex(text[j:]); m != nil {
				end(j + m[1])
				todo = append(todo, blanks(j+m[1]))
			}
		}
		todo = append(todo, k) // another line break may follow
	}
	return ends
}

var enhancedFraming = regexp.MustCompile(`^\d\.\d{1,3}\.\d{1,3}[\t ]`)

func isDigit(c byte) bool { return c >= '0' && c <= '9' }

// transcript holds the session's output until the session ends, so that the
// server text in all of it, and in the Result, is redacted together (see
// redactor.session) before any of it is written.
type transcript struct {
	w      io.Writer
	redact redactor
	lines  []tline
	heard  []heardText       // the server text the lines show, in order
	shown  map[string]string // what may be shown of it, once flushed
}

// tline is one transcript line; text renders it once the session's server
// text is redacted.
type tline struct {
	stage Stage
	dir   string
	text  func() string
}

func (t *transcript) add(stage Stage, dir string, text func() string) {
	t.lines = append(t.lines, tline{stage, dir, text})
}

func (t *transcript) client(stage Stage, text string) {
	t.add(stage, "C:", func() string { return text })
}

// note adds a local note, which may quote what an error quotes of a reply.
func (t *transcript) note(stage Stage, text string) {
	t.add(stage, "--", func() string { return t.requote(text) })
}

// server adds the lines of a reply that redactor.replyErr has masked, and
// the bytes the error that ended the read quotes, if any. A line cut off
// before its CRLF is printed after the lines; a rejected line is printed
// where the error is.
func (t *transcript) server(stage Stage, lines []string, err error) {
	for _, l := range lines {
		t.heard = append(t.heard, heardText{text: l})
		t.add(stage, "S:", func() string { return printable(t.shown[l]) })
	}
	switch e := err.(type) {
	case *badLineError:
		t.heard = append(t.heard, heardText{text: e.line, quoted: true})
	case *incompleteLineError:
		t.heard = append(t.heard, heardText{text: e.partial, quoted: true, cut: true})
		t.add(stage, "--", func() string { return "incomplete reply line (no CRLF): " + printable(t.shown[e.partial]) })
	}
}

// requote shows the reply bytes text quotes as an error does (%q) as
// session redaction left them.
func (t *transcript) requote(text string) string {
	for _, h := range t.heard {
		if shown, ok := t.shown[h.text]; ok && h.quoted && shown != h.text {
			text = strings.ReplaceAll(text, strconv.Quote(h.text), strconv.Quote(shown))
		}
	}
	return text
}

// flush masks the session's server text, the Result's repeat of it
// included, writes the transcript, and returns res with its server text
// shown as the transcript shows it.
func (t *transcript) flush(res Result) Result {
	heard := slices.Clone(t.heard)
	var lines []string
	if res.Reply != nil {
		lines = t.redact.reply(res.Reply.Lines)
		for _, l := range lines {
			heard = append(heard, heardText{text: l})
		}
	}
	for _, h := range t.heard {
		if h.quoted && strings.Contains(res.Detail, strconv.Quote(h.text)) {
			heard = append(heard, h)
		}
	}
	t.shown = t.redact.session(heard)
	for _, l := range t.lines {
		fmt.Fprintf(t.w, "%-11s %s %s\n", "["+string(l.stage)+"]", l.dir, t.redact.mask(l.text()))
	}
	if res.Reply != nil {
		reply := *res.Reply
		reply.Lines = make([]string, len(lines))
		for i, l := range lines {
			reply.Lines[i] = t.redact.mask(t.shown[l])
		}
		res.Reply = &reply
	}
	res.Detail = t.redact.mask(t.requote(res.Detail))
	res.Hint = t.redact.mask(res.Hint)
	return res
}

type session struct {
	cfg  Config
	t    *transcript
	conn net.Conn
	r    *bufio.Reader
}

func (s *session) run(ctx context.Context) Result {
	cfg := s.cfg
	addr := net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port))
	s.t.note(StageConnect, fmt.Sprintf("connecting to %s (from <%s>, to <%s>)", addr, cfg.From, cfg.To))

	dialCtx, cancel := context.WithTimeout(ctx, cfg.CommandTimeout)
	conn, err := cfg.Dial(dialCtx, "tcp", addr)
	cancel()
	if err != nil {
		return s.dialFailure(err)
	}
	defer conn.Close()
	stop := context.AfterFunc(ctx, func() { conn.Close() })
	defer stop()
	s.conn, s.r = conn, bufio.NewReader(conn)
	s.t.note(StageConnect, "connected to "+conn.RemoteAddr().String())

	// Banner.
	banner, err := s.read(StageConnect, cfg.CommandTimeout)
	if err != nil {
		return s.ioFailure(ctx, StageConnect, err, cfg.CommandTimeout)
	}
	if banner.Code != 220 {
		return s.replyFailure(StageConnect, banner)
	}

	caps, res, ok := s.ehlo(ctx)
	if !ok {
		return res
	}

	// STARTTLS is mandatory: credentials never go out unencrypted.
	if _, offered := caps["STARTTLS"]; !offered {
		s.quit()
		return Result{
			Stage:  StageSTARTTLS,
			Detail: "server did not advertise STARTTLS; refusing to send credentials without TLS",
			Hint:   hintNoSTARTTLS,
		}
	}
	if res, ok := s.startTLS(ctx); !ok {
		return res
	}
	if caps, res, ok = s.ehlo(ctx); !ok {
		return res
	}

	if !hasAuthMechanism(caps, "PLAIN") {
		s.t.note(StageAuth, "server did not advertise AUTH PLAIN after STARTTLS; trying it anyway")
	}
	reply, err := s.cmd(StageAuth, "AUTH PLAIN "+authPlainPayload(cfg.APIKey), `AUTH PLAIN <redacted: username "apikey", password = API key>`, cfg.CommandTimeout)
	if err != nil {
		return s.ioFailure(ctx, StageAuth, err, cfg.CommandTimeout)
	}
	if reply.Code == 334 {
		// A continuation challenge: cancel the exchange (RFC 4954) so the
		// QUIT that follows is read as a command, not as credentials.
		cancelReply, err := s.cmd(StageAuth, "*", "", cfg.CommandTimeout)
		if err != nil {
			s.t.note(StageAuth, "no reply to *: "+describeIOError(err, cfg.CommandTimeout))
		}
		if err != nil || cancelReply.Code == 334 || cancelReply.Code == 421 {
			return Result{Stage: StageAuth, Reply: &reply, Hint: hintForReply(StageAuth, reply)}
		}
	}
	if reply.Code != 235 {
		return s.replyFailure(StageAuth, reply)
	}

	reply, err = s.cmd(StageMailFrom, "MAIL FROM:<"+cfg.From+">", "", cfg.CommandTimeout)
	if err != nil {
		return s.ioFailure(ctx, StageMailFrom, err, cfg.CommandTimeout)
	}
	if reply.Code != 250 {
		return s.replyFailure(StageMailFrom, reply)
	}

	reply, err = s.cmd(StageRcptTo, "RCPT TO:<"+cfg.To+">", "", cfg.CommandTimeout)
	if err != nil {
		return s.ioFailure(ctx, StageRcptTo, err, cfg.CommandTimeout)
	}
	if reply.Code != 250 && reply.Code != 251 {
		return s.replyFailure(StageRcptTo, reply)
	}

	reply, err = s.cmd(StageData, "DATA", "", cfg.CommandTimeout)
	if err != nil {
		return s.ioFailure(ctx, StageData, err, cfg.CommandTimeout)
	}
	if reply.Code != 354 {
		return s.replyFailure(StageData, reply)
	}
	reply, err = s.sendMessage()
	if err != nil {
		return s.ioFailure(ctx, StageData, err, cfg.DataTimeout)
	}
	if reply.Code != 250 {
		return s.replyFailure(StageData, reply)
	}
	uuid := messageUUID(reply)
	if uuid == "" {
		s.quit()
		return Result{Stage: StageData, Reply: &reply, Hint: hintNoUUID}
	}

	s.quit()
	return Result{Verified: true, MessageUUID: uuid, Stage: StageQuit}
}

func (s *session) ehlo(ctx context.Context) (map[string]string, Result, bool) {
	reply, err := s.cmd(StageEHLO, "EHLO "+s.cfg.EHLOName, "", s.cfg.CommandTimeout)
	if err != nil {
		return nil, s.ioFailure(ctx, StageEHLO, err, s.cfg.CommandTimeout), false
	}
	if reply.Code != 250 {
		return nil, s.replyFailure(StageEHLO, reply), false
	}
	return capabilities(reply), Result{}, true
}

func (s *session) startTLS(ctx context.Context) (Result, bool) {
	timeout := s.cfg.CommandTimeout
	reply, err := s.cmd(StageSTARTTLS, "STARTTLS", "", timeout)
	if err != nil {
		return s.ioFailure(ctx, StageSTARTTLS, err, timeout), false
	}
	if reply.Code != 220 {
		return s.replyFailure(StageSTARTTLS, reply), false
	}
	// Anything already buffered arrived in plaintext after the 220 and
	// would be read as if it came over TLS (STARTTLS command injection).
	if n := s.r.Buffered(); n > 0 {
		return Result{
			Stage:  StageSTARTTLS,
			Detail: fmt.Sprintf("server sent %d unexpected plaintext bytes after its STARTTLS reply", n),
			Hint:   hintInjected,
		}, false
	}

	tlsCfg := &tls.Config{}
	if s.cfg.TLSConfig != nil {
		tlsCfg = s.cfg.TLSConfig.Clone()
	}
	if tlsCfg.ServerName == "" {
		tlsCfg.ServerName = s.cfg.Host
	}
	if tlsCfg.MinVersion == 0 {
		tlsCfg.MinVersion = tls.VersionTLS12
	}
	tc := tls.Client(s.conn, tlsCfg)
	_ = tc.SetDeadline(time.Now().Add(timeout))
	if err := tc.HandshakeContext(ctx); err != nil {
		var certErr *tls.CertificateVerificationError
		var recordErr tls.RecordHeaderError
		switch {
		case ctx.Err() == nil && errors.As(err, &certErr):
			return Result{Stage: StageSTARTTLS, Detail: "TLS handshake failed: " + err.Error(), Hint: hintTLS}, false
		case ctx.Err() == nil && errors.As(err, &recordErr):
			return Result{Stage: StageSTARTTLS, Detail: "TLS handshake failed: " + err.Error(), Hint: hintInjected}, false
		}
		res := s.ioFailure(ctx, StageSTARTTLS, err, timeout) // timeout, cancel, EOF, reset
		res.Detail = "TLS handshake: " + res.Detail
		return res, false
	}
	st := tc.ConnectionState()
	verified := fmt.Sprintf("certificate verified for %q", tlsCfg.ServerName)
	if tlsCfg.InsecureSkipVerify {
		verified = "certificate NOT verified"
	}
	s.t.note(StageSTARTTLS, fmt.Sprintf("TLS established: %s, %s, %s",
		tls.VersionName(st.Version), tls.CipherSuiteName(st.CipherSuite), verified))
	s.conn, s.r = tc, bufio.NewReader(tc)
	return Result{}, true
}

// sendMessage writes the dot-stuffed message and reads the final DATA reply.
func (s *session) sendMessage() (Reply, error) {
	msg, msgID := s.buildMessage()
	_ = s.conn.SetDeadline(time.Now().Add(s.cfg.DataTimeout))
	wireBytes := len(msg) + strings.Count(string(msg), "\n") // LF goes out as CRLF
	s.t.client(StageData, fmt.Sprintf("<message: %d bytes, Message-ID %s>", wireBytes, msgID))
	bw := bufio.NewWriter(s.conn)
	dw := textproto.NewWriter(bw).DotWriter()
	if _, err := dw.Write(msg); err != nil {
		return Reply{}, err
	}
	if err := dw.Close(); err != nil { // writes "." and flushes
		return Reply{}, err
	}
	s.t.client(StageData, ".")
	return s.receive(StageData)
}

func (s *session) buildMessage() ([]byte, string) {
	now := s.cfg.Now()
	var token [8]byte
	_, _ = rand.Read(token[:])
	domain := s.cfg.From[strings.LastIndex(s.cfg.From, "@")+1:]
	msgID := "<telnyx-smtp-verify." + hex.EncodeToString(token[:]) + "@" + domain + ">"
	var b strings.Builder
	fmt.Fprintf(&b, "From: <%s>\n", s.cfg.From)
	fmt.Fprintf(&b, "To: <%s>\n", s.cfg.To)
	fmt.Fprintf(&b, "Subject: Telnyx SMTP verify %s\n", now.UTC().Format(time.RFC3339))
	fmt.Fprintf(&b, "Date: %s\n", now.Format(time.RFC1123Z))
	fmt.Fprintf(&b, "Message-ID: %s\n", msgID)
	b.WriteString("MIME-Version: 1.0\n")
	b.WriteString("Content-Type: text/plain; charset=us-ascii\n")
	b.WriteString("Content-Transfer-Encoding: 7bit\n")
	b.WriteString("\n")
	fmt.Fprintf(&b, "This message was sent by \"telnyx email smtp verify\" to check SMTP\nsubmission through %s:%d. No action is needed.\n", s.cfg.Host, s.cfg.Port)
	return []byte(b.String()), msgID
}

// cmd sends one command line and reads its reply. display replaces the
// command text in the transcript (used to redact credentials).
func (s *session) cmd(stage Stage, line, display string, timeout time.Duration) (Reply, error) {
	if display == "" {
		display = line
	}
	_ = s.conn.SetDeadline(time.Now().Add(timeout))
	s.t.client(stage, display)
	if _, err := io.WriteString(s.conn, line+"\r\n"); err != nil {
		return Reply{}, err
	}
	return s.receive(stage)
}

func (s *session) read(stage Stage, timeout time.Duration) (Reply, error) {
	_ = s.conn.SetDeadline(time.Now().Add(timeout))
	return s.receive(stage)
}

// receive reads one reply and adds it to the transcript. Partial replies are
// shown too, including a line cut off before its CRLF: every reply byte is
// shown. The lines are redacted together with the bytes the error quotes,
// and the error returned quotes those bytes redacted, so no later message (a
// note, the verdict) holds a piece of a secret that rejoins the lines; the
// transcript then redacts them with the rest of the session's server text.
func (s *session) receive(stage Stage) (Reply, error) {
	reply, err := readReply(s.r)
	lines, err := s.t.redact.replyErr(reply.Lines, err)
	s.t.server(stage, lines, err)
	return reply, err
}

// quit ends the session politely; its outcome never changes the verdict.
func (s *session) quit() {
	timeout := min(s.cfg.CommandTimeout, quitTimeout)
	if _, err := s.cmd(StageQuit, "QUIT", "", timeout); err != nil {
		s.t.note(StageQuit, "no reply to QUIT: "+describeIOError(err, timeout))
	}
}

func (s *session) replyFailure(stage Stage, r Reply) Result {
	// 421 means the server is closing the channel; anything else gets a QUIT.
	if r.Code != 421 {
		s.quit()
	}
	return Result{Stage: stage, Reply: &r, Hint: hintForReply(stage, r)}
}

func (s *session) ioFailure(ctx context.Context, stage Stage, err error, timeout time.Duration) Result {
	if ctx.Err() != nil {
		return Result{Stage: stage, Detail: "canceled: " + ctx.Err().Error(), Hint: "The command was interrupted."}
	}
	var cut *incompleteLineError
	if errors.As(err, &cut) {
		return Result{Stage: stage, Detail: fmt.Sprintf("incomplete reply %q (no CRLF): %s", cut.partial, describeIOError(cut.err, timeout)), Hint: hintIncomplete}
	}
	var ne net.Error
	if errors.As(err, &ne) && ne.Timeout() {
		return Result{Stage: stage, Detail: describeIOError(err, timeout), Hint: hintTimeout(timeout)}
	}
	if errors.Is(err, errMalformedReply) {
		return Result{Stage: stage, Detail: err.Error(), Hint: "The server's reply is not valid SMTP. Check --host and --port."}
	}
	return Result{Stage: stage, Detail: describeIOError(err, timeout), Hint: "The connection failed mid-session. Retry; if it persists, check connectivity to the host."}
}

func (s *session) dialFailure(err error) Result {
	var dnsErr *net.DNSError
	if errors.As(err, &dnsErr) {
		return Result{Stage: StageConnect, Detail: err.Error(), Hint: hintDNS}
	}
	var ne net.Error
	if errors.As(err, &ne) && ne.Timeout() {
		return Result{Stage: StageConnect, Detail: fmt.Sprintf("timed out after %s connecting", s.cfg.CommandTimeout), Hint: hintDial}
	}
	return Result{Stage: StageConnect, Detail: err.Error(), Hint: hintDial}
}

func describeIOError(err error, timeout time.Duration) string {
	var ne net.Error
	switch {
	case errors.As(err, &ne) && ne.Timeout():
		return fmt.Sprintf("timed out after %s waiting for the server", timeout)
	case errors.Is(err, io.EOF), errors.Is(err, io.ErrUnexpectedEOF):
		return "connection closed by the server"
	default:
		return err.Error()
	}
}
