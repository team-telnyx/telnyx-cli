// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestAIToolsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tools", "create",
			"--display-name", "display_name",
			"--type", "type",
			"--client-side-tool", "{foo: bar}",
			"--function", "{foo: bar}",
			"--handoff", "{foo: bar}",
			"--invite", "{foo: bar}",
			"--pay", "{connector_name: connector_name, connector_parameters: {campaignid: bar, customer_ref: bar}, currency: currency, description: description, inter_digit_timeout_millis: 1, max_attempts: 1, min_postal_code_length: 5, payment_method: payment_method, postal_code: false, prompts: {bank-account-number: x, bank-routing-number: x, expiration-date: x, payment-card-number: x, postal-code: x, security-code: x}, timeout_millis: 12000, valid_card_types: [visa, mastercard]}",
			"--retrieval", "{foo: bar}",
			"--timeout-ms", "0",
			"--update-dynamic-variables", "{description: Collect caller details into conversation variables., name: collect_details, updatable_variables: [{name: customer_name, description: The caller's full name., type: string}]}",
			"--vision", "{max_frame_age_secs: 1}",
			"--webhook", "{foo: bar}",
			"--idempotency-key", "8e03978e-40d5-43e8-bc93-6894a57f9326",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(aiToolsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tools", "create",
			"--display-name", "display_name",
			"--type", "type",
			"--client-side-tool", "{foo: bar}",
			"--function", "{foo: bar}",
			"--handoff", "{foo: bar}",
			"--invite", "{foo: bar}",
			"--pay.connector-name", "connector_name",
			"--pay.connector-parameters", "{campaignid: bar, customer_ref: bar}",
			"--pay.currency", "currency",
			"--pay.description", "description",
			"--pay.inter-digit-timeout-millis", "1",
			"--pay.max-attempts", "1",
			"--pay.min-postal-code-length", "5",
			"--pay.payment-method", "payment_method",
			"--pay.postal-code=false",
			"--pay.prompts", "{bank-account-number: x, bank-routing-number: x, expiration-date: x, payment-card-number: x, postal-code: x, security-code: x}",
			"--pay.timeout-millis", "12000",
			"--pay.valid-card-types", "[visa, mastercard]",
			"--retrieval", "{foo: bar}",
			"--timeout-ms", "0",
			"--update-dynamic-variables.description", "Collect caller details into conversation variables.",
			"--update-dynamic-variables.name", "collect_details",
			"--update-dynamic-variables.updatable-variables", "[{name: customer_name, description: The caller's full name., type: string}]",
			"--vision.max-frame-age-secs", "1",
			"--webhook", "{foo: bar}",
			"--idempotency-key", "8e03978e-40d5-43e8-bc93-6894a57f9326",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"display_name: display_name\n" +
			"type: type\n" +
			"client_side_tool:\n" +
			"  foo: bar\n" +
			"function:\n" +
			"  foo: bar\n" +
			"handoff:\n" +
			"  foo: bar\n" +
			"invite:\n" +
			"  foo: bar\n" +
			"pay:\n" +
			"  connector_name: connector_name\n" +
			"  connector_parameters:\n" +
			"    campaignid: bar\n" +
			"    customer_ref: bar\n" +
			"  currency: currency\n" +
			"  description: description\n" +
			"  inter_digit_timeout_millis: 1\n" +
			"  max_attempts: 1\n" +
			"  min_postal_code_length: 5\n" +
			"  payment_method: payment_method\n" +
			"  postal_code: false\n" +
			"  prompts:\n" +
			"    bank-account-number: x\n" +
			"    bank-routing-number: x\n" +
			"    expiration-date: x\n" +
			"    payment-card-number: x\n" +
			"    postal-code: x\n" +
			"    security-code: x\n" +
			"  timeout_millis: 12000\n" +
			"  valid_card_types:\n" +
			"    - visa\n" +
			"    - mastercard\n" +
			"retrieval:\n" +
			"  foo: bar\n" +
			"timeout_ms: 0\n" +
			"update_dynamic_variables:\n" +
			"  description: Collect caller details into conversation variables.\n" +
			"  name: collect_details\n" +
			"  updatable_variables:\n" +
			"    - name: customer_name\n" +
			"      description: The caller's full name.\n" +
			"      type: string\n" +
			"vision:\n" +
			"  max_frame_age_secs: 1\n" +
			"webhook:\n" +
			"  foo: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:tools", "create",
			"--idempotency-key", "8e03978e-40d5-43e8-bc93-6894a57f9326",
		)
	})
}

func TestAIToolsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tools", "retrieve",
			"--tool-id", "tool_id",
		)
	})
}

func TestAIToolsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tools", "update",
			"--tool-id", "tool_id",
			"--client-side-tool", "{foo: bar}",
			"--display-name", "display_name",
			"--function", "{foo: bar}",
			"--handoff", "{foo: bar}",
			"--invite", "{foo: bar}",
			"--pay", "{connector_name: connector_name, connector_parameters: {campaignid: bar, customer_ref: bar}, currency: currency, description: description, inter_digit_timeout_millis: 1, max_attempts: 1, min_postal_code_length: 5, payment_method: payment_method, postal_code: false, prompts: {bank-account-number: x, bank-routing-number: x, expiration-date: x, payment-card-number: x, postal-code: x, security-code: x}, timeout_millis: 12000, valid_card_types: [visa, mastercard]}",
			"--retrieval", "{foo: bar}",
			"--timeout-ms", "0",
			"--type", "type",
			"--update-dynamic-variables", "{description: Collect caller details into conversation variables., name: collect_details, updatable_variables: [{name: customer_name, description: The caller's full name., type: string}]}",
			"--vision", "{max_frame_age_secs: 1}",
			"--webhook", "{foo: bar}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(aiToolsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tools", "update",
			"--tool-id", "tool_id",
			"--client-side-tool", "{foo: bar}",
			"--display-name", "display_name",
			"--function", "{foo: bar}",
			"--handoff", "{foo: bar}",
			"--invite", "{foo: bar}",
			"--pay.connector-name", "connector_name",
			"--pay.connector-parameters", "{campaignid: bar, customer_ref: bar}",
			"--pay.currency", "currency",
			"--pay.description", "description",
			"--pay.inter-digit-timeout-millis", "1",
			"--pay.max-attempts", "1",
			"--pay.min-postal-code-length", "5",
			"--pay.payment-method", "payment_method",
			"--pay.postal-code=false",
			"--pay.prompts", "{bank-account-number: x, bank-routing-number: x, expiration-date: x, payment-card-number: x, postal-code: x, security-code: x}",
			"--pay.timeout-millis", "12000",
			"--pay.valid-card-types", "[visa, mastercard]",
			"--retrieval", "{foo: bar}",
			"--timeout-ms", "0",
			"--type", "type",
			"--update-dynamic-variables.description", "Collect caller details into conversation variables.",
			"--update-dynamic-variables.name", "collect_details",
			"--update-dynamic-variables.updatable-variables", "[{name: customer_name, description: The caller's full name., type: string}]",
			"--vision.max-frame-age-secs", "1",
			"--webhook", "{foo: bar}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"client_side_tool:\n" +
			"  foo: bar\n" +
			"display_name: display_name\n" +
			"function:\n" +
			"  foo: bar\n" +
			"handoff:\n" +
			"  foo: bar\n" +
			"invite:\n" +
			"  foo: bar\n" +
			"pay:\n" +
			"  connector_name: connector_name\n" +
			"  connector_parameters:\n" +
			"    campaignid: bar\n" +
			"    customer_ref: bar\n" +
			"  currency: currency\n" +
			"  description: description\n" +
			"  inter_digit_timeout_millis: 1\n" +
			"  max_attempts: 1\n" +
			"  min_postal_code_length: 5\n" +
			"  payment_method: payment_method\n" +
			"  postal_code: false\n" +
			"  prompts:\n" +
			"    bank-account-number: x\n" +
			"    bank-routing-number: x\n" +
			"    expiration-date: x\n" +
			"    payment-card-number: x\n" +
			"    postal-code: x\n" +
			"    security-code: x\n" +
			"  timeout_millis: 12000\n" +
			"  valid_card_types:\n" +
			"    - visa\n" +
			"    - mastercard\n" +
			"retrieval:\n" +
			"  foo: bar\n" +
			"timeout_ms: 0\n" +
			"type: type\n" +
			"update_dynamic_variables:\n" +
			"  description: Collect caller details into conversation variables.\n" +
			"  name: collect_details\n" +
			"  updatable_variables:\n" +
			"    - name: customer_name\n" +
			"      description: The caller's full name.\n" +
			"      type: string\n" +
			"vision:\n" +
			"  max_frame_age_secs: 1\n" +
			"webhook:\n" +
			"  foo: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:tools", "update",
			"--tool-id", "tool_id",
		)
	})
}

func TestAIToolsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tools", "list",
			"--max-items", "10",
			"--filter-name", "filter[name]",
			"--filter-type", "filter[type]",
			"--page-number", "1",
			"--page-size", "1",
		)
	})
}

func TestAIToolsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tools", "delete",
			"--tool-id", "tool_id",
		)
	})
}
