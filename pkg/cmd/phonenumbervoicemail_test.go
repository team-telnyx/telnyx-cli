// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestPhoneNumbersVoicemailCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"phone-numbers:voicemail", "create",
			"--phone-number-id", "123455678900",
			"--enabled=true",
			"--greeting", "{media_name: my_voicemail_greeting, mode: custom_greeting}",
			"--pin", "1234",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(phoneNumbersVoicemailCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"phone-numbers:voicemail", "create",
			"--phone-number-id", "123455678900",
			"--enabled=true",
			"--greeting.media-name", "my_voicemail_greeting",
			"--greeting.mode", "custom_greeting",
			"--pin", "1234",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"enabled: true\n" +
			"greeting:\n" +
			"  media_name: my_voicemail_greeting\n" +
			"  mode: custom_greeting\n" +
			"pin: '1234'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"phone-numbers:voicemail", "create",
			"--phone-number-id", "123455678900",
		)
	})
}

func TestPhoneNumbersVoicemailRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"phone-numbers:voicemail", "retrieve",
			"--phone-number-id", "123455678900",
		)
	})
}

func TestPhoneNumbersVoicemailUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"phone-numbers:voicemail", "update",
			"--phone-number-id", "123455678900",
			"--enabled=true",
			"--greeting", "{media_name: my_voicemail_greeting, mode: custom_greeting}",
			"--pin", "1234",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(phoneNumbersVoicemailUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"phone-numbers:voicemail", "update",
			"--phone-number-id", "123455678900",
			"--enabled=true",
			"--greeting.media-name", "my_voicemail_greeting",
			"--greeting.mode", "custom_greeting",
			"--pin", "1234",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"enabled: true\n" +
			"greeting:\n" +
			"  media_name: my_voicemail_greeting\n" +
			"  mode: custom_greeting\n" +
			"pin: '1234'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"phone-numbers:voicemail", "update",
			"--phone-number-id", "123455678900",
		)
	})
}
