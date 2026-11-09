// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestAgentGatewaysTeamsToolPermissionsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-gateways:teams:tool-permissions", "create",
			"--instance-id", "32ce035d-3165-4031-b1a7-a7c0c259590a",
			"--team-id", "d7b21324-a2bf-406c-ba49-c6ec4d23d476",
			"--server", "{id: 88d5a97b-4080-4855-8bac-a39811b1dc24, tools: [get_order]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(agentGatewaysTeamsToolPermissionsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-gateways:teams:tool-permissions", "create",
			"--instance-id", "32ce035d-3165-4031-b1a7-a7c0c259590a",
			"--team-id", "d7b21324-a2bf-406c-ba49-c6ec4d23d476",
			"--server.id", "88d5a97b-4080-4855-8bac-a39811b1dc24",
			"--server.tools", "[get_order]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"servers:\n" +
			"  - id: 88d5a97b-4080-4855-8bac-a39811b1dc24\n" +
			"    tools:\n" +
			"      - get_order\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agent-gateways:teams:tool-permissions", "create",
			"--instance-id", "32ce035d-3165-4031-b1a7-a7c0c259590a",
			"--team-id", "d7b21324-a2bf-406c-ba49-c6ec4d23d476",
		)
	})
}

func TestAgentGatewaysTeamsToolPermissionsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-gateways:teams:tool-permissions", "list",
			"--instance-id", "32ce035d-3165-4031-b1a7-a7c0c259590a",
			"--team-id", "d7b21324-a2bf-406c-ba49-c6ec4d23d476",
		)
	})
}
