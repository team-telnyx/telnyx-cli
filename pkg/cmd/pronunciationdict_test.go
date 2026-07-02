// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestPronunciationDictsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pronunciation-dicts", "create",
			"--item", "{alias: tel-nicks, text: Telnyx, type: alias}",
			"--name", "Brand Names",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"items:\n" +
			"  - alias: tel-nicks\n" +
			"    text: Telnyx\n" +
			"    type: alias\n" +
			"name: Brand Names\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"pronunciation-dicts", "create",
		)
	})
}

func TestPronunciationDictsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pronunciation-dicts", "retrieve",
			"--id", "c215a3e1-be41-4701-97e8-1d3c22f9a5b7",
		)
	})
}

func TestPronunciationDictsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pronunciation-dicts", "update",
			"--id", "c215a3e1-be41-4701-97e8-1d3c22f9a5b7",
			"--item", "{alias: tel-nicks, text: Telnyx, type: alias}",
			"--name", "Updated Brand Names",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"items:\n" +
			"  - alias: tel-nicks\n" +
			"    text: Telnyx\n" +
			"    type: alias\n" +
			"name: Updated Brand Names\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"pronunciation-dicts", "update",
			"--id", "c215a3e1-be41-4701-97e8-1d3c22f9a5b7",
		)
	})
}

func TestPronunciationDictsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pronunciation-dicts", "list",
			"--max-items", "10",
			"--page-number", "1",
			"--page-size", "1",
		)
	})
}

func TestPronunciationDictsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pronunciation-dicts", "delete",
			"--id", "c215a3e1-be41-4701-97e8-1d3c22f9a5b7",
		)
	})
}
