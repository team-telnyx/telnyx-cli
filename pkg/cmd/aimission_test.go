// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIMissionsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:missions", "create",
			"--name", "name",
			"--description", "description",
			"--execution-mode", "external",
			"--instructions", "instructions",
			"--metadata", "{foo: bar}",
			"--model", "model",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: name\n" +
			"description: description\n" +
			"execution_mode: external\n" +
			"instructions: instructions\n" +
			"metadata:\n" +
			"  foo: bar\n" +
			"model: model\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:missions", "create",
		)
	})
}

func TestAIMissionsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:missions", "retrieve",
			"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAIMissionsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:missions", "list",
			"--max-items", "10",
			"--page-number", "1",
			"--page-size", "1",
		)
	})
}

func TestAIMissionsCloneMission(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:missions", "clone-mission",
			"--mission-id", "mission_id",
		)
	})
}

func TestAIMissionsDeleteMission(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:missions", "delete-mission",
			"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAIMissionsListEvents(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:missions", "list-events",
			"--max-items", "10",
			"--page-number", "1",
			"--page-size", "1",
			"--type", "type",
		)
	})
}

func TestAIMissionsUpdateMission(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:missions", "update-mission",
			"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--description", "description",
			"--execution-mode", "external",
			"--instructions", "instructions",
			"--metadata", "{foo: bar}",
			"--model", "model",
			"--name", "name",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"description: description\n" +
			"execution_mode: external\n" +
			"instructions: instructions\n" +
			"metadata:\n" +
			"  foo: bar\n" +
			"model: model\n" +
			"name: name\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:missions", "update-mission",
			"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
