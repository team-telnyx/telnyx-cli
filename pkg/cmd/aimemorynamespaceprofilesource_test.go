// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIMemoryNamespacesProfilesSourcesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:memory:namespaces:profiles:sources", "retrieve",
			"--namespace", "namespace",
			"--profile-id", "profile_id",
			"--source-id", "source_id",
		)
	})
}

func TestAIMemoryNamespacesProfilesSourcesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:memory:namespaces:profiles:sources", "list",
			"--max-items", "10",
			"--namespace", "namespace",
			"--profile-id", "profile_id",
			"--page-number", "1",
			"--page-size", "1",
			"--session-id", "session_id",
		)
	})
}

func TestAIMemoryNamespacesProfilesSourcesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:memory:namespaces:profiles:sources", "delete",
			"--namespace", "namespace",
			"--profile-id", "profile_id",
			"--source-id", "source_id",
		)
	})
}
