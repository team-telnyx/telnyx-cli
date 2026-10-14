// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestLlmTokenGatewayTokenGroupsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"llm-token-gateway:token-groups", "create",
			"--allowed-model", "Kimi-K3",
			"--allowed-model", "GLM-5.3-Flash",
			"--name", "support-backend",
			"--idempotency-key", "Idempotency-Key",
			"--blocked=true",
			"--budget-duration", "1d",
			"--cache-ttl", "1",
			"--fallbacks", "{Kimi-K3: [GLM-5.3-Flash]}",
			"--guardrails", "{dlp: {profiles: [financial], prompt: ignore, response: ignore}, secrets: {prompt: ignore, response: ignore}, streaming: buffered}",
			"--max-budget", "0",
			"--max-parallel-requests", "20",
			"--provider-key-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--retries", "{max_attempts: 2, backoff: exponential, request_timeout_ms: 10000, retry_delay_ms: 250}",
			"--rpm-limit", "0",
			"--tpm-limit", "0",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(llmTokenGatewayTokenGroupsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"llm-token-gateway:token-groups", "create",
			"--allowed-model", "Kimi-K3",
			"--allowed-model", "GLM-5.3-Flash",
			"--name", "support-backend",
			"--idempotency-key", "Idempotency-Key",
			"--blocked=true",
			"--budget-duration", "1d",
			"--cache-ttl", "1",
			"--fallbacks", "{Kimi-K3: [GLM-5.3-Flash]}",
			"--guardrails.dlp", "{profiles: [financial], prompt: ignore, response: ignore}",
			"--guardrails.secrets", "{prompt: ignore, response: ignore}",
			"--guardrails.streaming", "buffered",
			"--max-budget", "0",
			"--max-parallel-requests", "20",
			"--provider-key-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--retries.max-attempts", "2",
			"--retries.backoff", "exponential",
			"--retries.request-timeout-ms", "10000",
			"--retries.retry-delay-ms", "250",
			"--rpm-limit", "0",
			"--tpm-limit", "0",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"allowed_models:\n" +
			"  - Kimi-K3\n" +
			"  - GLM-5.3-Flash\n" +
			"name: support-backend\n" +
			"blocked: true\n" +
			"budget_duration: 1d\n" +
			"cache_ttl: 1\n" +
			"fallbacks:\n" +
			"  Kimi-K3:\n" +
			"    - GLM-5.3-Flash\n" +
			"guardrails:\n" +
			"  dlp:\n" +
			"    profiles:\n" +
			"      - financial\n" +
			"    prompt: ignore\n" +
			"    response: ignore\n" +
			"  secrets:\n" +
			"    prompt: ignore\n" +
			"    response: ignore\n" +
			"  streaming: buffered\n" +
			"max_budget: 0\n" +
			"max_parallel_requests: 20\n" +
			"provider_key_ids:\n" +
			"  - 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"retries:\n" +
			"  max_attempts: 2\n" +
			"  backoff: exponential\n" +
			"  request_timeout_ms: 10000\n" +
			"  retry_delay_ms: 250\n" +
			"rpm_limit: 0\n" +
			"tpm_limit: 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"llm-token-gateway:token-groups", "create",
			"--idempotency-key", "Idempotency-Key",
		)
	})
}

func TestLlmTokenGatewayTokenGroupsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"llm-token-gateway:token-groups", "retrieve",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestLlmTokenGatewayTokenGroupsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"llm-token-gateway:token-groups", "update",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--idempotency-key", "Idempotency-Key",
			"--if-match", "If-Match",
			"--allowed-model", "Kimi-K3",
			"--allowed-model", "GLM-5.3-Flash",
			"--blocked=true",
			"--budget-duration", "1d",
			"--cache-ttl", "1",
			"--fallbacks", "{Kimi-K3: [GLM-5.3-Flash]}",
			"--guardrails", "{dlp: {profiles: [financial], prompt: ignore, response: ignore}, secrets: {prompt: ignore, response: ignore}, streaming: buffered}",
			"--max-budget", "0",
			"--max-parallel-requests", "20",
			"--name", "support-backend",
			"--provider-key-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--retries", "{max_attempts: 2, backoff: exponential, request_timeout_ms: 10000, retry_delay_ms: 250}",
			"--rpm-limit", "0",
			"--tpm-limit", "0",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(llmTokenGatewayTokenGroupsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"llm-token-gateway:token-groups", "update",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--idempotency-key", "Idempotency-Key",
			"--if-match", "If-Match",
			"--allowed-model", "Kimi-K3",
			"--allowed-model", "GLM-5.3-Flash",
			"--blocked=true",
			"--budget-duration", "1d",
			"--cache-ttl", "1",
			"--fallbacks", "{Kimi-K3: [GLM-5.3-Flash]}",
			"--guardrails.dlp", "{profiles: [financial], prompt: ignore, response: ignore}",
			"--guardrails.secrets", "{prompt: ignore, response: ignore}",
			"--guardrails.streaming", "buffered",
			"--max-budget", "0",
			"--max-parallel-requests", "20",
			"--name", "support-backend",
			"--provider-key-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--retries.max-attempts", "2",
			"--retries.backoff", "exponential",
			"--retries.request-timeout-ms", "10000",
			"--retries.retry-delay-ms", "250",
			"--rpm-limit", "0",
			"--tpm-limit", "0",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"allowed_models:\n" +
			"  - Kimi-K3\n" +
			"  - GLM-5.3-Flash\n" +
			"blocked: true\n" +
			"budget_duration: 1d\n" +
			"cache_ttl: 1\n" +
			"fallbacks:\n" +
			"  Kimi-K3:\n" +
			"    - GLM-5.3-Flash\n" +
			"guardrails:\n" +
			"  dlp:\n" +
			"    profiles:\n" +
			"      - financial\n" +
			"    prompt: ignore\n" +
			"    response: ignore\n" +
			"  secrets:\n" +
			"    prompt: ignore\n" +
			"    response: ignore\n" +
			"  streaming: buffered\n" +
			"max_budget: 0\n" +
			"max_parallel_requests: 20\n" +
			"name: support-backend\n" +
			"provider_key_ids:\n" +
			"  - 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"retries:\n" +
			"  max_attempts: 2\n" +
			"  backoff: exponential\n" +
			"  request_timeout_ms: 10000\n" +
			"  retry_delay_ms: 250\n" +
			"rpm_limit: 0\n" +
			"tpm_limit: 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"llm-token-gateway:token-groups", "update",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--idempotency-key", "Idempotency-Key",
			"--if-match", "If-Match",
		)
	})
}

func TestLlmTokenGatewayTokenGroupsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"llm-token-gateway:token-groups", "list",
			"--max-items", "10",
			"--page-number", "1",
			"--page-size", "1",
			"--page-snapshot", "page[snapshot]",
		)
	})
}

func TestLlmTokenGatewayTokenGroupsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"llm-token-gateway:token-groups", "delete",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--idempotency-key", "Idempotency-Key",
			"--if-match", "If-Match",
		)
	})
}
