// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestPortingOrdersPhoneNumberConfigurationsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders:phone-number-configurations", "create",
		"--phone-number-configuration", "{porting_phone_number_id: 927f4687-318c-44bc-9f2f-22a5898143a4, user_bundle_id: ff901545-3e27-462a-ba9d-2b34654cab82}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(portingOrdersPhoneNumberConfigurationsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders:phone-number-configurations", "create",
		"--phone-number-configuration.porting-phone-number-id", "927f4687-318c-44bc-9f2f-22a5898143a4",
		"--phone-number-configuration.user-bundle-id", "ff901545-3e27-462a-ba9d-2b34654cab82",
	)
}

func TestPortingOrdersPhoneNumberConfigurationsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders:phone-number-configurations", "list",
		"--filter", "{porting_order: {status: [in-process]}, porting_phone_number: [5d6f7ede-1961-4717-bfb5-db392c5efc2d], user_bundle_id: [5d6f7ede-1961-4717-bfb5-db392c5efc2d]}",
		"--page-number", "0",
		"--page-size", "0",
		"--sort", "{value: created_at}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(portingOrdersPhoneNumberConfigurationsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders:phone-number-configurations", "list",
		"--filter.porting-order", "{status: [in-process]}",
		"--filter.porting-phone-number", "[5d6f7ede-1961-4717-bfb5-db392c5efc2d]",
		"--filter.user-bundle-id", "[5d6f7ede-1961-4717-bfb5-db392c5efc2d]",
		"--page-number", "0",
		"--page-size", "0",
		"--sort.value", "created_at",
	)
}
