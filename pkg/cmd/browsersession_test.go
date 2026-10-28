// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestBrowserSessionsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"browser-sessions", "create",
			"--policy", "{allowed_domains: [example.com]}",
			"--idempotency-key", "x",
			"--limits", "{idle_timeout_s: 30, lifetime_s: 60}",
			"--metadata", "{}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(browserSessionsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"browser-sessions", "create",
			"--policy.allowed-domains", "[example.com]",
			"--idempotency-key", "x",
			"--limits.idle-timeout-s", "30",
			"--limits.lifetime-s", "60",
			"--metadata", "{}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"policy:\n" +
			"  allowed_domains:\n" +
			"    - example.com\n" +
			"idempotency_key: x\n" +
			"limits:\n" +
			"  idle_timeout_s: 30\n" +
			"  lifetime_s: 60\n" +
			"metadata: {}\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"browser-sessions", "create",
		)
	})
}

func TestBrowserSessionsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"browser-sessions", "retrieve",
			"--id", "id",
		)
	})
}

func TestBrowserSessionsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"browser-sessions", "list",
			"--status", "provisioning",
		)
	})
}

func TestBrowserSessionsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"browser-sessions", "delete",
			"--id", "id",
		)
	})
}

func TestBrowserSessionsRetrieveArtifact(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"browser-sessions", "retrieve-artifact",
			"--id", "id",
			"--artifact-id", "artifactId",
		)
	})
}
