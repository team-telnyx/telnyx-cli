// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestWebhookDeliveriesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "webhook-deliveries", "retrieve",
			"--api-key", "string",
			"--id", "C9C0797E-901D-4349-A33C-C2C8F31A92C2",
		)
	})
}

func TestWebhookDeliveriesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "webhook-deliveries", "list",
			"--api-key", "string",
			"--filter", "{attempts: {contains: https://fallback.example.com/webhooks}, event_type: 'call_initiated,call.initiated', finished_at: {gte: '2019-03-29T11:10:00Z', lte: '2019-03-29T11:10:00Z'}, started_at: {gte: '2019-03-29T11:10:00Z', lte: '2019-03-29T11:10:00Z'}, status: {eq: delivered}, webhook: {contains: call.initiated}}",
			"--page-number", "0",
			"--page-size", "0",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(webhookDeliveriesList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "webhook-deliveries", "list",
			"--api-key", "string",
			"--filter.attempts", "{contains: https://fallback.example.com/webhooks}",
			"--filter.event-type", "call_initiated,call.initiated",
			"--filter.finished-at", "{gte: '2019-03-29T11:10:00Z', lte: '2019-03-29T11:10:00Z'}",
			"--filter.started-at", "{gte: '2019-03-29T11:10:00Z', lte: '2019-03-29T11:10:00Z'}",
			"--filter.status", "{eq: delivered}",
			"--filter.webhook", "{contains: call.initiated}",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}
