// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestAIFineTuningJobsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:fine-tuning:jobs", "create",
			"--api-key", "string",
			"--model", "model",
			"--training-file", "training_file",
			"--hyperparameters", "{n_epochs: 1}",
			"--suffix", "suffix",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(aiFineTuningJobsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "ai:fine-tuning:jobs", "create",
			"--api-key", "string",
			"--model", "model",
			"--training-file", "training_file",
			"--hyperparameters.n-epochs", "1",
			"--suffix", "suffix",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"model: model\n" +
			"training_file: training_file\n" +
			"hyperparameters:\n" +
			"  n_epochs: 1\n" +
			"suffix: suffix\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "ai:fine-tuning:jobs", "create",
			"--api-key", "string",
		)
	})
}

func TestAIFineTuningJobsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:fine-tuning:jobs", "retrieve",
			"--api-key", "string",
			"--job-id", "job_id",
		)
	})
}

func TestAIFineTuningJobsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:fine-tuning:jobs", "list",
			"--api-key", "string",
		)
	})
}

func TestAIFineTuningJobsCancel(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:fine-tuning:jobs", "cancel",
			"--api-key", "string",
			"--job-id", "job_id",
		)
	})
}
