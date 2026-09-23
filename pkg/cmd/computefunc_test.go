// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestComputeFuncsRetrieveLogs(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"compute:funcs", "retrieve-logs",
			"--id", "id",
			"--end-time", "'2019-12-27T18:11:19.117Z'",
			"--limit", "1",
			"--start-time", "'2019-12-27T18:11:19.117Z'",
			"--type", "runtime",
		)
	})
}

func TestComputeFuncsRetrieveMetricAggregates(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"compute:funcs", "retrieve-metric-aggregates",
			"--id", "id",
			"--end-time", "'2019-12-27T18:11:19.117Z'",
			"--start-time", "'2019-12-27T18:11:19.117Z'",
			"--filter-edge-site", "filter[edge_site]",
			"--filter-namespace", "filter[namespace]",
			"--page-number", "0",
			"--page-size", "1",
		)
	})
}

func TestComputeFuncsRetrieveRevisions(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"compute:funcs", "retrieve-revisions",
			"--id", "id",
			"--page-number", "1",
			"--page-size", "1",
		)
	})
}

func TestComputeFuncsRetrieveShipInspection(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"compute:funcs", "retrieve-ship-inspection",
			"--id", "id",
		)
	})
}
