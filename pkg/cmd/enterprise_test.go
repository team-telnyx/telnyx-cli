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
			"--billing-address", "{administrative_area: IL, city: Chicago, country: US, postal_code: '60601', street_address: 100 Main St, extended_address: Suite 504}",
			"--billing-contact", "{email: billing@run065.example.com, first_name: Alex, last_name: Bill, phone_number: '+13125550001'}",
			"--country-code", "US",
			"--doing-business-as", "Run 065 Debug",
			"--fein", "12-3456789",
			"--industry", "technology",
			"--jurisdiction-of-incorporation", "Delaware",
			"--legal-name", "Run 065 Debug Co",
			"--number-of-employees", "51-200",
			"--organization-contact", "{email: org@run065.example.com, first_name: Sam, job_title: Compliance Lead, last_name: Org, phone_number: '+13125550000'}",
			"--organization-legal-type", "llc",
			"--organization-physical-address", "{administrative_area: IL, city: Chicago, country: US, postal_code: '60601', street_address: 100 Main St, extended_address: Suite 504}",
			"--organization-type", "commercial",
			"--website", "https://run065.example.com",
			"--corporate-registration-number", "corporate_registration_number",
			"--customer-reference", "internal-id-12345",
			"--dun-bradstreet-number", "dun_bradstreet_number",
			"--primary-business-domain-sic-code", "primary_business_domain_sic_code",
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
			"--billing-address.administrative-area", "IL",
			"--billing-address.city", "Chicago",
			"--billing-address.country", "US",
			"--billing-address.postal-code", "60601",
			"--billing-address.street-address", "100 Main St",
			"--billing-address.extended-address", "Suite 504",
			"--billing-contact.email", "billing@run065.example.com",
			"--billing-contact.first-name", "Alex",
			"--billing-contact.last-name", "Bill",
			"--billing-contact.phone-number", "+13125550001",
			"--country-code", "US",
			"--doing-business-as", "Run 065 Debug",
			"--fein", "12-3456789",
			"--industry", "technology",
			"--jurisdiction-of-incorporation", "Delaware",
			"--legal-name", "Run 065 Debug Co",
			"--number-of-employees", "51-200",
			"--organization-contact.email", "org@run065.example.com",
			"--organization-contact.first-name", "Sam",
			"--organization-contact.job-title", "Compliance Lead",
			"--organization-contact.last-name", "Org",
			"--organization-contact.phone-number", "+13125550000",
			"--organization-legal-type", "llc",
			"--organization-physical-address.administrative-area", "IL",
			"--organization-physical-address.city", "Chicago",
			"--organization-physical-address.country", "US",
			"--organization-physical-address.postal-code", "60601",
			"--organization-physical-address.street-address", "100 Main St",
			"--organization-physical-address.extended-address", "Suite 504",
			"--organization-type", "commercial",
			"--website", "https://run065.example.com",
			"--corporate-registration-number", "corporate_registration_number",
			"--customer-reference", "internal-id-12345",
			"--dun-bradstreet-number", "dun_bradstreet_number",
			"--primary-business-domain-sic-code", "primary_business_domain_sic_code",
			"--professional-license-number", "professional_license_number",
			"--role-type", "enterprise",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"billing_address:\n" +
			"  administrative_area: IL\n" +
			"  city: Chicago\n" +
			"  country: US\n" +
			"  postal_code: '60601'\n" +
			"  street_address: 100 Main St\n" +
			"  extended_address: Suite 504\n" +
			"billing_contact:\n" +
			"  email: billing@run065.example.com\n" +
			"  first_name: Alex\n" +
			"  last_name: Bill\n" +
			"  phone_number: '+13125550001'\n" +
			"country_code: US\n" +
			"doing_business_as: Run 065 Debug\n" +
			"fein: 12-3456789\n" +
			"industry: technology\n" +
			"jurisdiction_of_incorporation: Delaware\n" +
			"legal_name: Run 065 Debug Co\n" +
			"number_of_employees: 51-200\n" +
			"organization_contact:\n" +
			"  email: org@run065.example.com\n" +
			"  first_name: Sam\n" +
			"  job_title: Compliance Lead\n" +
			"  last_name: Org\n" +
			"  phone_number: '+13125550000'\n" +
			"organization_legal_type: llc\n" +
			"organization_physical_address:\n" +
			"  administrative_area: IL\n" +
			"  city: Chicago\n" +
			"  country: US\n" +
			"  postal_code: '60601'\n" +
			"  street_address: 100 Main St\n" +
			"  extended_address: Suite 504\n" +
			"organization_type: commercial\n" +
			"website: https://run065.example.com\n" +
			"corporate_registration_number: corporate_registration_number\n" +
			"customer_reference: internal-id-12345\n" +
			"dun_bradstreet_number: dun_bradstreet_number\n" +
			"primary_business_domain_sic_code: primary_business_domain_sic_code\n" +
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
			"--enterprise-id", "4a6192a4-573d-446d-b3ce-aff9117272a6",
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
			"--enterprise-id", "4a6192a4-573d-446d-b3ce-aff9117272a6",
			"--billing-address", "{administrative_area: IL, city: Chicago, country: US, postal_code: '60601', street_address: 100 Main St, extended_address: Suite 504}",
			"--billing-contact", "{email: billing@acmeplumbing.example.com, first_name: Alex, last_name: Bill, phone_number: '+13125550001'}",
			"--corporate-registration-number", "corporate_registration_number",
			"--customer-reference", "internal-ref-2026Q2",
			"--doing-business-as", "Acme Plumbing",
			"--dun-bradstreet-number", "dun_bradstreet_number",
			"--fein", "12-3456789",
			"--industry", "business",
			"--jurisdiction-of-incorporation", "Delaware",
			"--legal-name", "Acme Plumbing LLC",
			"--number-of-employees", "51-200",
			"--organization-contact", "{email: sam@acmeplumbing.example.com, first_name: Sam, job_title: Compliance Lead, last_name: Owner, phone_number: '+13125550000'}",
			"--organization-legal-type", "llc",
			"--organization-physical-address", "{administrative_area: IL, city: Chicago, country: US, postal_code: '60601', street_address: 100 Main St, extended_address: Suite 504}",
			"--primary-business-domain-sic-code", "primary_business_domain_sic_code",
			"--professional-license-number", "professional_license_number",
			"--website", "https://acmeplumbing.example.com",
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
			"--enterprise-id", "4a6192a4-573d-446d-b3ce-aff9117272a6",
			"--billing-address.administrative-area", "IL",
			"--billing-address.city", "Chicago",
			"--billing-address.country", "US",
			"--billing-address.postal-code", "60601",
			"--billing-address.street-address", "100 Main St",
			"--billing-address.extended-address", "Suite 504",
			"--billing-contact.email", "billing@acmeplumbing.example.com",
			"--billing-contact.first-name", "Alex",
			"--billing-contact.last-name", "Bill",
			"--billing-contact.phone-number", "+13125550001",
			"--corporate-registration-number", "corporate_registration_number",
			"--customer-reference", "internal-ref-2026Q2",
			"--doing-business-as", "Acme Plumbing",
			"--dun-bradstreet-number", "dun_bradstreet_number",
			"--fein", "12-3456789",
			"--industry", "business",
			"--jurisdiction-of-incorporation", "Delaware",
			"--legal-name", "Acme Plumbing LLC",
			"--number-of-employees", "51-200",
			"--organization-contact.email", "sam@acmeplumbing.example.com",
			"--organization-contact.first-name", "Sam",
			"--organization-contact.job-title", "Compliance Lead",
			"--organization-contact.last-name", "Owner",
			"--organization-contact.phone-number", "+13125550000",
			"--organization-legal-type", "llc",
			"--organization-physical-address.administrative-area", "IL",
			"--organization-physical-address.city", "Chicago",
			"--organization-physical-address.country", "US",
			"--organization-physical-address.postal-code", "60601",
			"--organization-physical-address.street-address", "100 Main St",
			"--organization-physical-address.extended-address", "Suite 504",
			"--primary-business-domain-sic-code", "primary_business_domain_sic_code",
			"--professional-license-number", "professional_license_number",
			"--website", "https://acmeplumbing.example.com",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"billing_address:\n" +
			"  administrative_area: IL\n" +
			"  city: Chicago\n" +
			"  country: US\n" +
			"  postal_code: '60601'\n" +
			"  street_address: 100 Main St\n" +
			"  extended_address: Suite 504\n" +
			"billing_contact:\n" +
			"  email: billing@acmeplumbing.example.com\n" +
			"  first_name: Alex\n" +
			"  last_name: Bill\n" +
			"  phone_number: '+13125550001'\n" +
			"corporate_registration_number: corporate_registration_number\n" +
			"customer_reference: internal-ref-2026Q2\n" +
			"doing_business_as: Acme Plumbing\n" +
			"dun_bradstreet_number: dun_bradstreet_number\n" +
			"fein: 12-3456789\n" +
			"industry: business\n" +
			"jurisdiction_of_incorporation: Delaware\n" +
			"legal_name: Acme Plumbing LLC\n" +
			"number_of_employees: 51-200\n" +
			"organization_contact:\n" +
			"  email: sam@acmeplumbing.example.com\n" +
			"  first_name: Sam\n" +
			"  job_title: Compliance Lead\n" +
			"  last_name: Owner\n" +
			"  phone_number: '+13125550000'\n" +
			"organization_legal_type: llc\n" +
			"organization_physical_address:\n" +
			"  administrative_area: IL\n" +
			"  city: Chicago\n" +
			"  country: US\n" +
			"  postal_code: '60601'\n" +
			"  street_address: 100 Main St\n" +
			"  extended_address: Suite 504\n" +
			"primary_business_domain_sic_code: primary_business_domain_sic_code\n" +
			"professional_license_number: professional_license_number\n" +
			"website: https://acmeplumbing.example.com\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"enterprises", "update",
			"--enterprise-id", "4a6192a4-573d-446d-b3ce-aff9117272a6",
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
			"--filter-legal-name-contains", "Acme",
			"--legal-name", "Acme",
			"--page-number", "1",
			"--page-size", "10",
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
			"--enterprise-id", "4a6192a4-573d-446d-b3ce-aff9117272a6",
		)
	})
}

func TestEnterprisesActivateBrandedCalling(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"enterprises", "activate-branded-calling",
			"--enterprise-id", "4a6192a4-573d-446d-b3ce-aff9117272a6",
		)
	})
}
