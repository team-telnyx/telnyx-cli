// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestRequirementTypesRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"requirement-types", "retrieve",
		"--id", "a38c217a-8019-48f8-bff6-0fdd9939075b",
	)
}

func TestRequirementTypesList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"requirement-types", "list",
		"--filter", "{name: {contains: utility bill}}",
		"--sort", "name",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(requirementTypesList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"requirement-types", "list",
		"--filter.name", "{contains: utility bill}",
		"--sort", "name",
	)
}
