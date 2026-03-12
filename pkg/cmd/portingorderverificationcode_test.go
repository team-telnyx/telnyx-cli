// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestPortingOrdersVerificationCodesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "porting-orders:verification-codes", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--filter", "{verified: true}",
			"--page-number", "0",
			"--page-size", "0",
			"--sort", "{value: created_at}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(portingOrdersVerificationCodesList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "porting-orders:verification-codes", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--filter.verified=true",
			"--page-number", "0",
			"--page-size", "0",
			"--sort.value", "created_at",
		)
	})
}

func TestPortingOrdersVerificationCodesSend(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "porting-orders:verification-codes", "send",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--phone-number", "+61424000001",
			"--phone-number", "+61424000002",
			"--verification-method", "sms",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"phone_numbers:\n" +
			"  - '+61424000001'\n" +
			"  - '+61424000002'\n" +
			"verification_method: sms\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "porting-orders:verification-codes", "send",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPortingOrdersVerificationCodesVerify(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "porting-orders:verification-codes", "verify",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--verification-code", "{code: '12345', phone_number: '+61424000001'}",
			"--verification-code", "{code: '54321', phone_number: '+61424000002'}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(portingOrdersVerificationCodesVerify)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "porting-orders:verification-codes", "verify",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--verification-code.code", "12345",
			"--verification-code.phone-number", "+61424000001",
			"--verification-code.code", "54321",
			"--verification-code.phone-number", "+61424000002",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"verification_codes:\n" +
			"  - code: '12345'\n" +
			"    phone_number: '+61424000001'\n" +
			"  - code: '54321'\n" +
			"    phone_number: '+61424000002'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "porting-orders:verification-codes", "verify",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
