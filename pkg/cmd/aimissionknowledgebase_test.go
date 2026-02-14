// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIMissionsKnowledgeBasesCreateKnowledgeBase(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:missions:knowledge-bases", "create-knowledge-base",
		"--mission-id", "mission_id",
	)
}

func TestAIMissionsKnowledgeBasesDeleteKnowledgeBase(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:missions:knowledge-bases", "delete-knowledge-base",
		"--mission-id", "mission_id",
		"--knowledge-base-id", "knowledge_base_id",
	)
}

func TestAIMissionsKnowledgeBasesGetKnowledgeBase(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:missions:knowledge-bases", "get-knowledge-base",
		"--mission-id", "mission_id",
		"--knowledge-base-id", "knowledge_base_id",
	)
}

func TestAIMissionsKnowledgeBasesListKnowledgeBases(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:missions:knowledge-bases", "list-knowledge-bases",
		"--mission-id", "mission_id",
	)
}

func TestAIMissionsKnowledgeBasesUpdateKnowledgeBase(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:missions:knowledge-bases", "update-knowledge-base",
		"--mission-id", "mission_id",
		"--knowledge-base-id", "knowledge_base_id",
	)
}
