// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestNotificationChannelsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"notification-channels", "create",
			"--channel-destination", "+13125550000",
			"--channel-type-id", "sms",
			"--notification-profile-id", "12455643-3cf1-4683-ad23-1cd32f7d5e0a",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"channel_destination: '+13125550000'\n" +
			"channel_type_id: sms\n" +
			"notification_profile_id: 12455643-3cf1-4683-ad23-1cd32f7d5e0a\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"notification-channels", "create",
		)
	})
}

func TestNotificationChannelsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"notification-channels", "retrieve",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestNotificationChannelsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"notification-channels", "update",
			"--notification-channel-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--channel-destination", "+13125550000",
			"--channel-type-id", "sms",
			"--notification-profile-id", "12455643-3cf1-4683-ad23-1cd32f7d5e0a",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"channel_destination: '+13125550000'\n" +
			"channel_type_id: sms\n" +
			"notification_profile_id: 12455643-3cf1-4683-ad23-1cd32f7d5e0a\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"notification-channels", "update",
			"--notification-channel-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestNotificationChannelsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"notification-channels", "list",
			"--max-items", "10",
			"--filter", "{associated_record_type: {eq: phone_number}, channel_type_id: {eq: webhook}, notification_channel: {eq: 12455643-3cf1-4683-ad23-1cd32f7d5e0a}, notification_event_condition_id: {eq: 12455643-3cf1-4683-ad23-1cd32f7d5e0a}, notification_profile_id: {eq: 12455643-3cf1-4683-ad23-1cd32f7d5e0a}, status: {eq: enable-received}}",
			"--page-number", "0",
			"--page-size", "0",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(notificationChannelsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"notification-channels", "list",
			"--max-items", "10",
			"--filter.associated-record-type", "{eq: phone_number}",
			"--filter.channel-type-id", "{eq: webhook}",
			"--filter.notification-channel", "{eq: 12455643-3cf1-4683-ad23-1cd32f7d5e0a}",
			"--filter.notification-event-condition-id", "{eq: 12455643-3cf1-4683-ad23-1cd32f7d5e0a}",
			"--filter.notification-profile-id", "{eq: 12455643-3cf1-4683-ad23-1cd32f7d5e0a}",
			"--filter.status", "{eq: enable-received}",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

func TestNotificationChannelsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"notification-channels", "delete",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
