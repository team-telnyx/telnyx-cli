// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestCustomStorageCredentialsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"custom-storage-credentials", "create",
		"--connection-id", "connection_id",
		"--backend", "gcs",
		"--configuration", "{backend: gcs, bucket: example-bucket, credentials: OPAQUE_CREDENTIALS_TOKEN}",
	)
}

func TestCustomStorageCredentialsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"custom-storage-credentials", "retrieve",
		"--connection-id", "connection_id",
	)
}

func TestCustomStorageCredentialsUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"custom-storage-credentials", "update",
		"--connection-id", "connection_id",
		"--backend", "gcs",
		"--configuration", "{backend: gcs, bucket: example-bucket, credentials: OPAQUE_CREDENTIALS_TOKEN}",
	)
}

func TestCustomStorageCredentialsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"custom-storage-credentials", "delete",
		"--connection-id", "connection_id",
	)
}
