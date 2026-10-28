package smtpverify

import (
	"bufio"
	"context"
	"errors"
	"io"
	"net"
	"os"
	"slices"
	"strconv"
	"strings"
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadReply(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		wire         string
		wantCode     int
		wantEnhanced string
		wantLines    []string
		wantErr      error
	}{
		{
			name:         "single line with enhanced code",
			wire:         "535 5.7.8 Authentication credentials invalid\r\n",
			wantCode:     535,
			wantEnhanced: "5.7.8",
			wantLines:    []string{"535 5.7.8 Authentication credentials invalid"},
		},
		{
			name:         "two-digit enhanced detail",
			wire:         "538 5.7.11 Encryption required\r\n",
			wantCode:     538,
			wantEnhanced: "5.7.11",
			wantLines:    []string{"538 5.7.11 Encryption required"},
		},
		{
			name:      "no enhanced code",
			wire:      "221 Bye\r\n",
			wantCode:  221,
			wantLines: []string{"221 Bye"},
		},
		{
			name:      "bare code",
			wire:      "250\r\n",
			wantCode:  250,
			wantLines: []string{"250"},
		},
		{
			name:      "enhanced class must match reply class",
			wire:      "250 5.0.0 odd\r\n",
			wantCode:  250,
			wantLines: []string{"250 5.0.0 odd"},
		},
		{
			name:     "multi-line EHLO",
			wire:     "250-email-submission.telnyx.com\r\n250-SIZE 25000000\r\n250 STARTTLS\r\n",
			wantCode: 250,
			wantLines: []string{
				"250-email-submission.telnyx.com",
				"250-SIZE 25000000",
				"250 STARTTLS",
			},
		},
		{
			name:         "bare LF tolerated",
			wire:         "250 2.0.0 Ok: queued as 0b6c5b5e-3a0f-4c3e-9d6a-6f1f2b8e4c21\n",
			wantCode:     250,
			wantEnhanced: "2.0.0",
			wantLines:    []string{"250 2.0.0 Ok: queued as 0b6c5b5e-3a0f-4c3e-9d6a-6f1f2b8e4c21"},
		},
		{name: "too short", wire: "25\r\n", wantErr: errMalformedReply},
		{name: "non-digit code", wire: "2x0 Ok\r\n", wantErr: errMalformedReply},
		{name: "code out of range", wire: "199 Ok\r\n", wantErr: errMalformedReply},
		{name: "bad separator", wire: "250_Ok\r\n", wantErr: errMalformedReply},
		{name: "code changes mid-reply", wire: "250-a\r\n251 b\r\n", wantErr: errMalformedReply},
		{name: "line too long", wire: "250 " + strings.Repeat("x", maxReplyLineBytes) + "\r\n", wantErr: errMalformedReply},
		{name: "EOF before reply", wire: "", wantErr: io.EOF},
		{name: "EOF mid multi-line reply", wire: "250-a\r\n", wantErr: io.ErrUnexpectedEOF},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := readReply(bufio.NewReader(strings.NewReader(tt.wire)))
			if tt.wantErr != nil {
				require.Error(t, err)
				assert.True(t, errors.Is(err, tt.wantErr), "got %v, want %v", err, tt.wantErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.wantCode, got.Code)
			assert.Equal(t, tt.wantEnhanced, got.Enhanced)
			assert.Equal(t, tt.wantLines, got.Lines)
		})
	}
}

