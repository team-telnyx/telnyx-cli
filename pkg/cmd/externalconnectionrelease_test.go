// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestExternalConnectionsReleasesRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"external-connections:releases", "retrieve",
		"--id", "id",
		"--release-id", "7b6a6449-b055-45a6-81f6-f6f0dffa4cc6",
	)
}

func TestExternalConnectionsReleasesList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"external-connections:releases", "list",
		"--id", "id",
		"--filter", "{civic_address_id: {eq: '19990261512338516954'}, location_id: {eq: '19995665508264022121'}, phone_number: {contains: '+123', eq: '+1234567890'}, status: {eq: [pending, in_progress]}}",
		"--page", "{number: 1, size: 1}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(externalConnectionsReleasesList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"external-connections:releases", "list",
		"--id", "id",
		"--filter.civic-address-id", "{eq: '19990261512338516954'}",
		"--filter.location-id", "{eq: '19995665508264022121'}",
		"--filter.phone-number", "{contains: '+123', eq: '+1234567890'}",
		"--filter.status", "{eq: [pending, in_progress]}",
		"--page.number", "1",
		"--page.size", "1",
	)
}
