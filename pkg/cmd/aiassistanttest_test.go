// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestAIAssistantsTestsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:assistants:tests", "create",
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

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(aiAssistantsTestsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:assistants:tests", "create",
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
}

func TestAIAssistantsTestsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:assistants:tests", "retrieve",
		"--test-id", "test_id",
	)
}

func TestAIAssistantsTestsUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:assistants:tests", "update",
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

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(aiAssistantsTestsUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:assistants:tests", "update",
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
}

func TestAIAssistantsTestsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:assistants:tests", "list",
		"--destination", "destination",
		"--page-number", "0",
		"--page-size", "0",
		"--telnyx-conversation-channel", "telnyx_conversation_channel",
		"--test-suite", "test_suite",
	)
}

func TestAIAssistantsTestsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:assistants:tests", "delete",
		"--test-id", "test_id",
	)
}
