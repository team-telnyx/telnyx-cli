// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestVerifyProfilesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"verify-profiles", "create",
			"--name", "Test Profile",
			"--call", "{app_name: Example Secure App, code_length: 6, default_verification_timeout_secs: 300, messaging_template_id: 0abb5b4f-459f-445a-bfcd-488998b7572d, whitelisted_destinations: [US, CA]}",
			"--daily-spend-limit", "100",
			"--daily-spend-limit-enabled=true",
			"--flashcall", "{app_name: Example Secure App, default_verification_timeout_secs: 300, whitelisted_destinations: [US, CA]}",
			"--language", "en-US",
			"--sms", "{alpha_sender: sqF, app_name: Example Secure App, code_length: 6, default_verification_timeout_secs: 300, messaging_template_id: 0abb5b4f-459f-445a-bfcd-488998b7572d, whitelisted_destinations: [US, CA]}",
			"--webhook-failover-url", "http://example.com/webhook/failover",
			"--webhook-url", "http://example.com/webhook",
			"--whatsapp", "{default_verification_timeout_secs: 300, sender_phone_number: '+13035551234', template_id: authentication_template_name, waba_id: '1234567890', whitelisted_destinations: [US, CA]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(verifyProfilesCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"verify-profiles", "create",
			"--name", "Test Profile",
			"--call.app-name", "Example Secure App",
			"--call.code-length", "6",
			"--call.default-verification-timeout-secs", "300",
			"--call.messaging-template-id", "0abb5b4f-459f-445a-bfcd-488998b7572d",
			"--call.whitelisted-destinations", "[US, CA]",
			"--daily-spend-limit", "100",
			"--daily-spend-limit-enabled=true",
			"--flashcall.app-name", "Example Secure App",
			"--flashcall.default-verification-timeout-secs", "300",
			"--flashcall.whitelisted-destinations", "[US, CA]",
			"--language", "en-US",
			"--sms.alpha-sender", "sqF",
			"--sms.app-name", "Example Secure App",
			"--sms.code-length", "6",
			"--sms.default-verification-timeout-secs", "300",
			"--sms.messaging-template-id", "0abb5b4f-459f-445a-bfcd-488998b7572d",
			"--sms.whitelisted-destinations", "[US, CA]",
			"--webhook-failover-url", "http://example.com/webhook/failover",
			"--webhook-url", "http://example.com/webhook",
			"--whatsapp.default-verification-timeout-secs", "300",
			"--whatsapp.sender-phone-number", "+13035551234",
			"--whatsapp.template-id", "authentication_template_name",
			"--whatsapp.waba-id", "1234567890",
			"--whatsapp.whitelisted-destinations", "[US, CA]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: Test Profile\n" +
			"call:\n" +
			"  app_name: Example Secure App\n" +
			"  code_length: 6\n" +
			"  default_verification_timeout_secs: 300\n" +
			"  messaging_template_id: 0abb5b4f-459f-445a-bfcd-488998b7572d\n" +
			"  whitelisted_destinations:\n" +
			"    - US\n" +
			"    - CA\n" +
			"daily_spend_limit: 100\n" +
			"daily_spend_limit_enabled: true\n" +
			"flashcall:\n" +
			"  app_name: Example Secure App\n" +
			"  default_verification_timeout_secs: 300\n" +
			"  whitelisted_destinations:\n" +
			"    - US\n" +
			"    - CA\n" +
			"language: en-US\n" +
			"sms:\n" +
			"  alpha_sender: sqF\n" +
			"  app_name: Example Secure App\n" +
			"  code_length: 6\n" +
			"  default_verification_timeout_secs: 300\n" +
			"  messaging_template_id: 0abb5b4f-459f-445a-bfcd-488998b7572d\n" +
			"  whitelisted_destinations:\n" +
			"    - US\n" +
			"    - CA\n" +
			"webhook_failover_url: http://example.com/webhook/failover\n" +
			"webhook_url: http://example.com/webhook\n" +
			"whatsapp:\n" +
			"  default_verification_timeout_secs: 300\n" +
			"  sender_phone_number: '+13035551234'\n" +
			"  template_id: authentication_template_name\n" +
			"  waba_id: '1234567890'\n" +
			"  whitelisted_destinations:\n" +
			"    - US\n" +
			"    - CA\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"verify-profiles", "create",
		)
	})
}

func TestVerifyProfilesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"verify-profiles", "retrieve",
			"--verify-profile-id", "12ade33a-21c0-473b-b055-b3c836e1c292",
		)
	})
}

func TestVerifyProfilesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"verify-profiles", "update",
			"--verify-profile-id", "12ade33a-21c0-473b-b055-b3c836e1c292",
			"--call", "{app_name: Example Secure App, code_length: 6, default_verification_timeout_secs: 300, messaging_template_id: 0abb5b4f-459f-445a-bfcd-488998b7572d, whitelisted_destinations: [US, CA]}",
			"--daily-spend-limit", "100",
			"--daily-spend-limit-enabled=true",
			"--language", "en-US",
			"--name", "Test Profile",
			"--sms", "{alpha_sender: sqF, app_name: Example Secure App, code_length: 6, default_verification_timeout_secs: 300, messaging_template_id: 0abb5b4f-459f-445a-bfcd-488998b7572d, whitelisted_destinations: [US, CA]}",
			"--webhook-failover-url", "http://example.com/webhook/failover",
			"--webhook-url", "http://example.com/webhook",
			"--whatsapp", "{default_verification_timeout_secs: 300, sender_phone_number: '+13035551234', template_id: authentication_template_name, waba_id: '1234567890', whitelisted_destinations: [US, CA]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(verifyProfilesUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"verify-profiles", "update",
			"--verify-profile-id", "12ade33a-21c0-473b-b055-b3c836e1c292",
			"--call.app-name", "Example Secure App",
			"--call.code-length", "6",
			"--call.default-verification-timeout-secs", "300",
			"--call.messaging-template-id", "0abb5b4f-459f-445a-bfcd-488998b7572d",
			"--call.whitelisted-destinations", "[US, CA]",
			"--daily-spend-limit", "100",
			"--daily-spend-limit-enabled=true",
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
			"--whatsapp.default-verification-timeout-secs", "300",
			"--whatsapp.sender-phone-number", "+13035551234",
			"--whatsapp.template-id", "authentication_template_name",
			"--whatsapp.waba-id", "1234567890",
			"--whatsapp.whitelisted-destinations", "[US, CA]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"call:\n" +
			"  app_name: Example Secure App\n" +
			"  code_length: 6\n" +
			"  default_verification_timeout_secs: 300\n" +
			"  messaging_template_id: 0abb5b4f-459f-445a-bfcd-488998b7572d\n" +
			"  whitelisted_destinations:\n" +
			"    - US\n" +
			"    - CA\n" +
			"daily_spend_limit: 100\n" +
			"daily_spend_limit_enabled: true\n" +
			"language: en-US\n" +
			"name: Test Profile\n" +
			"sms:\n" +
			"  alpha_sender: sqF\n" +
			"  app_name: Example Secure App\n" +
			"  code_length: 6\n" +
			"  default_verification_timeout_secs: 300\n" +
			"  messaging_template_id: 0abb5b4f-459f-445a-bfcd-488998b7572d\n" +
			"  whitelisted_destinations:\n" +
			"    - US\n" +
			"    - CA\n" +
			"webhook_failover_url: http://example.com/webhook/failover\n" +
			"webhook_url: http://example.com/webhook\n" +
			"whatsapp:\n" +
			"  default_verification_timeout_secs: 300\n" +
			"  sender_phone_number: '+13035551234'\n" +
			"  template_id: authentication_template_name\n" +
			"  waba_id: '1234567890'\n" +
			"  whitelisted_destinations:\n" +
			"    - US\n" +
			"    - CA\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"verify-profiles", "update",
			"--verify-profile-id", "12ade33a-21c0-473b-b055-b3c836e1c292",
		)
	})
}

func TestVerifyProfilesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"verify-profiles", "list",
			"--max-items", "10",
			"--filter", "{name: name}",
			"--page-number", "0",
			"--page-size", "0",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(verifyProfilesList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"verify-profiles", "list",
			"--max-items", "10",
			"--filter.name", "name",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

func TestVerifyProfilesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"verify-profiles", "delete",
			"--verify-profile-id", "12ade33a-21c0-473b-b055-b3c836e1c292",
		)
	})
}

func TestVerifyProfilesCreateTemplate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"verify-profiles", "create-template",
			"--text", "Your {{app_name}} verification code is: {{code}}.",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("text: 'Your {{app_name}} verification code is: {{code}}.'")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"verify-profiles", "create-template",
		)
	})
}

func TestVerifyProfilesRetrieveTemplates(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"verify-profiles", "retrieve-templates",
		)
	})
}

func TestVerifyProfilesUpdateTemplate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"verify-profiles", "update-template",
			"--template-id", "12ade33a-21c0-473b-b055-b3c836e1c292",
			"--text", "Your {{app_name}} verification code is: {{code}}.",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("text: 'Your {{app_name}} verification code is: {{code}}.'")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"verify-profiles", "update-template",
			"--template-id", "12ade33a-21c0-473b-b055-b3c836e1c292",
		)
	})
}
