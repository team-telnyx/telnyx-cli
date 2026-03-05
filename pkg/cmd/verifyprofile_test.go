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
			t, "verify-profiles", "create",
			"--api-key", "string",
			"--name", "Test Profile",
			"--call", "{app_name: Example Secure App, code_length: 6, default_verification_timeout_secs: 300, messaging_template_id: 0abb5b4f-459f-445a-bfcd-488998b7572d, whitelisted_destinations: [US, CA]}",
			"--flashcall", "{app_name: Example Secure App, default_verification_timeout_secs: 300, whitelisted_destinations: [US, CA]}",
			"--language", "en-US",
			"--rcs", "{app_name: Example Secure App, code_length: 6, default_verification_timeout_secs: 300, messaging_template_id: 0abb5b4f-459f-445a-bfcd-488998b7572d, sms_fallback: true, whitelisted_destinations: [US, CA]}",
			"--sms", "{alpha_sender: sqF, app_name: Example Secure App, code_length: 6, default_verification_timeout_secs: 300, messaging_template_id: 0abb5b4f-459f-445a-bfcd-488998b7572d, whitelisted_destinations: [US, CA]}",
			"--webhook-failover-url", "http://example.com/webhook/failover",
			"--webhook-url", "http://example.com/webhook",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(verifyProfilesCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "verify-profiles", "create",
			"--api-key", "string",
			"--name", "Test Profile",
			"--call.app-name", "Example Secure App",
			"--call.code-length", "6",
			"--call.default-verification-timeout-secs", "300",
			"--call.messaging-template-id", "0abb5b4f-459f-445a-bfcd-488998b7572d",
			"--call.whitelisted-destinations", "[US, CA]",
			"--flashcall.app-name", "Example Secure App",
			"--flashcall.default-verification-timeout-secs", "300",
			"--flashcall.whitelisted-destinations", "[US, CA]",
			"--language", "en-US",
			"--rcs.app-name", "Example Secure App",
			"--rcs.code-length", "6",
			"--rcs.default-verification-timeout-secs", "300",
			"--rcs.messaging-template-id", "0abb5b4f-459f-445a-bfcd-488998b7572d",
			"--rcs.sms-fallback=true",
			"--rcs.whitelisted-destinations", "[US, CA]",
			"--sms.alpha-sender", "sqF",
			"--sms.app-name", "Example Secure App",
			"--sms.code-length", "6",
			"--sms.default-verification-timeout-secs", "300",
			"--sms.messaging-template-id", "0abb5b4f-459f-445a-bfcd-488998b7572d",
			"--sms.whitelisted-destinations", "[US, CA]",
			"--webhook-failover-url", "http://example.com/webhook/failover",
			"--webhook-url", "http://example.com/webhook",
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
			"flashcall:\n" +
			"  app_name: Example Secure App\n" +
			"  default_verification_timeout_secs: 300\n" +
			"  whitelisted_destinations:\n" +
			"    - US\n" +
			"    - CA\n" +
			"language: en-US\n" +
			"rcs:\n" +
			"  app_name: Example Secure App\n" +
			"  code_length: 6\n" +
			"  default_verification_timeout_secs: 300\n" +
			"  messaging_template_id: 0abb5b4f-459f-445a-bfcd-488998b7572d\n" +
			"  sms_fallback: true\n" +
			"  whitelisted_destinations:\n" +
			"    - US\n" +
			"    - CA\n" +
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
			"webhook_url: http://example.com/webhook\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "verify-profiles", "create",
			"--api-key", "string",
		)
	})
}

func TestVerifyProfilesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "verify-profiles", "retrieve",
			"--api-key", "string",
			"--verify-profile-id", "12ade33a-21c0-473b-b055-b3c836e1c292",
		)
	})
}

func TestVerifyProfilesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "verify-profiles", "update",
			"--api-key", "string",
			"--verify-profile-id", "12ade33a-21c0-473b-b055-b3c836e1c292",
			"--call", "{app_name: Example Secure App, code_length: 6, default_verification_timeout_secs: 300, messaging_template_id: 0abb5b4f-459f-445a-bfcd-488998b7572d, whitelisted_destinations: [US, CA]}",
			"--flashcall", "{app_name: Example Secure App, default_verification_timeout_secs: 300, whitelisted_destinations: [US, CA]}",
			"--language", "en-US",
			"--name", "Test Profile",
			"--rcs", "{app_name: Example Secure App, code_length: 6, default_verification_timeout_secs: 300, messaging_template_id: 0abb5b4f-459f-445a-bfcd-488998b7572d, sms_fallback: true, whitelisted_destinations: [US, CA]}",
			"--sms", "{alpha_sender: sqF, app_name: Example Secure App, code_length: 6, default_verification_timeout_secs: 300, messaging_template_id: 0abb5b4f-459f-445a-bfcd-488998b7572d, whitelisted_destinations: [US, CA]}",
			"--webhook-failover-url", "http://example.com/webhook/failover",
			"--webhook-url", "http://example.com/webhook",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(verifyProfilesUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "verify-profiles", "update",
			"--api-key", "string",
			"--verify-profile-id", "12ade33a-21c0-473b-b055-b3c836e1c292",
			"--call.app-name", "Example Secure App",
			"--call.code-length", "6",
			"--call.default-verification-timeout-secs", "300",
			"--call.messaging-template-id", "0abb5b4f-459f-445a-bfcd-488998b7572d",
			"--call.whitelisted-destinations", "[US, CA]",
			"--flashcall.app-name", "Example Secure App",
			"--flashcall.default-verification-timeout-secs", "300",
			"--flashcall.whitelisted-destinations", "[US, CA]",
			"--language", "en-US",
			"--name", "Test Profile",
			"--rcs.app-name", "Example Secure App",
			"--rcs.code-length", "6",
			"--rcs.default-verification-timeout-secs", "300",
			"--rcs.messaging-template-id", "0abb5b4f-459f-445a-bfcd-488998b7572d",
			"--rcs.sms-fallback=true",
			"--rcs.whitelisted-destinations", "[US, CA]",
			"--sms.alpha-sender", "sqF",
			"--sms.app-name", "Example Secure App",
			"--sms.code-length", "6",
			"--sms.default-verification-timeout-secs", "300",
			"--sms.messaging-template-id", "0abb5b4f-459f-445a-bfcd-488998b7572d",
			"--sms.whitelisted-destinations", "[US, CA]",
			"--webhook-failover-url", "http://example.com/webhook/failover",
			"--webhook-url", "http://example.com/webhook",
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
			"flashcall:\n" +
			"  app_name: Example Secure App\n" +
			"  default_verification_timeout_secs: 300\n" +
			"  whitelisted_destinations:\n" +
			"    - US\n" +
			"    - CA\n" +
			"language: en-US\n" +
			"name: Test Profile\n" +
			"rcs:\n" +
			"  app_name: Example Secure App\n" +
			"  code_length: 6\n" +
			"  default_verification_timeout_secs: 300\n" +
			"  messaging_template_id: 0abb5b4f-459f-445a-bfcd-488998b7572d\n" +
			"  sms_fallback: true\n" +
			"  whitelisted_destinations:\n" +
			"    - US\n" +
			"    - CA\n" +
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
			"webhook_url: http://example.com/webhook\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "verify-profiles", "update",
			"--api-key", "string",
			"--verify-profile-id", "12ade33a-21c0-473b-b055-b3c836e1c292",
		)
	})
}

func TestVerifyProfilesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "verify-profiles", "list",
			"--api-key", "string",
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
			t, "verify-profiles", "list",
			"--api-key", "string",
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
			t, "verify-profiles", "delete",
			"--api-key", "string",
			"--verify-profile-id", "12ade33a-21c0-473b-b055-b3c836e1c292",
		)
	})
}

func TestVerifyProfilesCreateTemplate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "verify-profiles", "create-template",
			"--api-key", "string",
			"--text", "Your {{app_name}} verification code is: {{code}}.",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("text: 'Your {{app_name}} verification code is: {{code}}.'")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "verify-profiles", "create-template",
			"--api-key", "string",
		)
	})
}

func TestVerifyProfilesRetrieveTemplates(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "verify-profiles", "retrieve-templates",
			"--api-key", "string",
		)
	})
}

func TestVerifyProfilesUpdateTemplate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "verify-profiles", "update-template",
			"--api-key", "string",
			"--template-id", "12ade33a-21c0-473b-b055-b3c836e1c292",
			"--text", "Your {{app_name}} verification code is: {{code}}.",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("text: 'Your {{app_name}} verification code is: {{code}}.'")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "verify-profiles", "update-template",
			"--api-key", "string",
			"--template-id", "12ade33a-21c0-473b-b055-b3c836e1c292",
		)
	})
}
