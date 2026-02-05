// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestPortingOrdersVerificationCodesList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders:verification-codes", "list",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--filter", "{verified: true}",
		"--page-number", "0",
		"--page-size", "0",
		"--sort", "{value: created_at}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(portingOrdersVerificationCodesList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders:verification-codes", "list",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--filter.verified=true",
		"--page-number", "0",
		"--page-size", "0",
		"--sort.value", "created_at",
	)
}

func TestPortingOrdersVerificationCodesSend(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders:verification-codes", "send",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--phone-number", "+61424000001",
		"--phone-number", "+61424000002",
		"--verification-method", "sms",
	)
}

func TestPortingOrdersVerificationCodesVerify(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders:verification-codes", "verify",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--verification-code", "{code: '12345', phone_number: '+61424000001'}",
		"--verification-code", "{code: '54321', phone_number: '+61424000002'}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(portingOrdersVerificationCodesVerify)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders:verification-codes", "verify",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--verification-code.code", "12345",
		"--verification-code.phone-number", "+61424000001",
		"--verification-code.code", "54321",
		"--verification-code.phone-number", "+61424000002",
	)
}
