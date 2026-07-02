// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestDirRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"dir", "retrieve",
			"--dir-id", "16635d38-75a6-4481-82e8-69af60e05011",
		)
	})
}

func TestDirUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"dir", "update",
			"--dir-id", "16635d38-75a6-4481-82e8-69af60e05011",
			"--authorizer-email", "dev@stainless.com",
			"--authorizer-name", "authorizer_name",
			"--call-reason", "Appointment reminders",
			"--call-reason", "Billing inquiries",
			"--call-reason", "Lab results",
			"--certify-brand-is-accurate=true",
			"--certify-ip-ownership=true",
			"--certify-no-shaft-content=true",
			"--display-name", "Acme Plumbing & Wellness",
			"--document", "{document_id: 2a7e8337-e803-4057-a4ae-26c40eb0bc6c, document_type: business_registration, description: Certificate of incorporation.}",
			"--logo-url", "https://acmeplumbing.example.com/logo-v2-256.bmp",
			"--reselling=true",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(dirUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"dir", "update",
			"--dir-id", "16635d38-75a6-4481-82e8-69af60e05011",
			"--authorizer-email", "dev@stainless.com",
			"--authorizer-name", "authorizer_name",
			"--call-reason", "Appointment reminders",
			"--call-reason", "Billing inquiries",
			"--call-reason", "Lab results",
			"--certify-brand-is-accurate=true",
			"--certify-ip-ownership=true",
			"--certify-no-shaft-content=true",
			"--display-name", "Acme Plumbing & Wellness",
			"--document.document-id", "2a7e8337-e803-4057-a4ae-26c40eb0bc6c",
			"--document.document-type", "business_registration",
			"--document.description", "Certificate of incorporation.",
			"--logo-url", "https://acmeplumbing.example.com/logo-v2-256.bmp",
			"--reselling=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"authorizer_email: dev@stainless.com\n" +
			"authorizer_name: authorizer_name\n" +
			"call_reasons:\n" +
			"  - Appointment reminders\n" +
			"  - Billing inquiries\n" +
			"  - Lab results\n" +
			"certify_brand_is_accurate: true\n" +
			"certify_ip_ownership: true\n" +
			"certify_no_shaft_content: true\n" +
			"display_name: Acme Plumbing & Wellness\n" +
			"documents:\n" +
			"  - document_id: 2a7e8337-e803-4057-a4ae-26c40eb0bc6c\n" +
			"    document_type: business_registration\n" +
			"    description: Certificate of incorporation.\n" +
			"logo_url: https://acmeplumbing.example.com/logo-v2-256.bmp\n" +
			"reselling: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"dir", "update",
			"--dir-id", "16635d38-75a6-4481-82e8-69af60e05011",
		)
	})
}

func TestDirList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"dir", "list",
			"--max-items", "10",
			"--filter-call-reason-contains", "filter[call_reason][contains]",
			"--filter-display-name-contains", "filter[display_name][contains]",
			"--filter-enterprise-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--filter-expiring-at-gte", "'2019-12-27T18:11:19.117Z'",
			"--filter-expiring-at-lte", "'2019-12-27T18:11:19.117Z'",
			"--filter-status", "draft",
			"--page-number", "1",
			"--page-size", "20",
			"--sort", "created_at",
		)
	})
}

func TestDirDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"dir", "delete",
			"--dir-id", "16635d38-75a6-4481-82e8-69af60e05011",
		)
	})
}

func TestDirCreateLoa(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"dir", "create-loa",
			"--dir-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--phone-number", "+13125550000",
			"--agent", "{administrative_area: administrative_area, city: city, contact_email: dev@stainless.com, contact_name: contact_name, contact_phone: '+13125550000', contact_title: contact_title, country: US, legal_name: legal_name, postal_code: postal_code, street_address: street_address, dba: dba, extended_address: extended_address}",
			"--signature", "{image_base64: x, signer_name: signer_name}",
			"--output", "/dev/null",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(dirCreateLoa)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"dir", "create-loa",
			"--dir-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--phone-number", "+13125550000",
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
			"--signature.image-base64", "x",
			"--signature.signer-name", "signer_name",
			"--output", "/dev/null",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"phone_numbers:\n" +
			"  - '+13125550000'\n" +
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
			"  image_base64: x\n" +
			"  signer_name: signer_name\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"dir", "create-loa",
			"--dir-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--output", "/dev/null",
		)
	})
}

func TestDirListDocumentTypes(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"dir", "list-document-types",
		)
	})
}

func TestDirListInfringementClaims(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"dir", "list-infringement-claims",
			"--max-items", "10",
			"--dir-id", "16635d38-75a6-4481-82e8-69af60e05011",
			"--page-number", "1",
			"--page-size", "20",
		)
	})
}

func TestDirSubmit(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"dir", "submit",
			"--dir-id", "16635d38-75a6-4481-82e8-69af60e05011",
		)
	})
}

func TestDirUpdateInfringement(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"dir", "update-infringement",
			"--dir-id", "16635d38-75a6-4481-82e8-69af60e05011",
			"--certify-brand-is-accurate=true",
			"--certify-ip-ownership=true",
			"--certify-no-infringement=true",
			"--certify-no-shaft-content=true",
			"--infringement-resolution-notes", "Updated the display name to remove the disputed mark and re-uploaded the authorization.",
			"--call-reason", "[string]",
			"--display-name", "x",
			"--document", "[{document_id: 2a7e8337-e803-4057-a4ae-26c40eb0bc6c, document_type: business_registration, description: Certificate of incorporation.}]",
			"--logo-url", "logo_url",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(dirUpdateInfringement)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"dir", "update-infringement",
			"--dir-id", "16635d38-75a6-4481-82e8-69af60e05011",
			"--certify-brand-is-accurate=true",
			"--certify-ip-ownership=true",
			"--certify-no-infringement=true",
			"--certify-no-shaft-content=true",
			"--infringement-resolution-notes", "Updated the display name to remove the disputed mark and re-uploaded the authorization.",
			"--call-reason", "[string]",
			"--display-name", "x",
			"--document.document-id", "2a7e8337-e803-4057-a4ae-26c40eb0bc6c",
			"--document.document-type", "business_registration",
			"--document.description", "Certificate of incorporation.",
			"--logo-url", "logo_url",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"certify_brand_is_accurate: true\n" +
			"certify_ip_ownership: true\n" +
			"certify_no_infringement: true\n" +
			"certify_no_shaft_content: true\n" +
			"infringement_resolution_notes: >-\n" +
			"  Updated the display name to remove the disputed mark and re-uploaded the\n" +
			"  authorization.\n" +
			"call_reasons:\n" +
			"  - string\n" +
			"display_name: x\n" +
			"documents:\n" +
			"  - document_id: 2a7e8337-e803-4057-a4ae-26c40eb0bc6c\n" +
			"    document_type: business_registration\n" +
			"    description: Certificate of incorporation.\n" +
			"logo_url: logo_url\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"dir", "update-infringement",
			"--dir-id", "16635d38-75a6-4481-82e8-69af60e05011",
		)
	})
}
