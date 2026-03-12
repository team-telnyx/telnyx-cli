// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestPortingOrdersPhoneNumberConfigurationsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "porting-orders:phone-number-configurations", "create",
			"--api-key", "string",
			"--phone-number-configuration", "{porting_phone_number_id: 927f4687-318c-44bc-9f2f-22a5898143a4, user_bundle_id: ff901545-3e27-462a-ba9d-2b34654cab82}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(portingOrdersPhoneNumberConfigurationsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "porting-orders:phone-number-configurations", "create",
			"--api-key", "string",
			"--phone-number-configuration.porting-phone-number-id", "927f4687-318c-44bc-9f2f-22a5898143a4",
			"--phone-number-configuration.user-bundle-id", "ff901545-3e27-462a-ba9d-2b34654cab82",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"phone_number_configurations:\n" +
			"  - porting_phone_number_id: 927f4687-318c-44bc-9f2f-22a5898143a4\n" +
			"    user_bundle_id: ff901545-3e27-462a-ba9d-2b34654cab82\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "porting-orders:phone-number-configurations", "create",
			"--api-key", "string",
		)
	})
}

func TestPortingOrdersPhoneNumberConfigurationsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "porting-orders:phone-number-configurations", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--filter", "{porting_order: {status: [in-process]}, porting_phone_number: [5d6f7ede-1961-4717-bfb5-db392c5efc2d], user_bundle_id: [5d6f7ede-1961-4717-bfb5-db392c5efc2d]}",
			"--page-number", "0",
			"--page-size", "0",
			"--sort", "{value: created_at}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(portingOrdersPhoneNumberConfigurationsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "porting-orders:phone-number-configurations", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--filter.porting-order", "{status: [in-process]}",
			"--filter.porting-phone-number", "[5d6f7ede-1961-4717-bfb5-db392c5efc2d]",
			"--filter.user-bundle-id", "[5d6f7ede-1961-4717-bfb5-db392c5efc2d]",
			"--page-number", "0",
			"--page-size", "0",
			"--sort.value", "created_at",
		)
	})
}
