// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestAIAssistantsCanaryDeploysCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:assistants:canary-deploys", "create",
			"--assistant-id", "assistant_id",
			"--rule", "{serve: {rollout: [{version_id: Version Id, weight: 0}], version_id: Version Id}, match: [{attribute: Attribute, operator: in, values: [string]}]}",
			"--idempotency-key", "8e03978e-40d5-43e8-bc93-6894a57f9326",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(aiAssistantsCanaryDeploysCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:assistants:canary-deploys", "create",
			"--assistant-id", "assistant_id",
			"--rule.serve", "{rollout: [{version_id: Version Id, weight: 0}], version_id: Version Id}",
			"--rule.match", "[{attribute: Attribute, operator: in, values: [string]}]",
			"--idempotency-key", "8e03978e-40d5-43e8-bc93-6894a57f9326",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"rules:\n" +
			"  - serve:\n" +
			"      rollout:\n" +
			"        - version_id: Version Id\n" +
			"          weight: 0\n" +
			"      version_id: Version Id\n" +
			"    match:\n" +
			"      - attribute: Attribute\n" +
			"        operator: in\n" +
			"        values:\n" +
			"          - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:assistants:canary-deploys", "create",
			"--assistant-id", "assistant_id",
			"--idempotency-key", "8e03978e-40d5-43e8-bc93-6894a57f9326",
		)
	})
}

func TestAIAssistantsCanaryDeploysRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:assistants:canary-deploys", "retrieve",
			"--assistant-id", "assistant_id",
		)
	})
}

func TestAIAssistantsCanaryDeploysUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:assistants:canary-deploys", "update",
			"--assistant-id", "assistant_id",
			"--rule", "{serve: {rollout: [{version_id: Version Id, weight: 0}], version_id: Version Id}, match: [{attribute: Attribute, operator: in, values: [string]}]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(aiAssistantsCanaryDeploysUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:assistants:canary-deploys", "update",
			"--assistant-id", "assistant_id",
			"--rule.serve", "{rollout: [{version_id: Version Id, weight: 0}], version_id: Version Id}",
			"--rule.match", "[{attribute: Attribute, operator: in, values: [string]}]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"rules:\n" +
			"  - serve:\n" +
			"      rollout:\n" +
			"        - version_id: Version Id\n" +
			"          weight: 0\n" +
			"      version_id: Version Id\n" +
			"    match:\n" +
			"      - attribute: Attribute\n" +
			"        operator: in\n" +
			"        values:\n" +
			"          - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:assistants:canary-deploys", "update",
			"--assistant-id", "assistant_id",
		)
	})
}

func TestAIAssistantsCanaryDeploysDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:assistants:canary-deploys", "delete",
			"--assistant-id", "assistant_id",
		)
	})
}
