// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestDirPhoneNumbersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"dir:phone-numbers", "list",
			"--max-items", "10",
			"--dir-id", "16635d38-75a6-4481-82e8-69af60e05011",
			"--page-number", "1",
			"--page-size", "20",
			"--status", "submitted",
		)
	})
}

func TestDirPhoneNumbersAdd(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"dir:phone-numbers", "add",
			"--dir-id", "16635d38-75a6-4481-82e8-69af60e05011",
			"--document", "{document_id: 2a7e8337-e803-4057-a4ae-26c40eb0bc6c, document_type: letter_of_authorization, description: LOA authorising Telnyx to register these numbers under the DIR.}",
			"--phone-number", "+19493253498",
			"--phone-number", "+12134445566",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(dirPhoneNumbersAdd)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"dir:phone-numbers", "add",
			"--dir-id", "16635d38-75a6-4481-82e8-69af60e05011",
			"--document.document-id", "2a7e8337-e803-4057-a4ae-26c40eb0bc6c",
			"--document.document-type", "letter_of_authorization",
			"--document.description", "LOA authorising Telnyx to register these numbers under the DIR.",
			"--phone-number", "+19493253498",
			"--phone-number", "+12134445566",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"documents:\n" +
			"  - document_id: 2a7e8337-e803-4057-a4ae-26c40eb0bc6c\n" +
			"    document_type: letter_of_authorization\n" +
			"    description: LOA authorising Telnyx to register these numbers under the DIR.\n" +
			"phone_numbers:\n" +
			"  - '+19493253498'\n" +
			"  - '+12134445566'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"dir:phone-numbers", "add",
			"--dir-id", "16635d38-75a6-4481-82e8-69af60e05011",
		)
	})
}

func TestDirPhoneNumbersRemove(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"dir:phone-numbers", "remove",
			"--dir-id", "16635d38-75a6-4481-82e8-69af60e05011",
			"--phone-number", "+19493253498",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"phone_numbers:\n" +
			"  - '+19493253498'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"dir:phone-numbers", "remove",
			"--dir-id", "16635d38-75a6-4481-82e8-69af60e05011",
		)
	})
}
