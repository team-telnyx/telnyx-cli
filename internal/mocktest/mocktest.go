package mocktest

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var mockServerURL *url.URL

func init() {
	mockServerURL, _ = url.Parse("http://localhost:4010")
	if testURL := os.Getenv("TEST_API_BASE_URL"); testURL != "" {
		if parsed, err := url.Parse(testURL); err == nil {
			mockServerURL = parsed
		}
	}
}

// OnlyMockServerDialer only allows network connections to the mock server
type OnlyMockServerDialer struct{}

func (d *OnlyMockServerDialer) DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	if address == mockServerURL.Host {
		return (&net.Dialer{}).DialContext(ctx, network, address)
	}

	return nil, fmt.Errorf("BLOCKED: connection to %s not allowed (only allowed: %s)", address, mockServerURL.Host)
}

func blockNetworkExceptMockServer() (http.RoundTripper, http.RoundTripper) {
	restricted := &http.Transport{
		DialContext: (&OnlyMockServerDialer{}).DialContext,
	}

	origClient, origDefault := http.DefaultClient.Transport, http.DefaultTransport
	http.DefaultClient.Transport, http.DefaultTransport = restricted, restricted
	return origClient, origDefault
}

func restoreNetwork(origClient, origDefault http.RoundTripper) {
	http.DefaultClient.Transport, http.DefaultTransport = origClient, origDefault
}

// TestRunMockTestWithFlags runs a test against a mock server with the provided
// CLI flags and ensures it succeeds
func TestRunMockTestWithFlags(t *testing.T, flags ...string) {
	origClient, origDefault := blockNetworkExceptMockServer()
	defer restoreNetwork(origClient, origDefault)

	// Check if mock server is running
	conn, err := net.DialTimeout("tcp", mockServerURL.Host, 2*time.Second)
	if err != nil {
		require.Fail(t, "Mock server is not running on "+mockServerURL.Host+". Please start the mock server before running tests.")
	} else {
		conn.Close()
	}

	// Get the path to the main command
	_, filename, _, ok := runtime.Caller(0)
	require.True(t, ok, "Could not get current file path")
	dirPath := filepath.Dir(filename)
	project := filepath.Join(dirPath, "..", "..", "cmd", "...")

	args := []string{"run", project, "--base-url", mockServerURL.String()}
	args = append(args, flags...)

	t.Logf("Testing command: telnyx %s", strings.Join(args[4:], " "))

	cmd := exec.Command("go", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		assert.Fail(t, "Test failed", "Error: %v\nOutput: %s", err, output)
	}

	t.Logf("Test passed successfully with output:\n%s\n", output)
}

func TestFile(t *testing.T, contents string) string {
	tmpDir := t.TempDir()
	filename := filepath.Join(tmpDir, "file.txt")
	require.NoError(t, os.WriteFile(filename, []byte(contents), 0644))
	return filename
}
