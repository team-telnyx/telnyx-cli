// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestVerifiedNumbersCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"verified-numbers", "create",
		"--api-key", "string",
		"--phone-number", "+15551234567",
		"--verification-method", "sms",
		"--extension", "ww243w1",
	)
}

func TestVerifiedNumbersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"verified-numbers", "retrieve",
		"--api-key", "string",
		"--phone-number", "+15551234567",
	)
}

func TestVerifiedNumbersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"verified-numbers", "list",
		"--api-key", "string",
		"--page-number", "0",
		"--page-size", "0",
	)
}

func TestVerifiedNumbersDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"verified-numbers", "delete",
		"--api-key", "string",
		"--phone-number", "+15551234567",
	)
}
