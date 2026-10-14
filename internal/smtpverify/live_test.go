//go:build smtp_live

// Live end-to-end run against a real SMTP submission host. The build tag
// keeps it out of `go test ./...` and CI; it sends one real message.
//
//	TELNYX_SMTP_LIVE_HOST=<host> TELNYX_API_KEY=<key> \
//	TELNYX_SMTP_LIVE_FROM=<sender on a verified domain> [TELNYX_SMTP_LIVE_TO=<recipient>] \
//	go test -tags smtp_live -run TestLive -v ./internal/smtpverify/

package smtpverify_test

import (
	"bytes"
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/team-telnyx/telnyx-cli/internal/smtpverify"
)

func TestLive(t *testing.T) {
	host := os.Getenv("TELNYX_SMTP_LIVE_HOST")
	key := os.Getenv("TELNYX_API_KEY")
	from := os.Getenv("TELNYX_SMTP_LIVE_FROM")
	if host == "" || key == "" || from == "" {
		t.Skip("set TELNYX_SMTP_LIVE_HOST, TELNYX_API_KEY and TELNYX_SMTP_LIVE_FROM to run")
	}
	var out bytes.Buffer
	res, err := smtpverify.Run(context.Background(), smtpverify.Config{
		Host:   host,
		APIKey: key,
		From:   from,
		To:     os.Getenv("TELNYX_SMTP_LIVE_TO"),
	}, &out)
	t.Log("\n" + out.String())
	require.NoError(t, err)
	require.True(t, res.Verified, res.Verdict())
}
