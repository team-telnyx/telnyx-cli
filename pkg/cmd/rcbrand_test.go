// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestRcsBrandsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"rcs:brands", "create",
			"--addresses", "{primary: {administrative_area: IL, city: Chicago, country_code: US, line_1: 1 Main Street, postal_code: '60601', line_2: x}}",
			"--contacts", "{brand: {contact_type: BRAND, email: jane@example.com, first_name: Jane, last_name: Doe, phone_number: '+13125550100', title: Messaging Operations Manager}}",
			"--display-name", "Acme",
			"--identifiers", "{ein: {identifier_type: EIN, value: 12-3456789}, stock_symbol: {identifier_type: STOCK_SYMBOL, value: J!Q0Ok0bzJb7:pro}}",
			"--legal-entity-type", "LIMITED_LIABILITY_COMPANY",
			"--legal-name", "Acme LLC",
			"--organization-type", "PRIVATE_PROFIT",
			"--website-url", "https://www.example.com",
			"--profile-id", "40000000-0000-0000-0000-000000000001",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(rcsBrandsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"rcs:brands", "create",
			"--addresses", "{primary: {administrative_area: IL, city: Chicago, country_code: US, line_1: 1 Main Street, postal_code: '60601', line_2: x}}",
			"--contacts.brand", "{contact_type: BRAND, email: jane@example.com, first_name: Jane, last_name: Doe, phone_number: '+13125550100', title: Messaging Operations Manager}",
			"--display-name", "Acme",
			"--identifiers.ein", "{identifier_type: EIN, value: 12-3456789}",
			"--identifiers.stock-symbol", "{identifier_type: STOCK_SYMBOL, value: J!Q0Ok0bzJb7:pro}",
			"--legal-entity-type", "LIMITED_LIABILITY_COMPANY",
			"--legal-name", "Acme LLC",
			"--organization-type", "PRIVATE_PROFIT",
			"--website-url", "https://www.example.com",
			"--profile-id", "40000000-0000-0000-0000-000000000001",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"addresses:\n" +
			"  primary:\n" +
			"    administrative_area: IL\n" +
			"    city: Chicago\n" +
			"    country_code: US\n" +
			"    line_1: 1 Main Street\n" +
			"    postal_code: '60601'\n" +
			"    line_2: x\n" +
			"contacts:\n" +
			"  brand:\n" +
			"    contact_type: BRAND\n" +
			"    email: jane@example.com\n" +
			"    first_name: Jane\n" +
			"    last_name: Doe\n" +
			"    phone_number: '+13125550100'\n" +
			"    title: Messaging Operations Manager\n" +
			"display_name: Acme\n" +
			"identifiers:\n" +
			"  ein:\n" +
			"    identifier_type: EIN\n" +
			"    value: 12-3456789\n" +
			"  stock_symbol:\n" +
			"    identifier_type: STOCK_SYMBOL\n" +
			"    value: J!Q0Ok0bzJb7:pro\n" +
			"legal_entity_type: LIMITED_LIABILITY_COMPANY\n" +
			"legal_name: Acme LLC\n" +
			"organization_type: PRIVATE_PROFIT\n" +
			"website_url: https://www.example.com\n" +
			"profile_id: 40000000-0000-0000-0000-000000000001\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"rcs:brands", "create",
		)
	})
}

func TestRcsBrandsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"rcs:brands", "retrieve",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestRcsBrandsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"rcs:brands", "update",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--addresses", "{foo: {administrative_area: x, city: x, country_code: SE, line_1: x, postal_code: x, line_2: x}}",
			"--contacts", "{brand: {contact_type: BRAND, email: dev@stainless.com, first_name: x, last_name: x, phone_number: '+49605132', title: x}}",
			"--display-name", "Acme Communications",
			"--identifiers", "{ein: {identifier_type: EIN, value: 29-1051329}, stock_symbol: {identifier_type: STOCK_SYMBOL, value: J!Q0Ok0bzJb7:pro}}",
			"--legal-entity-type", "LIMITED_LIABILITY_COMPANY",
			"--legal-name", "x",
			"--organization-type", "PRIVATE_PROFIT",
			"--profile-id", "profile_id",
			"--website-url", "https://example.com",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(rcsBrandsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"rcs:brands", "update",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--addresses", "{foo: {administrative_area: x, city: x, country_code: SE, line_1: x, postal_code: x, line_2: x}}",
			"--contacts.brand", "{contact_type: BRAND, email: dev@stainless.com, first_name: x, last_name: x, phone_number: '+49605132', title: x}",
			"--display-name", "Acme Communications",
			"--identifiers.ein", "{identifier_type: EIN, value: 29-1051329}",
			"--identifiers.stock-symbol", "{identifier_type: STOCK_SYMBOL, value: J!Q0Ok0bzJb7:pro}",
			"--legal-entity-type", "LIMITED_LIABILITY_COMPANY",
			"--legal-name", "x",
			"--organization-type", "PRIVATE_PROFIT",
			"--profile-id", "profile_id",
			"--website-url", "https://example.com",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"addresses:\n" +
			"  foo:\n" +
			"    administrative_area: x\n" +
			"    city: x\n" +
			"    country_code: SE\n" +
			"    line_1: x\n" +
			"    postal_code: x\n" +
			"    line_2: x\n" +
			"contacts:\n" +
			"  brand:\n" +
			"    contact_type: BRAND\n" +
			"    email: dev@stainless.com\n" +
			"    first_name: x\n" +
			"    last_name: x\n" +
			"    phone_number: '+49605132'\n" +
			"    title: x\n" +
			"display_name: Acme Communications\n" +
			"identifiers:\n" +
			"  ein:\n" +
			"    identifier_type: EIN\n" +
			"    value: 29-1051329\n" +
			"  stock_symbol:\n" +
			"    identifier_type: STOCK_SYMBOL\n" +
			"    value: J!Q0Ok0bzJb7:pro\n" +
			"legal_entity_type: LIMITED_LIABILITY_COMPANY\n" +
			"legal_name: x\n" +
			"organization_type: PRIVATE_PROFIT\n" +
			"profile_id: profile_id\n" +
			"website_url: https://example.com\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"rcs:brands", "update",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestRcsBrandsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"rcs:brands", "list",
		)
	})
}

func TestRcsBrandsSubmit(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"rcs:brands", "submit",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
