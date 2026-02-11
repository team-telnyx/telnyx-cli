// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIMissionsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:missions", "create",
		"--name", "name",
		"--description", "description",
		"--execution-mode", "external",
		"--instructions", "instructions",
		"--metadata", "{foo: bar}",
		"--model", "model",
	)
}

func TestAIMissionsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:missions", "retrieve",
		"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestAIMissionsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:missions", "list",
		"--page-number", "1",
		"--page-size", "1",
	)
}

func TestAIMissionsCloneMission(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:missions", "clone-mission",
		"--mission-id", "mission_id",
	)
}

func TestAIMissionsDeleteMission(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:missions", "delete-mission",
		"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestAIMissionsListEvents(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:missions", "list-events",
		"--page-number", "1",
		"--page-size", "1",
		"--type", "type",
	)
}

func TestAIMissionsUpdateMission(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:missions", "update-mission",
		"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--description", "description",
		"--execution-mode", "external",
		"--instructions", "instructions",
		"--metadata", "{foo: bar}",
		"--model", "model",
		"--name", "name",
	)
}
