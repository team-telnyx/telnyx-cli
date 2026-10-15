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
// it, or the AUTH PLAIN payload, is redacted.
func Run(ctx context.Context, cfg Config, out io.Writer) (Result, error) {
	cfg, err := cfg.normalize()
	if err != nil {
		return Result{}, err
	}
	redact := newRedactor(cfg.APIKey)
	s := &session{cfg: cfg, t: transcript{out, redact}}
	res := redact.result(s.run(ctx))
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
// raw and as printable and %q render it.
type redactor struct{ r *strings.Replacer }

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
	return redactor{strings.NewReplacer(pairs...)}
}

func (rd redactor) mask(s string) string { return rd.r.Replace(s) }

// result redacts the server text and local detail a Result carries.
func (rd redactor) result(res Result) Result {
	if res.Reply != nil {
		reply := *res.Reply
		reply.Lines = make([]string, len(res.Reply.Lines))
		for i, l := range res.Reply.Lines {
			reply.Lines[i] = rd.mask(l)
		}
		res.Reply = &reply
	}
	res.Detail = rd.mask(res.Detail)
	res.Hint = rd.mask(res.Hint)
	return res
}

type transcript struct {
	w      io.Writer
	redact redactor
}

func (t transcript) line(stage Stage, dir, text string) {
	fmt.Fprintf(t.w, "%-11s %s %s\n", "["+string(stage)+"]", dir, t.redact.mask(text))
}

func (t transcript) client(stage Stage, text string) { t.line(stage, "C:", text) }
func (t transcript) note(stage Stage, text string)   { t.line(stage, "--", text) }
func (t transcript) server(stage Stage, r Reply) {
	for _, l := range r.Lines {
		t.line(stage, "S:", printable(l))
	}
}

type session struct {
	cfg  Config
	t    transcript
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

// receive reads one reply and prints it. Partial replies are printed too,
// including a line cut off before its CRLF: every reply byte is shown.
func (s *session) receive(stage Stage) (Reply, error) {
	reply, err := readReply(s.r)
	s.t.server(stage, reply)
	var cut *incompleteLineError
	if errors.As(err, &cut) {
		s.t.note(stage, "incomplete reply line (no CRLF): "+printable(cut.partial))
	}
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
