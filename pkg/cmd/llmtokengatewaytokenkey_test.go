// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestLlmTokenGatewayTokenKeysCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"llm-token-gateway:token-keys", "create",
			"--name", "support-service",
			"--token-group-id", "11111111-1111-4111-8111-111111111111",
			"--idempotency-key", "Idempotency-Key",
			"--allowed-model", "[J!Q0Ok0bzJb7]",
			"--blocked=true",
			"--budget-duration", "1d",
			"--expires-at", "'2019-12-27T18:11:19.117Z'",
			"--max-budget", "0.000001",
			"--max-parallel-requests", "5",
			"--required-end-user-id=true",
			"--rpm-limit", "1",
			"--token-user-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--tpm-limit", "1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: support-service\n" +
			"token_group_id: 11111111-1111-4111-8111-111111111111\n" +
			"allowed_models:\n" +
			"  - J!Q0Ok0bzJb7\n" +
			"blocked: true\n" +
			"budget_duration: 1d\n" +
			"expires_at: '2019-12-27T18:11:19.117Z'\n" +
			"max_budget: 0.000001\n" +
			"max_parallel_requests: 5\n" +
			"required_end_user_id: true\n" +
			"rpm_limit: 1\n" +
			"token_user_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"tpm_limit: 1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"llm-token-gateway:token-keys", "create",
			"--idempotency-key", "Idempotency-Key",
		)
	})
}

func TestLlmTokenGatewayTokenKeysRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"llm-token-gateway:token-keys", "retrieve",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestLlmTokenGatewayTokenKeysUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"llm-token-gateway:token-keys", "update",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--idempotency-key", "Idempotency-Key",
			"--if-match", "If-Match",
			"--allowed-model", "[J!Q0Ok0bzJb7]",
			"--blocked=true",
			"--budget-duration", "1d",
			"--expires-at", "'2019-12-27T18:11:19.117Z'",
			"--max-budget", "0.000001",
			"--max-parallel-requests", "5",
			"--name", "name",
			"--required-end-user-id=true",
			"--rpm-limit", "1",
			"--tpm-limit", "1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"allowed_models:\n" +
			"  - J!Q0Ok0bzJb7\n" +
			"blocked: true\n" +
			"budget_duration: 1d\n" +
			"expires_at: '2019-12-27T18:11:19.117Z'\n" +
			"max_budget: 0.000001\n" +
			"max_parallel_requests: 5\n" +
			"name: name\n" +
			"required_end_user_id: true\n" +
			"rpm_limit: 1\n" +
			"tpm_limit: 1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"llm-token-gateway:token-keys", "update",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--idempotency-key", "Idempotency-Key",
			"--if-match", "If-Match",
		)
	})
}

func TestLlmTokenGatewayTokenKeysList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"llm-token-gateway:token-keys", "list",
			"--max-items", "10",
			"--page-number", "1",
			"--page-size", "1",
			"--page-snapshot", "page[snapshot]",
			"--token-group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--token-user-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestLlmTokenGatewayTokenKeysDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"llm-token-gateway:token-keys", "delete",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--idempotency-key", "Idempotency-Key",
			"--if-match", "If-Match",
		)
	})
}
