// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestPortingOrdersCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders", "create",
		"--phone-number", "+13035550000",
		"--phone-number", "+13035550001",
		"--phone-number", "+13035550002",
		"--customer-group-reference", "Group-456",
		"--customer-reference", "Acct 123abc",
	)
}

func TestPortingOrdersRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders", "retrieve",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--include-phone-numbers=true",
	)
}

func TestPortingOrdersUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders", "update",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--activation-settings", "{foc_datetime_requested: '2021-03-19T10:07:15.527Z'}",
		"--customer-group-reference", "customer_group_reference",
		"--customer-reference", "customer_reference",
		"--documents", "{invoice: ce74b771-d23d-4960-81ec-8741b3862146, loa: 64ffb720-04c7-455b-92d6-20fcca92e935}",
		"--end-user", "{admin: {account_number: 123abc, auth_person_name: Porter McPortersen II, billing_phone_number: '13035551234', business_identifier: abc123, entity_name: Porter McPortersen, pin_passcode: '1234', tax_identifier: 1234abcd}, location: {administrative_area: TX, country_code: US, extended_address: 14th Floor, locality: Austin, postal_code: '78701', street_address: 600 Congress Avenue}}",
		"--messaging", "{enable_messaging: true}",
		"--misc", "{new_billing_phone_number: new_billing_phone_number, remaining_numbers_action: disconnect, type: full}",
		"--phone-number-configuration", "{billing_group_id: f1486bae-f067-460c-ad43-73a92848f902, connection_id: f1486bae-f067-460c-ad43-73a92848f902, emergency_address_id: f1486bae-f067-460c-ad43-73a92848f902, messaging_profile_id: f1486bae-f067-460c-ad43-73a92848f901, tags: [abc, '123']}",
		"--requirement-group-id", "DE748D99-06FA-4D90-9F9A-F4B62696BADA",
		"--requirement", "{field_value: 9787fb5f-cbe5-4de4-b765-3303774ee9fe, requirement_type_id: 59b0762a-b274-4f76-ac32-4d5cf0272e66}",
		"--user-feedback", "{user_comment: I loved my experience porting numbers with Telnyx, user_rating: 5}",
		"--webhook-url", "https://example.com",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(portingOrdersUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders", "update",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--activation-settings.foc-datetime-requested", "2021-03-19T10:07:15.527Z",
		"--customer-group-reference", "customer_group_reference",
		"--customer-reference", "customer_reference",
		"--documents.invoice", "ce74b771-d23d-4960-81ec-8741b3862146",
		"--documents.loa", "64ffb720-04c7-455b-92d6-20fcca92e935",
		"--end-user.admin", "{account_number: 123abc, auth_person_name: Porter McPortersen II, billing_phone_number: '13035551234', business_identifier: abc123, entity_name: Porter McPortersen, pin_passcode: '1234', tax_identifier: 1234abcd}",
		"--end-user.location", "{administrative_area: TX, country_code: US, extended_address: 14th Floor, locality: Austin, postal_code: '78701', street_address: 600 Congress Avenue}",
		"--messaging.enable-messaging=true",
		"--misc.new-billing-phone-number", "new_billing_phone_number",
		"--misc.remaining-numbers-action", "disconnect",
		"--misc.type", "full",
		"--phone-number-configuration.billing-group-id", "f1486bae-f067-460c-ad43-73a92848f902",
		"--phone-number-configuration.connection-id", "f1486bae-f067-460c-ad43-73a92848f902",
		"--phone-number-configuration.emergency-address-id", "f1486bae-f067-460c-ad43-73a92848f902",
		"--phone-number-configuration.messaging-profile-id", "f1486bae-f067-460c-ad43-73a92848f901",
		"--phone-number-configuration.tags", "[abc, '123']",
		"--requirement-group-id", "DE748D99-06FA-4D90-9F9A-F4B62696BADA",
		"--requirement.field-value", "9787fb5f-cbe5-4de4-b765-3303774ee9fe",
		"--requirement.requirement-type-id", "59b0762a-b274-4f76-ac32-4d5cf0272e66",
		"--user-feedback.user-comment", "I loved my experience porting numbers with Telnyx",
		"--user-feedback.user-rating", "5",
		"--webhook-url", "https://example.com",
	)
}

func TestPortingOrdersList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders", "list",
		"--filter", "{activation_settings: {fast_port_eligible: true, foc_datetime_requested: {gt: '2021-03-25T10:00:00.000Z', lt: '2021-03-25T10:00:00.000Z'}}, customer_group_reference: customer_group_reference, customer_reference: customer_reference, end_user: {admin: {auth_person_name: auth_person_name, entity_name: entity_name}}, misc: {type: full}, parent_support_key: parent_support_key, phone_numbers: {carrier_name: carrier_name, country_code: country_code, phone_number: {contains: contains}}}",
		"--include-phone-numbers=true",
		"--page-number", "0",
		"--page-size", "0",
		"--sort", "{value: created_at}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(portingOrdersList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders", "list",
		"--filter.activation-settings", "{fast_port_eligible: true, foc_datetime_requested: {gt: '2021-03-25T10:00:00.000Z', lt: '2021-03-25T10:00:00.000Z'}}",
		"--filter.customer-group-reference", "customer_group_reference",
		"--filter.customer-reference", "customer_reference",
		"--filter.end-user", "{admin: {auth_person_name: auth_person_name, entity_name: entity_name}}",
		"--filter.misc", "{type: full}",
		"--filter.parent-support-key", "parent_support_key",
		"--filter.phone-numbers", "{carrier_name: carrier_name, country_code: country_code, phone_number: {contains: contains}}",
		"--include-phone-numbers=true",
		"--page-number", "0",
		"--page-size", "0",
		"--sort.value", "created_at",
	)
}

func TestPortingOrdersDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders", "delete",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestPortingOrdersRetrieveAllowedFocWindows(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders", "retrieve-allowed-foc-windows",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestPortingOrdersRetrieveExceptionTypes(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders", "retrieve-exception-types",
	)
}

func TestPortingOrdersRetrieveRequirements(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders", "retrieve-requirements",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--page-number", "0",
		"--page-size", "0",
	)
}

func TestPortingOrdersRetrieveSubRequest(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders", "retrieve-sub-request",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}
