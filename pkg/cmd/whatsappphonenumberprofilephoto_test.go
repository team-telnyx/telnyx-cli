// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"strings"
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestWhatsappPhoneNumbersProfilePhotoRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"whatsapp:phone-numbers:profile:photo", "retrieve",
			"--phone-number", "phone_number",
		)
	})
}

func TestWhatsappPhoneNumbersProfilePhotoDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"whatsapp:phone-numbers:profile:photo", "delete",
			"--phone-number", "phone_number",
		)
	})
}

func TestWhatsappPhoneNumbersProfilePhotoUpload(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"whatsapp:phone-numbers:profile:photo", "upload",
			"--phone-number", "phone_number",
			"--file", mocktest.TestFile(t, "Example data"),
		)
	})

	t.Run("piping data", func(t *testing.T) {
		testFile := mocktest.TestFile(t, "Example data")
		// Test piping YAML data over stdin
		pipeDataStr := "file: Example data"
		pipeDataStr = strings.ReplaceAll(pipeDataStr, "Example data", testFile)
		pipeData := []byte(pipeDataStr)
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"whatsapp:phone-numbers:profile:photo", "upload",
			"--phone-number", "phone_number",
		)
	})
}
