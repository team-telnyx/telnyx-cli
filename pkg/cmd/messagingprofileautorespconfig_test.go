// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestMessagingProfilesAutorespConfigsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messaging-profiles:autoresp-configs", "create",
			"--profile-id", "profile_id",
			"--country-code", "US",
			"--keyword", "keyword1",
			"--keyword", "keyword2",
			"--op", "start",
			"--resp-text", "Thank you for your message",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"country_code: US\n" +
			"keywords:\n" +
			"  - keyword1\n" +
			"  - keyword2\n" +
			"op: start\n" +
			"resp_text: Thank you for your message\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"messaging-profiles:autoresp-configs", "create",
			"--profile-id", "profile_id",
		)
	})
}

func TestMessagingProfilesAutorespConfigsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messaging-profiles:autoresp-configs", "retrieve",
			"--profile-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--autoresp-cfg-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestMessagingProfilesAutorespConfigsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messaging-profiles:autoresp-configs", "update",
			"--profile-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--autoresp-cfg-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--country-code", "US",
			"--keyword", "keyword1",
			"--keyword", "keyword2",
			"--op", "start",
			"--resp-text", "Thank you for your message",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"country_code: US\n" +
			"keywords:\n" +
			"  - keyword1\n" +
			"  - keyword2\n" +
			"op: start\n" +
			"resp_text: Thank you for your message\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"messaging-profiles:autoresp-configs", "update",
			"--profile-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--autoresp-cfg-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestMessagingProfilesAutorespConfigsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messaging-profiles:autoresp-configs", "list",
			"--profile-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--country-code", "country_code",
			"--created-at", "{gte: gte, lte: lte}",
			"--updated-at", "{gte: gte, lte: lte}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(messagingProfilesAutorespConfigsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messaging-profiles:autoresp-configs", "list",
			"--profile-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--country-code", "country_code",
			"--created-at.gte", "gte",
			"--created-at.lte", "lte",
			"--updated-at.gte", "gte",
			"--updated-at.lte", "lte",
		)
	})
}

func TestMessagingProfilesAutorespConfigsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messaging-profiles:autoresp-configs", "delete",
			"--profile-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--autoresp-cfg-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
