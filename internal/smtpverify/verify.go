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
	"maps"
	"net"
	"net/mail"
	"net/textproto"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"sync"
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
// session ends, as all of its server text is redacted together. Canceling
// ctx ends the session early, with an "interrupted" verdict (a deadline
// passing, with a "canceled" one); the transcript so far is still written,
// redacted.
func Run(ctx context.Context, cfg Config, out io.Writer) (Result, error) {
	cfg, err := cfg.normalize()
	if err != nil {
		return Result{}, err
	}
	s := &session{cfg: cfg, t: &transcript{w: out, redact: newRedactor(cfg.APIKey)}}
	res := s.t.flush(ctx, s.run(ctx))
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
// end of the text, where a cut-off reply may stop inside it. grams holds
// each minFragment characters in a row of it, or all of it if shorter (see
// cover). whole and start are compiled when first used.
type wrappedSecret struct {
	secret       string
	whole, start func() *regexp.Regexp
	mask         string
	grams        map[string]bool
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
		grams := map[string]bool{}
		for n, j := min(minFragment, len(pairs[i])), 0; j+n <= len(pairs[i]); j++ {
			grams[pairs[i][j:j+n]] = true
		}
		wholeRE, startRE := whole.String(), start.String()
		rd.wrapped = append(rd.wrapped, wrappedSecret{
			secret: pairs[i],
			whole:  sync.OnceValue(func() *regexp.Regexp { return regexp.MustCompile(wholeRE) }),
			start:  sync.OnceValue(func() *regexp.Regexp { return regexp.MustCompile(startRE) }),
			mask:   pairs[i+1],
			grams:  grams,
		})
	}
	return rd
}

// mask masks each piece of minFragment or more characters of a secret in s
// (see cover), and each whole secret.
func (rd redactor) mask(s string) string {
	m := make([]int, len(s))
	rd.cover(s, m)
	return rd.r.Replace(rd.render(s, m))
}

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
		text = w.whole().ReplaceAllLiteralString(text, w.mask)
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
		if loc := w.start().FindStringIndex(text); loc != nil && strings.Contains(text[loc[0]:], "\n") {
			text = text[:loc[0]] + w.mask
		}
	}
	return strings.Split(text, "\n")
}

// minFragment is how many characters of a secret in a row are masked
// wherever they are (see cover), and how long a run of a secret that line
// breaks and framing split must be to be masked (see spanning). It is the
// shortest length that server text which is not this session's own secret
// all but never matches by chance. TestCoverageFalsePositives runs 40,000
// random keys of the Telnyx format ("KEY", 32 upper-case hex digits, "_", 22
// letters and digits), each with a session of server text far denser in
// UUIDs, upper-case hex queue IDs, dates, Message-IDs, addresses and MIME
// base64 than a real one: 7 masks nothing there (and text in 5 of 400,000
// such sessions); 6 masked text in 7 of the 40,000, and 5 in 176. Upper-case
// hex queue IDs and UUIDs, and digits in trace IDs, dates and
// "SIZE 25000000", matched the key's hex digits. Shorter pieces are shown:
// up to 6 characters of a key are what chance text holds too.
const minFragment = 7

// heardText is a string of server bytes that the output shows: a reply
// line, or the bytes a read error quotes.
type heardText struct {
	text   string
	quoted bool // quoted by an error, as %q renders it
	cut    bool // cut off before its CRLF: whatever followed never arrived
	lower  bool // shown lower-cased, as the message UUID is (see messageUUID)
}

