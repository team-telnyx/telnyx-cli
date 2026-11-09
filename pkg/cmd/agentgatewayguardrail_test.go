// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAgentGatewaysGuardrailsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-gateways:guardrails", "create",
			"--instance-id", "32ce035d-3165-4031-b1a7-a7c0c259590a",
			"--mode", "before_request",
			"--name", "mask-pii",
			"--type", "content_filter",
			"--always-on=true",
			"--settings", "{patterns: bar}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"modes:\n" +
			"  - before_request\n" +
			"name: mask-pii\n" +
			"type: content_filter\n" +
			"always_on: true\n" +
			"settings:\n" +
			"  patterns: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agent-gateways:guardrails", "create",
			"--instance-id", "32ce035d-3165-4031-b1a7-a7c0c259590a",
		)
	})
}

func TestAgentGatewaysGuardrailsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-gateways:guardrails", "update",
			"--instance-id", "32ce035d-3165-4031-b1a7-a7c0c259590a",
			"--guardrail-id", "9645e556-6532-4341-8cab-21b58cdef529",
			"--always-on=true",
			"--mode", "before_request",
			"--name", "name",
			"--settings", "{foo: bar}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"always_on: true\n" +
			"modes:\n" +
			"  - before_request\n" +
			"name: name\n" +
			"settings:\n" +
			"  foo: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agent-gateways:guardrails", "update",
			"--instance-id", "32ce035d-3165-4031-b1a7-a7c0c259590a",
			"--guardrail-id", "9645e556-6532-4341-8cab-21b58cdef529",
		)
	})
}

func TestAgentGatewaysGuardrailsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-gateways:guardrails", "list",
			"--instance-id", "32ce035d-3165-4031-b1a7-a7c0c259590a",
		)
	})
}

func TestAgentGatewaysGuardrailsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-gateways:guardrails", "delete",
			"--instance-id", "32ce035d-3165-4031-b1a7-a7c0c259590a",
			"--guardrail-id", "9645e556-6532-4341-8cab-21b58cdef529",
		)
	})
}
