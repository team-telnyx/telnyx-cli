// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestVoiceDesignsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "voice-designs", "create",
			"--api-key", "string",
			"--prompt", "Speak in a warm, friendly tone",
			"--text", "Hello, welcome to our service.",
			"--language", "Auto",
			"--max-new-tokens", "100",
			"--name", "friendly-narrator",
			"--repetition-penalty", "1",
			"--temperature", "0",
			"--top-k", "1",
			"--top-p", "0",
			"--voice-design-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"prompt: Speak in a warm, friendly tone\n" +
			"text: Hello, welcome to our service.\n" +
			"language: Auto\n" +
			"max_new_tokens: 100\n" +
			"name: friendly-narrator\n" +
			"repetition_penalty: 1\n" +
			"temperature: 0\n" +
			"top_k: 1\n" +
			"top_p: 0\n" +
			"voice_design_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "voice-designs", "create",
			"--api-key", "string",
		)
	})
}

func TestVoiceDesignsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "voice-designs", "retrieve",
			"--api-key", "string",
			"--id", "id",
			"--version", "1",
		)
	})
}

func TestVoiceDesignsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "voice-designs", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--filter-name", "filter[name]",
			"--page-number", "1",
			"--page-size", "1",
			"--sort", "name",
		)
	})
}

func TestVoiceDesignsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "voice-designs", "delete",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestVoiceDesignsDeleteVersion(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "voice-designs", "delete-version",
			"--api-key", "string",
			"--id", "id",
			"--version", "1",
		)
	})
}

func TestVoiceDesignsDownloadSample(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "voice-designs", "download-sample",
			"--api-key", "string",
			"--id", "id",
			"--version", "1",
			"--output", "/dev/null",
		)
	})
}

func TestVoiceDesignsRename(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "voice-designs", "rename",
			"--api-key", "string",
			"--id", "id",
			"--name", "updated-narrator",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("name: updated-narrator")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "voice-designs", "rename",
			"--api-key", "string",
			"--id", "id",
		)
	})
}
