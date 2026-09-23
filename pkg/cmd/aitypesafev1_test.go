// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAITypesafeV1Systemone(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:typesafe:v1", "systemone",
			"--questions", "{team: {criteria: {billing: Payments and refunds, technical_support: Service faults and technical problems, sales: New purchases}, instructions: Choose the team that should handle this incident., type: choice}, production_incident: {instructions: Does the message describe an active production incident?, type: noul, criteria: {'false': 'false', 'true': 'true'}}, urgency: {criteria: [Low, Normal, High, Critical], instructions: Rate operational urgency., type: score}}",
			"--state", "Our production calls are failing. Every customer is affected.",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"questions:\n" +
			"  team:\n" +
			"    criteria:\n" +
			"      billing: Payments and refunds\n" +
			"      technical_support: Service faults and technical problems\n" +
			"      sales: New purchases\n" +
			"    instructions: Choose the team that should handle this incident.\n" +
			"    type: choice\n" +
			"  production_incident:\n" +
			"    instructions: Does the message describe an active production incident?\n" +
			"    type: noul\n" +
			"    criteria:\n" +
			"      'false': 'false'\n" +
			"      'true': 'true'\n" +
			"  urgency:\n" +
			"    criteria:\n" +
			"      - Low\n" +
			"      - Normal\n" +
			"      - High\n" +
			"      - Critical\n" +
			"    instructions: Rate operational urgency.\n" +
			"    type: score\n" +
			"state: Our production calls are failing. Every customer is affected.\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:typesafe:v1", "systemone",
		)
	})
}