// RFC 5321 §4.2: a reply line ends in CRLF. A fragment the server cut off
// (by closing or stalling) is an error, never a line, however complete it
// looks: an unterminated acceptance must not become VERIFIED.
func TestReadReplyRejectsUnterminatedLine(t *testing.T) {
	t.Parallel()
	const accepted = "250 2.0.0 Ok: queued as 0b6c5b5e-3a0f-4c3e-9d6a-6f1f2b8e4c21"
	tests := []struct {
		name        string
		wire        io.Reader
		wantLines   []string // complete lines before the fragment
		wantPartial string
		wantTimeout bool // else the cause is io.ErrUnexpectedEOF
	}{
		{name: "EOF after partial 250 acceptance", wire: strings.NewReader(accepted), wantPartial: accepted},
		{name: "EOF after CR without LF", wire: strings.NewReader(accepted + "\r"), wantPartial: accepted + "\r"},
		{name: "EOF mid final line of a multi-line reply", wire: strings.NewReader("250-a\r\n" + accepted), wantLines: []string{"250-a"}, wantPartial: accepted},
		{
			name:        "timeout after partial 250 acceptance",
			wire:        io.MultiReader(strings.NewReader(accepted), iotest.ErrReader(os.ErrDeadlineExceeded)),
			wantPartial: accepted,
			wantTimeout: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := readReply(bufio.NewReader(tt.wire))
			var cut *incompleteLineError
			require.True(t, errors.As(err, &cut), "got %v, want an incomplete line", err)
			assert.Equal(t, tt.wantPartial, cut.partial)
			if tt.wantTimeout {
				var ne net.Error
				assert.True(t, errors.As(err, &ne) && ne.Timeout(), "timeout cause kept: %v", err)
			} else {
				assert.True(t, errors.Is(err, io.ErrUnexpectedEOF), "got %v", err)
			}
			assert.Equal(t, tt.wantLines, got.Lines, "the fragment is not a reply line")
			assert.Empty(t, messageUUID(got))
		})
	}
}

func TestReadReplyConsumesExactlyOneReply(t *testing.T) {
	t.Parallel()
	r := bufio.NewReader(strings.NewReader("250-a\r\n250 b\r\n221 Bye\r\n"))
	first, err := readReply(r)
	require.NoError(t, err)
	assert.Equal(t, []string{"250-a", "250 b"}, first.Lines)
	second, err := readReply(r)
	require.NoError(t, err)
	assert.Equal(t, []string{"221 Bye"}, second.Lines)
}

// Only the contract's acceptance reply carries the message UUID; any other
// UUID-shaped token is someone else's ID and must not be promoted.
func TestMessageUUID(t *testing.T) {
	t.Parallel()
	const id = "0b6c5b5e-3a0f-4c3e-9d6a-6f1f2b8e4c21"
	tests := []struct {
		name  string
		lines []string
		want  string
	}{
		{"contract shape", []string{"250 2.0.0 Ok: queued as " + id}, id},
		{"upper-case hex normalized", []string{"250 2.0.0 Ok: queued as 0B6C5B5E-3A0F-4C3E-9D6A-6F1F2B8E4C21"}, id},
		{"no uuid", []string{"250 Ok"}, ""},
		{"not a uuid", []string{"250 2.0.0 Ok: queued as 1234ABCD"}, ""},
		{"trace id", []string{"250 2.0.0 accepted; trace " + id + "@mta"}, ""},
		{"bare uuid", []string{"250 2.0.0 " + id}, ""},
		{"uuid on a later line", []string{"250-accepted", "250 id " + id}, ""},
		{"multi-line reply ending in the shape", []string{"250-accepted", "250 2.0.0 Ok: queued as " + id}, ""},
		{"phrase after other text", []string{"250 trace 11111111-1111-4111-8111-111111111111 queued as " + id}, ""},
		{"trailing text", []string{"250 2.0.0 Ok: queued as " + id + " (trace 11111111-1111-4111-8111-111111111111)"}, ""},
		{"uuid run on", []string{"250 2.0.0 Ok: queued as " + id + "0"}, ""},
		{"missing enhanced code", []string{"250 Ok: queued as " + id}, ""},
		{"other enhanced code", []string{"250 2.6.0 Ok: queued as " + id}, ""},
		{"phrase case differs", []string{"250 2.0.0 OK: QUEUED AS " + id}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, messageUUID(Reply{Code: 250, Lines: tt.lines}))
		})
	}
}

