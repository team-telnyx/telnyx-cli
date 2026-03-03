// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestSessionAnalysisMetadataRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"session-analysis:metadata", "retrieve",
		"--api-key", "string",
	)
}

func TestSessionAnalysisMetadataRetrieveRecordType(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"session-analysis:metadata", "retrieve-record-type",
		"--api-key", "string",
		"--record-type", "record_type",
	)
}
