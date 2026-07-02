// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestDocumentsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"documents", "retrieve",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestDocumentsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"documents", "update",
			"--document-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--customer-reference", "MY REF 001",
			"--filename", "test-document.pdf",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"customer_reference: MY REF 001\n" +
			"filename: test-document.pdf\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"documents", "update",
			"--document-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestDocumentsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"documents", "list",
			"--max-items", "10",
			"--filter", "{created_at: {gt: '2021-01-01T00:00:00Z', lt: '2021-04-09T22:25:27.521Z'}, customer_reference: {eq: MY REF 001, in: [REF001, REF002]}, filename: {contains: invoice}}",
			"--page-number", "0",
			"--page-size", "0",
			"--sort", "filename",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(documentsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"documents", "list",
			"--max-items", "10",
			"--filter.created-at", "{gt: '2021-01-01T00:00:00Z', lt: '2021-04-09T22:25:27.521Z'}",
			"--filter.customer-reference", "{eq: MY REF 001, in: [REF001, REF002]}",
			"--filter.filename", "{contains: invoice}",
			"--page-number", "0",
			"--page-size", "0",
			"--sort", "filename",
		)
	})
}

func TestDocumentsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"documents", "delete",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestDocumentsDownload(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"documents", "download",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--output", "/dev/null",
		)
	})
}

func TestDocumentsGenerateDownloadLink(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"documents", "generate-download-link",
			"--id", "550e8400-e29b-41d4-a716-446655440000",
		)
	})
}

func TestDocumentsUpload(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"documents", "upload",
			"--document", "{customer_reference: MY REF 001, file: ZXhhbXBsZSBvZiBlbmNvZGVkIGNvbnRlbnQ=, filename: test-document.pdf, url: https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(documentsUpload)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"documents", "upload",
			"--document.customer-reference", "MY REF 001",
			"--document.file", "ZXhhbXBsZSBvZiBlbmNvZGVkIGNvbnRlbnQ=",
			"--document.filename", "test-document.pdf",
			"--document.url", "https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"customer_reference: MY REF 001\n" +
			"file: ZXhhbXBsZSBvZiBlbmNvZGVkIGNvbnRlbnQ=\n" +
			"filename: test-document.pdf\n" +
			"url: https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"documents", "upload",
		)
	})
}

func TestDocumentsUploadJson(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"documents", "upload-json",
			"--document", "{customer_reference: MY REF 001, file: ZXhhbXBsZSBvZiBlbmNvZGVkIGNvbnRlbnQ=, filename: test-document.pdf, url: https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(documentsUploadJson)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"documents", "upload-json",
			"--document.customer-reference", "MY REF 001",
			"--document.file", "ZXhhbXBsZSBvZiBlbmNvZGVkIGNvbnRlbnQ=",
			"--document.filename", "test-document.pdf",
			"--document.url", "https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"customer_reference: MY REF 001\n" +
			"file: ZXhhbXBsZSBvZiBlbmNvZGVkIGNvbnRlbnQ=\n" +
			"filename: test-document.pdf\n" +
			"url: https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"documents", "upload-json",
		)
	})
}