// Whatever the server echoes back — the key, or the AUTH PLAIN payload that
// carries it — is masked, raw or as printable and %q render it.
func TestRedactor(t *testing.T) {
	t.Parallel()
	const key = "KEYtest_0123456789abcdef"
	payload := authPlainPayload(key)
	tests := []struct {
		name, key, in, want string
	}{
		{"AUTH PLAIN payload", key, "535 5.7.8 invalid: AUTH PLAIN " + payload, "535 5.7.8 invalid: AUTH PLAIN " + redactedPayload},
		{"raw key", key, "235 ok for apikey:" + key + ".", "235 ok for apikey:" + redactedKey + "."},
		{"key in decoded payload", key, printable("\x00apikey\x00" + key), `\x00apikey\x00` + redactedKey},
		{"every occurrence", key, key + " " + payload + " " + key, redactedKey + " " + redactedPayload + " " + redactedKey},
		{"key as printable escapes it", "KEY\x01x", printable("x KEY\x01x"), "x " + redactedKey},
		{"key as %q escapes it", `KEY"x\`, strconv.Quote(`a KEY"x\ b`), `"a ` + redactedKey + ` b"`},
		{"nothing to mask", key, "250 2.1.0 Ok", "250 2.1.0 Ok"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, newRedactor(tt.key).mask(tt.in))
		})
	}
}

// A secret wrapped across a reply's lines is masked as one span, line breaks
// and the next line's framing included; lines it does not span are kept.
func TestRedactorReply(t *testing.T) {
	t.Parallel()
	const key = "KEYtest_0123456789abcdef"
	payload := authPlainPayload(key)
	tests := []struct {
		name     string
		in, want []string
	}{
		{"no lines", nil, nil},
		{"nothing to mask", []string{"250-a", "250 b"}, []string{"250-a", "250 b"}},
		{"payload over two lines", []string{"535-x " + payload[:5], "535 " + payload[5:] + " y"}, []string{"535-x " + redactedPayload + " y"}},
		{"key over CR, blanks and enhanced codes", []string{"220 hi", "535-5.7.8 " + key[:3] + " \r", "535-5.7.8  " + key[3:7], "535 5.7.8 " + key[7:], "unrelated"}, []string{"220 hi", "535-5.7.8 " + redactedKey, "unrelated"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, newRedactor(key).reply(tt.in))
		})
	}
}

