// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestMessagingHostedNumberOrdersCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messaging-hosted-number-orders", "create",
			"--messaging-profile-id", "dc8f39ac-953d-4520-b93b-786ae87db0da",
			"--phone-number", "+18665550001",
			"--phone-number", "+18665550002",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"messaging_profile_id: dc8f39ac-953d-4520-b93b-786ae87db0da\n" +
			"phone_numbers:\n" +
			"  - '+18665550001'\n" +
			"  - '+18665550002'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"messaging-hosted-number-orders", "create",
		)
	})
}

func TestMessagingHostedNumberOrdersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messaging-hosted-number-orders", "retrieve",
			"--id", "id",
		)
	})
}

func TestMessagingHostedNumberOrdersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messaging-hosted-number-orders", "list",
			"--max-items", "10",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

func TestMessagingHostedNumberOrdersDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messaging-hosted-number-orders", "delete",
			"--id", "id",
		)
	})
}

func TestMessagingHostedNumberOrdersCheckEligibility(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messaging-hosted-number-orders", "check-eligibility",
			"--phone-number", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"phone_numbers:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"messaging-hosted-number-orders", "check-eligibility",
		)
	})
}

func TestMessagingHostedNumberOrdersCreateVerificationCodes(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messaging-hosted-number-orders", "create-verification-codes",
			"--id", "id",
			"--phone-number", "string",
			"--verification-method", "sms",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"phone_numbers:\n" +
			"  - string\n" +
			"verification_method: sms\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"messaging-hosted-number-orders", "create-verification-codes",
			"--id", "id",
		)
	})
}

func TestMessagingHostedNumberOrdersValidateCodes(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messaging-hosted-number-orders", "validate-codes",
			"--id", "id",
			"--verification-code", "{code: code, phone_number: phone_number}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(messagingHostedNumberOrdersValidateCodes)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messaging-hosted-number-orders", "validate-codes",
			"--id", "id",
			"--verification-code.code", "code",
			"--verification-code.phone-number", "phone_number",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"verification_codes:\n" +
			"  - code: code\n" +
			"    phone_number: phone_number\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"messaging-hosted-number-orders", "validate-codes",
			"--id", "id",
		)
	})
}
