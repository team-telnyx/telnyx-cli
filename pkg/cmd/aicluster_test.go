// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIClustersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:clusters", "retrieve",
			"--task-id", "task_id",
			"--show-subclusters=true",
			"--top-n-nodes", "0",
		)
	})
}

func TestAIClustersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:clusters", "list",
			"--max-items", "10",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

func TestAIClustersDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:clusters", "delete",
			"--task-id", "task_id",
		)
	})
}

func TestAIClustersCompute(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:clusters", "compute",
			"--bucket", "bucket",
			"--file", "string",
			"--min-cluster-size", "0",
			"--min-subcluster-size", "0",
			"--prefix", "prefix",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"bucket: bucket\n" +
			"files:\n" +
			"  - string\n" +
			"min_cluster_size: 0\n" +
			"min_subcluster_size: 0\n" +
			"prefix: prefix\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:clusters", "compute",
		)
	})
}

func TestAIClustersFetchGraph(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:clusters", "fetch-graph",
			"--task-id", "task_id",
			"--cluster-id", "0",
			"--output", "/dev/null",
		)
	})
}
