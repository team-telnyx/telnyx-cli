// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIToolsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tools", "create",
			"--display-name", "display_name",
			"--type", "type",
			"--function", "{foo: bar}",
			"--handoff", "{foo: bar}",
			"--invite", "{foo: bar}",
			"--retrieval", "{foo: bar}",
			"--timeout-ms", "0",
			"--webhook", "{foo: bar}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"display_name: display_name\n" +
			"type: type\n" +
			"function:\n" +
			"  foo: bar\n" +
			"handoff:\n" +
			"  foo: bar\n" +
			"invite:\n" +
			"  foo: bar\n" +
			"retrieval:\n" +
			"  foo: bar\n" +
			"timeout_ms: 0\n" +
			"webhook:\n" +
			"  foo: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:tools", "create",
		)
	})
}

func TestAIToolsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tools", "retrieve",
			"--tool-id", "tool_id",
		)
	})
}

func TestAIToolsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tools", "update",
			"--tool-id", "tool_id",
			"--display-name", "display_name",
			"--function", "{foo: bar}",
			"--handoff", "{foo: bar}",
			"--invite", "{foo: bar}",
			"--retrieval", "{foo: bar}",
			"--timeout-ms", "0",
			"--type", "type",
			"--webhook", "{foo: bar}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"display_name: display_name\n" +
			"function:\n" +
			"  foo: bar\n" +
			"handoff:\n" +
			"  foo: bar\n" +
			"invite:\n" +
			"  foo: bar\n" +
			"retrieval:\n" +
			"  foo: bar\n" +
			"timeout_ms: 0\n" +
			"type: type\n" +
			"webhook:\n" +
			"  foo: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:tools", "update",
			"--tool-id", "tool_id",
		)
	})
}

func TestAIToolsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tools", "list",
			"--max-items", "10",
			"--filter-name", "filter[name]",
			"--filter-type", "filter[type]",
			"--page-number", "1",
			"--page-size", "1",
		)
	})
}

func TestAIToolsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tools", "delete",
			"--tool-id", "tool_id",
		)
	})
}
