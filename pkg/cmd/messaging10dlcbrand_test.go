// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestMessaging10dlcBrandCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "messaging-10dlc:brand", "create",
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"country: US\n" +
			"displayName: ABC Mobile\n" +
			"email: email\n" +
			"entityType: PRIVATE_PROFIT\n" +
			"vertical: TECHNOLOGY\n" +
			"businessContactEmail: name@example.com\n" +
			"city: New York\n" +
			"companyName: ABC Inc.\n" +
			"ein: '111111111'\n" +
			"firstName: John\n" +
			"ipAddress: ipAddress\n" +
			"isReseller: true\n" +
			"lastName: Smith\n" +
			"mobilePhone: '+12024567890'\n" +
			"mock: true\n" +
			"phone: '+12024567890'\n" +
			"postalCode: '10001'\n" +
			"state: NY\n" +
			"stockExchange: NASDAQ\n" +
			"stockSymbol: ABC\n" +
			"street: '123'\n" +
			"webhookFailoverURL: https://webhook.com/9010a453-4df8-4be6-a551-1070892888d6\n" +
			"webhookURL: https://webhook.com/67ea78a8-9f32-4d04-b62d-f9502e8e5f93\n" +
			"website: http://www.abcmobile.com\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "messaging-10dlc:brand", "create",
			"--api-key", "string",
		)
	})
}

func TestMessaging10dlcBrandRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "messaging-10dlc:brand", "retrieve",
			"--api-key", "string",
			"--brand-id", "brandId",
		)
	})
}

func TestMessaging10dlcBrandUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "messaging-10dlc:brand", "update",
			"--api-key", "string",
			"--brand-id", "brandId",
			"--country", "US",
			"--display-name", "ABC Mobile",
			"--email", "email",
			"--entity-type", "PRIVATE_PROFIT",
			"--vertical", "TECHNOLOGY",
			"--alt-business-id", "altBusinessId",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"country: US\n" +
			"displayName: ABC Mobile\n" +
			"email: email\n" +
			"entityType: PRIVATE_PROFIT\n" +
			"vertical: TECHNOLOGY\n" +
			"altBusinessId: altBusinessId\n" +
			"altBusinessIdType: NONE\n" +
			"businessContactEmail: name@example.com\n" +
			"city: New York\n" +
			"companyName: ABC Inc.\n" +
			"ein: '111111111'\n" +
			"firstName: John\n" +
			"identityStatus: VERIFIED\n" +
			"ipAddress: ipAddress\n" +
			"isReseller: true\n" +
			"lastName: Smith\n" +
			"phone: '+12024567890'\n" +
			"postalCode: '10001'\n" +
			"state: NY\n" +
			"stockExchange: NASDAQ\n" +
			"stockSymbol: ABC\n" +
			"street: '123'\n" +
			"webhookFailoverURL: https://webhook.com/9010a453-4df8-4be6-a551-1070892888d6\n" +
			"webhookURL: https://webhook.com/67ea78a8-9f32-4d04-b62d-f9502e8e5f93\n" +
			"website: http://www.abcmobile.com\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "messaging-10dlc:brand", "update",
			"--api-key", "string",
			"--brand-id", "brandId",
		)
	})
}

func TestMessaging10dlcBrandList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "messaging-10dlc:brand", "list",
			"--api-key", "string",
			"--max-items", "10",
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
	})
}

func TestMessaging10dlcBrandDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "messaging-10dlc:brand", "delete",
			"--api-key", "string",
			"--brand-id", "brandId",
		)
	})
}

func TestMessaging10dlcBrandGetFeedback(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "messaging-10dlc:brand", "get-feedback",
			"--api-key", "string",
			"--brand-id", "brandId",
		)
	})
}

func TestMessaging10dlcBrandGetSMSOtpByReference(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "messaging-10dlc:brand", "get-sms-otp-by-reference",
			"--api-key", "string",
			"--reference-id", "OTP4B2001",
			"--brand-id", "B123ABC",
		)
	})
}

func TestMessaging10dlcBrandResend2faEmail(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "messaging-10dlc:brand", "resend-2fa-email",
			"--api-key", "string",
			"--brand-id", "brandId",
		)
	})
}

func TestMessaging10dlcBrandRetrieveSMSOtpStatus(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "messaging-10dlc:brand", "retrieve-sms-otp-status",
			"--api-key", "string",
			"--brand-id", "4b20019b-043a-78f8-0657-b3be3f4b4002",
		)
	})
}

func TestMessaging10dlcBrandRevet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "messaging-10dlc:brand", "revet",
			"--api-key", "string",
			"--brand-id", "brandId",
		)
	})
}

func TestMessaging10dlcBrandTriggerSMSOtp(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "messaging-10dlc:brand", "trigger-sms-otp",
			"--api-key", "string",
			"--brand-id", "4b20019b-043a-78f8-0657-b3be3f4b4002",
			"--pin-sms", "Your PIN is @OTP_PIN@",
			"--success-sms", "Verification successful!",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"pinSms: Your PIN is @OTP_PIN@\n" +
			"successSms: Verification successful!\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "messaging-10dlc:brand", "trigger-sms-otp",
			"--api-key", "string",
			"--brand-id", "4b20019b-043a-78f8-0657-b3be3f4b4002",
		)
	})
}

func TestMessaging10dlcBrandVerifySMSOtp(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "messaging-10dlc:brand", "verify-sms-otp",
			"--api-key", "string",
			"--brand-id", "4b20019b-043a-78f8-0657-b3be3f4b4002",
			"--otp-pin", "123456",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("otpPin: '123456'")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "messaging-10dlc:brand", "verify-sms-otp",
			"--api-key", "string",
			"--brand-id", "4b20019b-043a-78f8-0657-b3be3f4b4002",
		)
	})
}
