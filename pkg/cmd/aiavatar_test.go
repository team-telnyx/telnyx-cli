// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"strings"
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIAvatarsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:avatars", "create",
			"--image-file", mocktest.TestFile(t, "Example data"),
			"--name", "Customer portrait",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		testFile := mocktest.TestFile(t, "Example data")
		// Test piping YAML data over stdin
		pipeDataStr := "" +
			"image_file: Example data\n" +
			"name: Customer portrait\n"
		pipeDataStr = strings.ReplaceAll(pipeDataStr, "Example data", testFile)
		pipeData := []byte(pipeDataStr)
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:avatars", "create",
		)
	})
}

func TestAIAvatarsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:avatars", "retrieve",
			"--avatar-id", "avatar_0123456789abcdef0123456789abcdef",
		)
	})
}

func TestAIAvatarsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:avatars", "update",
			"--avatar-id", "avatar_0123456789abcdef0123456789abcdef",
			"--name", "Updated portrait",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("name: Updated portrait")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:avatars", "update",
			"--avatar-id", "avatar_0123456789abcdef0123456789abcdef",
		)
	})
}

func TestAIAvatarsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:avatars", "list",
			"--max-items", "10",
			"--filter-avatar-type", "stock",
			"--filter-id", "avatar_0123456789abcdef0123456789abcdef",
			"--filter-name", "x",
			"--page-number", "1",
			"--page-size", "1",
		)
	})
}
