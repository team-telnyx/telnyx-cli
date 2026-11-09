// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestAgentGatewaysPoliciesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-gateways:policies", "create",
			"--instance-id", "32ce035d-3165-4031-b1a7-a7c0c259590a",
			"--name", "baseline",
			"--base", "base",
			"--description", "Mask personal data",
			"--guardrails", "{add: [mask-pii], remove: [string]}",
			"--model", "model",
			"--pipeline", "{hook: before_request, steps: [{guardrail: guardrail, message: message, on_error: next, on_fail: next, on_pass: next, pass_data: true}]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(agentGatewaysPoliciesCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-gateways:policies", "create",
			"--instance-id", "32ce035d-3165-4031-b1a7-a7c0c259590a",
			"--name", "baseline",
			"--base", "base",
			"--description", "Mask personal data",
			"--guardrails.add", "[mask-pii]",
			"--guardrails.remove", "[string]",
			"--model", "model",
			"--pipeline.hook", "before_request",
			"--pipeline.steps", "[{guardrail: guardrail, message: message, on_error: next, on_fail: next, on_pass: next, pass_data: true}]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: baseline\n" +
			"base: base\n" +
			"description: Mask personal data\n" +
			"guardrails:\n" +
			"  add:\n" +
			"    - mask-pii\n" +
			"  remove:\n" +
			"    - string\n" +
			"model: model\n" +
			"pipeline:\n" +
			"  hook: before_request\n" +
			"  steps:\n" +
			"    - guardrail: guardrail\n" +
			"      message: message\n" +
			"      on_error: next\n" +
			"      on_fail: next\n" +
			"      on_pass: next\n" +
			"      pass_data: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agent-gateways:policies", "create",
			"--instance-id", "32ce035d-3165-4031-b1a7-a7c0c259590a",
		)
	})
}

func TestAgentGatewaysPoliciesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-gateways:policies", "retrieve",
			"--instance-id", "32ce035d-3165-4031-b1a7-a7c0c259590a",
			"--name", "strict",
		)
	})
}

func TestAgentGatewaysPoliciesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-gateways:policies", "update",
			"--instance-id", "32ce035d-3165-4031-b1a7-a7c0c259590a",
			"--name", "strict",
			"--base", "base",
			"--description", "Baseline plus a word check",
			"--guardrails", "{add: [string], remove: [string]}",
			"--model", "model",
			"--pipeline", "{hook: before_request, steps: [{guardrail: guardrail, message: message, on_error: next, on_fail: next, on_pass: next, pass_data: true}]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(agentGatewaysPoliciesUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-gateways:policies", "update",
			"--instance-id", "32ce035d-3165-4031-b1a7-a7c0c259590a",
			"--name", "strict",
			"--base", "base",
			"--description", "Baseline plus a word check",
			"--guardrails.add", "[string]",
			"--guardrails.remove", "[string]",
			"--model", "model",
			"--pipeline.hook", "before_request",
			"--pipeline.steps", "[{guardrail: guardrail, message: message, on_error: next, on_fail: next, on_pass: next, pass_data: true}]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"base: base\n" +
			"description: Baseline plus a word check\n" +
			"guardrails:\n" +
			"  add:\n" +
			"    - string\n" +
			"  remove:\n" +
			"    - string\n" +
			"model: model\n" +
			"pipeline:\n" +
			"  hook: before_request\n" +
			"  steps:\n" +
			"    - guardrail: guardrail\n" +
			"      message: message\n" +
			"      on_error: next\n" +
			"      on_fail: next\n" +
			"      on_pass: next\n" +
			"      pass_data: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agent-gateways:policies", "update",
			"--instance-id", "32ce035d-3165-4031-b1a7-a7c0c259590a",
			"--name", "strict",
		)
	})
}

func TestAgentGatewaysPoliciesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-gateways:policies", "list",
			"--instance-id", "32ce035d-3165-4031-b1a7-a7c0c259590a",
		)
	})
}

func TestAgentGatewaysPoliciesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-gateways:policies", "delete",
			"--instance-id", "32ce035d-3165-4031-b1a7-a7c0c259590a",
			"--name", "strict",
		)
	})
}
