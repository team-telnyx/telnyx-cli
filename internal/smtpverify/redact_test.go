package smtpverify

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"math/rand/v2"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/team-telnyx/telnyx-cli/internal/smtpverify/smtpmock"
)

// randomKey is a key of the Telnyx format: "KEY", 32 upper-case hex digits,
// "_", 22 letters and digits.
func randomKey(r *rand.Rand) string {
	const hex = "0123456789ABCDEF"
	const alnum = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	return "KEY" + randomString(r, hex, 32) + "_" + randomString(r, alnum, 22)
}

func randomString(r *rand.Rand, alphabet string, n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = alphabet[r.IntN(len(alphabet))]
	}
	return string(b)
}

// noisySession is a session's server text far denser than a real one in
// what could match a key by chance: UUIDs in lower and upper case, upper-case
// hex queue IDs, hex trace IDs, Message-IDs, addresses, MIME base64 and
// dates, around the contract's own replies.
func noisySession(r *rand.Rand) []string {
	const lower, upper, word = "0123456789abcdef", "0123456789ABCDEF", "abcdefghijklmnopqrstuvwxyz0123456789"
	uuid := func(hex string) string {
		return randomString(r, hex, 8) + "-" + randomString(r, hex, 4) + "-" + randomString(r, hex, 4) + "-" + randomString(r, hex, 4) + "-" + randomString(r, hex, 12)
	}
	lines := []string{
		smtpmock.Banner,
		"250-email-submission.telnyx.com", "250-SIZE 25000000", "250-8BITMIME", "250-SMTPUTF8", "250-PIPELINING", "250 AUTH PLAIN LOGIN",
		smtpmock.ReplyAuthOK, "250 2.1.0 Ok", "250 2.1.5 Ok",
	}
	for range 4 {
		mime := make([]byte, 57)
		for i := range mime {
			mime[i] = byte(r.IntN(256))
		}
		lines = append(lines,
			"250 2.0.0 Ok: queued as "+uuid(lower),
			"250 2.0.0 Ok: queued as "+uuid(upper),
			"250 2.0.0 Ok: queued as "+randomString(r, upper, 10+r.IntN(3)),
			"250-trace "+randomString(r, lower, 32),
			"250-Message-ID <telnyx-smtp-verify."+randomString(r, lower, 16)+"@"+randomString(r, word, 8)+".com>",
			"550 5.1.1 <"+randomString(r, word, 6)+"."+randomString(r, word, 5)+"+"+randomString(r, word, 3)+"@"+randomString(r, word, 7)+".example.net>: Recipient address rejected",
			"250-"+base64.StdEncoding.EncodeToString(mime),
			fmt.Sprintf("250-Date: Thu, %02d Oct 2026 %02d:%02d:%02d -0700 id %d", r.IntN(28)+1, r.IntN(24), r.IntN(60), r.IntN(60), r.IntN(1e9)),
		)
	}
	return append(lines, smtpmock.ReplyQuit)
}

// The false-positive control for minFragment and the gap bounds: 40,000
// random keys, each against a noisy session, its message UUIDs also shown
// lower-cased as VERIFIED shows them, with every rule of session and mask,
// mask nothing. The keys run in shards, each a seeded stream of its own, in
// parallel. (minFragment = 6 masks text in some of these sessions; see
// minFragment.)
func TestCoverageFalsePositives(t *testing.T) {
	t.Parallel()
	const shards, keys = 8, 40000
	for shard := range shards {
		t.Run(strconv.Itoa(shard), func(t *testing.T) {
			t.Parallel()
			r := rand.New(rand.NewPCG(2353, uint64(shard)))
			for range keys / shards {
				key := randomKey(r)
				rd := newRedactor(key)
				noise := noisySession(r)
				var heard []heardText
				for _, l := range noise {
					heard = append(heard, heardText{text: l})
				}
				heard[len(heard)-1].cut = true
				for _, l := range noise {
					if uuid := messageUUID(Reply{Code: 250, Lines: []string{l}}); uuid != "" {
						heard = append(heard, heardText{text: l[len(l)-len(uuid):], lower: true})
					}
				}
				shown, err := rd.session(context.Background(), heard)
				require.NoError(t, err)
				for _, h := range heard {
					if want := map[bool]string{false: h.text, true: lowerASCII(h.text)}[h.lower]; shown[h.text] != want || rd.mask(h.text) != h.text {
						t.Errorf("key %s: %q shown as %q, masked as %q", key, h.text, shown[h.text], rd.mask(h.text))
					}
				}
			}
		})
	}
}

// hostileBanner is a banner's lines, about size bytes of them, each 2,045
// bytes or less as the parser accepts, of body over and over.
func hostileBanner(body string, size int) []heardText {
	var heard []heardText
	for n := 0; n < size; {
		l := ("220-" + strings.Repeat(body, 2045/len(body)+1))[:2045]
		heard = append(heard, heardText{text: l})
		n += len(l)
	}
	return heard
}

