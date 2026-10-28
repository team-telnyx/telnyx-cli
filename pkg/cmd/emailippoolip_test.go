// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestEmailIPPoolsIPsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email:ip-pools:ips", "create",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--dedicated-ip-id", "6a7f9a30-1234-4b5c-8d9e-0f1a2b3c4d5e",
			"--weight", "100",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"dedicated_ip_id: 6a7f9a30-1234-4b5c-8d9e-0f1a2b3c4d5e\n" +
			"weight: 100\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"email:ip-pools:ips", "create",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestEmailIPPoolsIPsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email:ip-pools:ips", "delete",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--ip-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--drain=true",
			"--force=true",
		)
	})
}
