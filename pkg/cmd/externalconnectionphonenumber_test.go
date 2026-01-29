// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestExternalConnectionsPhoneNumbersRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"external-connections:phone-numbers", "retrieve",
		"--id", "id",
		"--phone-number-id", "1234567889",
	)
}

func TestExternalConnectionsPhoneNumbersUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"external-connections:phone-numbers", "update",
		"--id", "id",
		"--phone-number-id", "1234567889",
		"--location-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestExternalConnectionsPhoneNumbersList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"external-connections:phone-numbers", "list",
		"--id", "id",
		"--filter", "{civic_address_id: {eq: '19990261512338516954'}, location_id: {eq: '19995665508264022121'}, phone_number: {contains: '+1970', eq: '+19705555098'}}",
		"--page", "{number: 1, size: 1}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(externalConnectionsPhoneNumbersList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"external-connections:phone-numbers", "list",
		"--id", "id",
		"--filter.civic-address-id", "{eq: '19990261512338516954'}",
		"--filter.location-id", "{eq: '19995665508264022121'}",
		"--filter.phone-number", "{contains: '+1970', eq: '+19705555098'}",
		"--page.number", "1",
		"--page.size", "1",
	)
}
