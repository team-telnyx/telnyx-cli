// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"strings"
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestVoiceClonesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"voice-clones", "create",
			"--gender", "male",
			"--language", "en",
			"--name", "clone-narrator",
			"--voice-design-id", "550e8400-e29b-41d4-a716-446655440000",
			"--provider", "telnyx",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"gender: male\n" +
			"language: en\n" +
			"name: clone-narrator\n" +
			"voice_design_id: 550e8400-e29b-41d4-a716-446655440000\n" +
			"provider: telnyx\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"voice-clones", "create",
		)
	})
}

func TestVoiceClonesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"voice-clones", "update",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--name", "updated-clone",
			"--gender", "male",
			"--language", "language",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: updated-clone\n" +
			"gender: male\n" +
			"language: language\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"voice-clones", "update",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestVoiceClonesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"voice-clones", "list",
			"--max-items", "10",
			"--filter-name", "filter[name]",
			"--filter-provider", "telnyx",
			"--page-number", "1",
			"--page-size", "1",
			"--sort", "name",
		)
	})
}

func TestVoiceClonesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"voice-clones", "delete",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestVoiceClonesCreateFromUpload(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"voice-clones", "create-from-upload",
			"--audio-file", mocktest.TestFile(t, "Example data"),
			"--language", "lkf-Lz1vLbBu-9uDh-9AHaOS2D-Cbf",
			"--name", "name",
			"--gender", "male",
			"--label", "label",
			"--provider", "telnyx",
			"--ref-text", "ref_text",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		testFile := mocktest.TestFile(t, "Example data")
		// Test piping YAML data over stdin
		pipeDataStr := "" +
			"audio_file: Example data\n" +
			"language: lkf-Lz1vLbBu-9uDh-9AHaOS2D-Cbf\n" +
			"name: name\n" +
			"gender: male\n" +
			"label: label\n" +
			"provider: telnyx\n" +
			"ref_text: ref_text\n"
		pipeDataStr = strings.ReplaceAll(pipeDataStr, "Example data", testFile)
		pipeData := []byte(pipeDataStr)
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"voice-clones", "create-from-upload",
		)
	})
}

func TestVoiceClonesDownloadSample(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"voice-clones", "download-sample",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--output", "/dev/null",
		)
	})
}
