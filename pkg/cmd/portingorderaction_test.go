// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestPortingOrdersActionsActivate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "porting-orders:actions", "activate",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPortingOrdersActionsCancel(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "porting-orders:actions", "cancel",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPortingOrdersActionsConfirm(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "porting-orders:actions", "confirm",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPortingOrdersActionsShare(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "porting-orders:actions", "share",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--expires-in-seconds", "3600",
			"--permissions", "porting_order.document.read",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"expires_in_seconds: 3600\n" +
			"permissions: porting_order.document.read\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "porting-orders:actions", "share",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
