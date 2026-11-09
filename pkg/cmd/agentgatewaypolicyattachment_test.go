// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAgentGatewaysPolicyAttachmentsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-gateways:policy-attachments", "create",
			"--instance-id", "32ce035d-3165-4031-b1a7-a7c0c259590a",
			"--policy", "strict",
			"--default=true",
			"--model", "GLM-5.3*",
			"--priority", "10",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"policy: strict\n" +
			"default: true\n" +
			"models:\n" +
			"  - GLM-5.3*\n" +
			"priority: 10\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agent-gateways:policy-attachments", "create",
			"--instance-id", "32ce035d-3165-4031-b1a7-a7c0c259590a",
		)
	})
}

func TestAgentGatewaysPolicyAttachmentsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-gateways:policy-attachments", "list",
			"--instance-id", "32ce035d-3165-4031-b1a7-a7c0c259590a",
		)
	})
}

func TestAgentGatewaysPolicyAttachmentsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-gateways:policy-attachments", "delete",
			"--instance-id", "32ce035d-3165-4031-b1a7-a7c0c259590a",
			"--attachment-id", "3394712f-2ad5-428a-8a3d-7a08d4b2906d",
		)
	})
}
