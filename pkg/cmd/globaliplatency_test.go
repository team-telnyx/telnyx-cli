// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestGlobalIPLatencyRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"global-ip-latency", "retrieve",
		"--filter", "{global_ip_id: string}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(globalIPLatencyRetrieve)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"global-ip-latency", "retrieve",
		"--filter.global-ip-id", "string",
	)
}
