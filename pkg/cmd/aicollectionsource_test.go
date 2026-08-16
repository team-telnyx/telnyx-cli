// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestAICollectionsSourcesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:collections:sources", "create",
			"--uuid", "6a09ccbd-8f9b-4c3a-9b0e-2f1d3c4b5a6e",
			"--source-type", "voice",
			"--bucket-id", "policy-docs",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"source_type: voice\n" +
			"bucket_id: policy-docs\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:collections:sources", "create",
			"--uuid", "6a09ccbd-8f9b-4c3a-9b0e-2f1d3c4b5a6e",
		)
	})
}

func TestAICollectionsSourcesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:collections:sources", "list",
			"--uuid", "6a09ccbd-8f9b-4c3a-9b0e-2f1d3c4b5a6e",
		)
	})
}

func TestAICollectionsSourcesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:collections:sources", "delete",
			"--uuid", "6a09ccbd-8f9b-4c3a-9b0e-2f1d3c4b5a6e",
			"--source-id", "42",
		)
	})
}

func TestAICollectionsSourcesReplace(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:collections:sources", "replace",
			"--uuid", "6a09ccbd-8f9b-4c3a-9b0e-2f1d3c4b5a6e",
			"--source", "{source_type: voice, bucket_id: policy-docs}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(aiCollectionsSourcesReplace)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:collections:sources", "replace",
			"--uuid", "6a09ccbd-8f9b-4c3a-9b0e-2f1d3c4b5a6e",
			"--source.source-type", "voice",
			"--source.bucket-id", "policy-docs",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"sources:\n" +
			"  - source_type: voice\n" +
			"    bucket_id: policy-docs\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:collections:sources", "replace",
			"--uuid", "6a09ccbd-8f9b-4c3a-9b0e-2f1d3c4b5a6e",
		)
	})
}