// session masks all of a session's server text together, so that no split
// of a secret leaves pieces to rejoin: across the lines of a reply, across
// replies, between a reply and the bytes an error quotes, or between the
// transcript and the Result, the message UUID that VERIFIED prints
// included. heard is that text in the order the output shows it, the
// Result's repeat of it last. Joined line by line, it is searched twice: as
// the server sent it, and with a text shown lower-cased, the message UUID,
// lower-cased wherever it is. VERIFIED shows the UUID lower-cased, after the
// QUIT reply, and the reply that carried it shows it as sent, before the
// QUIT reply, so a reader can join a run of a secret into the UUID or out of
// it in either case. Each search looks for
//   - each piece of minFragment or more characters of a secret, or all of a
//     shorter secret, wherever it is and whatever is around it (see cover);
//   - each run of minFragment or more characters of a secret that a gap
//     splits, its pieces shorter perhaps (see spanning);
//   - each start of a secret that runs to the end of a cut-off text, the
//     rest never having arrived, and crosses a line break on the way. Blanks
//     or a CR may follow it before the cut. A start within the cut-off line
//     alone may be chance: a cut-off upper-case UUID ending in "A" would
//     begin the AUTH PLAIN payload.
//
// What either search finds is masked: a piece or run's characters, and the
// gaps between them within a line; the framing that opens a line is kept. A
// text shown more than once is masked the same way each time, by what was
// found in any of its showings, and so is a text shown lower-cased wherever
// it is, in either case: one showing left unmasked would give away what
// another masked. session returns what may be shown of each text. It gives
// up with errRedactBudget once ctx is done, the two searches sharing its
// budget.
func (rd redactor) session(ctx context.Context, heard []heardText) (map[string]string, error) {
	var b strings.Builder
	var cuts []int
	lower := map[string]bool{}
	for _, h := range heard {
		b.WriteString(h.text)
		if h.cut {
			cuts = append(cuts, b.Len())
		}
		b.WriteByte('\n')
		if h.lower && h.text != "" {
			lower[h.text] = true
		}
	}
	text := b.String()
	lowered := slices.Sorted(maps.Keys(lower))
	asShown := text // lower-casing keeps every byte where it is
	for _, s := range lowered {
		asShown = strings.ReplaceAll(asShown, s, lowerASCII(s))
	}
	views := []string{text}
	if asShown != text {
		views = append(views, asShown)
	}
	masks := make([]int, len(text)) // 1 + the index in rd.wrapped of the secret masked at each byte, or 0
	for _, v := range views {
		rd.cover(v, masks)
	}
	for i, w := range rd.wrapped {
		for _, v := range views {
			in, err := spanning(ctx, v, w.secret, cuts)
			if err != nil {
				return nil, err
			}
			for p, ok := range in {
				if ok && (masks[p] == 0 || i+1 < masks[p]) { // the first secret wins, as in mask
					masks[p] = i + 1
				}
			}
		}
	}
	// A text shown lower-cased is masked alike wherever it shows, in either
	// case.
	for _, s := range lowered {
		s = lowerASCII(s)
		var at []int
		for p := 0; ; p += len(s) {
			i := strings.Index(asShown[p:], s)
			if i < 0 {
				break
			}
			p += i
			at = append(at, p)
		}
		m := make([]int, len(s))
		for _, p := range at {
			merge(m, masks[p:p+len(s)])
		}
		for _, p := range at {
			merge(masks[p:p+len(s)], m)
		}
	}
	union := map[string][]int{}
	at := 0
	for _, h := range heard {
		m := masks[at : at+len(h.text)]
		at += len(h.text) + 1
		if u, ok := union[h.text]; ok {
			merge(u, m)
		} else {
			union[h.text] = slices.Clone(m)
		}
	}
	shown := make(map[string]string, len(union))
	for s, m := range union {
		if lower[s] {
			shown[s] = rd.render(lowerASCII(s), m)
		} else {
			shown[s] = rd.render(s, m)
		}
	}
	return shown, nil
}

// cover sets m[i] for each byte of text that is part of a piece of
// minFragment or more characters of a secret, all of a shorter one, to 1 +
// the index in rd.wrapped of that secret; the first secret wins, as in mask.
// Those bytes are the windows of text that are minFragment characters of a
// secret, run together: every piece that long of a secret is masked wherever
// it is, whatever text is around it, in whatever order and however often the
// pieces come. Each window is looked up in the set of the secret's own, so
// cover takes time linear in the length of text.
func (rd redactor) cover(text string, m []int) {
	for i := len(rd.wrapped) - 1; i >= 0; i-- {
		w := rd.wrapped[i]
		n := min(minFragment, len(w.secret))
		for j, end := 0, 0; j+n <= len(text); j++ {
			if w.grams[text[j:j+n]] {
				for k := max(j, end); k < j+n; k++ {
					m[k] = i + 1
				}
				end = j + n
			}
		}
	}
}

