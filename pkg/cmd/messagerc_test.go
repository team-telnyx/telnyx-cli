// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestMessagesRcsGenerateDeeplink(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messages:rcs", "generate-deeplink",
		"--agent-id", "agent_id",
		"--body", "body",
		"--phone-number", "phone_number",
	)
}

func TestMessagesRcsSend(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messages:rcs", "send",
		"--agent-id", "Agent007",
		"--agent-message", "{content_message: {content_info: {file_url: https://example.com/elephant.jpg, force_refresh: true, thumbnail_url: thumbnail_url}, rich_card: {carousel_card: {card_contents: [{description: description, media: {content_info: {file_url: https://example.com/elephant.jpg, force_refresh: true, thumbnail_url: thumbnail_url}, height: MEDIUM}, suggestions: [{action: {create_calendar_event_action: {description: description, end_time: '2024-10-02T15:02:31Z', start_time: '2024-10-02T15:01:23Z', title: title}, dial_action: {phone_number: '+13125551234'}, fallback_url: fallback_url, open_url_action: {application: BROWSER, url: http://example.com, webview_view_mode: HALF, description: description}, postback_data: postback_data, share_location_action: {foo: bar}, text: Hello world, view_location_action: {label: label, lat_long: {latitude: 41.8, longitude: -87.6}, query: query}}, reply: {postback_data: postback_data, text: text}}], title: Elephant}], card_width: SMALL}, standalone_card: {card_content: {description: description, media: {content_info: {file_url: https://example.com/elephant.jpg, force_refresh: true, thumbnail_url: thumbnail_url}, height: MEDIUM}, suggestions: [{action: {create_calendar_event_action: {description: description, end_time: '2024-10-02T15:02:31Z', start_time: '2024-10-02T15:01:23Z', title: title}, dial_action: {phone_number: '+13125551234'}, fallback_url: fallback_url, open_url_action: {application: BROWSER, url: http://example.com, webview_view_mode: HALF, description: description}, postback_data: postback_data, share_location_action: {foo: bar}, text: Hello world, view_location_action: {label: label, lat_long: {latitude: 41.8, longitude: -87.6}, query: query}}, reply: {postback_data: postback_data, text: text}}], title: Elephant}, card_orientation: HORIZONTAL, thumbnail_image_alignment: LEFT}}, suggestions: [{action: {create_calendar_event_action: {description: description, end_time: '2024-10-02T15:02:31Z', start_time: '2024-10-02T15:01:23Z', title: title}, dial_action: {phone_number: '+13125551234'}, fallback_url: fallback_url, open_url_action: {application: BROWSER, url: http://example.com, webview_view_mode: HALF, description: description}, postback_data: postback_data, share_location_action: {foo: bar}, text: Hello world, view_location_action: {label: label, lat_long: {latitude: 41.8, longitude: -87.6}, query: query}}, reply: {postback_data: postback_data, text: text}}], text: Hello world!}, event: {event_type: IS_TYPING}, expire_time: '2024-10-02T15:01:23Z', ttl: 10.5s}",
		"--messaging-profile-id", "messaging_profile_id",
		"--to", "+13125551234",
		"--mms-fallback", "{from: '+13125551234', media_urls: [string], subject: Test Message, text: Hello world!}",
		"--sms-fallback", "{from: '+13125551234', text: Hello world!}",
		"--type", "RCS",
		"--webhook-url", "webhook_url",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(messagesRcsSend)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"messages:rcs", "send",
		"--agent-id", "Agent007",
		"--agent-message.content-message", "{content_info: {file_url: https://example.com/elephant.jpg, force_refresh: true, thumbnail_url: thumbnail_url}, rich_card: {carousel_card: {card_contents: [{description: description, media: {content_info: {file_url: https://example.com/elephant.jpg, force_refresh: true, thumbnail_url: thumbnail_url}, height: MEDIUM}, suggestions: [{action: {create_calendar_event_action: {description: description, end_time: '2024-10-02T15:02:31Z', start_time: '2024-10-02T15:01:23Z', title: title}, dial_action: {phone_number: '+13125551234'}, fallback_url: fallback_url, open_url_action: {application: BROWSER, url: http://example.com, webview_view_mode: HALF, description: description}, postback_data: postback_data, share_location_action: {foo: bar}, text: Hello world, view_location_action: {label: label, lat_long: {latitude: 41.8, longitude: -87.6}, query: query}}, reply: {postback_data: postback_data, text: text}}], title: Elephant}], card_width: SMALL}, standalone_card: {card_content: {description: description, media: {content_info: {file_url: https://example.com/elephant.jpg, force_refresh: true, thumbnail_url: thumbnail_url}, height: MEDIUM}, suggestions: [{action: {create_calendar_event_action: {description: description, end_time: '2024-10-02T15:02:31Z', start_time: '2024-10-02T15:01:23Z', title: title}, dial_action: {phone_number: '+13125551234'}, fallback_url: fallback_url, open_url_action: {application: BROWSER, url: http://example.com, webview_view_mode: HALF, description: description}, postback_data: postback_data, share_location_action: {foo: bar}, text: Hello world, view_location_action: {label: label, lat_long: {latitude: 41.8, longitude: -87.6}, query: query}}, reply: {postback_data: postback_data, text: text}}], title: Elephant}, card_orientation: HORIZONTAL, thumbnail_image_alignment: LEFT}}, suggestions: [{action: {create_calendar_event_action: {description: description, end_time: '2024-10-02T15:02:31Z', start_time: '2024-10-02T15:01:23Z', title: title}, dial_action: {phone_number: '+13125551234'}, fallback_url: fallback_url, open_url_action: {application: BROWSER, url: http://example.com, webview_view_mode: HALF, description: description}, postback_data: postback_data, share_location_action: {foo: bar}, text: Hello world, view_location_action: {label: label, lat_long: {latitude: 41.8, longitude: -87.6}, query: query}}, reply: {postback_data: postback_data, text: text}}], text: Hello world!}",
		"--agent-message.event", "{event_type: IS_TYPING}",
		"--agent-message.expire-time", "2024-10-02T15:01:23Z",
		"--agent-message.ttl", "10.5s",
		"--messaging-profile-id", "messaging_profile_id",
		"--to", "+13125551234",
		"--mms-fallback.from", "+13125551234",
		"--mms-fallback.media-urls", "[string]",
		"--mms-fallback.subject", "Test Message",
		"--mms-fallback.text", "Hello world!",
		"--sms-fallback.from", "+13125551234",
		"--sms-fallback.text", "Hello world!",
		"--type", "RCS",
		"--webhook-url", "webhook_url",
	)
}
