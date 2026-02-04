// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestMessaging10dlcBrandCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-10dlc:brand", "create",
		"--country", "US",
		"--display-name", "ABC Mobile",
		"--email", "email",
		"--entity-type", "PRIVATE_PROFIT",
		"--vertical", "TECHNOLOGY",
		"--business-contact-email", "name@example.com",
		"--city", "New York",
		"--company-name", "ABC Inc.",
		"--ein", "111111111",
		"--first-name", "John",
		"--ip-address", "ipAddress",
		"--is-reseller=true",
		"--last-name", "Smith",
		"--mobile-phone", "+12024567890",
		"--mock=true",
		"--phone", "+12024567890",
		"--postal-code", "10001",
		"--state", "NY",
		"--stock-exchange", "NASDAQ",
		"--stock-symbol", "ABC",
		"--street", "123",
		"--webhook-failover-url", "https://webhook.com/9010a453-4df8-4be6-a551-1070892888d6",
		"--webhook-url", "https://webhook.com/67ea78a8-9f32-4d04-b62d-f9502e8e5f93",
		"--website", "http://www.abcmobile.com",
	)
}

func TestMessaging10dlcBrandRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-10dlc:brand", "retrieve",
		"--brand-id", "brandId",
	)
}

func TestMessaging10dlcBrandUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-10dlc:brand", "update",
		"--brand-id", "brandId",
		"--country", "US",
		"--display-name", "ABC Mobile",
		"--email", "email",
		"--entity-type", "PRIVATE_PROFIT",
		"--vertical", "TECHNOLOGY",
		"--alt-business-id", "altBusiness_id",
		"--alt-business-id-type", "NONE",
		"--business-contact-email", "name@example.com",
		"--city", "New York",
		"--company-name", "ABC Inc.",
		"--ein", "111111111",
		"--first-name", "John",
		"--identity-status", "VERIFIED",
		"--ip-address", "ipAddress",
		"--is-reseller=true",
		"--last-name", "Smith",
		"--phone", "+12024567890",
		"--postal-code", "10001",
		"--state", "NY",
		"--stock-exchange", "NASDAQ",
		"--stock-symbol", "ABC",
		"--street", "123",
		"--webhook-failover-url", "https://webhook.com/9010a453-4df8-4be6-a551-1070892888d6",
		"--webhook-url", "https://webhook.com/67ea78a8-9f32-4d04-b62d-f9502e8e5f93",
		"--website", "http://www.abcmobile.com",
	)
}

func TestMessaging10dlcBrandList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-10dlc:brand", "list",
		"--brand-id", "826ef77a-348c-445b-81a5-a9b13c68fbfe",
		"--country", "country",
		"--display-name", "displayName",
		"--entity-type", "entityType",
		"--page", "1",
		"--records-per-page", "0",
		"--sort", "assignedCampaignsCount",
		"--state", "state",
		"--tcr-brand-id", "BBAND1",
	)
}

func TestMessaging10dlcBrandDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-10dlc:brand", "delete",
		"--brand-id", "brandId",
	)
}

func TestMessaging10dlcBrandGetFeedback(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-10dlc:brand", "get-feedback",
		"--brand-id", "brandId",
	)
}

func TestMessaging10dlcBrandResend2faEmail(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-10dlc:brand", "resend-2fa-email",
		"--brand-id", "brandId",
	)
}

func TestMessaging10dlcBrandRetrieveSMSOtpStatus(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-10dlc:brand", "retrieve-sms-otp-status",
		"--brand-id", "4b20019b-043a-78f8-0657-b3be3f4b4002",
	)
}

func TestMessaging10dlcBrandRevet(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-10dlc:brand", "revet",
		"--brand-id", "brandId",
	)
}

func TestMessaging10dlcBrandTriggerSMSOtp(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-10dlc:brand", "trigger-sms-otp",
		"--brand-id", "4b20019b-043a-78f8-0657-b3be3f4b4002",
		"--pin-sms", "Your PIN is @OTP_PIN@",
		"--success-sms", "Verification successful!",
	)
}

func TestMessaging10dlcBrandVerifySMSOtp(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-10dlc:brand", "verify-sms-otp",
		"--brand-id", "4b20019b-043a-78f8-0657-b3be3f4b4002",
		"--otp-pin", "123456",
	)
}
