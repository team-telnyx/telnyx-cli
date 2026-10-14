// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIMemoryNamespacesProfilesMemoriesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:memory:namespaces:profiles:memories", "retrieve",
			"--namespace", "namespace",
			"--profile-id", "profile_id",
			"--memory-id", "memory_id",
		)
	})
}

func TestAIMemoryNamespacesProfilesMemoriesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:memory:namespaces:profiles:memories", "list",
			"--max-items", "10",
			"--namespace", "namespace",
			"--profile-id", "profile_id",
			"--page-number", "1",
			"--page-size", "1",
			"--session-id", "session_id",
			"--source-id", "source_id",
		)
	})
}
