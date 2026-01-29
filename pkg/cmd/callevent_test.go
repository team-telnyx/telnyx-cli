// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestCallEventsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"call-events", "list",
		"--filter", "{application_name: {contains: contains}, application_session_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, connection_id: connection_id, failed: false, from: '+12025550142', leg_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, name: name, occurred_at: {eq: '2019-03-29T11:10:00Z', gt: '2019-03-29T11:10:00Z', gte: '2019-03-29T11:10:00Z', lt: '2019-03-29T11:10:00Z', lte: '2019-03-29T11:10:00Z'}, outbound.outbound_voice_profile_id: outbound.outbound_voice_profile_id, product: texml, status: init, to: '+12025550142', type: webhook}",
		"--page", "{after: after, before: before, limit: 1}",
		"--page-number", "0",
		"--page-size", "0",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(callEventsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"call-events", "list",
		"--filter.application-name", "{contains: contains}",
		"--filter.application-session-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--filter.connection-id", "connection_id",
		"--filter.failed=false",
		"--filter.from", "+12025550142",
		"--filter.leg-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--filter.name", "name",
		"--filter.occurred-at", "{eq: '2019-03-29T11:10:00Z', gt: '2019-03-29T11:10:00Z', gte: '2019-03-29T11:10:00Z', lt: '2019-03-29T11:10:00Z', lte: '2019-03-29T11:10:00Z'}",
		"--filter.outbound-outbound-voice-profile-id", "outbound.outbound_voice_profile_id",
		"--filter.product", "texml",
		"--filter.status", "init",
		"--filter.to", "+12025550142",
		"--filter.type", "webhook",
		"--page.after", "after",
		"--page.before", "before",
		"--page.limit", "1",
		"--page-number", "0",
		"--page-size", "0",
	)
}