// A secret wrapped across a reply's lines can run on into the bytes that end
// the read: a rejected line, framed or not, or a line cut off after the
// secret or inside it. Those bytes are masked together with the lines, and
// the error comes back quoting them masked, of the same type and cause.
func TestRedactorReplyErr(t *testing.T) {
	t.Parallel()
	const key = "KEYtest_0123456789abcdef"
	payload := authPlainPayload(key)
	tests := []struct {
		name      string
		lines     []string
		err       error
		wantLines []string
		wantErr   string
		wantCause error
	}{
		{
			name:      "rejected line completes the payload",
			lines:     []string{"535-x " + payload[:5]},
			err:       &badLineError{why: "code changed mid-reply", line: "550 " + payload[5:] + " y"},
			wantLines: []string{},
			wantErr:   `malformed reply: code changed mid-reply in "535-x ` + redactedPayload + ` y"`,
			wantCause: errMalformedReply,
		},
		{
			name:      "rejected line with no reply code completes the key",
			lines:     []string{"220-a", "535-x " + key[:9]},
			err:       &badLineError{why: "reply code 123 out of range", line: key[9:]},
			wantLines: []string{"220-a"},
			wantErr:   `malformed reply: reply code 123 out of range in "535-x ` + redactedKey + `"`,
			wantCause: errMalformedReply,
		},
		{
			name:      "cut-off line completes the key",
			lines:     []string{"535-x " + key[:3]},
			err:       &incompleteLineError{partial: "535 5.7.8 " + key[3:] + " (ch", err: io.ErrUnexpectedEOF},
			wantLines: []string{},
			wantErr:   `incomplete reply "535-x ` + redactedKey + ` (ch" (no CRLF): unexpected EOF`,
			wantCause: io.ErrUnexpectedEOF,
		},
		{
			name:      "cut inside the payload",
			lines:     []string{"250-a", "535-x " + payload[:5]},
			err:       &incompleteLineError{partial: "535 " + payload[5:20], err: os.ErrDeadlineExceeded},
			wantLines: []string{"250-a"},
			wantErr:   `incomplete reply "535-x ` + redactedPayload + `" (no CRLF): i/o timeout`,
			wantCause: os.ErrDeadlineExceeded,
		},
		{
			name:      "cut after the next line's framing",
			lines:     []string{"535-x " + key[:5]},
			err:       &incompleteLineError{partial: "535-", err: io.ErrUnexpectedEOF},
			wantLines: []string{},
			wantErr:   `incomplete reply "535-x ` + redactedKey + `" (no CRLF): unexpected EOF`,
			wantCause: io.ErrUnexpectedEOF,
		},
		{
			name:      "start within the cut-off line alone is kept",
			err:       &incompleteLineError{partial: "250 2.0.0 Ok: queued as 0B6C5B5E-3A", err: io.ErrUnexpectedEOF},
			wantLines: []string{},
			wantErr:   `incomplete reply "250 2.0.0 Ok: queued as 0B6C5B5E-3A" (no CRLF): unexpected EOF`,
			wantCause: io.ErrUnexpectedEOF,
		},
		{
			name:      "nothing to mask",
			lines:     []string{"250-a"},
			err:       &badLineError{why: "code changed mid-reply", line: "251 b"},
			wantLines: []string{"250-a"},
			wantErr:   `malformed reply: code changed mid-reply in "251 b"`,
			wantCause: errMalformedReply,
		},
		{
			name:      "error quoting no bytes",
			lines:     []string{"535-x " + key[:5], "535 " + key[5:]},
			err:       io.ErrUnexpectedEOF,
			wantLines: []string{"535-x " + redactedKey},
			wantErr:   "unexpected EOF",
			wantCause: io.ErrUnexpectedEOF,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			lines, err := newRedactor(key).replyErr(tt.lines, tt.err)
			assert.Equal(t, tt.wantLines, lines)
			require.Error(t, err)
			assert.IsType(t, tt.err, err)
			assert.Equal(t, tt.wantErr, err.Error())
			assert.True(t, errors.Is(err, tt.wantCause), "cause kept: %v", err)
		})
	}
}

