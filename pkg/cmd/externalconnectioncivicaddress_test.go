// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestExternalConnectionsCivicAddressesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"external-connections:civic-addresses", "retrieve",
			"--id", "1293384261075731499",
			"--address-id", "318fb664-d341-44d2-8405-e6bfb9ced6d9",
		)
	})
}

func TestExternalConnectionsCivicAddressesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"external-connections:civic-addresses", "list",
			"--id", "1293384261075731499",
			"--filter", "{country: [US, CA, MX, BR]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(externalConnectionsCivicAddressesList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"external-connections:civic-addresses", "list",
			"--id", "1293384261075731499",
			"--filter.country", "[US, CA, MX, BR]",
		)
	})
}
