// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIMissionsRunsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:missions:runs", "create",
			"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--input", "{foo: bar}",
			"--metadata", "{foo: bar}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"input:\n" +
			"  foo: bar\n" +
			"metadata:\n" +
			"  foo: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:missions:runs", "create",
			"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAIMissionsRunsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:missions:runs", "retrieve",
			"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--run-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAIMissionsRunsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:missions:runs", "update",
			"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--run-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--error", "error",
			"--metadata", "{foo: bar}",
			"--result-payload", "{foo: bar}",
			"--result-summary", "result_summary",
			"--status", "pending",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"error: error\n" +
			"metadata:\n" +
			"  foo: bar\n" +
			"result_payload:\n" +
			"  foo: bar\n" +
			"result_summary: result_summary\n" +
			"status: pending\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:missions:runs", "update",
			"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--run-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAIMissionsRunsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:missions:runs", "list",
			"--max-items", "10",
			"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--page-number", "1",
			"--page-size", "1",
			"--status", "status",
		)
	})
}

func TestAIMissionsRunsCancelRun(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:missions:runs", "cancel-run",
			"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--run-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAIMissionsRunsListRuns(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:missions:runs", "list-runs",
			"--max-items", "10",
			"--page-number", "1",
			"--page-size", "1",
			"--status", "status",
		)
	})
}

func TestAIMissionsRunsPauseRun(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:missions:runs", "pause-run",
			"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--run-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAIMissionsRunsResumeRun(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:missions:runs", "resume-run",
			"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--run-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
