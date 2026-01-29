// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestVerifyProfilesCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"verify-profiles", "create",
		"--name", "Test Profile",
		"--call", "{app_name: Example Secure App, code_length: 6, default_verification_timeout_secs: 300, messaging_template_id: 0abb5b4f-459f-445a-bfcd-488998b7572d, whitelisted_destinations: [US, CA]}",
		"--flashcall", "{default_verification_timeout_secs: 300, whitelisted_destinations: [US, CA]}",
		"--language", "en-US",
		"--sms", "{whitelisted_destinations: [US, CA], alpha_sender: sqF, app_name: Example Secure App, code_length: 6, default_verification_timeout_secs: 300, messaging_template_id: 0abb5b4f-459f-445a-bfcd-488998b7572d}",
		"--webhook-failover-url", "http://example.com/webhook/failover",
		"--webhook-url", "http://example.com/webhook",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(verifyProfilesCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"verify-profiles", "create",
		"--name", "Test Profile",
		"--call.app-name", "Example Secure App",
		"--call.code-length", "6",
		"--call.default-verification-timeout-secs", "300",
		"--call.messaging-template-id", "0abb5b4f-459f-445a-bfcd-488998b7572d",
		"--call.whitelisted-destinations", "[US, CA]",
		"--flashcall.default-verification-timeout-secs", "300",
		"--flashcall.whitelisted-destinations", "[US, CA]",
		"--language", "en-US",
		"--sms.whitelisted-destinations", "[US, CA]",
		"--sms.alpha-sender", "sqF",
		"--sms.app-name", "Example Secure App",
		"--sms.code-length", "6",
		"--sms.default-verification-timeout-secs", "300",
		"--sms.messaging-template-id", "0abb5b4f-459f-445a-bfcd-488998b7572d",
		"--webhook-failover-url", "http://example.com/webhook/failover",
		"--webhook-url", "http://example.com/webhook",
	)
}

func TestVerifyProfilesRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"verify-profiles", "retrieve",
		"--verify-profile-id", "12ade33a-21c0-473b-b055-b3c836e1c292",
	)
}

func TestVerifyProfilesUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"verify-profiles", "update",
		"--verify-profile-id", "12ade33a-21c0-473b-b055-b3c836e1c292",
		"--call", "{app_name: Example Secure App, code_length: 6, default_verification_timeout_secs: 300, messaging_template_id: 0abb5b4f-459f-445a-bfcd-488998b7572d, whitelisted_destinations: [US, CA]}",
		"--flashcall", "{default_verification_timeout_secs: 300, whitelisted_destinations: [US, CA]}",
		"--language", "en-US",
		"--name", "Test Profile",
		"--sms", "{alpha_sender: sqF, app_name: Example Secure App, code_length: 6, default_verification_timeout_secs: 300, messaging_template_id: 0abb5b4f-459f-445a-bfcd-488998b7572d, whitelisted_destinations: [US, CA]}",
		"--webhook-failover-url", "http://example.com/webhook/failover",
		"--webhook-url", "http://example.com/webhook",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(verifyProfilesUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"verify-profiles", "update",
		"--verify-profile-id", "12ade33a-21c0-473b-b055-b3c836e1c292",
		"--call.app-name", "Example Secure App",
		"--call.code-length", "6",
		"--call.default-verification-timeout-secs", "300",
		"--call.messaging-template-id", "0abb5b4f-459f-445a-bfcd-488998b7572d",
		"--call.whitelisted-destinations", "[US, CA]",
		"--flashcall.default-verification-timeout-secs", "300",
		"--flashcall.whitelisted-destinations", "[US, CA]",
		"--language", "en-US",
		"--name", "Test Profile",
		"--sms.alpha-sender", "sqF",
		"--sms.app-name", "Example Secure App",
		"--sms.code-length", "6",
		"--sms.default-verification-timeout-secs", "300",
		"--sms.messaging-template-id", "0abb5b4f-459f-445a-bfcd-488998b7572d",
		"--sms.whitelisted-destinations", "[US, CA]",
		"--webhook-failover-url", "http://example.com/webhook/failover",
		"--webhook-url", "http://example.com/webhook",
	)
}

func TestVerifyProfilesList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"verify-profiles", "list",
		"--filter", "{name: name}",
		"--page-number", "0",
		"--page-size", "0",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(verifyProfilesList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"verify-profiles", "list",
		"--filter.name", "name",
		"--page-number", "0",
		"--page-size", "0",
	)
}

func TestVerifyProfilesDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"verify-profiles", "delete",
		"--verify-profile-id", "12ade33a-21c0-473b-b055-b3c836e1c292",
	)
}

func TestVerifyProfilesCreateTemplate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"verify-profiles", "create-template",
		"--text", "Your {{app_name}} verification code is: {{code}}.",
	)
}

func TestVerifyProfilesRetrieveTemplates(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"verify-profiles", "retrieve-templates",
	)
}

func TestVerifyProfilesUpdateTemplate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"verify-profiles", "update-template",
		"--template-id", "12ade33a-21c0-473b-b055-b3c836e1c292",
		"--text", "Your {{app_name}} verification code is: {{code}}.",
	)
}
