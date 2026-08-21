// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestEmailBlocksCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-blocks", "create",
			"--to", "spammer@bad.tld",
			"--domain-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--expires-at", "'2026-12-31T23:59:59Z'",
			"--from", "from",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"to: spammer@bad.tld\n" +
			"domain_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"expires_at: '2026-12-31T23:59:59Z'\n" +
			"from: from\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"email-blocks", "create",
		)
	})
}

func TestEmailBlocksRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-blocks", "retrieve",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestEmailBlocksList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-blocks", "list",
			"--max-items", "10",
			"--filter-created-after", "'2019-12-27T18:11:19.117Z'",
			"--filter-created-before", "'2019-12-27T18:11:19.117Z'",
			"--filter-domain-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--filter-reason", "hard_bounce",
			"--page-after", "page[after]",
			"--page-before", "page[before]",
			"--page-number", "1",
			"--page-size", "1",
			"--sort", "created_at",
		)
	})
}

func TestEmailBlocksDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-blocks", "delete",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestEmailBlocksRetrieveEvents(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-blocks", "retrieve-events",
			"--max-items", "10",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--page-number", "1",
			"--page-size", "1",
		)
	})
}

func TestEmailBlocksRetrieveExport(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-blocks", "retrieve-export",
			"--filter-created-after", "'2019-12-27T18:11:19.117Z'",
			"--filter-created-before", "'2019-12-27T18:11:19.117Z'",
			"--filter-domain-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--filter-reason", "hard_bounce",
			"--page-number", "1",
			"--page-size", "1",
			"--sort", "created_at",
		)
	})
}