// merge marks in dst each byte that src marks; the first secret wins, as in
// mask.
func merge(dst, src []int) {
	for i, v := range src {
		if v != 0 && (dst[i] == 0 || v < dst[i]) {
			dst[i] = v
		}
	}
}

// lowerASCII lower-cases the ASCII letters of s, byte for byte.
func lowerASCII(s string) string {
	b := []byte(s)
	for i, c := range b {
		if 'A' <= c && c <= 'Z' {
			b[i] = c + 'a' - 'A'
		}
	}
	return string(b)
}

// render shows s with each run of bytes that m marks for the same secret
// replaced by that secret's mask.
func (rd redactor) render(s string, m []int) string {
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
	return out.String()
}

// errRedactBudget is session giving up: masking the server text took longer
// than its budget (see redactBudget).
var errRedactBudget = errors.New("redaction budget exceeded")

// redactBudget is how long session may take for n bytes of server text: 1 s
// per MB, at least 1 s. Masking takes time linear in n, a few ms per 100 KB,
// so only a fault reaches it; the transcript is then withheld (see flush).
var redactBudget = func(n int) time.Duration {
	return max(time.Second, time.Duration(n)*time.Second/(1<<20))
}

// maxSpanStates bounds the states spanning keeps (about 9 bytes each): the
// text's length times how often a character occurs in the secret.
const maxSpanStates = 1 << 24

