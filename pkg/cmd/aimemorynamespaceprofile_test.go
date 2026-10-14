// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIMemoryNamespacesProfilesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:memory:namespaces:profiles", "list",
			"--max-items", "10",
			"--namespace", "namespace",
			"--page-number", "1",
			"--page-size", "1",
		)
	})
}

func TestAIMemoryNamespacesProfilesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:memory:namespaces:profiles", "delete",
			"--namespace", "namespace",
			"--profile-id", "profile_id",
		)
	})
}

func TestAIMemoryNamespacesProfilesIngest(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:memory:namespaces:profiles", "ingest",
			"--namespace", "namespace",
			"--profile-id", "profile_id",
			"--body", "{foo: bar}",
			"--session-id", "session_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("foo: bar")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:memory:namespaces:profiles", "ingest",
			"--namespace", "namespace",
			"--profile-id", "profile_id",
			"--session-id", "session_id",
		)
	})
}

func TestAIMemoryNamespacesProfilesPreview(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:memory:namespaces:profiles", "preview",
			"--namespace", "namespace",
			"--profile-id", "profile_id",
			"--instructions", "Lead with billing status. Prefer plain language over jargon.",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("instructions: Lead with billing status. Prefer plain language over jargon.")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:memory:namespaces:profiles", "preview",
			"--namespace", "namespace",
			"--profile-id", "profile_id",
		)
	})
}

func TestAIMemoryNamespacesProfilesRecall(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:memory:namespaces:profiles", "recall",
			"--namespace", "namespace",
			"--profile-id", "profile_id",
			"--query", "where do invoices go?",
			"--top-k", "5",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"query: where do invoices go?\n" +
			"top_k: 5\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:memory:namespaces:profiles", "recall",
			"--namespace", "namespace",
			"--profile-id", "profile_id",
		)
	})
}

func TestAIMemoryNamespacesProfilesRemember(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:memory:namespaces:profiles", "remember",
			"--namespace", "namespace",
			"--profile-id", "profile_id",
			"--text", "Prefers window seats and flies out of ORD",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("text: Prefers window seats and flies out of ORD")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:memory:namespaces:profiles", "remember",
			"--namespace", "namespace",
			"--profile-id", "profile_id",
		)
	})
}

func TestAIMemoryNamespacesProfilesRetrieveSummary(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:memory:namespaces:profiles", "retrieve-summary",
			"--namespace", "namespace",
			"--profile-id", "profile_id",
		)
	})
}
