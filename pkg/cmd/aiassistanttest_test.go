// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestAIAssistantsTestsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:assistants:tests", "create",
			"--api-key", "string",
			"--destination", "+15551234567",
			"--instructions", "Act as a frustrated customer who received a damaged product. Ask for a refund and escalate if not satisfied with the initial response.",
			"--name", "Customer Support Bot Test",
			"--rubric", "{criteria: Assistant responds within 30 seconds, name: Response Time}",
			"--rubric", "{criteria: Provides correct product information, name: Accuracy}",
			"--description", "description",
			"--max-duration-seconds", "1",
			"--telnyx-conversation-channel", "web_chat",
			"--test-suite", "test_suite",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(aiAssistantsTestsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "ai:assistants:tests", "create",
			"--api-key", "string",
			"--destination", "+15551234567",
			"--instructions", "Act as a frustrated customer who received a damaged product. Ask for a refund and escalate if not satisfied with the initial response.",
			"--name", "Customer Support Bot Test",
			"--rubric.criteria", "Assistant responds within 30 seconds",
			"--rubric.name", "Response Time",
			"--rubric.criteria", "Provides correct product information",
			"--rubric.name", "Accuracy",
			"--description", "description",
			"--max-duration-seconds", "1",
			"--telnyx-conversation-channel", "web_chat",
			"--test-suite", "test_suite",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"destination: '+15551234567'\n" +
			"instructions: >-\n" +
			"  Act as a frustrated customer who received a damaged product. Ask for a refund\n" +
			"  and escalate if not satisfied with the initial response.\n" +
			"name: Customer Support Bot Test\n" +
			"rubric:\n" +
			"  - criteria: Assistant responds within 30 seconds\n" +
			"    name: Response Time\n" +
			"  - criteria: Provides correct product information\n" +
			"    name: Accuracy\n" +
			"description: description\n" +
			"max_duration_seconds: 1\n" +
			"telnyx_conversation_channel: web_chat\n" +
			"test_suite: test_suite\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "ai:assistants:tests", "create",
			"--api-key", "string",
		)
	})
}

func TestAIAssistantsTestsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:assistants:tests", "retrieve",
			"--api-key", "string",
			"--test-id", "test_id",
		)
	})
}

func TestAIAssistantsTestsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:assistants:tests", "update",
			"--api-key", "string",
			"--test-id", "test_id",
			"--description", "description",
			"--destination", "x",
			"--instructions", "x",
			"--max-duration-seconds", "1",
			"--name", "x",
			"--rubric", "{criteria: criteria, name: name}",
			"--telnyx-conversation-channel", "phone_call",
			"--test-suite", "test_suite",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(aiAssistantsTestsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "ai:assistants:tests", "update",
			"--api-key", "string",
			"--test-id", "test_id",
			"--description", "description",
			"--destination", "x",
			"--instructions", "x",
			"--max-duration-seconds", "1",
			"--name", "x",
			"--rubric.criteria", "criteria",
			"--rubric.name", "name",
			"--telnyx-conversation-channel", "phone_call",
			"--test-suite", "test_suite",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"description: description\n" +
			"destination: x\n" +
			"instructions: x\n" +
			"max_duration_seconds: 1\n" +
			"name: x\n" +
			"rubric:\n" +
			"  - criteria: criteria\n" +
			"    name: name\n" +
			"telnyx_conversation_channel: phone_call\n" +
			"test_suite: test_suite\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "ai:assistants:tests", "update",
			"--api-key", "string",
			"--test-id", "test_id",
		)
	})
}

func TestAIAssistantsTestsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:assistants:tests", "list",
			"--api-key", "string",
			"--destination", "destination",
			"--page-number", "0",
			"--page-size", "0",
			"--telnyx-conversation-channel", "telnyx_conversation_channel",
			"--test-suite", "test_suite",
		)
	})
}

func TestAIAssistantsTestsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:assistants:tests", "delete",
			"--api-key", "string",
			"--test-id", "test_id",
		)
	})
}
