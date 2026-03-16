// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestPhoneNumbersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"phone-numbers", "retrieve",
			"--id", "1293384261075731499",
		)
	})
}

func TestPhoneNumbersUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"phone-numbers", "update",
			"--phone-number-id", "1293384261075731499",
			"--address-id", "dc8f39ac-953d-4520-b93b-786ae87db0da",
			"--billing-group-id", "dc8e4d67-33a0-4cbb-af74-7b58f05bd494",
			"--connection-id", "dc8e4d67-33a0-4cbb-af74-7b58f05bd494",
			"--customer-reference", "customer-reference",
			"--external-pin", "1234",
			"--hd-voice-enabled=true",
			"--tag", "tag",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"address_id: dc8f39ac-953d-4520-b93b-786ae87db0da\n" +
			"billing_group_id: dc8e4d67-33a0-4cbb-af74-7b58f05bd494\n" +
			"connection_id: dc8e4d67-33a0-4cbb-af74-7b58f05bd494\n" +
			"customer_reference: customer-reference\n" +
			"external_pin: '1234'\n" +
			"hd_voice_enabled: true\n" +
			"tags:\n" +
			"  - tag\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"phone-numbers", "update",
			"--phone-number-id", "1293384261075731499",
		)
	})
}

func TestPhoneNumbersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"phone-numbers", "list",
			"--max-items", "10",
			"--filter", "{billing_group_id: 62e4bf2e-c278-4282-b524-488d9c9c43b2, connection_id: '1521916448077776306', country_iso_alpha2: US, customer_reference: customer_reference, emergency_address_id: '9102160989215728032', number_type: {eq: local}, phone_number: phone_number, source: ported, status: active, tag: tag, voice.connection_name: {contains: test, ends_with: test, eq: test, starts_with: test}, voice.usage_payment_method: channel, without_tags: 'true'}",
			"--handle-messaging-profile-error", "false",
			"--page-number", "0",
			"--page-size", "0",
			"--sort", "connection_name",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(phoneNumbersList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"phone-numbers", "list",
			"--max-items", "10",
			"--filter.billing-group-id", "62e4bf2e-c278-4282-b524-488d9c9c43b2",
			"--filter.connection-id", "1521916448077776306",
			"--filter.country-iso-alpha2", "US",
			"--filter.customer-reference", "customer_reference",
			"--filter.emergency-address-id", "9102160989215728032",
			"--filter.number-type", "{eq: local}",
			"--filter.phone-number", "phone_number",
			"--filter.source", "ported",
			"--filter.status", "active",
			"--filter.tag", "tag",
			"--filter.voice-connection-name", "{contains: test, ends_with: test, eq: test, starts_with: test}",
			"--filter.voice-usage-payment-method", "channel",
			"--filter.without-tags", "true",
			"--handle-messaging-profile-error", "false",
			"--page-number", "0",
			"--page-size", "0",
			"--sort", "connection_name",
		)
	})
}

func TestPhoneNumbersDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"phone-numbers", "delete",
			"--id", "1293384261075731499",
		)
	})
}

func TestPhoneNumbersSlimList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"phone-numbers", "slim-list",
			"--max-items", "10",
			"--filter", "{billing_group_id: 62e4bf2e-c278-4282-b524-488d9c9c43b2, connection_id: '1521916448077776306', country_iso_alpha2: US, customer_reference: customer_reference, emergency_address_id: '9102160989215728032', number_type: {eq: local}, phone_number: phone_number, source: ported, status: active, tag: tag, voice.connection_name: {contains: test, ends_with: test, eq: test, starts_with: test}, voice.usage_payment_method: channel}",
			"--include-connection=true",
			"--include-tags=true",
			"--page-number", "0",
			"--page-size", "0",
			"--sort", "connection_name",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(phoneNumbersSlimList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"phone-numbers", "slim-list",
			"--max-items", "10",
			"--filter.billing-group-id", "62e4bf2e-c278-4282-b524-488d9c9c43b2",
			"--filter.connection-id", "1521916448077776306",
			"--filter.country-iso-alpha2", "US",
			"--filter.customer-reference", "customer_reference",
			"--filter.emergency-address-id", "9102160989215728032",
			"--filter.number-type", "{eq: local}",
			"--filter.phone-number", "phone_number",
			"--filter.source", "ported",
			"--filter.status", "active",
			"--filter.tag", "tag",
			"--filter.voice-connection-name", "{contains: test, ends_with: test, eq: test, starts_with: test}",
			"--filter.voice-usage-payment-method", "channel",
			"--include-connection=true",
			"--include-tags=true",
			"--page-number", "0",
			"--page-size", "0",
			"--sort", "connection_name",
		)
	})
}
