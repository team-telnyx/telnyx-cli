// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestWebSearchCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"web-search", "create",
			"--query", "latest AI agent frameworks",
			"--count", "10",
			"--country", "US",
			"--exclude-domain", "pinterest.com",
			"--freshness", "week",
			"--include-domain", "arxiv.org",
			"--include-domain", "github.com",
			"--livecrawl=false",
			"--safesearch", "moderate",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"query: latest AI agent frameworks\n" +
			"count: 10\n" +
			"country: US\n" +
			"exclude_domains:\n" +
			"  - pinterest.com\n" +
			"freshness: week\n" +
			"include_domains:\n" +
			"  - arxiv.org\n" +
			"  - github.com\n" +
			"livecrawl: false\n" +
			"safesearch: moderate\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"web-search", "create",
		)
	})
}

func TestWebSearchContents(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"web-search", "contents",
			"--url", "https://en.wikipedia.org/wiki/Artificial_intelligence",
			"--crawl-timeout", "10",
			"--format", "markdown",
			"--format", "metadata",
			"--max-age", "null",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"urls:\n" +
			"  - https://en.wikipedia.org/wiki/Artificial_intelligence\n" +
			"crawl_timeout: 10\n" +
			"formats:\n" +
			"  - markdown\n" +
			"  - metadata\n" +
			"max_age: null\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"web-search", "contents",
		)
	})
}
