// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestExternalRequirementsSubNumberOrdersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"external-requirements:sub-number-orders", "retrieve",
			"--regulatory-requirement-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--sub-number-order-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestExternalRequirementsSubNumberOrdersUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"external-requirements:sub-number-orders", "update",
			"--regulatory-requirement-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--sub-number-order-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--requirement", "{first_name: Jane, last_name: Doe}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(externalRequirementsSubNumberOrdersUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"external-requirements:sub-number-orders", "update",
			"--regulatory-requirement-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--sub-number-order-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--requirement.first-name", "Jane",
			"--requirement.last-name", "Doe",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"requirement:\n" +
			"  first_name: Jane\n" +
			"  last_name: Doe\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"external-requirements:sub-number-orders", "update",
			"--regulatory-requirement-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--sub-number-order-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