// session masks a session's server text together: a secret split across
// replies, framing kept; 7 characters of a secret in a row on their own,
// not 6, whatever text is around them and in whatever order they come; a run
// of 7 that framing splits, over at most 8 line breaks; the start of a
// secret running into a cut across a line break, with a CR at the cut or
// not, but not one within the cut-off line alone; a message UUID as sent
// and as shown, lower-cased; and the same text the same way wherever it is
// shown.
func TestRedactorSession(t *testing.T) {
	t.Parallel()
	const key = "KEYtest_0123456789abcdef"
	payload := authPlainPayload(key)
	lines := func(texts ...string) []heardText {
		var heard []heardText
		for _, s := range texts {
			heard = append(heard, heardText{text: s})
		}
		return heard
	}
	cut := func(s string) heardText { return heardText{text: s, quoted: true, cut: true} }
	tests := []struct {
		name  string
		heard []heardText
		want  []string // what is shown of each text, in order
	}{
		{"nothing to mask", lines("220 hi", "250-a", "250 b"), []string{"220 hi", "250-a", "250 b"}},
		{"payload halves in two replies", lines("535 "+payload[:19], "221 "+payload[19:]), []string{"535 " + redactedPayload, "221 " + redactedPayload}},
		{"key halves under 16 characters", lines("535 5.7.8 for apikey:"+key[:12], "221 2.0.0 "+key[12:]), []string{"535 5.7.8 for apikey:" + redactedKey, "221 2.0.0 " + redactedKey}},
		{"reply code that is part of the key", lines("535-x "+key[:9], key[9:]), []string{"535-x " + redactedKey, redactedKey}},
		{"framing that could also be the next character", lines("535-x "+key[:15], "700 "+key[15:]), []string{"535-x " + redactedKey, "700 " + redactedKey}},
		{"line of framing alone between", lines("535-"+key[:8], "535-", "535 5.7.8 "+key[8:]), []string{"535-" + redactedKey, "535-", "535 5.7.8 " + redactedKey}},
		{"bare CR inside a line", lines("535 " + key[:12] + "\r" + key[12:] + " y"), []string{"535 " + redactedKey + " y"}},
		{"6 characters on their own are kept", lines("250 x" + key[3:9] + "y"), []string{"250 x" + key[3:9] + "y"}},
		{"7 characters on their own", lines("250 x" + key[3:10] + "y"), []string{"250 x" + redactedKey + "y"}},
		{"pieces around other text", lines("535 5.7.8 for apikey:"+key[:8], "221 Bye "+key[8:16], "250 ok "+key[16:]+" done"), []string{"535 5.7.8 for apikey:" + redactedKey, "221 Bye " + redactedKey, "250 ok " + redactedKey + " done"}},
		{"pieces out of order", lines("250 a "+key[16:], "250 b "+key[:8], "250 c "+key[8:16]), []string{"250 a " + redactedKey, "250 b " + redactedKey, "250 c " + redactedKey}},
		{"6-character pieces around other text are kept", lines("250 a "+key[:6]+" b", "250 c "+key[6:12]+" d"), []string{"250 a " + key[:6] + " b", "250 c " + key[6:12] + " d"}},
		{"run that framing splits, its pieces under 7 characters", lines("535-x "+key[:3], "535 "+key[3:7]), []string{"535-x " + redactedKey, "535 " + redactedKey}},
		{"run over 8 line breaks", lines(slices.Concat([]string{"535-x " + key[:4]}, slices.Repeat([]string{"535-"}, 7), []string{"535 " + key[4:9]})...), slices.Concat([]string{"535-x " + redactedKey}, slices.Repeat([]string{"535-"}, 7), []string{"535 " + redactedKey})},
		{"run over 9 line breaks is not joined", lines(slices.Concat([]string{"535-x " + key[:4]}, slices.Repeat([]string{"535-"}, 8), []string{"535 " + key[4:9]})...), slices.Concat([]string{"535-x " + key[:4]}, slices.Repeat([]string{"535-"}, 8), []string{"535 " + key[4:9]})},
		{"payload tail on its own", lines("535 x " + payload[19:]), []string{"535 x " + redactedPayload}},
		{"start of the key at a cut after a CR", []heardText{{text: "535-x " + key[:4]}, cut("535 " + key[4:6] + "\r")}, []string{"535-x " + redactedKey, "535 " + redactedKey + "\r"}},
		{"start of the key across replies at a cut", []heardText{{text: "535 " + key[:4]}, cut("221 " + key[4:6])}, []string{"535 " + redactedKey, "221 " + redactedKey}},
		{"start not at a cut is kept", lines("535-x "+key[:4], "535 "+key[4:6]), []string{"535-x " + key[:4], "535 " + key[4:6]}},
		{"start within the cut-off line alone is kept", []heardText{cut("250 2.0.0 Ok: queued as 0B6C5B5E-3A")}, []string{"250 2.0.0 Ok: queued as 0B6C5B5E-3A"}},
		{"start within the cut-off line alone is kept, CR at the cut", []heardText{cut("250 2.0.0 Ok: queued as 0B6C5B5E-3A\r")}, []string{"250 2.0.0 Ok: queued as 0B6C5B5E-3A\r"}},
		{"message UUID as sent, shown lower-cased", []heardText{{text: "0B6C5B5E-3A0F-4C3E-9D6A-6F1F2B8E4C21", lower: true}}, []string{"0b6c5b5e-3a0f-4c3e-9d6a-6f1f2b8e4c21"}},
		{"message UUID holding the key's digits", []heardText{{text: "01234567-3A0F-4C3E-9D6A-6F1F2B8E4C21", lower: true}}, []string{redactedKey + "-3a0f-4c3e-9d6a-6f1f2b8e4c21"}},
		{"message UUID that holds the key once lower-cased", []heardText{{text: "9ABCDEF0-3A0F-4C3E-9D6A-6F1F2B8E4C21", lower: true}}, []string{redactedKey + "0-3a0f-4c3e-9d6a-6f1f2b8e4c21"}},
		{"message UUID ending a run of the key", []heardText{{text: "221 Bye " + key[:4]}, {text: key[4:7] + "-3A0F-4C3E-9D6A-6F1F2B8E4C21", lower: true}}, []string{"221 Bye " + redactedKey, redactedKey + "-3a0f-4c3e-9d6a-6f1f2b8e4c21"}},
		{"message UUID ending a run of the key once lower-cased", []heardText{{text: "221 Bye 789"}, {text: "ABCDEF00-3A0F-4C3E-9D6A-6F1F2B8E4C21", lower: true}}, []string{"221 Bye " + redactedKey, redactedKey + "00-3a0f-4c3e-9d6a-6f1f2b8e4c21"}},
		{
			name:  "message UUID starting a run of the key once lower-cased, in the reply that carried it",
			heard: []heardText{{text: "250 2.0.0 Ok: queued as 3A0F9999-3A0F-4C3E-9D6A-000000789ABC"}, {text: "221 2.0.0 def"}, {text: "3A0F9999-3A0F-4C3E-9D6A-000000789ABC", lower: true}},
			want:  []string{"250 2.0.0 Ok: queued as 3A0F9999-3A0F-4C3E-9D6A-000000" + redactedKey, "221 2.0.0 " + redactedKey, "3a0f9999-3a0f-4c3e-9d6a-000000" + redactedKey},
		},
		{
			name:  "message UUID masked alike in the reply that carried it",
			heard: []heardText{{text: "250 2.0.0 Ok: queued as BCDEF000-3A0F-4C3E-9D6A-6F1F2B8E4C21"}, {text: "221 Bye 789a"}, {text: "BCDEF000-3A0F-4C3E-9D6A-6F1F2B8E4C21", lower: true}},
			want:  []string{"250 2.0.0 Ok: queued as " + redactedKey + "000-3A0F-4C3E-9D6A-6F1F2B8E4C21", "221 Bye " + redactedKey, redactedKey + "000-3a0f-4c3e-9d6a-6f1f2b8e4c21"},
		},
		{
			name:  "second half first, repeated after the first",
			heard: lines("535 5.7.8 "+key[12:], "221 "+key[:12], "535 5.7.8 "+key[12:]),
			want:  []string{"535 5.7.8 " + redactedKey, "221 " + redactedKey, "535 5.7.8 " + redactedKey},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			shown, err := newRedactor(key).session(context.Background(), tt.heard)
			require.NoError(t, err)
			var got []string
			for _, h := range tt.heard {
				got = append(got, shown[h.text])
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPrintable(t *testing.T) {
	t.Parallel()
	tests := []struct{ in, want string }{
		{"250 2.0.0 Ok", "250 2.0.0 Ok"},
		{"220 x\x1b[2K\rfake", `220 x\x1b[2K\x0dfake`},
		{"tab\tkept", "tab\tkept"},
		{"del\x7f", `del\x7f`},
		{"c1 \u009b", `c1 \u009b`},
		{"bad utf8 \xff", `bad utf8 \xff`},
		{"unicode ok ü", "unicode ok ü"},
		{"bidi \u202eoverride", `bidi \u202eoverride`},
		{"zero\u200bwidth", `zero\u200bwidth`},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, printable(tt.in), "%q", tt.in)
	}
}

func TestCapabilities(t *testing.T) {
	t.Parallel()
	caps := capabilities(Reply{Code: 250, Lines: []string{
		"250-email-submission.telnyx.com",
		"250-SIZE 25000000",
		"250-starttls",
		"250 AUTH PLAIN LOGIN",
	}})
	assert.Equal(t, "25000000", caps["SIZE"])
	_, hasTLS := caps["STARTTLS"]
	assert.True(t, hasTLS, "keywords are case-insensitive")
	_, hasGreeting := caps["EMAIL-SUBMISSION.TELNYX.COM"]
	assert.False(t, hasGreeting, "the greeting line is not a capability")
	assert.True(t, hasAuthMechanism(caps, "PLAIN"))
	assert.True(t, hasAuthMechanism(caps, "login"))
	assert.False(t, hasAuthMechanism(caps, "CRAM-MD5"))
	assert.False(t, hasAuthMechanism(map[string]string{}, "PLAIN"))
}

func TestHintForReply(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		stage Stage
		reply string
		want  string // substring
	}{
		{"bad key", StageAuth, "535 5.7.8 Authentication credentials invalid", `literal "apikey"`},
		{"missing capability at AUTH", StageAuth, "535 5.7.1 Authorization failed: email send capability required", "lacks the email send capability"},
		{"intake authorization at DATA", StageData, "535 5.7.1 Authorization failed on message submission", "when the message was submitted"},
		{"pre-TLS AUTH", StageAuth, "538 5.7.11 Encryption required", "STARTTLS"},
		{"throttled", StageAuth, "421 4.7.0 Temporary authentication failure, retry later", "throttled"},
		{"session expired", StageMailFrom, "421 4.4.2 Session expired, close and retry", "maximum lifetime"},
		{"domain refused", StageMailFrom, "550 5.7.1 Domain is not verified. Complete DNS setup before sending email.", "verified sending domain"},
		{"recipient cap", StageRcptTo, "452 4.5.3 Too many recipients", "Too many recipients"},
		{"quota at DATA", StageData, "452 4.5.3 Mail quota exceeded, try again later", "quota"},
		{"size", StageData, "552 5.3.4 Message size exceeds fixed maximum message size", "SIZE"},
		{"unknown 4xx", StageRcptTo, "450 4.2.0 Mailbox busy", "Temporary failure"},
		{"unknown 5xx", StageRcptTo, "550 5.1.1 No such user", "Permanent failure"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			r, err := readReply(bufio.NewReader(strings.NewReader(tt.reply + "\r\n")))
			require.NoError(t, err)
			assert.Contains(t, hintForReply(tt.stage, r), tt.want)
		})
	}
}

func TestVerdict(t *testing.T) {
	t.Parallel()
	reply := Reply{Code: 535, Enhanced: "5.7.8", Lines: []string{"535 5.7.8 Authentication credentials invalid"}}
	multi := Reply{Code: 554, Lines: []string{"554-first", "554 second"}}
	tests := []struct {
		name string
		res  Result
		want string
	}{
		{"verified", Result{Verified: true, MessageUUID: "0b6c5b5e-3a0f-4c3e-9d6a-6f1f2b8e4c21"}, "VERIFIED: message UUID 0b6c5b5e-3a0f-4c3e-9d6a-6f1f2b8e4c21"},
		{"server reply", Result{Stage: StageAuth, Reply: &reply}, "FAILED at AUTH: 535 5.7.8 Authentication credentials invalid"},
		{"multi-line reply kept whole", Result{Stage: StageData, Reply: &multi}, "FAILED at DATA: 554-first | 554 second"},
		{"local failure", Result{Stage: StageSTARTTLS, Detail: "server did not advertise STARTTLS"}, "FAILED at STARTTLS: server did not advertise STARTTLS"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, tt.res.Verdict())
		})
	}
}
