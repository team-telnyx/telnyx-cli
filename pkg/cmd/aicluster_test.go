// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIClustersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:clusters", "retrieve",
		"--api-key", "string",
		"--task-id", "task_id",
		"--show-subclusters=true",
		"--top-n-nodes", "0",
	)
}

func TestAIClustersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:clusters", "list",
		"--api-key", "string",
		"--page-number", "0",
		"--page-size", "0",
	)
}

func TestAIClustersDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:clusters", "delete",
		"--api-key", "string",
		"--task-id", "task_id",
	)
}

func TestAIClustersCompute(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:clusters", "compute",
		"--api-key", "string",
		"--bucket", "bucket",
		"--file", "string",
		"--min-cluster-size", "0",
		"--min-subcluster-size", "0",
		"--prefix", "prefix",
	)
}

func TestAIClustersFetchGraph(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:clusters", "fetch-graph",
		"--api-key", "string",
		"--task-id", "task_id",
		"--cluster-id", "0",
		"--output", "/dev/null",
	)
}
