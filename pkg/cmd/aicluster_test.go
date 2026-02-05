// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIClustersRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:clusters", "retrieve",
		"--task-id", "task_id",
		"--show-subclusters=true",
		"--top-n-nodes", "0",
	)
}

func TestAIClustersList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:clusters", "list",
		"--page-number", "0",
		"--page-size", "0",
	)
}

func TestAIClustersDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:clusters", "delete",
		"--task-id", "task_id",
	)
}

func TestAIClustersCompute(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:clusters", "compute",
		"--bucket", "bucket",
		"--file", "string",
		"--min-cluster-size", "0",
		"--min-subcluster-size", "0",
		"--prefix", "prefix",
	)
}
