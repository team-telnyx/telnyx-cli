// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestExternalConnectionsPhoneNumbersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"external-connections:phone-numbers", "retrieve",
			"--id", "1293384261075731499",
			"--phone-number-id", "1234567889",
		)
	})
}

func TestExternalConnectionsPhoneNumbersUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"external-connections:phone-numbers", "update",
			"--id", "1293384261075731499",
			"--phone-number-id", "1234567889",
			"--location-id", "3fa85f64-5717-4562-b3fc-2c963f66afa6",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("location_id: 3fa85f64-5717-4562-b3fc-2c963f66afa6")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"external-connections:phone-numbers", "update",
			"--id", "1293384261075731499",
			"--phone-number-id", "1234567889",
		)
	})
}

func TestExternalConnectionsPhoneNumbersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"external-connections:phone-numbers", "list",
			"--max-items", "10",
			"--id", "1293384261075731499",
			"--filter", "{civic_address_id: {eq: '19990261512338516954'}, location_id: {eq: '19995665508264022121'}, phone_number: {contains: '+1970', eq: '+19705555098'}}",
			"--page-number", "0",
			"--page-size", "0",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(externalConnectionsPhoneNumbersList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"external-connections:phone-numbers", "list",
			"--max-items", "10",
			"--id", "1293384261075731499",
			"--filter.civic-address-id", "{eq: '19990261512338516954'}",
			"--filter.location-id", "{eq: '19995665508264022121'}",
			"--filter.phone-number", "{contains: '+1970', eq: '+19705555098'}",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}
