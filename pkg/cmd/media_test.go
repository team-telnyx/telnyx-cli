// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestMediaRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"media", "retrieve",
		"--media-name", "media_name",
	)
}

func TestMediaUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"media", "update",
		"--media-name", "media_name",
		"--media-url", "http://www.example.com/audio.mp3",
		"--ttl-secs", "86400",
	)
}

func TestMediaList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"media", "list",
		"--filter", "{content_type: [application_xml]}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(mediaList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"media", "list",
		"--filter.content-type", "[application_xml]",
	)
}

func TestMediaDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"media", "delete",
		"--media-name", "media_name",
	)
}

func TestMediaUpload(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"media", "upload",
		"--media-url", "http://www.example.com/audio.mp3",
		"--media-name", "my-file",
		"--ttl-secs", "86400",
	)
}
