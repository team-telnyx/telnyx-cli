// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestNotificationSettingsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"notification-settings", "create",
		"--notification-channel-id", "12455643-3cf1-4683-ad23-1cd32f7d5e0a",
		"--notification-event-condition-id", "70c7c5cb-dce2-4124-accb-870d39dbe852",
		"--notification-profile-id", "12455643-3cf1-4683-ad23-1cd32f7d5e0a",
		"--parameter", "{name: phone_number, value: '+13125550000'}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(notificationSettingsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"notification-settings", "create",
		"--notification-channel-id", "12455643-3cf1-4683-ad23-1cd32f7d5e0a",
		"--notification-event-condition-id", "70c7c5cb-dce2-4124-accb-870d39dbe852",
		"--notification-profile-id", "12455643-3cf1-4683-ad23-1cd32f7d5e0a",
		"--parameter.name", "phone_number",
		"--parameter.value", "+13125550000",
	)
}

func TestNotificationSettingsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"notification-settings", "retrieve",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestNotificationSettingsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"notification-settings", "list",
		"--filter", "{associated_record_type: {eq: phone_number}, channel_type_id: {eq: webhook}, notification_channel: {eq: 12455643-3cf1-4683-ad23-1cd32f7d5e0a}, notification_event_condition_id: {eq: 12455643-3cf1-4683-ad23-1cd32f7d5e0a}, notification_profile_id: {eq: 12455643-3cf1-4683-ad23-1cd32f7d5e0a}, status: {eq: enable-received}}",
		"--page", "{number: 1, size: 1}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(notificationSettingsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"notification-settings", "list",
		"--filter.associated-record-type", "{eq: phone_number}",
		"--filter.channel-type-id", "{eq: webhook}",
		"--filter.notification-channel", "{eq: 12455643-3cf1-4683-ad23-1cd32f7d5e0a}",
		"--filter.notification-event-condition-id", "{eq: 12455643-3cf1-4683-ad23-1cd32f7d5e0a}",
		"--filter.notification-profile-id", "{eq: 12455643-3cf1-4683-ad23-1cd32f7d5e0a}",
		"--filter.status", "{eq: enable-received}",
		"--page.number", "1",
		"--page.size", "1",
	)
}

func TestNotificationSettingsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"notification-settings", "delete",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}
