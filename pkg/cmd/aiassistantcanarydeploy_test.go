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
			t, "ai:assistants:canary-deploys", "create",
			"--api-key", "string",
			"--assistant-id", "assistant_id",
			"--version", "{percentage: 1, version_id: version_id}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(aiAssistantsCanaryDeploysCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "ai:assistants:canary-deploys", "create",
			"--api-key", "string",
			"--assistant-id", "assistant_id",
			"--version.percentage", "1",
			"--version.version-id", "version_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"versions:\n" +
			"  - percentage: 1\n" +
			"    version_id: version_id\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "ai:assistants:canary-deploys", "create",
			"--api-key", "string",
			"--assistant-id", "assistant_id",
		)
	})
}

func TestAIAssistantsCanaryDeploysRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:assistants:canary-deploys", "retrieve",
			"--api-key", "string",
			"--assistant-id", "assistant_id",
		)
	})
}

func TestAIAssistantsCanaryDeploysUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:assistants:canary-deploys", "update",
			"--api-key", "string",
			"--assistant-id", "assistant_id",
			"--version", "{percentage: 1, version_id: version_id}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(aiAssistantsCanaryDeploysUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "ai:assistants:canary-deploys", "update",
			"--api-key", "string",
			"--assistant-id", "assistant_id",
			"--version.percentage", "1",
			"--version.version-id", "version_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"versions:\n" +
			"  - percentage: 1\n" +
			"    version_id: version_id\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "ai:assistants:canary-deploys", "update",
			"--api-key", "string",
			"--assistant-id", "assistant_id",
		)
	})
}

func TestAIAssistantsCanaryDeploysDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:assistants:canary-deploys", "delete",
			"--api-key", "string",
			"--assistant-id", "assistant_id",
		)
	})
}
