// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIMissionsKnowledgeBasesCreateKnowledgeBase(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:missions:knowledge-bases", "create-knowledge-base",
		"--api-key", "string",
		"--mission-id", "mission_id",
	)
}

func TestAIMissionsKnowledgeBasesDeleteKnowledgeBase(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:missions:knowledge-bases", "delete-knowledge-base",
		"--api-key", "string",
		"--mission-id", "mission_id",
		"--knowledge-base-id", "knowledge_base_id",
	)
}

func TestAIMissionsKnowledgeBasesGetKnowledgeBase(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:missions:knowledge-bases", "get-knowledge-base",
		"--api-key", "string",
		"--mission-id", "mission_id",
		"--knowledge-base-id", "knowledge_base_id",
	)
}

func TestAIMissionsKnowledgeBasesListKnowledgeBases(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:missions:knowledge-bases", "list-knowledge-bases",
		"--api-key", "string",
		"--mission-id", "mission_id",
	)
}

func TestAIMissionsKnowledgeBasesUpdateKnowledgeBase(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:missions:knowledge-bases", "update-knowledge-base",
		"--api-key", "string",
		"--mission-id", "mission_id",
		"--knowledge-base-id", "knowledge_base_id",
	)
}
