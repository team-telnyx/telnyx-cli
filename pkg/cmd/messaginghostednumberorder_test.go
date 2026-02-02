// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestMessagingHostedNumberOrdersCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-hosted-number-orders", "create",
		"--messaging-profile-id", "dc8f39ac-953d-4520-b93b-786ae87db0da",
		"--phone-number", "+18665550001",
		"--phone-number", "+18665550002",
	)
}

func TestMessagingHostedNumberOrdersRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-hosted-number-orders", "retrieve",
		"--id", "id",
	)
}

func TestMessagingHostedNumberOrdersList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-hosted-number-orders", "list",
		"--page", "{number: 1, size: 1}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(messagingHostedNumberOrdersList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-hosted-number-orders", "list",
		"--page.number", "1",
		"--page.size", "1",
	)
}

func TestMessagingHostedNumberOrdersDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-hosted-number-orders", "delete",
		"--id", "id",
	)
}

func TestMessagingHostedNumberOrdersCheckEligibility(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-hosted-number-orders", "check-eligibility",
		"--phone-number", "string",
	)
}

func TestMessagingHostedNumberOrdersCreateVerificationCodes(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-hosted-number-orders", "create-verification-codes",
		"--id", "id",
		"--phone-number", "string",
		"--verification-method", "sms",
	)
}

func TestMessagingHostedNumberOrdersValidateCodes(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-hosted-number-orders", "validate-codes",
		"--id", "id",
		"--verification-code", "{code: code, phone_number: phone_number}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(messagingHostedNumberOrdersValidateCodes)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-hosted-number-orders", "validate-codes",
		"--id", "id",
		"--verification-code.code", "code",
		"--verification-code.phone-number", "phone_number",
	)
}