// Server text that makes gaps long and ambiguous: Astra R6's 16 KB banner of
// lines of "250\r" (framing after framing), the same with framing opening
// each line so gaps could run on from line to line, blanks and line breaks as
// long as a gap takes, and pieces of the secrets between framing.
var hostileBodies = map[string]string{
	"framing over and over":        "250\r",
	"framing that opens each line": "\r250",
	"blanks and breaks":            "0" + blanks + strings.Repeat("\r\n", maxGapBlanks/2) + "535-" + blanks + "5.7.8" + blanks,
	"pieces between framing":       authPlainPayload(hostileKey)[:30] + "\r250 " + hostileKey[:20] + "\r",
}

// blanks is as many blanks in a row as a gap takes.
var blanks = strings.Repeat(" \t", maxGapBlanks/2)

const hostileKey = "KEY01892F3C4A5B6D7E8F9A0B1C2D3E4F_Vb7Nq2Wx9Rt4Ys6Uk8Jm1Lp3Hd5G"

// redactCost is how long session takes over heard, and how many bytes it
// allocates.
func redactCost(t *testing.T, heard []heardText) (time.Duration, uint64) {
	t.Helper()
	rd := newRedactor(hostileKey)
	var before, after runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&before)
	start := time.Now()
	_, err := rd.session(context.Background(), heard)
	elapsed := time.Since(start)
	runtime.ReadMemStats(&after)
	require.NoError(t, err)
	return elapsed, after.TotalAlloc - before.TotalAlloc
}

// upperUUID is a message UUID that the server sent in upper case. VERIFIED
// shows it lower-cased, so session searches the text twice.
var upperUUID = heardText{text: "0B6C5B5E-3A0F-4C3E-9D6A-6F1F2B8E4C21", lower: true}

// Astra R6: a 16 KB banner of lines of "250\r" took 10 s and 7.7 GB to redact,
// and twice the text four times the memory. A gap is bounded now, and
// nothing is cached per byte, so redaction takes time and memory linear in
// the text: hostile banners take well under 250 ms and a few MB, and twice
// the text takes about twice the memory. So they do when the session ends
// VERIFIED with an upper-case message UUID, and the text is searched twice.
func TestSessionRedactionIsLinear(t *testing.T) {
	// Not parallel: timed, and it measures allocation.
	for name, body := range hostileBodies {
		for _, verified := range []bool{false, true} {
			banner := func(size int) []heardText {
				heard := hostileBanner(body, size)
				if verified {
					heard = append(heard, upperUUID)
				}
				return heard
			}
			t.Run(fmt.Sprintf("%s, verified %t", name, verified), func(t *testing.T) {
				elapsed, alloc := redactCost(t, banner(16<<10))
				_, alloc2 := redactCost(t, banner(32<<10))
				t.Logf("16 KB: %v, %.1f MiB; 32 KB: %.1f MiB", elapsed, float64(alloc)/(1<<20), float64(alloc2)/(1<<20))
				assert.Less(t, elapsed, 250*time.Millisecond)
				assert.Less(t, alloc, uint64(16<<20))
				assert.Less(t, float64(alloc2)/float64(alloc), 2.5, "twice the text, about twice the memory")
			})
		}
	}
}

func BenchmarkSessionRedaction(b *testing.B) {
	for name, body := range hostileBodies {
		heard := hostileBanner(body, 16<<10)
		size := 0
		for _, h := range heard {
			size += len(h.text)
		}
		rd := newRedactor(hostileKey)
		b.Run(name, func(b *testing.B) {
			b.SetBytes(int64(size))
			b.ReportAllocs()
			for b.Loop() {
				if _, err := rd.session(context.Background(), heard); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

// Past its budget, redaction gives up: nothing of the transcript is written,
// and the verdict says why in plain words, quoting none of it.
func TestRunWithholdsTranscriptPastBudget(t *testing.T) {
	// Not parallel: it changes redactBudget.
	defer func(budget func(int) time.Duration) { redactBudget = budget }(redactBudget)
	redactBudget = func(int) time.Duration { return 0 }
	const key = "KEYtest_0123456789abcdef"
	srv := smtpmock.Start(t, smtpmock.Script{AuthReply: "535 5.7.8 Authentication credentials invalid for apikey:" + key[:12]})
	var out bytes.Buffer
	res, err := Run(context.Background(), Config{
		Host: srv.Host, Port: srv.Port, APIKey: key, From: "verify@example.com",
		TLSConfig: srv.ClientTLS, CommandTimeout: 3 * time.Second, DataTimeout: 3 * time.Second,
	}, &out)
	require.NoError(t, err)
	assert.False(t, res.Verified)
	assert.Nil(t, res.Reply)
	assert.Equal(t, StageAuth, res.Stage)
	assert.Regexp(t, `^FAILED at AUTH: transcript withheld: the server's replies \(\d+ bytes\) could not be checked for the API key within 0s$`, res.Verdict())
	assert.Equal(t, "\nhint: "+hintWithheld+"\n"+res.Verdict()+"\n", out.String(), "nothing but the hint and the verdict")
}
