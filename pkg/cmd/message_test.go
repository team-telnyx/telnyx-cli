// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestMessagesRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messages", "retrieve",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestMessagesCancelScheduled(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messages", "cancel-scheduled",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestMessagesSchedule(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messages", "schedule",
		"--to", "+18445550001",
		"--auto-detect=true",
		"--from", "+18445550001",
		"--media-url", "string",
		"--messaging-profile-id", "abc85f64-5717-4562-b3fc-2c9600000000",
		"--send-at", "2019-01-23T18:30:00Z",
		"--subject", "From Telnyx!",
		"--text", "Hello, World!",
		"--type", "SMS",
		"--use-profile-webhooks=true",
		"--webhook-failover-url", "https://backup.example.com/hooks",
		"--webhook-url", "http://example.com/webhooks",
	)
}

func TestMessagesSend(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messages", "send",
		"--to", "+18445550001",
		"--auto-detect=true",
		"--encoding", "auto",
		"--from", "+18445550001",
		"--media-url", "http://example.com",
		"--messaging-profile-id", "abc85f64-5717-4562-b3fc-2c9600000000",
		"--send-at", "2019-12-27T18:11:19.117Z",
		"--subject", "From Telnyx!",
		"--text", "Hello, World!",
		"--type", "MMS",
		"--use-profile-webhooks=true",
		"--webhook-failover-url", "https://backup.example.com/hooks",
		"--webhook-url", "http://example.com/webhooks",
	)
}

func TestMessagesSendGroupMms(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messages", "send-group-mms",
		"--from", "+13125551234",
		"--to", "+18655551234",
		"--to", "+14155551234",
		"--media-url", "http://example.com",
		"--subject", "From Telnyx!",
		"--text", "Hello, World!",
		"--use-profile-webhooks=true",
		"--webhook-failover-url", "https://backup.example.com/hooks",
		"--webhook-url", "http://example.com/webhooks",
	)
}

func TestMessagesSendLongCode(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messages", "send-long-code",
		"--from", "+18445550001",
		"--to", "+13125550002",
		"--auto-detect=true",
		"--encoding", "auto",
		"--media-url", "http://example.com",
		"--subject", "From Telnyx!",
		"--text", "Hello, World!",
		"--type", "MMS",
		"--use-profile-webhooks=true",
		"--webhook-failover-url", "https://backup.example.com/hooks",
		"--webhook-url", "http://example.com/webhooks",
	)
}

func TestMessagesSendNumberPool(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messages", "send-number-pool",
		"--messaging-profile-id", "abc85f64-5717-4562-b3fc-2c9600000000",
		"--to", "+13125550002",
		"--auto-detect=true",
		"--encoding", "auto",
		"--media-url", "http://example.com",
		"--subject", "From Telnyx!",
		"--text", "Hello, World!",
		"--type", "MMS",
		"--use-profile-webhooks=true",
		"--webhook-failover-url", "https://backup.example.com/hooks",
		"--webhook-url", "http://example.com/webhooks",
	)
}

func TestMessagesSendShortCode(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messages", "send-short-code",
		"--from", "+18445550001",
		"--to", "+18445550001",
		"--auto-detect=true",
		"--encoding", "auto",
		"--media-url", "http://example.com",
		"--subject", "From Telnyx!",
		"--text", "Hello, World!",
		"--type", "MMS",
		"--use-profile-webhooks=true",
		"--webhook-failover-url", "https://backup.example.com/hooks",
		"--webhook-url", "http://example.com/webhooks",
	)
}

