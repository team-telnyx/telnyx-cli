// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestPortingOrdersActivationJobsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders:activation-jobs", "retrieve",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--activation-job-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestPortingOrdersActivationJobsUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders:activation-jobs", "update",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--activation-job-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--activate-at", "2019-01-01T00:00:00Z",
	)
}

func TestPortingOrdersActivationJobsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders:activation-jobs", "list",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--page-number", "0",
		"--page-size", "0",
	)
}
