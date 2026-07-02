// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestEnterprisesReputationLoaUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"enterprises:reputation:loa", "update",
			"--enterprise-id", "4a6192a4-573d-446d-b3ce-aff9117272a6",
			"--loa-document-id", "2a7e8337-e803-4057-a4ae-26c40eb0bc6c",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("loa_document_id: 2a7e8337-e803-4057-a4ae-26c40eb0bc6c")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"enterprises:reputation:loa", "update",
			"--enterprise-id", "4a6192a4-573d-446d-b3ce-aff9117272a6",
		)
	})
}

func TestEnterprisesReputationLoaRender(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"enterprises:reputation:loa", "render",
			"--enterprise-id", "4a6192a4-573d-446d-b3ce-aff9117272a6",
			"--agent", "{administrative_area: administrative_area, city: city, contact_email: dev@stainless.com, contact_name: contact_name, contact_phone: '+13125550000', contact_title: contact_title, country: US, legal_name: legal_name, postal_code: postal_code, street_address: street_address, dba: dba, extended_address: extended_address}",
			"--signature", "{image_base64: image_base64, signer_name: signer_name}",
			"--output", "/dev/null",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(enterprisesReputationLoaRender)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"enterprises:reputation:loa", "render",
			"--enterprise-id", "4a6192a4-573d-446d-b3ce-aff9117272a6",
			"--agent.administrative-area", "administrative_area",
			"--agent.city", "city",
			"--agent.contact-email", "dev@stainless.com",
			"--agent.contact-name", "contact_name",
			"--agent.contact-phone", "+13125550000",
			"--agent.contact-title", "contact_title",
			"--agent.country", "US",
			"--agent.legal-name", "legal_name",
			"--agent.postal-code", "postal_code",
			"--agent.street-address", "street_address",
			"--agent.dba", "dba",
			"--agent.extended-address", "extended_address",
			"--signature.image-base64", "image_base64",
			"--signature.signer-name", "signer_name",
			"--output", "/dev/null",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"agent:\n" +
			"  administrative_area: administrative_area\n" +
			"  city: city\n" +
			"  contact_email: dev@stainless.com\n" +
			"  contact_name: contact_name\n" +
			"  contact_phone: '+13125550000'\n" +
			"  contact_title: contact_title\n" +
			"  country: US\n" +
			"  legal_name: legal_name\n" +
			"  postal_code: postal_code\n" +
			"  street_address: street_address\n" +
			"  dba: dba\n" +
			"  extended_address: extended_address\n" +
			"signature:\n" +
			"  image_base64: image_base64\n" +
			"  signer_name: signer_name\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"enterprises:reputation:loa", "render",
			"--enterprise-id", "4a6192a4-573d-446d-b3ce-aff9117272a6",
			"--output", "/dev/null",
		)
	})
}
