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
			t,
			"--api-key", "string",
			"ai:fine-tuning:jobs", "create",
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
			t,
			"--api-key", "string",
			"ai:fine-tuning:jobs", "create",
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
			t, pipeData,
			"--api-key", "string",
			"ai:fine-tuning:jobs", "create",
		)
	})
}

func TestAIFineTuningJobsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:fine-tuning:jobs", "retrieve",
			"--job-id", "job_id",
		)
	})
}

func TestAIFineTuningJobsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:fine-tuning:jobs", "list",
		)
	})
}

func TestAIFineTuningJobsCancel(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:fine-tuning:jobs", "cancel",
			"--job-id", "job_id",
		)
	})
}
