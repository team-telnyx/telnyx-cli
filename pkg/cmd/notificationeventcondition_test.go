// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestNotificationEventConditionsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"notification-event-conditions", "list",
		"--filter", "{associated_record_type: {eq: phone_number}, channel_type_id: {eq: webhook}, notification_channel: {eq: 12455643-3cf1-4683-ad23-1cd32f7d5e0a}, notification_event_condition_id: {eq: 12455643-3cf1-4683-ad23-1cd32f7d5e0a}, notification_profile_id: {eq: 12455643-3cf1-4683-ad23-1cd32f7d5e0a}, status: {eq: enable-received}}",
		"--page-number", "0",
		"--page-size", "0",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(notificationEventConditionsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"notification-event-conditions", "list",
		"--filter.associated-record-type", "{eq: phone_number}",
		"--filter.channel-type-id", "{eq: webhook}",
		"--filter.notification-channel", "{eq: 12455643-3cf1-4683-ad23-1cd32f7d5e0a}",
		"--filter.notification-event-condition-id", "{eq: 12455643-3cf1-4683-ad23-1cd32f7d5e0a}",
		"--filter.notification-profile-id", "{eq: 12455643-3cf1-4683-ad23-1cd32f7d5e0a}",
		"--filter.status", "{eq: enable-received}",
		"--page-number", "0",
		"--page-size", "0",
	)
}
