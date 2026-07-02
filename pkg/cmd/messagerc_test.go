// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestMessagesRcsGenerateDeeplink(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messages:rcs", "generate-deeplink",
			"--agent-id", "agent_id",
			"--body", "body",
			"--phone-number", "phone_number",
		)
	})
}

func TestMessagesRcsSend(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(messagesRcsSend)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"agent_id: Agent007\n" +
			"agent_message:\n" +
			"  content_message:\n" +
			"    content_info:\n" +
			"      file_url: https://example.com/elephant.jpg\n" +
			"      force_refresh: true\n" +
			"      thumbnail_url: thumbnail_url\n" +
			"    rich_card:\n" +
			"      carousel_card:\n" +
			"        card_contents:\n" +
			"          - description: description\n" +
			"            media:\n" +
			"              content_info:\n" +
			"                file_url: https://example.com/elephant.jpg\n" +
			"                force_refresh: true\n" +
			"                thumbnail_url: thumbnail_url\n" +
			"              height: MEDIUM\n" +
			"            suggestions:\n" +
			"              - action:\n" +
			"                  create_calendar_event_action:\n" +
			"                    description: description\n" +
			"                    end_time: '2024-10-02T15:02:31Z'\n" +
			"                    start_time: '2024-10-02T15:01:23Z'\n" +
			"                    title: title\n" +
			"                  dial_action:\n" +
			"                    phone_number: '+13125551234'\n" +
			"                  fallback_url: fallback_url\n" +
			"                  open_url_action:\n" +
			"                    application: BROWSER\n" +
			"                    url: http://example.com\n" +
			"                    webview_view_mode: HALF\n" +
			"                    description: description\n" +
			"                  postback_data: postback_data\n" +
			"                  share_location_action:\n" +
			"                    foo: bar\n" +
			"                  text: Hello world\n" +
			"                  view_location_action:\n" +
			"                    label: label\n" +
			"                    lat_long:\n" +
			"                      latitude: 41.8\n" +
			"                      longitude: -87.6\n" +
			"                    query: query\n" +
			"                reply:\n" +
			"                  postback_data: postback_data\n" +
			"                  text: text\n" +
			"            title: Elephant\n" +
			"        card_width: SMALL\n" +
			"      standalone_card:\n" +
			"        card_content:\n" +
			"          description: description\n" +
			"          media:\n" +
			"            content_info:\n" +
			"              file_url: https://example.com/elephant.jpg\n" +
			"              force_refresh: true\n" +
			"              thumbnail_url: thumbnail_url\n" +
			"            height: MEDIUM\n" +
			"          suggestions:\n" +
			"            - action:\n" +
			"                create_calendar_event_action:\n" +
			"                  description: description\n" +
			"                  end_time: '2024-10-02T15:02:31Z'\n" +
			"                  start_time: '2024-10-02T15:01:23Z'\n" +
			"                  title: title\n" +
			"                dial_action:\n" +
			"                  phone_number: '+13125551234'\n" +
			"                fallback_url: fallback_url\n" +
			"                open_url_action:\n" +
			"                  application: BROWSER\n" +
			"                  url: http://example.com\n" +
			"                  webview_view_mode: HALF\n" +
			"                  description: description\n" +
			"                postback_data: postback_data\n" +
			"                share_location_action:\n" +
			"                  foo: bar\n" +
			"                text: Hello world\n" +
			"                view_location_action:\n" +
			"                  label: label\n" +
			"                  lat_long:\n" +
			"                    latitude: 41.8\n" +
			"                    longitude: -87.6\n" +
			"                  query: query\n" +
			"              reply:\n" +
			"                postback_data: postback_data\n" +
			"                text: text\n" +
			"          title: Elephant\n" +
			"        card_orientation: HORIZONTAL\n" +
			"        thumbnail_image_alignment: LEFT\n" +
			"    suggestions:\n" +
			"      - action:\n" +
			"          create_calendar_event_action:\n" +
			"            description: description\n" +
			"            end_time: '2024-10-02T15:02:31Z'\n" +
			"            start_time: '2024-10-02T15:01:23Z'\n" +
			"            title: title\n" +
			"          dial_action:\n" +
			"            phone_number: '+13125551234'\n" +
			"          fallback_url: fallback_url\n" +
			"          open_url_action:\n" +
			"            application: BROWSER\n" +
			"            url: http://example.com\n" +
			"            webview_view_mode: HALF\n" +
			"            description: description\n" +
			"          postback_data: postback_data\n" +
			"          share_location_action:\n" +
			"            foo: bar\n" +
			"          text: Hello world\n" +
			"          view_location_action:\n" +
			"            label: label\n" +
			"            lat_long:\n" +
			"              latitude: 41.8\n" +
			"              longitude: -87.6\n" +
			"            query: query\n" +
			"        reply:\n" +
			"          postback_data: postback_data\n" +
			"          text: text\n" +
			"    text: Hello world!\n" +
			"  event:\n" +
			"    event_type: IS_TYPING\n" +
			"  expire_time: '2024-10-02T15:01:23Z'\n" +
			"  ttl: 10.5s\n" +
			"messaging_profile_id: messaging_profile_id\n" +
			"to: '+13125551234'\n" +
			"mms_fallback:\n" +
			"  from: '+13125551234'\n" +
			"  media_urls:\n" +
			"    - string\n" +
			"  subject: Test Message\n" +
			"  text: Hello world!\n" +
			"sms_fallback:\n" +
			"  from: '+13125551234'\n" +
			"  text: Hello world!\n" +
			"type: RCS\n" +
			"webhook_url: webhook_url\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"messages:rcs", "send",
		)
	})
}
