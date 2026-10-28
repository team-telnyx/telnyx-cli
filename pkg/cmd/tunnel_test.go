// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestTunnelsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tunnels", "create",
			"--name", "chicago-dc",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("name: chicago-dc")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"tunnels", "create",
		)
	})
}

func TestTunnelsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tunnels", "retrieve",
			"--id", "62d438d2-8a51-4c5e-9d0c-1f6b2a7e4c39",
		)
	})
}

func TestTunnelsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tunnels", "list",
			"--filter", "{name: chicago-dc, status: healthy}",
			"--page", "{after: after, limit: 1}",
			"--sort", "-created_at",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(tunnelsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tunnels", "list",
			"--filter.name", "chicago-dc",
			"--filter.status", "healthy",
			"--page.after", "after",
			"--page.limit", "1",
			"--sort", "-created_at",
		)
	})
}

func TestTunnelsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tunnels", "delete",
			"--id", "62d438d2-8a51-4c5e-9d0c-1f6b2a7e4c39",
		)
	})
}
