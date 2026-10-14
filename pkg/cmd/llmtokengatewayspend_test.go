// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestLlmTokenGatewaySpendRetrieveEvents(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"llm-token-gateway:spend", "retrieve-events",
			"--end-date", "'2019-12-27'",
			"--end-user-id", "end_user_id",
			"--metadata", "{foo: J!Q0Ok0bzJb7}",
			"--page-number", "1",
			"--page-size", "1",
			"--page-snapshot", "page[snapshot]",
			"--start-date", "'2019-12-27'",
			"--token-group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--token-key-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--token-user-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestLlmTokenGatewaySpendRetrieveSummary(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"llm-token-gateway:spend", "retrieve-summary",
			"--end-date", "'2019-12-27'",
			"--group-by", "token_group",
			"--start-date", "'2019-12-27'",
			"--end-user-id", "end_user_id",
			"--metadata", "{foo: J!Q0Ok0bzJb7}",
			"--page-number", "1",
			"--page-size", "1",
			"--page-snapshot", "page[snapshot]",
			"--token-group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--token-key-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--token-user-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