// spanning reports, for each byte of text, whether it is part of a run of a
// secret that a gap splits, or of a start of one that runs into a cut; cuts
// are the ends of cut-off texts. Its pieces may each be shorter than
// minFragment, which cover would miss. A run is a path through states, each
// a byte of text read as a character of secret, from one state to the state
// reading the next character at the next byte or past a gap. The longest run
// through a state is the longest path into it joined to the longest path out
// of it.
//
// A state has at most one successor at the next byte and one at each end of
// the gap after it, and a gap has a bounded number of ends (see gapEnds), so
// each of the passes over the states, in order or in reverse, takes time
// linear in their number: at most the length of text times the number of
// times a character occurs in secret. Nothing is cached per byte. spanning
// gives up with errRedactBudget once ctx is done.
func spanning(ctx context.Context, text, secret string, cuts []int) ([]bool, error) {
	sc := newSpanScan(text, secret)
	if sc.first[len(text)] > maxSpanStates {
		return nil, errRedactBudget
	}
	n, states := len(text), sc.first[len(text)]
	into := make([]int32, states) // longest path into each state, the state included
	out := make([]int32, states)  // longest path out of each state, the state included
	flags := make([]uint8, states)
	const (
		crossed = 1 << iota // a path from secret[0] into the state crosses a gap
		head                // the state is on a longest reading, from where it starts...
		tail                // ...and on to where it ends
		gapIn               // a longest reading into the state crosses a gap
		gapOut              // a longest reading out of the state crosses a gap
		walked              // the state is on a start of the secret that runs into a cut
	)
	var next []edge
	expired := func(p int) bool { return p%4096 == 0 && ctx.Err() != nil }

	for p := range n {
		if expired(p) {
			return nil, errRedactBudget
		}
		sc.visit(p)
		for k, j := range sc.at[text[p]] {
			s := sc.first[p] + int32(k)
			into[s] = max(into[s], 1)
			fromStart := into[s] == j+1
			next = sc.next(next[:0], p, j)
			for _, e := range next {
				into[e.s] = max(into[e.s], into[s]+1)
				if fromStart && (e.gap || flags[s]&crossed != 0) {
					flags[e.s] |= crossed
				}
			}
		}
	}
	for p := n - 1; p >= 0; p-- {
		if expired(p) {
			return nil, errRedactBudget
		}
		sc.visit(p)
		for k, j := range sc.at[text[p]] {
			s := sc.first[p] + int32(k)
			out[s] = 1
			next = sc.next(next[:0], p, j)
			for _, e := range next {
				out[s] = max(out[s], out[e.s]+1)
			}
		}
	}

	// Where bytes read as a run in more than one way, as when the framing
	// that opens a line could also be read as characters of the secret, the
	// longest reading wins: a state is masked only on a longest path, one of
	// tight edges from a state nothing leads into (head) to a state that
	// leads nowhere (tail), and only if that path crosses a gap; cover masks
	// the rest.
	tight := func(s, s2 int32) bool { return into[s2] == into[s]+1 && out[s] == out[s2]+1 }
	for p := range n {
		if expired(p) {
			return nil, errRedactBudget
		}
		sc.visit(p)
		for k, j := range sc.at[text[p]] {
			s := sc.first[p] + int32(k)
			if into[s] == 1 {
				flags[s] |= head
			}
			if flags[s]&head == 0 {
				continue
			}
			next = sc.next(next[:0], p, j)
			for _, e := range next {
				if tight(s, e.s) {
					flags[e.s] |= head
					if e.gap || flags[s]&gapIn != 0 {
						flags[e.s] |= gapIn
					}
				}
			}
		}
	}

	in := make([]bool, n)
	// gap marks the bytes between two characters of a run read at p and r,
	// unless they end a line: the framing of the next line is kept.
	gap := func(p, r int) {
		if strings.IndexByte(text[p+1:r], '\n') < 0 {
			for q := p + 1; q < r; q++ {
				in[q] = true
			}
		}
	}
	// A start of the secret runs into the cut at e if only blanks lie
	// between its last character and e, after a gap or not. A CR at the cut
	// ends a line, so is a gap. blanks[i] is where the blanks before cuts[i]
	// begin, and blanksCR[i] where the blanks and CRs do.
	blanks, blanksCR := make([]int, len(cuts)), make([]int, len(cuts))
	for i, e := range cuts {
		blanks[i] = len(strings.TrimRight(text[:e], " \t"))
		blanksCR[i] = len(strings.TrimRight(text[:e], " \t\r"))
	}
	atCut := func(p int, s int32) bool {
		for i, e := range cuts {
			if p >= e {
				continue
			}
			blank := p+1 >= blanks[i]
			for _, g := range sc.ends {
				blank = blank || blanks[i] <= g && g <= e
			}
			// The start must cross a line break: a gap on its way, or a line
			// break before the blanks and CR at the cut.
			if blank && (flags[s]&crossed != 0 || strings.ContainsAny(text[p+1:max(p+1, blanksCR[i])], "\r\n")) {
				return true
			}
		}
		return false
	}
	long := int32(min(minFragment, len(secret)))
	for p := n - 1; p >= 0; p-- {
		if expired(p) {
			return nil, errRedactBudget
		}
		sc.visit(p)
		for k, j := range sc.at[text[p]] {
			s := sc.first[p] + int32(k)
			next = sc.next(next[:0], p, j)
			if out[s] == 1 {
				flags[s] |= tail
			}
			for _, e := range next {
				if flags[e.s]&tail != 0 && tight(s, e.s) {
					flags[s] |= tail
					if e.gap || flags[e.s]&gapOut != 0 {
						flags[s] |= gapOut
					}
				}
			}
			if f := flags[s]; f&head != 0 && f&tail != 0 && f&(gapIn|gapOut) != 0 && into[s]+out[s]-1 >= long {
				in[p] = true
				for _, e := range next {
					if e.gap && flags[e.s]&tail != 0 && tight(s, e.s) {
						gap(p, e.r)
					}
				}
			}
			// A start of the secret that runs into a cut is masked, and every
			// path from secret[0] into it, walked back from the cut.
			if into[s] != j+1 {
				continue
			}
			if atCut(p, s) {
				flags[s] |= walked
			}
			for _, e := range next {
				if flags[e.s]&walked != 0 {
					flags[s] |= walked
					gap(p, e.r)
				}
			}
			if flags[s]&walked != 0 {
				in[p] = true
			}
		}
	}
	return in, nil
}

