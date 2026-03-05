// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestMessagingProfilesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "messaging-profiles", "create",
			"--api-key", "string",
			"--name", "My name",
			"--whitelisted-destination", "US",
			"--ai-assistant-id", "ai_assistant_id",
			"--alpha-sender", "sqF",
			"--daily-spend-limit", "269125115713",
			"--daily-spend-limit-enabled=true",
			"--enabled=true",
			"--health-webhook-url", "health_webhook_url",
			"--mms-fall-back-to-sms=true",
			"--mms-transcoding=true",
			"--mobile-only=true",
			"--number-pool-settings", "{long_code_weight: 1, skip_unhealthy: true, toll_free_weight: 10, geomatch: false, sticky_sender: false}",
			"--resource-group-id", "resource_group_id",
			"--smart-encoding=true",
			"--url-shortener-settings", "{domain: example.ex, prefix: '', replace_blacklist_only: true, send_webhooks: false}",
			"--webhook-api-version", "2",
			"--webhook-failover-url", "https://backup.example.com/hooks",
			"--webhook-url", "https://www.example.com/hooks",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(messagingProfilesCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "messaging-profiles", "create",
			"--api-key", "string",
			"--name", "My name",
			"--whitelisted-destination", "US",
			"--ai-assistant-id", "ai_assistant_id",
			"--alpha-sender", "sqF",
			"--daily-spend-limit", "269125115713",
			"--daily-spend-limit-enabled=true",
			"--enabled=true",
			"--health-webhook-url", "health_webhook_url",
			"--mms-fall-back-to-sms=true",
			"--mms-transcoding=true",
			"--mobile-only=true",
			"--number-pool-settings.long-code-weight", "1",
			"--number-pool-settings.skip-unhealthy=true",
			"--number-pool-settings.toll-free-weight", "10",
			"--number-pool-settings.geomatch=false",
			"--number-pool-settings.sticky-sender=false",
			"--resource-group-id", "resource_group_id",
			"--smart-encoding=true",
			"--url-shortener-settings.domain", "example.ex",
			"--url-shortener-settings.prefix", "",
			"--url-shortener-settings.replace-blacklist-only=true",
			"--url-shortener-settings.send-webhooks=false",
			"--webhook-api-version", "2",
			"--webhook-failover-url", "https://backup.example.com/hooks",
			"--webhook-url", "https://www.example.com/hooks",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: My name\n" +
			"whitelisted_destinations:\n" +
			"  - US\n" +
			"ai_assistant_id: ai_assistant_id\n" +
			"alpha_sender: sqF\n" +
			"daily_spend_limit: '269125115713'\n" +
			"daily_spend_limit_enabled: true\n" +
			"enabled: true\n" +
			"health_webhook_url: health_webhook_url\n" +
			"mms_fall_back_to_sms: true\n" +
			"mms_transcoding: true\n" +
			"mobile_only: true\n" +
			"number_pool_settings:\n" +
			"  long_code_weight: 1\n" +
			"  skip_unhealthy: true\n" +
			"  toll_free_weight: 10\n" +
			"  geomatch: false\n" +
			"  sticky_sender: false\n" +
			"resource_group_id: resource_group_id\n" +
			"smart_encoding: true\n" +
			"url_shortener_settings:\n" +
			"  domain: example.ex\n" +
			"  prefix: ''\n" +
			"  replace_blacklist_only: true\n" +
			"  send_webhooks: false\n" +
			"webhook_api_version: '2'\n" +
			"webhook_failover_url: https://backup.example.com/hooks\n" +
			"webhook_url: https://www.example.com/hooks\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "messaging-profiles", "create",
			"--api-key", "string",
		)
	})
}

func TestMessagingProfilesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "messaging-profiles", "retrieve",
			"--api-key", "string",
			"--messaging-profile-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestMessagingProfilesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "messaging-profiles", "update",
			"--api-key", "string",
			"--messaging-profile-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--alpha-sender", "sqF",
			"--daily-spend-limit", "269125115713",
			"--daily-spend-limit-enabled=true",
			"--enabled=true",
			"--mms-fall-back-to-sms=true",
			"--mms-transcoding=true",
			"--mobile-only=true",
			"--name", "Updated Profile for Messages",
			"--number-pool-settings", "{long_code_weight: 2, skip_unhealthy: false, toll_free_weight: 10, geomatch: false, sticky_sender: true}",
			"--smart-encoding=true",
			"--url-shortener-settings", "{domain: example.ex, prefix: cmpny, replace_blacklist_only: true, send_webhooks: false}",
			"--v1-secret", "rP1VamejkU2v0qIUxntqLW2c",
			"--webhook-api-version", "2",
			"--webhook-failover-url", "https://backup.example.com/hooks",
			"--webhook-url", "https://www.example.com/hooks",
			"--whitelisted-destination", "US",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(messagingProfilesUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "messaging-profiles", "update",
			"--api-key", "string",
			"--messaging-profile-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--alpha-sender", "sqF",
			"--daily-spend-limit", "269125115713",
			"--daily-spend-limit-enabled=true",
			"--enabled=true",
			"--mms-fall-back-to-sms=true",
			"--mms-transcoding=true",
			"--mobile-only=true",
			"--name", "Updated Profile for Messages",
			"--number-pool-settings.long-code-weight", "2",
			"--number-pool-settings.skip-unhealthy=false",
			"--number-pool-settings.toll-free-weight", "10",
			"--number-pool-settings.geomatch=false",
			"--number-pool-settings.sticky-sender=true",
			"--smart-encoding=true",
			"--url-shortener-settings.domain", "example.ex",
			"--url-shortener-settings.prefix", "cmpny",
			"--url-shortener-settings.replace-blacklist-only=true",
			"--url-shortener-settings.send-webhooks=false",
			"--v1-secret", "rP1VamejkU2v0qIUxntqLW2c",
			"--webhook-api-version", "2",
			"--webhook-failover-url", "https://backup.example.com/hooks",
			"--webhook-url", "https://www.example.com/hooks",
			"--whitelisted-destination", "US",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"alpha_sender: sqF\n" +
			"daily_spend_limit: '269125115713'\n" +
			"daily_spend_limit_enabled: true\n" +
			"enabled: true\n" +
			"mms_fall_back_to_sms: true\n" +
			"mms_transcoding: true\n" +
			"mobile_only: true\n" +
			"name: Updated Profile for Messages\n" +
			"number_pool_settings:\n" +
			"  long_code_weight: 2\n" +
			"  skip_unhealthy: false\n" +
			"  toll_free_weight: 10\n" +
			"  geomatch: false\n" +
			"  sticky_sender: true\n" +
			"smart_encoding: true\n" +
			"url_shortener_settings:\n" +
			"  domain: example.ex\n" +
			"  prefix: cmpny\n" +
			"  replace_blacklist_only: true\n" +
			"  send_webhooks: false\n" +
			"v1_secret: rP1VamejkU2v0qIUxntqLW2c\n" +
			"webhook_api_version: '2'\n" +
			"webhook_failover_url: https://backup.example.com/hooks\n" +
			"webhook_url: https://www.example.com/hooks\n" +
			"whitelisted_destinations:\n" +
			"  - US\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "messaging-profiles", "update",
			"--api-key", "string",
			"--messaging-profile-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestMessagingProfilesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "messaging-profiles", "list",
			"--api-key", "string",
			"--filter", "{name: name}",
			"--filter-name-contains", "filter[name][contains]",
			"--filter-name-eq", "filter[name][eq]",
			"--page-number", "0",
			"--page-size", "0",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(messagingProfilesList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "messaging-profiles", "list",
			"--api-key", "string",
			"--filter.name", "name",
			"--filter-name-contains", "filter[name][contains]",
			"--filter-name-eq", "filter[name][eq]",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

func TestMessagingProfilesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "messaging-profiles", "delete",
			"--api-key", "string",
			"--messaging-profile-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestMessagingProfilesListAlphanumericSenderIDs(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "messaging-profiles", "list-alphanumeric-sender-ids",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

func TestMessagingProfilesListPhoneNumbers(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "messaging-profiles", "list-phone-numbers",
			"--api-key", "string",
			"--messaging-profile-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

func TestMessagingProfilesListShortCodes(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "messaging-profiles", "list-short-codes",
			"--api-key", "string",
			"--messaging-profile-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

func TestMessagingProfilesRetrieveMetrics(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "messaging-profiles", "retrieve-metrics",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--time-frame", "1h",
		)
	})
}
