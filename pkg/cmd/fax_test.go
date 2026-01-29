// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestFaxesCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"faxes", "create",
		"--connection-id", "234423",
		"--from", "+13125790015",
		"--to", "+13127367276",
		"--black-threshold", "1",
		"--client-state", "aGF2ZSBhIG5pY2UgZGF5ID1d",
		"--from-display-name", "Company Name",
		"--media-name", "my_media_uploaded_to_media_storage_api",
		"--media-url", "https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf",
		"--monochrome=true",
		"--preview-format", "pdf",
		"--quality", "high",
		"--store-media=true",
		"--store-preview=true",
		"--t38-enabled=true",
		"--webhook-url", "https://www.example.com/server-b/",
	)
}

func TestFaxesRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"faxes", "retrieve",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestFaxesList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"faxes", "list",
		"--filter", "{created_at: {gt: '2020-02-02T22:25:27.521992Z', gte: '2020-02-02T22:25:27.521992Z', lt: '2020-02-02T22:25:27.521992Z', lte: '2020-02-02T22:25:27.521992Z'}, direction: {eq: inbound}, from: {eq: '+13127367276'}, to: {eq: '+13127367276'}}",
		"--page-number", "0",
		"--page-size", "0",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(faxesList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"faxes", "list",
		"--filter.created-at", "{gt: '2020-02-02T22:25:27.521992Z', gte: '2020-02-02T22:25:27.521992Z', lt: '2020-02-02T22:25:27.521992Z', lte: '2020-02-02T22:25:27.521992Z'}",
		"--filter.direction", "{eq: inbound}",
		"--filter.from", "{eq: '+13127367276'}",
		"--filter.to", "{eq: '+13127367276'}",
		"--page-number", "0",
		"--page-size", "0",
	)
}

func TestFaxesDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"faxes", "delete",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}