// spanScan holds what spanning needs to step from state to state. The states
// of text[p] are numbered from first[p], one per index in at[text[p]].
type spanScan struct {
	text, secret string
	at           [256][]int32 // the indexes in secret of each byte value
	rank         []int32      // where each index of secret is in at[secret[j]]
	first        []int32
	ends         []int // where a gap after the byte being visited may end
}

// edge leads to state s, which reads text[r], past a gap or not.
type edge struct {
	s   int32
	r   int
	gap bool
}

func newSpanScan(text, secret string) *spanScan {
	sc := &spanScan{text: text, secret: secret, rank: make([]int32, len(secret)), first: make([]int32, len(text)+1)}
	for j := 0; j < len(secret); j++ {
		sc.rank[j] = int32(len(sc.at[secret[j]]))
		sc.at[secret[j]] = append(sc.at[secret[j]], int32(j))
	}
	for p := 0; p < len(text); p++ {
		sc.first[p+1] = min(sc.first[p]+int32(len(sc.at[text[p]])), maxSpanStates+1)
	}
	return sc
}

// visit readies next for the states of text[p].
func (sc *spanScan) visit(p int) { sc.ends = gapEnds(sc.ends[:0], sc.text, p+1) }

// next appends to edges each state that can follow the state reading
// secret[j] at text[p], the byte last visited: the state reading secret[j+1]
// at the next byte, or at the end of a gap.
func (sc *spanScan) next(edges []edge, p int, j int32) []edge {
	if int(j)+1 == len(sc.secret) {
		return edges
	}
	c, k := sc.secret[j+1], sc.rank[j+1]
	if p+1 < len(sc.text) && sc.text[p+1] == c {
		edges = append(edges, edge{sc.first[p+1] + k, p + 1, false})
	}
	for _, r := range sc.ends {
		if r < len(sc.text) && sc.text[r] == c {
			edges = append(edges, edge{sc.first[r] + k, r, true})
		}
	}
	return edges
}

// maxGapBreaks and maxGapBlanks bound a gap: at most 8 line breaks, and at
// most 16 blanks or line breaks in a row. One echo is never split by more
// than a line break and a line's framing or two; a server that splits a key
// into pieces between lines of framing alone must use more than 8 to keep
// them apart. The bounds keep each gap, and so the work at each byte,
// constant however the server frames its text.
const (
	maxGapBreaks = 8
	maxGapBlanks = 16
)

// gapEnds appends to ends each i, in order, such that text[q:i] is a gap:
// what may sit between two characters of a secret that a server splits. That
// is up to maxGapBreaks line breaks, each with blanks around it and the
// framing that may open the next line: reply code, separator and enhanced
// status code, each optional, as in replyLineBreak. A break is a LF, or a CR:
// a server may end a line with a bare CR, which readLine keeps in the line it
// reads. A gap is at most maxGapBreaks*(4*maxGapBlanks+14) bytes long.
func gapEnds(ends []int, text string, q int) []int {
	i := q
	blanks := func(add bool) {
		for n := 0; n < maxGapBlanks && i < len(text) && (text[i] == ' ' || text[i] == '\t'); n++ {
			i++
			if add {
				ends = append(ends, i)
			}
		}
	}
	for range maxGapBreaks {
		blanks(false)
		if i == len(text) || text[i] != '\r' && text[i] != '\n' {
			break
		}
		for n := 0; n < maxGapBlanks && i < len(text) && strings.IndexByte(" \t\r\n", text[i]) >= 0; n++ {
			i++
			ends = append(ends, i)
		}
		if i+3 <= len(text) && isDigit(text[i]) && isDigit(text[i+1]) && isDigit(text[i+2]) {
			i += 3
			ends = append(ends, i)
			if i < len(text) && text[i] == '-' {
				i++
				ends = append(ends, i)
			}
			blanks(true)
		}
		if n := enhancedFraming(text[i:]); n > 0 {
			i += n
			ends = append(ends, i)
			blanks(true)
		}
	}
	return ends
}

