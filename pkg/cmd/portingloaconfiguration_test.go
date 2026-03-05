// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestPortingLoaConfigurationsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "porting:loa-configurations", "create",
			"--api-key", "string",
			"--address", "{city: Austin, country_code: US, state: TX, street_address: 600 Congress Avenue, zip_code: '78701', extended_address: 14th Floor}",
			"--company-name", "Telnyx",
			"--contact", "{email: testing@telnyx.com, phone_number: '+12003270001'}",
			"--logo", "{document_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e}",
			"--name", "My LOA Configuration",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(portingLoaConfigurationsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "porting:loa-configurations", "create",
			"--api-key", "string",
			"--address.city", "Austin",
			"--address.country-code", "US",
			"--address.state", "TX",
			"--address.street-address", "600 Congress Avenue",
			"--address.zip-code", "78701",
			"--address.extended-address", "14th Floor",
			"--company-name", "Telnyx",
			"--contact.email", "testing@telnyx.com",
			"--contact.phone-number", "+12003270001",
			"--logo.document-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--name", "My LOA Configuration",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"address:\n" +
			"  city: Austin\n" +
			"  country_code: US\n" +
			"  state: TX\n" +
			"  street_address: 600 Congress Avenue\n" +
			"  zip_code: '78701'\n" +
			"  extended_address: 14th Floor\n" +
			"company_name: Telnyx\n" +
			"contact:\n" +
			"  email: testing@telnyx.com\n" +
			"  phone_number: '+12003270001'\n" +
			"logo:\n" +
			"  document_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"name: My LOA Configuration\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "porting:loa-configurations", "create",
			"--api-key", "string",
		)
	})
}

func TestPortingLoaConfigurationsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "porting:loa-configurations", "retrieve",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPortingLoaConfigurationsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "porting:loa-configurations", "update",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--address", "{city: Austin, country_code: US, state: TX, street_address: 600 Congress Avenue, zip_code: '78701', extended_address: 14th Floor}",
			"--company-name", "Telnyx",
			"--contact", "{email: testing@telnyx.com, phone_number: '+12003270001'}",
			"--logo", "{document_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e}",
			"--name", "My LOA Configuration",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(portingLoaConfigurationsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "porting:loa-configurations", "update",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--address.city", "Austin",
			"--address.country-code", "US",
			"--address.state", "TX",
			"--address.street-address", "600 Congress Avenue",
			"--address.zip-code", "78701",
			"--address.extended-address", "14th Floor",
			"--company-name", "Telnyx",
			"--contact.email", "testing@telnyx.com",
			"--contact.phone-number", "+12003270001",
			"--logo.document-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--name", "My LOA Configuration",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"address:\n" +
			"  city: Austin\n" +
			"  country_code: US\n" +
			"  state: TX\n" +
			"  street_address: 600 Congress Avenue\n" +
			"  zip_code: '78701'\n" +
			"  extended_address: 14th Floor\n" +
			"company_name: Telnyx\n" +
			"contact:\n" +
			"  email: testing@telnyx.com\n" +
			"  phone_number: '+12003270001'\n" +
			"logo:\n" +
			"  document_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"name: My LOA Configuration\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "porting:loa-configurations", "update",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPortingLoaConfigurationsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "porting:loa-configurations", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

func TestPortingLoaConfigurationsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "porting:loa-configurations", "delete",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPortingLoaConfigurationsPreview0(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "porting:loa-configurations", "preview-0",
			"--api-key", "string",
			"--address", "{city: Austin, country_code: US, state: TX, street_address: 600 Congress Avenue, zip_code: '78701', extended_address: 14th Floor}",
			"--company-name", "Telnyx",
			"--contact", "{email: testing@telnyx.com, phone_number: '+12003270001'}",
			"--logo", "{document_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e}",
			"--name", "My LOA Configuration",
			"--output", "/dev/null",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(portingLoaConfigurationsPreview0)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "porting:loa-configurations", "preview-0",
			"--api-key", "string",
			"--address.city", "Austin",
			"--address.country-code", "US",
			"--address.state", "TX",
			"--address.street-address", "600 Congress Avenue",
			"--address.zip-code", "78701",
			"--address.extended-address", "14th Floor",
			"--company-name", "Telnyx",
			"--contact.email", "testing@telnyx.com",
			"--contact.phone-number", "+12003270001",
			"--logo.document-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--name", "My LOA Configuration",
			"--output", "/dev/null",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"address:\n" +
			"  city: Austin\n" +
			"  country_code: US\n" +
			"  state: TX\n" +
			"  street_address: 600 Congress Avenue\n" +
			"  zip_code: '78701'\n" +
			"  extended_address: 14th Floor\n" +
			"company_name: Telnyx\n" +
			"contact:\n" +
			"  email: testing@telnyx.com\n" +
			"  phone_number: '+12003270001'\n" +
			"logo:\n" +
			"  document_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"name: My LOA Configuration\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "porting:loa-configurations", "preview-0",
			"--api-key", "string",
			"--output", "/dev/null",
		)
	})
}

func TestPortingLoaConfigurationsPreview1(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "porting:loa-configurations", "preview-1",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--output", "/dev/null",
		)
	})
}
