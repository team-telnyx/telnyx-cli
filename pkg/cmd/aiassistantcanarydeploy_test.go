// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestAIAssistantsCanaryDeploysCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:assistants:canary-deploys", "create",
		"--assistant-id", "assistant_id",
		"--version", "{percentage: 1, version_id: version_id}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(aiAssistantsCanaryDeploysCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:assistants:canary-deploys", "create",
		"--assistant-id", "assistant_id",
		"--version.percentage", "1",
		"--version.version-id", "version_id",
	)
}

func TestAIAssistantsCanaryDeploysRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:assistants:canary-deploys", "retrieve",
		"--assistant-id", "assistant_id",
	)
}

func TestAIAssistantsCanaryDeploysUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:assistants:canary-deploys", "update",
		"--assistant-id", "assistant_id",
		"--version", "{percentage: 1, version_id: version_id}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(aiAssistantsCanaryDeploysUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:assistants:canary-deploys", "update",
		"--assistant-id", "assistant_id",
		"--version.percentage", "1",
		"--version.version-id", "version_id",
	)
}

func TestAIAssistantsCanaryDeploysDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:assistants:canary-deploys", "delete",
		"--assistant-id", "assistant_id",
	)
}
