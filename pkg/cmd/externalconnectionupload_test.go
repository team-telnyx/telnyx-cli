// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestExternalConnectionsUploadsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "external-connections:uploads", "create",
			"--api-key", "string",
			"--id", "1293384261075731499",
			"--number-id", "3920457616934164700",
			"--number-id", "3920457616934164701",
			"--number-id", "3920457616934164702",
			"--number-id", "3920457616934164703",
			"--additional-usage", "calling_user_assignment",
			"--civic-address-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--location-id", "67ea7693-9cd5-4a68-8c76-abb3aa5bf5d2",
			"--usage", "first_party_app_assignment",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"number_ids:\n" +
			"  - '3920457616934164700'\n" +
			"  - '3920457616934164701'\n" +
			"  - '3920457616934164702'\n" +
			"  - '3920457616934164703'\n" +
			"additional_usages:\n" +
			"  - calling_user_assignment\n" +
			"civic_address_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"location_id: 67ea7693-9cd5-4a68-8c76-abb3aa5bf5d2\n" +
			"usage: first_party_app_assignment\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "external-connections:uploads", "create",
			"--api-key", "string",
			"--id", "1293384261075731499",
		)
	})
}

func TestExternalConnectionsUploadsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "external-connections:uploads", "retrieve",
			"--api-key", "string",
			"--id", "1293384261075731499",
			"--ticket-id", "7b6a6449-b055-45a6-81f6-f6f0dffa4cc6",
		)
	})
}

func TestExternalConnectionsUploadsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "external-connections:uploads", "list",
			"--api-key", "string",
			"--id", "1293384261075731499",
			"--filter", "{civic_address_id: {eq: '19990261512338516954'}, location_id: {eq: '19995665508264022121'}, phone_number: {contains: '+1970', eq: '+19705555098'}, status: {eq: [pending_upload, pending]}}",
			"--page-number", "0",
			"--page-size", "0",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(externalConnectionsUploadsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "external-connections:uploads", "list",
			"--api-key", "string",
			"--id", "1293384261075731499",
			"--filter.civic-address-id", "{eq: '19990261512338516954'}",
			"--filter.location-id", "{eq: '19995665508264022121'}",
			"--filter.phone-number", "{contains: '+1970', eq: '+19705555098'}",
			"--filter.status", "{eq: [pending_upload, pending]}",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

func TestExternalConnectionsUploadsPendingCount(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "external-connections:uploads", "pending-count",
			"--api-key", "string",
			"--id", "1293384261075731499",
		)
	})
}

func TestExternalConnectionsUploadsRefreshStatus(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "external-connections:uploads", "refresh-status",
			"--api-key", "string",
			"--id", "1293384261075731499",
		)
	})
}

func TestExternalConnectionsUploadsRetry(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "external-connections:uploads", "retry",
			"--api-key", "string",
			"--id", "1293384261075731499",
			"--ticket-id", "7b6a6449-b055-45a6-81f6-f6f0dffa4cc6",
		)
	})
}
