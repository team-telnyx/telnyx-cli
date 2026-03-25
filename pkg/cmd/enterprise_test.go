// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestEnterprisesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"enterprises", "create",
			"--billing-address", "{administrative_area: Illinois, city: Chicago, country: United States, postal_code: '60601', street_address: 123 Main St, extended_address: Suite 400}",
			"--billing-contact", "{email: billing@acme.com, first_name: John, last_name: Doe, phone_number: '15551234568'}",
			"--country-code", "US",
			"--doing-business-as", "Acme",
			"--fein", "12-3456789",
			"--industry", "technology",
			"--legal-name", "Acme Corp Inc.",
			"--number-of-employees", "51-200",
			"--organization-contact", "{email: jane.smith@acme.com, first_name: Jane, job_title: VP of Engineering, last_name: Smith, phone: '+16035551234'}",
			"--organization-legal-type", "corporation",
			"--organization-physical-address", "{administrative_area: Illinois, city: Chicago, country: United States, postal_code: '60601', street_address: 123 Main St, extended_address: Suite 400}",
			"--organization-type", "commercial",
			"--website", "https://acme.com",
			"--corporate-registration-number", "corporate_registration_number",
			"--customer-reference", "customer_reference",
			"--dun-bradstreet-number", "dun_bradstreet_number",
			"--primary-business-domain-sic-code", "7372",
			"--professional-license-number", "professional_license_number",
			"--role-type", "enterprise",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(enterprisesCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"enterprises", "create",
			"--billing-address.administrative-area", "Illinois",
			"--billing-address.city", "Chicago",
			"--billing-address.country", "United States",
			"--billing-address.postal-code", "60601",
			"--billing-address.street-address", "123 Main St",
			"--billing-address.extended-address", "Suite 400",
			"--billing-contact.email", "billing@acme.com",
			"--billing-contact.first-name", "John",
			"--billing-contact.last-name", "Doe",
			"--billing-contact.phone-number", "15551234568",
			"--country-code", "US",
			"--doing-business-as", "Acme",
			"--fein", "12-3456789",
			"--industry", "technology",
			"--legal-name", "Acme Corp Inc.",
			"--number-of-employees", "51-200",
			"--organization-contact.email", "jane.smith@acme.com",
			"--organization-contact.first-name", "Jane",
			"--organization-contact.job-title", "VP of Engineering",
			"--organization-contact.last-name", "Smith",
			"--organization-contact.phone", "+16035551234",
			"--organization-legal-type", "corporation",
			"--organization-physical-address.administrative-area", "Illinois",
			"--organization-physical-address.city", "Chicago",
			"--organization-physical-address.country", "United States",
			"--organization-physical-address.postal-code", "60601",
			"--organization-physical-address.street-address", "123 Main St",
			"--organization-physical-address.extended-address", "Suite 400",
			"--organization-type", "commercial",
			"--website", "https://acme.com",
			"--corporate-registration-number", "corporate_registration_number",
			"--customer-reference", "customer_reference",
			"--dun-bradstreet-number", "dun_bradstreet_number",
			"--primary-business-domain-sic-code", "7372",
			"--professional-license-number", "professional_license_number",
			"--role-type", "enterprise",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"billing_address:\n" +
			"  administrative_area: Illinois\n" +
			"  city: Chicago\n" +
			"  country: United States\n" +
			"  postal_code: '60601'\n" +
			"  street_address: 123 Main St\n" +
			"  extended_address: Suite 400\n" +
			"billing_contact:\n" +
			"  email: billing@acme.com\n" +
			"  first_name: John\n" +
			"  last_name: Doe\n" +
			"  phone_number: '15551234568'\n" +
			"country_code: US\n" +
			"doing_business_as: Acme\n" +
			"fein: 12-3456789\n" +
			"industry: technology\n" +
			"legal_name: Acme Corp Inc.\n" +
			"number_of_employees: 51-200\n" +
			"organization_contact:\n" +
			"  email: jane.smith@acme.com\n" +
			"  first_name: Jane\n" +
			"  job_title: VP of Engineering\n" +
			"  last_name: Smith\n" +
			"  phone: '+16035551234'\n" +
			"organization_legal_type: corporation\n" +
			"organization_physical_address:\n" +
			"  administrative_area: Illinois\n" +
			"  city: Chicago\n" +
			"  country: United States\n" +
			"  postal_code: '60601'\n" +
			"  street_address: 123 Main St\n" +
			"  extended_address: Suite 400\n" +
			"organization_type: commercial\n" +
			"website: https://acme.com\n" +
			"corporate_registration_number: corporate_registration_number\n" +
			"customer_reference: customer_reference\n" +
			"dun_bradstreet_number: dun_bradstreet_number\n" +
			"primary_business_domain_sic_code: '7372'\n" +
			"professional_license_number: professional_license_number\n" +
			"role_type: enterprise\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"enterprises", "create",
		)
	})
}

func TestEnterprisesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"enterprises", "retrieve",
			"--enterprise-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestEnterprisesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"enterprises", "update",
			"--enterprise-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--billing-address", "{administrative_area: Illinois, city: Chicago, country: United States, postal_code: '60601', street_address: 123 Main St, extended_address: Suite 400}",
			"--billing-contact", "{email: billing@acme.com, first_name: John, last_name: Doe, phone_number: '15551234568'}",
			"--corporate-registration-number", "corporate_registration_number",
			"--customer-reference", "customer_reference",
			"--doing-business-as", "doing_business_as",
			"--dun-bradstreet-number", "dun_bradstreet_number",
			"--fein", "fein",
			"--industry", "industry",
			"--legal-name", "xxx",
			"--number-of-employees", "1-10",
			"--organization-contact", "{email: jane.smith@acme.com, first_name: Jane, job_title: VP of Engineering, last_name: Smith, phone: '+16035551234'}",
			"--organization-legal-type", "corporation",
			"--organization-physical-address", "{administrative_area: Illinois, city: Chicago, country: United States, postal_code: '60601', street_address: 123 Main St, extended_address: Suite 400}",
			"--primary-business-domain-sic-code", "primary_business_domain_sic_code",
			"--professional-license-number", "professional_license_number",
			"--website", "website",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(enterprisesUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"enterprises", "update",
			"--enterprise-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--billing-address.administrative-area", "Illinois",
			"--billing-address.city", "Chicago",
			"--billing-address.country", "United States",
			"--billing-address.postal-code", "60601",
			"--billing-address.street-address", "123 Main St",
			"--billing-address.extended-address", "Suite 400",
			"--billing-contact.email", "billing@acme.com",
			"--billing-contact.first-name", "John",
			"--billing-contact.last-name", "Doe",
			"--billing-contact.phone-number", "15551234568",
			"--corporate-registration-number", "corporate_registration_number",
			"--customer-reference", "customer_reference",
			"--doing-business-as", "doing_business_as",
			"--dun-bradstreet-number", "dun_bradstreet_number",
			"--fein", "fein",
			"--industry", "industry",
			"--legal-name", "xxx",
			"--number-of-employees", "1-10",
			"--organization-contact.email", "jane.smith@acme.com",
			"--organization-contact.first-name", "Jane",
			"--organization-contact.job-title", "VP of Engineering",
			"--organization-contact.last-name", "Smith",
			"--organization-contact.phone", "+16035551234",
			"--organization-legal-type", "corporation",
			"--organization-physical-address.administrative-area", "Illinois",
			"--organization-physical-address.city", "Chicago",
			"--organization-physical-address.country", "United States",
			"--organization-physical-address.postal-code", "60601",
			"--organization-physical-address.street-address", "123 Main St",
			"--organization-physical-address.extended-address", "Suite 400",
			"--primary-business-domain-sic-code", "primary_business_domain_sic_code",
			"--professional-license-number", "professional_license_number",
			"--website", "website",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"billing_address:\n" +
			"  administrative_area: Illinois\n" +
			"  city: Chicago\n" +
			"  country: United States\n" +
			"  postal_code: '60601'\n" +
			"  street_address: 123 Main St\n" +
			"  extended_address: Suite 400\n" +
			"billing_contact:\n" +
			"  email: billing@acme.com\n" +
			"  first_name: John\n" +
			"  last_name: Doe\n" +
			"  phone_number: '15551234568'\n" +
			"corporate_registration_number: corporate_registration_number\n" +
			"customer_reference: customer_reference\n" +
			"doing_business_as: doing_business_as\n" +
			"dun_bradstreet_number: dun_bradstreet_number\n" +
			"fein: fein\n" +
			"industry: industry\n" +
			"legal_name: xxx\n" +
			"number_of_employees: 1-10\n" +
			"organization_contact:\n" +
			"  email: jane.smith@acme.com\n" +
			"  first_name: Jane\n" +
			"  job_title: VP of Engineering\n" +
			"  last_name: Smith\n" +
			"  phone: '+16035551234'\n" +
			"organization_legal_type: corporation\n" +
			"organization_physical_address:\n" +
			"  administrative_area: Illinois\n" +
			"  city: Chicago\n" +
			"  country: United States\n" +
			"  postal_code: '60601'\n" +
			"  street_address: 123 Main St\n" +
			"  extended_address: Suite 400\n" +
			"primary_business_domain_sic_code: primary_business_domain_sic_code\n" +
			"professional_license_number: professional_license_number\n" +
			"website: website\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"enterprises", "update",
			"--enterprise-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestEnterprisesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"enterprises", "list",
			"--max-items", "10",
			"--legal-name", "Acme",
			"--page-number", "1",
			"--page-size", "1",
		)
	})
}

func TestEnterprisesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"enterprises", "delete",
			"--enterprise-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}
