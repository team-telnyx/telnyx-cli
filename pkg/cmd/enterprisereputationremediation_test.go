// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestEnterprisesReputationRemediationRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"enterprises:reputation:remediation", "retrieve",
			"--enterprise-id", "4a6192a4-573d-446d-b3ce-aff9117272a6",
			"--remediation-id", "b7c1f1c0-7a9d-4f0a-9d3e-2f6a1c4b8e21",
		)
	})
}

func TestEnterprisesReputationRemediationList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"enterprises:reputation:remediation", "list",
			"--max-items", "10",
			"--enterprise-id", "4a6192a4-573d-446d-b3ce-aff9117272a6",
			"--filter-created-at-gte", "'2026-01-01T00:00:00Z'",
			"--filter-created-at-lte", "'2026-12-31T23:59:59Z'",
			"--filter-status", "in_progress",
			"--page-number", "1",
			"--page-size", "20",
		)
	})
}

func TestEnterprisesReputationRemediationSubmit(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"enterprises:reputation:remediation", "submit",
			"--enterprise-id", "4a6192a4-573d-446d-b3ce-aff9117272a6",
			"--call-purpose", "Appointment reminders for our dental clinic.",
			"--phone-number", "+19493253498",
			"--phone-number", "+12134445566",
			"--contact-email", "ops@example.com",
			"--webhook-url", "https://example.com/webhooks/remediation",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"call_purpose: Appointment reminders for our dental clinic.\n" +
			"phone_numbers:\n" +
			"  - '+19493253498'\n" +
			"  - '+12134445566'\n" +
			"contact_email: ops@example.com\n" +
			"webhook_url: https://example.com/webhooks/remediation\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"enterprises:reputation:remediation", "submit",
			"--enterprise-id", "4a6192a4-573d-446d-b3ce-aff9117272a6",
		)
	})
}
