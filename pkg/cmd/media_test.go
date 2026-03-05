// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestMediaRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "media", "retrieve",
			"--api-key", "string",
			"--media-name", "media_name",
		)
	})
}

func TestMediaUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "media", "update",
			"--api-key", "string",
			"--media-name", "media_name",
			"--media-url", "http://www.example.com/audio.mp3",
			"--ttl-secs", "86400",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"media_url: http://www.example.com/audio.mp3\n" +
			"ttl_secs: 86400\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "media", "update",
			"--api-key", "string",
			"--media-name", "media_name",
		)
	})
}

func TestMediaList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "media", "list",
			"--api-key", "string",
			"--filter", "{content_type: [application_xml]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(mediaList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "media", "list",
			"--api-key", "string",
			"--filter.content-type", "[application_xml]",
		)
	})
}

func TestMediaDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "media", "delete",
			"--api-key", "string",
			"--media-name", "media_name",
		)
	})
}

func TestMediaDownload(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "media", "download",
			"--api-key", "string",
			"--media-name", "media_name",
			"--output", "/dev/null",
		)
	})
}

func TestMediaUpload(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "media", "upload",
			"--api-key", "string",
			"--media-url", "http://www.example.com/audio.mp3",
			"--media-name", "my-file",
			"--ttl-secs", "86400",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"media_url: http://www.example.com/audio.mp3\n" +
			"media_name: my-file\n" +
			"ttl_secs: 86400\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "media", "upload",
			"--api-key", "string",
		)
	})
}
