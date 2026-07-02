// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestPortingOrdersActivationJobsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"porting-orders:activation-jobs", "retrieve",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--activation-job-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPortingOrdersActivationJobsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"porting-orders:activation-jobs", "update",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--activation-job-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--activate-at", "'2019-01-01T00:00:00Z'",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("activate_at: '2019-01-01T00:00:00Z'")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"porting-orders:activation-jobs", "update",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--activation-job-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPortingOrdersActivationJobsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"porting-orders:activation-jobs", "list",
			"--max-items", "10",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}
