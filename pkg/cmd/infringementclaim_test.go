// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestInfringementClaimsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"infringement-claims", "retrieve",
			"--claim-id", "e379fbc8-cd83-4bef-a280-a0ac9d00dcf8",
		)
	})
}

func TestInfringementClaimsContest(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"infringement-claims", "contest",
			"--claim-id", "e379fbc8-cd83-4bef-a280-a0ac9d00dcf8",
			"--contest-notes", "We own the trademark outright; our registration precedes the claimant by three years. See attached certificate.",
			"--document", "{document_id: 2a7e8337-e803-4057-a4ae-26c40eb0bc6c, document_type: trademark_registration, description: USPTO trademark certificate.}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(infringementClaimsContest)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"infringement-claims", "contest",
			"--claim-id", "e379fbc8-cd83-4bef-a280-a0ac9d00dcf8",
			"--contest-notes", "We own the trademark outright; our registration precedes the claimant by three years. See attached certificate.",
			"--document.document-id", "2a7e8337-e803-4057-a4ae-26c40eb0bc6c",
			"--document.document-type", "trademark_registration",
			"--document.description", "USPTO trademark certificate.",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"contest_notes: >-\n" +
			"  We own the trademark outright; our registration precedes the claimant by three\n" +
			"  years. See attached certificate.\n" +
			"documents:\n" +
			"  - document_id: 2a7e8337-e803-4057-a4ae-26c40eb0bc6c\n" +
			"    document_type: trademark_registration\n" +
			"    description: USPTO trademark certificate.\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"infringement-claims", "contest",
			"--claim-id", "e379fbc8-cd83-4bef-a280-a0ac9d00dcf8",
		)
	})
}
