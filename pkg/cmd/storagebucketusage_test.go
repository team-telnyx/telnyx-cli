// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestStorageBucketsUsageGetAPIUsage(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"storage:buckets:usage", "get-api-usage",
		"--bucket-name", "",
		"--filter", "{end_time: '2019-12-27T18:11:19.117Z', start_time: '2019-12-27T18:11:19.117Z'}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(storageBucketsUsageGetAPIUsage)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"storage:buckets:usage", "get-api-usage",
		"--bucket-name", "",
		"--filter.end-time", "2019-12-27T18:11:19.117Z",
		"--filter.start-time", "2019-12-27T18:11:19.117Z",
	)
}

func TestStorageBucketsUsageGetBucketUsage(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"storage:buckets:usage", "get-bucket-usage",
		"--bucket-name", "",
	)
}
