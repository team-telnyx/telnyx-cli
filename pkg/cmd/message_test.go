// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestMessagesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messages", "retrieve",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestMessagesCancelScheduled(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messages", "cancel-scheduled",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestMessagesRetrieveGroupMessages(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messages", "retrieve-group-messages",
			"--message-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestMessagesSchedule(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messages", "schedule",
			"--to", "+18445550001",
			"--auto-detect=true",
			"--from", "+18445550001",
			"--media-url", "string",
			"--messaging-profile-id", "abc85f64-5717-4562-b3fc-2c9600000000",
			"--send-at", "'2019-01-23T18:30:00Z'",
			"--subject", "From Telnyx!",
			"--text", "Hello, World!",
			"--type", "SMS",
			"--use-profile-webhooks=true",
			"--webhook-failover-url", "https://backup.example.com/hooks",
			"--webhook-url", "http://example.com/webhooks",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"to: '+18445550001'\n" +
			"auto_detect: true\n" +
			"from: '+18445550001'\n" +
			"media_urls:\n" +
			"  - string\n" +
			"messaging_profile_id: abc85f64-5717-4562-b3fc-2c9600000000\n" +
			"send_at: '2019-01-23T18:30:00Z'\n" +
			"subject: From Telnyx!\n" +
			"text: Hello, World!\n" +
			"type: SMS\n" +
			"use_profile_webhooks: true\n" +
			"webhook_failover_url: https://backup.example.com/hooks\n" +
			"webhook_url: http://example.com/webhooks\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"messages", "schedule",
		)
	})
}

func TestMessagesSend(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messages", "send",
			"--to", "+18445550001",
			"--auto-detect=true",
			"--encoding", "auto",
			"--from", "+18445550001",
			"--media-url", "http://example.com",
			"--messaging-profile-id", "abc85f64-5717-4562-b3fc-2c9600000000",
			"--send-at", "'2019-12-27T18:11:19.117Z'",
			"--subject", "From Telnyx!",
			"--text", "Hello, World!",
			"--type", "MMS",
			"--use-profile-webhooks=true",
			"--webhook-failover-url", "https://backup.example.com/hooks",
			"--webhook-url", "http://example.com/webhooks",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"to: '+18445550001'\n" +
			"auto_detect: true\n" +
			"encoding: auto\n" +
			"from: '+18445550001'\n" +
			"media_urls:\n" +
			"  - http://example.com\n" +
			"messaging_profile_id: abc85f64-5717-4562-b3fc-2c9600000000\n" +
			"send_at: '2019-12-27T18:11:19.117Z'\n" +
			"subject: From Telnyx!\n" +
			"text: Hello, World!\n" +
			"type: MMS\n" +
			"use_profile_webhooks: true\n" +
			"webhook_failover_url: https://backup.example.com/hooks\n" +
			"webhook_url: http://example.com/webhooks\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"messages", "send",
		)
	})
}

func TestMessagesSendGroupMms(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"from: '+13125551234'\n" +
			"to:\n" +
			"  - '+18655551234'\n" +
			"  - '+14155551234'\n" +
			"media_urls:\n" +
			"  - http://example.com\n" +
			"subject: From Telnyx!\n" +
			"text: Hello, World!\n" +
			"use_profile_webhooks: true\n" +
			"webhook_failover_url: https://backup.example.com/hooks\n" +
			"webhook_url: http://example.com/webhooks\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"messages", "send-group-mms",
		)
	})
}

func TestMessagesSendLongCode(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"from: '+18445550001'\n" +
			"to: '+13125550002'\n" +
			"auto_detect: true\n" +
			"encoding: auto\n" +
			"media_urls:\n" +
			"  - http://example.com\n" +
			"subject: From Telnyx!\n" +
			"text: Hello, World!\n" +
			"type: MMS\n" +
			"use_profile_webhooks: true\n" +
			"webhook_failover_url: https://backup.example.com/hooks\n" +
			"webhook_url: http://example.com/webhooks\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"messages", "send-long-code",
		)
	})
}

func TestMessagesSendNumberPool(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"messaging_profile_id: abc85f64-5717-4562-b3fc-2c9600000000\n" +
			"to: '+13125550002'\n" +
			"auto_detect: true\n" +
			"encoding: auto\n" +
			"media_urls:\n" +
			"  - http://example.com\n" +
			"subject: From Telnyx!\n" +
			"text: Hello, World!\n" +
			"type: MMS\n" +
			"use_profile_webhooks: true\n" +
			"webhook_failover_url: https://backup.example.com/hooks\n" +
			"webhook_url: http://example.com/webhooks\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"messages", "send-number-pool",
		)
	})
}

func TestMessagesSendShortCode(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"from: '+18445550001'\n" +
			"to: '+18445550001'\n" +
			"auto_detect: true\n" +
			"encoding: auto\n" +
			"media_urls:\n" +
			"  - http://example.com\n" +
			"subject: From Telnyx!\n" +
			"text: Hello, World!\n" +
			"type: MMS\n" +
			"use_profile_webhooks: true\n" +
			"webhook_failover_url: https://backup.example.com/hooks\n" +
			"webhook_url: http://example.com/webhooks\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"messages", "send-short-code",
		)
	})
}

func TestMessagesSendWithAlphanumericSender(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messages", "send-with-alphanumeric-sender",
			"--from", "MyCompany",
			"--messaging-profile-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--text", "text",
			"--to", "+E.164",
			"--use-profile-webhooks=true",
			"--webhook-failover-url", "webhook_failover_url",
			"--webhook-url", "webhook_url",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"from: MyCompany\n" +
			"messaging_profile_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"text: text\n" +
			"to: +E.164\n" +
			"use_profile_webhooks: true\n" +
			"webhook_failover_url: webhook_failover_url\n" +
			"webhook_url: webhook_url\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"messages", "send-with-alphanumeric-sender",
		)
	})
}
