// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestRcsAgentsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"rcs:agents", "create",
			"--brand-id", "11111111-1111-4111-8111-111111111111",
			"--configuration", "{basics: {email: {address: support@example.com, label: Support}, brand_color: '#123456', description: Order confirmations and delivery updates, hero_url: https://www.example.com/rcs/hero.png, logo_url: https://www.example.com/rcs/logo.png, phone_number: {label: x, number: '+49605132'}, privacy_policy_url: https://www.example.com/privacy, terms_and_conditions_url: https://www.example.com/terms, website: {label: x, url: https://example.com}}, campaign: {company_overview: x, additional_information: x, agent_overview: x, consent_settings: {call_to_action: x, double_opt_in: true, help_response: x, opt_in_message: x, opt_in_methods: [{method_type: SMS, description: x}], opt_out_response: x, call_to_action_media_url: https://example.com, call_to_action_url: https://example.com, double_opt_in_message: x}, interactions: [{interaction_type: TRANSACTIONAL_UPDATES, description: x}], message_examples: [x]}, testing: {test_url: https://example.com, additional_information: x, message_id: x}}",
			"--display-name", "Acme Order Updates",
			"--use-case", "TRANSACTIONAL",
			"--idempotency-key", "Idempotency-Key",
			"--hosting-region", "hosting_region",
			"--profile-id", "profile_id",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(rcsAgentsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"rcs:agents", "create",
			"--brand-id", "11111111-1111-4111-8111-111111111111",
			"--configuration.basics", "{email: {address: support@example.com, label: Support}, brand_color: '#123456', description: Order confirmations and delivery updates, hero_url: https://www.example.com/rcs/hero.png, logo_url: https://www.example.com/rcs/logo.png, phone_number: {label: x, number: '+49605132'}, privacy_policy_url: https://www.example.com/privacy, terms_and_conditions_url: https://www.example.com/terms, website: {label: x, url: https://example.com}}",
			"--configuration.campaign", "{company_overview: x, additional_information: x, agent_overview: x, consent_settings: {call_to_action: x, double_opt_in: true, help_response: x, opt_in_message: x, opt_in_methods: [{method_type: SMS, description: x}], opt_out_response: x, call_to_action_media_url: https://example.com, call_to_action_url: https://example.com, double_opt_in_message: x}, interactions: [{interaction_type: TRANSACTIONAL_UPDATES, description: x}], message_examples: [x]}",
			"--configuration.testing", "{test_url: https://example.com, additional_information: x, message_id: x}",
			"--display-name", "Acme Order Updates",
			"--use-case", "TRANSACTIONAL",
			"--idempotency-key", "Idempotency-Key",
			"--hosting-region", "hosting_region",
			"--profile-id", "profile_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"brand_id: 11111111-1111-4111-8111-111111111111\n" +
			"configuration:\n" +
			"  basics:\n" +
			"    email:\n" +
			"      address: support@example.com\n" +
			"      label: Support\n" +
			"    brand_color: '#123456'\n" +
			"    description: Order confirmations and delivery updates\n" +
			"    hero_url: https://www.example.com/rcs/hero.png\n" +
			"    logo_url: https://www.example.com/rcs/logo.png\n" +
			"    phone_number:\n" +
			"      label: x\n" +
			"      number: '+49605132'\n" +
			"    privacy_policy_url: https://www.example.com/privacy\n" +
			"    terms_and_conditions_url: https://www.example.com/terms\n" +
			"    website:\n" +
			"      label: x\n" +
			"      url: https://example.com\n" +
			"  campaign:\n" +
			"    company_overview: x\n" +
			"    additional_information: x\n" +
			"    agent_overview: x\n" +
			"    consent_settings:\n" +
			"      call_to_action: x\n" +
			"      double_opt_in: true\n" +
			"      help_response: x\n" +
			"      opt_in_message: x\n" +
			"      opt_in_methods:\n" +
			"        - method_type: SMS\n" +
			"          description: x\n" +
			"      opt_out_response: x\n" +
			"      call_to_action_media_url: https://example.com\n" +
			"      call_to_action_url: https://example.com\n" +
			"      double_opt_in_message: x\n" +
			"    interactions:\n" +
			"      - interaction_type: TRANSACTIONAL_UPDATES\n" +
			"        description: x\n" +
			"    message_examples:\n" +
			"      - x\n" +
			"  testing:\n" +
			"    test_url: https://example.com\n" +
			"    additional_information: x\n" +
			"    message_id: x\n" +
			"display_name: Acme Order Updates\n" +
			"use_case: TRANSACTIONAL\n" +
			"hosting_region: hosting_region\n" +
			"profile_id: profile_id\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"rcs:agents", "create",
			"--idempotency-key", "Idempotency-Key",
		)
	})
}

func TestRcsAgentsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"rcs:agents", "retrieve",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestRcsAgentsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"rcs:agents", "update",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--configuration", "{basics: {phone_number: {label: x, number: '+49605132'}, brand_color: '#2FDCd1', description: x, email: {address: dev@stainless.com, label: x}, hero_url: https://example.com, logo_url: https://example.com, privacy_policy_url: https://example.com, terms_and_conditions_url: https://example.com, website: {label: x, url: https://example.com}}, campaign: {company_overview: x, additional_information: x, agent_overview: x, consent_settings: {call_to_action: x, double_opt_in: true, help_response: x, opt_in_message: x, opt_in_methods: [{method_type: SMS, description: x}], opt_out_response: x, call_to_action_media_url: https://example.com, call_to_action_url: https://example.com, double_opt_in_message: x}, interactions: [{interaction_type: TRANSACTIONAL_UPDATES, description: x}], message_examples: [x]}, testing: {test_url: https://example.com, additional_information: x, message_id: x}}",
			"--display-name", "Acme Delivery Updates",
			"--hosting-region", "hosting_region",
			"--profile-id", "profile_id",
			"--use-case", "MULTI_USE",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(rcsAgentsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"rcs:agents", "update",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--configuration.basics", "{phone_number: {label: x, number: '+49605132'}, brand_color: '#2FDCd1', description: x, email: {address: dev@stainless.com, label: x}, hero_url: https://example.com, logo_url: https://example.com, privacy_policy_url: https://example.com, terms_and_conditions_url: https://example.com, website: {label: x, url: https://example.com}}",
			"--configuration.campaign", "{company_overview: x, additional_information: x, agent_overview: x, consent_settings: {call_to_action: x, double_opt_in: true, help_response: x, opt_in_message: x, opt_in_methods: [{method_type: SMS, description: x}], opt_out_response: x, call_to_action_media_url: https://example.com, call_to_action_url: https://example.com, double_opt_in_message: x}, interactions: [{interaction_type: TRANSACTIONAL_UPDATES, description: x}], message_examples: [x]}",
			"--configuration.testing", "{test_url: https://example.com, additional_information: x, message_id: x}",
			"--display-name", "Acme Delivery Updates",
			"--hosting-region", "hosting_region",
			"--profile-id", "profile_id",
			"--use-case", "MULTI_USE",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"configuration:\n" +
			"  basics:\n" +
			"    phone_number:\n" +
			"      label: x\n" +
			"      number: '+49605132'\n" +
			"    brand_color: '#2FDCd1'\n" +
			"    description: x\n" +
			"    email:\n" +
			"      address: dev@stainless.com\n" +
			"      label: x\n" +
			"    hero_url: https://example.com\n" +
			"    logo_url: https://example.com\n" +
			"    privacy_policy_url: https://example.com\n" +
			"    terms_and_conditions_url: https://example.com\n" +
			"    website:\n" +
			"      label: x\n" +
			"      url: https://example.com\n" +
			"  campaign:\n" +
			"    company_overview: x\n" +
			"    additional_information: x\n" +
			"    agent_overview: x\n" +
			"    consent_settings:\n" +
			"      call_to_action: x\n" +
			"      double_opt_in: true\n" +
			"      help_response: x\n" +
			"      opt_in_message: x\n" +
			"      opt_in_methods:\n" +
			"        - method_type: SMS\n" +
			"          description: x\n" +
			"      opt_out_response: x\n" +
			"      call_to_action_media_url: https://example.com\n" +
			"      call_to_action_url: https://example.com\n" +
			"      double_opt_in_message: x\n" +
			"    interactions:\n" +
			"      - interaction_type: TRANSACTIONAL_UPDATES\n" +
			"        description: x\n" +
			"    message_examples:\n" +
			"      - x\n" +
			"  testing:\n" +
			"    test_url: https://example.com\n" +
			"    additional_information: x\n" +
			"    message_id: x\n" +
			"display_name: Acme Delivery Updates\n" +
			"hosting_region: hosting_region\n" +
			"profile_id: profile_id\n" +
			"use_case: MULTI_USE\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"rcs:agents", "update",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestRcsAgentsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"rcs:agents", "list",
			"--brand-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestRcsAgentsLaunch(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"rcs:agents", "launch",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--campaign", "{company_overview: Acme provides online retail services., additional_information: x, agent_overview: The agent sends order confirmations and delivery updates., consent_settings: {call_to_action: Select RCS updates during checkout., double_opt_in: false, help_response: Contact support@example.com for help., opt_in_message: You are subscribed to Acme order updates., opt_in_methods: [{method_type: WEBSITE, description: x}], opt_out_response: You will receive no more messages., call_to_action_media_url: https://www.example.com/rcs/opt-in.png, call_to_action_url: https://www.example.com/checkout, double_opt_in_message: x}, interactions: [{interaction_type: TRANSACTIONAL_UPDATES, description: x}], message_examples: [Your Acme order is confirmed., Your Acme order has shipped., Your Acme order was delivered.]}",
			"--testing", "{test_url: https://www.example.com/rcs/test-video, additional_information: 'Demonstrates START, STOP, HELP, and an order-status interaction.', message_id: x}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(rcsAgentsLaunch)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"rcs:agents", "launch",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--campaign", "{company_overview: Acme provides online retail services., additional_information: x, agent_overview: The agent sends order confirmations and delivery updates., consent_settings: {call_to_action: Select RCS updates during checkout., double_opt_in: false, help_response: Contact support@example.com for help., opt_in_message: You are subscribed to Acme order updates., opt_in_methods: [{method_type: WEBSITE, description: x}], opt_out_response: You will receive no more messages., call_to_action_media_url: https://www.example.com/rcs/opt-in.png, call_to_action_url: https://www.example.com/checkout, double_opt_in_message: x}, interactions: [{interaction_type: TRANSACTIONAL_UPDATES, description: x}], message_examples: [Your Acme order is confirmed., Your Acme order has shipped., Your Acme order was delivered.]}",
			"--testing.test-url", "https://www.example.com/rcs/test-video",
			"--testing.additional-information", "Demonstrates START, STOP, HELP, and an order-status interaction.",
			"--testing.message-id", "x",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"campaign:\n" +
			"  company_overview: Acme provides online retail services.\n" +
			"  additional_information: x\n" +
			"  agent_overview: The agent sends order confirmations and delivery updates.\n" +
			"  consent_settings:\n" +
			"    call_to_action: Select RCS updates during checkout.\n" +
			"    double_opt_in: false\n" +
			"    help_response: Contact support@example.com for help.\n" +
			"    opt_in_message: You are subscribed to Acme order updates.\n" +
			"    opt_in_methods:\n" +
			"      - method_type: WEBSITE\n" +
			"        description: x\n" +
			"    opt_out_response: You will receive no more messages.\n" +
			"    call_to_action_media_url: https://www.example.com/rcs/opt-in.png\n" +
			"    call_to_action_url: https://www.example.com/checkout\n" +
			"    double_opt_in_message: x\n" +
			"  interactions:\n" +
			"    - interaction_type: TRANSACTIONAL_UPDATES\n" +
			"      description: x\n" +
			"  message_examples:\n" +
			"    - Your Acme order is confirmed.\n" +
			"    - Your Acme order has shipped.\n" +
			"    - Your Acme order was delivered.\n" +
			"testing:\n" +
			"  test_url: https://www.example.com/rcs/test-video\n" +
			"  additional_information: Demonstrates START, STOP, HELP, and an order-status interaction.\n" +
			"  message_id: x\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"rcs:agents", "launch",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestRcsAgentsRetrieveCarrierApprovals(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"rcs:agents", "retrieve-carrier-approvals",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestRcsAgentsSubmit(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"rcs:agents", "submit",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
