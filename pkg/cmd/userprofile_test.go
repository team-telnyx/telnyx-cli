// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestUserProfileList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"user-profile", "list",
		)
	})
}

func TestUserProfilePatchAll(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"user-profile", "patch-all",
			"--address-1", "123 Example Street",
			"--address-2", "Suite 100",
			"--business-name", "Example Company",
			"--city", "Chicago",
			"--company-size", "11-50",
			"--country", "US",
			"--data-locality", "DEU",
			"--default-managed-account-allow-custom-pricing=false",
			"--default-managed-account-inherit-emergency-terms-of-service=false",
			"--default-managed-account-rollup-billing=false",
			"--default-notify-manager-account=true",
			"--default-numbers-external-pin", "1234",
			"--first-name", "Alex",
			"--got-free-number-at-signup=false",
			"--last-name", "Taylor",
			"--low-balance-notification-frequency", "12",
			"--low-balance-threshold-amount", "10",
			"--notify-on-invoice-available=true",
			"--notify-on-invoice-payment=false",
			"--npac-user=false",
			"--onboarding-build-method", "API",
			"--onboarding-industry", "software",
			"--onboarding-monthly-budget", "1000-5000",
			"--onboarding-product-interest", "voice",
			"--onboarding-project-type", "contact center",
			"--onboarding-reasons-for-consideration", "API integration",
			"--onboarding-referrer", "search engine",
			"--onboarding-use-cases-other", "customer support",
			"--phone-number", "+13125550100",
			"--port-out-pin-validation=false",
			"--primary-use-case", "configure sip trunks",
			"--role", "account manager",
			"--state", "IL",
			"--ui-theme", "1",
			"--use-case", "voice",
			"--user-subscription-opt-in=false",
			"--zip", "60601",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"address_1: 123 Example Street\n" +
			"address_2: Suite 100\n" +
			"business_name: Example Company\n" +
			"city: Chicago\n" +
			"company_size: 11-50\n" +
			"country: US\n" +
			"data_locality: DEU\n" +
			"default_managed_account_allow_custom_pricing: false\n" +
			"default_managed_account_inherit_emergency_terms_of_service: false\n" +
			"default_managed_account_rollup_billing: false\n" +
			"default_notify_manager_account: true\n" +
			"default_numbers_external_pin: '1234'\n" +
			"first_name: Alex\n" +
			"got_free_number_at_signup: false\n" +
			"last_name: Taylor\n" +
			"low_balance_notification_frequency: 12\n" +
			"low_balance_threshold_amount: 10\n" +
			"notify_on_invoice_available: true\n" +
			"notify_on_invoice_payment: false\n" +
			"npac_user: false\n" +
			"onboarding_build_method:\n" +
			"  - API\n" +
			"onboarding_industries:\n" +
			"  - software\n" +
			"onboarding_monthly_budget:\n" +
			"  - 1000-5000\n" +
			"onboarding_product_interest:\n" +
			"  - voice\n" +
			"onboarding_project_types:\n" +
			"  - contact center\n" +
			"onboarding_reasons_for_consideration:\n" +
			"  - API integration\n" +
			"onboarding_referrer:\n" +
			"  - search engine\n" +
			"onboarding_use_cases_other:\n" +
			"  - customer support\n" +
			"phone_number: '+13125550100'\n" +
			"port_out_pin_validation: false\n" +
			"primary_use_case: configure sip trunks\n" +
			"role: account manager\n" +
			"state: IL\n" +
			"ui_theme: 1\n" +
			"use_cases:\n" +
			"  - voice\n" +
			"user_subscription_opt_in: false\n" +
			"zip: '60601'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"user-profile", "patch-all",
		)
	})
}
