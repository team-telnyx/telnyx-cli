// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestPhoneNumbersVoicemailCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "phone-numbers:voicemail", "create",
			"--api-key", "string",
			"--phone-number-id", "123455678900",
			"--enabled=true",
			"--pin", "1234",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"enabled: true\n" +
			"pin: '1234'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "phone-numbers:voicemail", "create",
			"--api-key", "string",
			"--phone-number-id", "123455678900",
		)
	})
}

func TestPhoneNumbersVoicemailRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "phone-numbers:voicemail", "retrieve",
			"--api-key", "string",
			"--phone-number-id", "123455678900",
		)
	})
}

func TestPhoneNumbersVoicemailUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "phone-numbers:voicemail", "update",
			"--api-key", "string",
			"--phone-number-id", "123455678900",
			"--enabled=true",
			"--pin", "1234",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"enabled: true\n" +
			"pin: '1234'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "phone-numbers:voicemail", "update",
			"--api-key", "string",
			"--phone-number-id", "123455678900",
		)
	})
}