// enhancedFraming returns the length of the enhanced status code and blank
// that s begins with ("5.7.8 "), or 0.
func enhancedFraming(s string) int {
	if len(s) < 2 || !isDigit(s[0]) || s[1] != '.' {
		return 0
	}
	i := 2
	digits := func() bool { // one to three
		n := 0
		for n < 3 && i < len(s) && isDigit(s[i]) {
			i, n = i+1, n+1
		}
		return n > 0
	}
	if !digits() || i == len(s) || s[i] != '.' {
		return 0
	}
	i++
	if !digits() || i == len(s) || s[i] != ' ' && s[i] != '\t' {
		return 0
	}
	return i + 1
}

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

// flush masks the session's server text, the Result's repeat of it and its
// message UUID included, writes the transcript, and returns res with its
// server text shown as the transcript shows it, the message UUID
// lower-cased (run leaves it as the server sent it). It runs once the
// session ends, however it ends: canceled too, as by Ctrl-C, so masking is
// bound by its own budget (see redactBudget), not by ctx's cancellation.
// Past the budget, nothing of the transcript is written, and res becomes a
// failure that quotes none of it.
func (t *transcript) flush(ctx context.Context, res Result) Result {
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
	if res.Verified {
		// The verdict shows the UUID last, lower-cased. A real one holds no
		// key material, so masking can only mangle a hostile one.
		heard = append(heard, heardText{text: res.MessageUUID, lower: true})
	}
	size := 0
	for _, h := range heard {
		size += len(h.text)
	}
	budget := redactBudget(size)
	ctx, cancel := context.WithTimeout(context.WithoutCancel(ctx), budget)
	defer cancel()
	shown, err := t.redact.session(ctx, heard)
	if err != nil {
		return Result{
			Stage:  res.Stage,
			Detail: fmt.Sprintf("transcript withheld: the server's replies (%d bytes) could not be checked for the API key within %s", size, budget),
			Hint:   hintWithheld,
		}
	}
	t.shown = shown
	for _, l := range t.lines {
		// Client lines are built here, never from server bytes, and the AUTH
		// line shows a placeholder: cover could only mask a Message-ID's
		// random digits that happen to match.
		text := l.text()
		if l.dir == "C:" {
			text = t.redact.r.Replace(text)
		} else {
			text = t.redact.mask(text)
		}
		fmt.Fprintf(t.w, "%-11s %s %s\n", "["+string(l.stage)+"]", l.dir, text)
	}
	if res.Reply != nil {
		reply := *res.Reply
		reply.Lines = make([]string, len(lines))
		for i, l := range lines {
			reply.Lines[i] = t.redact.mask(t.shown[l])
		}
		res.Reply = &reply
	}
	if res.Verified {
		res.MessageUUID = t.shown[res.MessageUUID]
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
	if err != nil && ctx.Err() != nil {
		return canceled(ctx, StageConnect)
	}
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
	// Left as the server sent it: flush masks it as sent and as shown,
	// lower-cased.
	uuid = reply.Lines[0][len(reply.Lines[0])-len(uuid):]
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

// canceled is the Result of a session that ctx ended: "interrupted" if it
// was canceled, as by Ctrl-C.
func canceled(ctx context.Context, stage Stage) Result {
	detail := "canceled: " + ctx.Err().Error()
	if errors.Is(ctx.Err(), context.Canceled) {
		detail = "interrupted"
	}
	return Result{Stage: stage, Detail: detail, Hint: "The command was interrupted."}
}

func (s *session) ioFailure(ctx context.Context, stage Stage, err error, timeout time.Duration) Result {
	if ctx.Err() != nil {
		return canceled(ctx, stage)
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
