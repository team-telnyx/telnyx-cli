// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestPortingLoaConfigurationsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting:loa-configurations", "create",
		"--address", "{city: Austin, country_code: US, state: TX, street_address: 600 Congress Avenue, zip_code: '78701', extended_address: 14th Floor}",
		"--company-name", "Telnyx",
		"--contact", "{email: testing@telnyx.com, phone_number: '+12003270001'}",
		"--logo", "{document_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e}",
		"--name", "My LOA Configuration",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(portingLoaConfigurationsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting:loa-configurations", "create",
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
}

func TestPortingLoaConfigurationsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting:loa-configurations", "retrieve",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestPortingLoaConfigurationsUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting:loa-configurations", "update",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--address", "{city: Austin, country_code: US, state: TX, street_address: 600 Congress Avenue, zip_code: '78701', extended_address: 14th Floor}",
		"--company-name", "Telnyx",
		"--contact", "{email: testing@telnyx.com, phone_number: '+12003270001'}",
		"--logo", "{document_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e}",
		"--name", "My LOA Configuration",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(portingLoaConfigurationsUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting:loa-configurations", "update",
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
}

func TestPortingLoaConfigurationsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting:loa-configurations", "list",
		"--page", "{number: 1, size: 1}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(portingLoaConfigurationsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting:loa-configurations", "list",
		"--page.number", "1",
		"--page.size", "1",
	)
}

func TestPortingLoaConfigurationsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting:loa-configurations", "delete",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}
