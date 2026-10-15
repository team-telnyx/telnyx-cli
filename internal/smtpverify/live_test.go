//go:build smtp_live

// Live end-to-end run against a real SMTP submission host. The build tag
// keeps it out of `go test ./...` and CI; it sends one real message.
//
//	TELNYX_API_KEY=<key> TELNYX_SMTP_LIVE_FROM=<sender on a verified domain> \
//	[TELNYX_SMTP_LIVE_HOST=<host; default the dev front door>] [TELNYX_SMTP_LIVE_TO=<recipient>] \
//	go test -tags smtp_live -run TestLive -v ./internal/smtpverify/
//
// A run whose prerequisites are not met yet is classified GATED and
// skipped (exit status 0): a gated skip is an honest skip, neither a pass
// nor a failure. Past the gate, every failure fails the test.

package smtpverify_test

import (
	"bytes"
	"context"
	"errors"
	"net"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/team-telnyx/telnyx-cli/internal/smtpverify"
)

// devFrontDoor is the dev SMTP submission front door; its DNS record is
// tracked in internal #436.
const devFrontDoor = "smtp.dev.telnyx.tech"

func TestLive(t *testing.T) {
	key := os.Getenv("TELNYX_API_KEY")
	from := os.Getenv("TELNYX_SMTP_LIVE_FROM")
	var missing []string
	if key == "" {
		missing = append(missing, "TELNYX_API_KEY")
	}
	if from == "" {
		missing = append(missing, "TELNYX_SMTP_LIVE_FROM")
	}
	if len(missing) > 0 {
		t.Skipf("GATED: %s not set — skipping live verification", strings.Join(missing, " and "))
	}
	host := os.Getenv("TELNYX_SMTP_LIVE_HOST")
	if host == "" {
		host = devFrontDoor
	}
	if host == devFrontDoor {
		gateOnDNS(t, host)
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

// gateOnDNS classifies the run GATED while the dev front door has no DNS
// record (internal #436). Any other lookup failure is a real failure, as is
// a custom TELNYX_SMTP_LIVE_HOST that does not resolve (not gated).
func gateOnDNS(t *testing.T, host string) {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := net.DefaultResolver.LookupHost(ctx, host)
	var dnsErr *net.DNSError
	switch {
	case err == nil:
	case errors.As(err, &dnsErr) && dnsErr.IsNotFound:
		t.Skipf("GATED: dev front door DNS record not yet available (internal #436): %s does not resolve — skipping live verification", host)
	default:
		t.Fatalf("FAIL: resolving the dev front door %s: %v", host, err)
	}
}
