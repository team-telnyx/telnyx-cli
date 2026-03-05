// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestAIMissionsRunsPlanCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:missions:runs:plan", "create",
			"--api-key", "string",
			"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--run-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--step", "{description: description, sequence: 0, step_id: step_id, metadata: {foo: bar}, parent_step_id: parent_step_id}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(aiMissionsRunsPlanCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "ai:missions:runs:plan", "create",
			"--api-key", "string",
			"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--run-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--step.description", "description",
			"--step.sequence", "0",
			"--step.step-id", "step_id",
			"--step.metadata", "{foo: bar}",
			"--step.parent-step-id", "parent_step_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"steps:\n" +
			"  - description: description\n" +
			"    sequence: 0\n" +
			"    step_id: step_id\n" +
			"    metadata:\n" +
			"      foo: bar\n" +
			"    parent_step_id: parent_step_id\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "ai:missions:runs:plan", "create",
			"--api-key", "string",
			"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--run-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAIMissionsRunsPlanRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:missions:runs:plan", "retrieve",
			"--api-key", "string",
			"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--run-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAIMissionsRunsPlanAddStepsToPlan(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:missions:runs:plan", "add-steps-to-plan",
			"--api-key", "string",
			"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--run-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--step", "{description: description, sequence: 0, step_id: step_id, metadata: {foo: bar}, parent_step_id: parent_step_id}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(aiMissionsRunsPlanAddStepsToPlan)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "ai:missions:runs:plan", "add-steps-to-plan",
			"--api-key", "string",
			"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--run-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--step.description", "description",
			"--step.sequence", "0",
			"--step.step-id", "step_id",
			"--step.metadata", "{foo: bar}",
			"--step.parent-step-id", "parent_step_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"steps:\n" +
			"  - description: description\n" +
			"    sequence: 0\n" +
			"    step_id: step_id\n" +
			"    metadata:\n" +
			"      foo: bar\n" +
			"    parent_step_id: parent_step_id\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "ai:missions:runs:plan", "add-steps-to-plan",
			"--api-key", "string",
			"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--run-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAIMissionsRunsPlanGetStepDetails(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:missions:runs:plan", "get-step-details",
			"--api-key", "string",
			"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--run-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--step-id", "step_id",
		)
	})
}

func TestAIMissionsRunsPlanUpdateStep(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:missions:runs:plan", "update-step",
			"--api-key", "string",
			"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--run-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--step-id", "step_id",
			"--metadata", "{foo: bar}",
			"--status", "pending",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"metadata:\n" +
			"  foo: bar\n" +
			"status: pending\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "ai:missions:runs:plan", "update-step",
			"--api-key", "string",
			"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--run-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--step-id", "step_id",
		)
	})
}