func TestMessagesSendWhatsapp(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messages", "send-whatsapp",
		"--from", "+13125551234",
		"--to", "+13125551234",
		"--whatsapp-message", "{audio: {caption: caption, filename: filename, link: http://example.com/media.jpg, voice: true}, biz_opaque_callback_data: biz_opaque_callback_data, contacts: [{addresses: [{city: city, country: country, country_code: country_code, state: state, street: street, type: type, zip: zip}], birthday: birthday, emails: [{email: email, type: type}], name: name, org: {company: company, department: department, title: title}, phones: [{phone: phone, type: type, wa_id: wa_id}], urls: [{type: type, url: url}]}], document: {caption: caption, filename: filename, link: http://example.com/media.jpg, voice: true}, image: {caption: caption, filename: filename, link: http://example.com/media.jpg, voice: true}, interactive: {action: {button: button, buttons: [{reply: {id: id, title: title}, type: reply}], cards: [{action: {catalog_id: catalog_id, product_retailer_id: product_retailer_id}, body: {text: text}, card_index: 0, header: {image: {caption: caption, filename: filename, link: http://example.com/media.jpg, voice: true}, type: image, video: {caption: caption, filename: filename, link: http://example.com/media.jpg, voice: true}}, type: cta_url}], catalog_id: catalog_id, mode: mode, name: name, parameters: {display_text: display_text, url: url}, product_retailer_id: product_retailer_id, sections: [{product_items: [{product_retailer_id: product_retailer_id}], rows: [{id: id, description: description, title: title}], title: title}]}, body: {text: text}, footer: {text: text}, header: {document: {caption: caption, filename: filename, link: http://example.com/media.jpg, voice: true}, image: {caption: caption, filename: filename, link: http://example.com/media.jpg, voice: true}, sub_text: sub_text, text: text, video: {caption: caption, filename: filename, link: http://example.com/media.jpg, voice: true}}, type: cta_url}, location: {address: address, latitude: latitude, longitude: longitude, name: name}, reaction: {emoji: emoji, message_id: message_id}, sticker: {caption: caption, filename: filename, link: http://example.com/media.jpg, voice: true}, type: audio, video: {caption: caption, filename: filename, link: http://example.com/media.jpg, voice: true}}",
		"--type", "WHATSAPP",
		"--webhook-url", "webhook_url",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(messagesSendWhatsapp)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"messages", "send-whatsapp",
		"--from", "+13125551234",
		"--to", "+13125551234",
		"--whatsapp-message.audio", "{caption: caption, filename: filename, link: http://example.com/media.jpg, voice: true}",
		"--whatsapp-message.biz-opaque-callback-data", "biz_opaque_callback_data",
		"--whatsapp-message.contacts", "[{addresses: [{city: city, country: country, country_code: country_code, state: state, street: street, type: type, zip: zip}], birthday: birthday, emails: [{email: email, type: type}], name: name, org: {company: company, department: department, title: title}, phones: [{phone: phone, type: type, wa_id: wa_id}], urls: [{type: type, url: url}]}]",
		"--whatsapp-message.document", "{caption: caption, filename: filename, link: http://example.com/media.jpg, voice: true}",
		"--whatsapp-message.image", "{caption: caption, filename: filename, link: http://example.com/media.jpg, voice: true}",
		"--whatsapp-message.interactive", "{action: {button: button, buttons: [{reply: {id: id, title: title}, type: reply}], cards: [{action: {catalog_id: catalog_id, product_retailer_id: product_retailer_id}, body: {text: text}, card_index: 0, header: {image: {caption: caption, filename: filename, link: http://example.com/media.jpg, voice: true}, type: image, video: {caption: caption, filename: filename, link: http://example.com/media.jpg, voice: true}}, type: cta_url}], catalog_id: catalog_id, mode: mode, name: name, parameters: {display_text: display_text, url: url}, product_retailer_id: product_retailer_id, sections: [{product_items: [{product_retailer_id: product_retailer_id}], rows: [{id: id, description: description, title: title}], title: title}]}, body: {text: text}, footer: {text: text}, header: {document: {caption: caption, filename: filename, link: http://example.com/media.jpg, voice: true}, image: {caption: caption, filename: filename, link: http://example.com/media.jpg, voice: true}, sub_text: sub_text, text: text, video: {caption: caption, filename: filename, link: http://example.com/media.jpg, voice: true}}, type: cta_url}",
		"--whatsapp-message.location", "{address: address, latitude: latitude, longitude: longitude, name: name}",
		"--whatsapp-message.reaction", "{emoji: emoji, message_id: message_id}",
		"--whatsapp-message.sticker", "{caption: caption, filename: filename, link: http://example.com/media.jpg, voice: true}",
		"--whatsapp-message.type", "audio",
		"--whatsapp-message.video", "{caption: caption, filename: filename, link: http://example.com/media.jpg, voice: true}",
		"--type", "WHATSAPP",
		"--webhook-url", "webhook_url",
	)
}
