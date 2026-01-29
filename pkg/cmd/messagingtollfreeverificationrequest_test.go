// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestMessagingTollfreeVerificationRequestsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-tollfree:verification:requests", "create",
		"--additional-information", "additionalInformation",
		"--business-addr1", "600 Congress Avenue",
		"--business-city", "Austin",
		"--business-contact-email", "email@example.com",
		"--business-contact-first-name", "John",
		"--business-contact-last-name", "Doe",
		"--business-contact-phone", "+18005550100",
		"--business-name", "Telnyx LLC",
		"--business-state", "Texas",
		"--business-zip", "78701",
		"--corporate-website", "http://example.com",
		"--isv-reseller", "isvReseller",
		"--message-volume", "100,000",
		"--opt-in-workflow", "User signs into the Telnyx portal, enters a number and is prompted to select whether they want to use 2FA verification for security purposes. If they've opted in a confirmation message is sent out to the handset",
		"--opt-in-workflow-image-url", "{url: https://telnyx.com/sign-up}",
		"--opt-in-workflow-image-url", "{url: https://telnyx.com/company/data-privacy}",
		"--phone-number", "{phoneNumber: '+18773554398'}",
		"--phone-number", "{phoneNumber: '+18773554399'}",
		"--production-message-content", "Your Telnyx OTP is XXXX",
		"--use-case", "2FA",
		"--use-case-summary", "This is a use case where Telnyx sends out 2FA codes to portal users to verify their identity in order to sign into the portal",
		"--age-gated-content=true",
		"--business-addr2", "14th Floor",
		"--business-registration-country", "US",
		"--business-registration-number", "12-3456789",
		"--business-registration-type", "EIN",
		"--campaign-verify-authorization-token", "cv_token_abc123xyz",
		"--doing-business-as", "Acme Services",
		"--entity-type", "SOLE_PROPRIETOR",
		"--help-message-response", "Reply HELP for assistance or STOP to unsubscribe. Contact: support@example.com",
		"--opt-in-confirmation-response", "You have successfully opted in to receive messages from Acme Corp",
		"--opt-in-keywords", "START, YES, SUBSCRIBE",
		"--privacy-policy-url", "https://example.com/privacy",
		"--terms-and-condition-url", "https://example.com/terms",
		"--webhook-url", "http://example-webhook.com",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(messagingTollfreeVerificationRequestsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-tollfree:verification:requests", "create",
		"--additional-information", "additionalInformation",
		"--business-addr1", "600 Congress Avenue",
		"--business-city", "Austin",
		"--business-contact-email", "email@example.com",
		"--business-contact-first-name", "John",
		"--business-contact-last-name", "Doe",
		"--business-contact-phone", "+18005550100",
		"--business-name", "Telnyx LLC",
		"--business-state", "Texas",
		"--business-zip", "78701",
		"--corporate-website", "http://example.com",
		"--isv-reseller", "isvReseller",
		"--message-volume", "100,000",
		"--opt-in-workflow", "User signs into the Telnyx portal, enters a number and is prompted to select whether they want to use 2FA verification for security purposes. If they've opted in a confirmation message is sent out to the handset",
		"--opt-in-workflow-image-url.url", "https://telnyx.com/sign-up",
		"--opt-in-workflow-image-url.url", "https://telnyx.com/company/data-privacy",
		"--phone-number.phone-number", "+18773554398",
		"--phone-number.phone-number", "+18773554399",
		"--production-message-content", "Your Telnyx OTP is XXXX",
		"--use-case", "2FA",
		"--use-case-summary", "This is a use case where Telnyx sends out 2FA codes to portal users to verify their identity in order to sign into the portal",
		"--age-gated-content=true",
		"--business-addr2", "14th Floor",
		"--business-registration-country", "US",
		"--business-registration-number", "12-3456789",
		"--business-registration-type", "EIN",
		"--campaign-verify-authorization-token", "cv_token_abc123xyz",
		"--doing-business-as", "Acme Services",
		"--entity-type", "SOLE_PROPRIETOR",
		"--help-message-response", "Reply HELP for assistance or STOP to unsubscribe. Contact: support@example.com",
		"--opt-in-confirmation-response", "You have successfully opted in to receive messages from Acme Corp",
		"--opt-in-keywords", "START, YES, SUBSCRIBE",
		"--privacy-policy-url", "https://example.com/privacy",
		"--terms-and-condition-url", "https://example.com/terms",
		"--webhook-url", "http://example-webhook.com",
	)
}

func TestMessagingTollfreeVerificationRequestsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-tollfree:verification:requests", "retrieve",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestMessagingTollfreeVerificationRequestsUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-tollfree:verification:requests", "update",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--additional-information", "additionalInformation",
		"--business-addr1", "600 Congress Avenue",
		"--business-city", "Austin",
		"--business-contact-email", "email@example.com",
		"--business-contact-first-name", "John",
		"--business-contact-last-name", "Doe",
		"--business-contact-phone", "+18005550100",
		"--business-name", "Telnyx LLC",
		"--business-state", "Texas",
		"--business-zip", "78701",
		"--corporate-website", "http://example.com",
		"--isv-reseller", "isvReseller",
		"--message-volume", "100,000",
		"--opt-in-workflow", "User signs into the Telnyx portal, enters a number and is prompted to select whether they want to use 2FA verification for security purposes. If they've opted in a confirmation message is sent out to the handset",
		"--opt-in-workflow-image-url", "{url: https://telnyx.com/sign-up}",
		"--opt-in-workflow-image-url", "{url: https://telnyx.com/company/data-privacy}",
		"--phone-number", "{phoneNumber: '+18773554398'}",
		"--phone-number", "{phoneNumber: '+18773554399'}",
		"--production-message-content", "Your Telnyx OTP is XXXX",
		"--use-case", "2FA",
		"--use-case-summary", "This is a use case where Telnyx sends out 2FA codes to portal users to verify their identity in order to sign into the portal",
		"--age-gated-content=true",
		"--business-addr2", "14th Floor",
		"--business-registration-country", "US",
		"--business-registration-number", "12-3456789",
		"--business-registration-type", "EIN",
		"--campaign-verify-authorization-token", "cv_token_abc123xyz",
		"--doing-business-as", "Acme Services",
		"--entity-type", "SOLE_PROPRIETOR",
		"--help-message-response", "Reply HELP for assistance or STOP to unsubscribe. Contact: support@example.com",
		"--opt-in-confirmation-response", "You have successfully opted in to receive messages from Acme Corp",
		"--opt-in-keywords", "START, YES, SUBSCRIBE",
		"--privacy-policy-url", "https://example.com/privacy",
		"--terms-and-condition-url", "https://example.com/terms",
		"--webhook-url", "http://example-webhook.com",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(messagingTollfreeVerificationRequestsUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-tollfree:verification:requests", "update",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--additional-information", "additionalInformation",
		"--business-addr1", "600 Congress Avenue",
		"--business-city", "Austin",
		"--business-contact-email", "email@example.com",
		"--business-contact-first-name", "John",
		"--business-contact-last-name", "Doe",
		"--business-contact-phone", "+18005550100",
		"--business-name", "Telnyx LLC",
		"--business-state", "Texas",
		"--business-zip", "78701",
		"--corporate-website", "http://example.com",
		"--isv-reseller", "isvReseller",
		"--message-volume", "100,000",
		"--opt-in-workflow", "User signs into the Telnyx portal, enters a number and is prompted to select whether they want to use 2FA verification for security purposes. If they've opted in a confirmation message is sent out to the handset",
		"--opt-in-workflow-image-url.url", "https://telnyx.com/sign-up",
		"--opt-in-workflow-image-url.url", "https://telnyx.com/company/data-privacy",
		"--phone-number.phone-number", "+18773554398",
		"--phone-number.phone-number", "+18773554399",
		"--production-message-content", "Your Telnyx OTP is XXXX",
		"--use-case", "2FA",
		"--use-case-summary", "This is a use case where Telnyx sends out 2FA codes to portal users to verify their identity in order to sign into the portal",
		"--age-gated-content=true",
		"--business-addr2", "14th Floor",
		"--business-registration-country", "US",
		"--business-registration-number", "12-3456789",
		"--business-registration-type", "EIN",
		"--campaign-verify-authorization-token", "cv_token_abc123xyz",
		"--doing-business-as", "Acme Services",
		"--entity-type", "SOLE_PROPRIETOR",
		"--help-message-response", "Reply HELP for assistance or STOP to unsubscribe. Contact: support@example.com",
		"--opt-in-confirmation-response", "You have successfully opted in to receive messages from Acme Corp",
		"--opt-in-keywords", "START, YES, SUBSCRIBE",
		"--privacy-policy-url", "https://example.com/privacy",
		"--terms-and-condition-url", "https://example.com/terms",
		"--webhook-url", "http://example-webhook.com",
	)
}

func TestMessagingTollfreeVerificationRequestsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-tollfree:verification:requests", "list",
		"--page", "1",
		"--page-size", "1",
		"--date-end", "2019-12-27T18:11:19.117Z",
		"--date-start", "2019-12-27T18:11:19.117Z",
		"--phone-number", "phone_number",
		"--status", "Verified",
	)
}

func TestMessagingTollfreeVerificationRequestsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-tollfree:verification:requests", "delete",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}
