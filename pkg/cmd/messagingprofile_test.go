// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestMessagingProfilesCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-profiles", "create",
		"--name", "My name",
		"--whitelisted-destination", "US",
		"--alpha-sender", "sqF",
		"--daily-spend-limit", "269125115713",
		"--daily-spend-limit-enabled=true",
		"--enabled=true",
		"--mms-fall-back-to-sms=true",
		"--mms-transcoding=true",
		"--mobile-only=true",
		"--number-pool-settings", "{long_code_weight: 1, skip_unhealthy: true, toll_free_weight: 10, geomatch: false, sticky_sender: false}",
		"--smart-encoding=true",
		"--url-shortener-settings", "{domain: example.ex, prefix: '', replace_blacklist_only: true, send_webhooks: false}",
		"--webhook-api-version", "2",
		"--webhook-failover-url", "https://backup.example.com/hooks",
		"--webhook-url", "https://www.example.com/hooks",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(messagingProfilesCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-profiles", "create",
		"--name", "My name",
		"--whitelisted-destination", "US",
		"--alpha-sender", "sqF",
		"--daily-spend-limit", "269125115713",
		"--daily-spend-limit-enabled=true",
		"--enabled=true",
		"--mms-fall-back-to-sms=true",
		"--mms-transcoding=true",
		"--mobile-only=true",
		"--number-pool-settings.long-code-weight", "1",
		"--number-pool-settings.skip-unhealthy=true",
		"--number-pool-settings.toll-free-weight", "10",
		"--number-pool-settings.geomatch=false",
		"--number-pool-settings.sticky-sender=false",
		"--smart-encoding=true",
		"--url-shortener-settings.domain", "example.ex",
		"--url-shortener-settings.prefix", "",
		"--url-shortener-settings.replace-blacklist-only=true",
		"--url-shortener-settings.send-webhooks=false",
		"--webhook-api-version", "2",
		"--webhook-failover-url", "https://backup.example.com/hooks",
		"--webhook-url", "https://www.example.com/hooks",
	)
}

func TestMessagingProfilesRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-profiles", "retrieve",
		"--messaging-profile-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestMessagingProfilesUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-profiles", "update",
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

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(messagingProfilesUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-profiles", "update",
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
}

func TestMessagingProfilesList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-profiles", "list",
		"--filter", "{name: name}",
		"--page", "{number: 1, size: 1}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(messagingProfilesList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-profiles", "list",
		"--filter.name", "name",
		"--page.number", "1",
		"--page.size", "1",
	)
}

func TestMessagingProfilesDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-profiles", "delete",
		"--messaging-profile-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestMessagingProfilesListPhoneNumbers(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-profiles", "list-phone-numbers",
		"--messaging-profile-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--page", "{number: 1, size: 1}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(messagingProfilesListPhoneNumbers)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-profiles", "list-phone-numbers",
		"--messaging-profile-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--page.number", "1",
		"--page.size", "1",
	)
}

func TestMessagingProfilesListShortCodes(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-profiles", "list-short-codes",
		"--messaging-profile-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--page", "{number: 1, size: 1}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(messagingProfilesListShortCodes)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-profiles", "list-short-codes",
		"--messaging-profile-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--page.number", "1",
		"--page.size", "1",
	)
}
