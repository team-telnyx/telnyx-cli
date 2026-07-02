// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"strings"
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestStorageBucketsSslCertificateCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"storage:buckets:ssl-certificate", "create",
			"--bucket-name", "",
			"--certificate", mocktest.TestFile(t, "Example data"),
			"--private-key", mocktest.TestFile(t, "Example data"),
		)
	})

	t.Run("piping data", func(t *testing.T) {
		testFile := mocktest.TestFile(t, "Example data")
		// Test piping YAML data over stdin
		pipeDataStr := "" +
			"certificate: Example data\n" +
			"private_key: Example data\n"
		pipeDataStr = strings.ReplaceAll(pipeDataStr, "Example data", testFile)
		pipeData := []byte(pipeDataStr)
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"storage:buckets:ssl-certificate", "create",
			"--bucket-name", "",
		)
	})
}

func TestStorageBucketsSslCertificateRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"storage:buckets:ssl-certificate", "retrieve",
			"--bucket-name", "",
		)
	})
}

func TestStorageBucketsSslCertificateDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"storage:buckets:ssl-certificate", "delete",
			"--bucket-name", "",
		)
	})
}
