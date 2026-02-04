// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestReportsListMdrs(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"reports", "list-mdrs",
		"--id", "e093fbe0-5bde-11eb-ae93-0242ac130002",
		"--cld", "+15551237654",
		"--cli", "+15551237654",
		"--direction", "INBOUND",
		"--end-date", "end_date",
		"--message-type", "SMS",
		"--profile", "My profile",
		"--start-date", "start_date",
		"--status", "DELIVERED",
	)
}

func TestReportsListWdrs(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"reports", "list-wdrs",
		"--id", "e093fbe0-5bde-11eb-ae93-0242ac130002",
		"--end-date", "2021-06-01T00:00:00Z",
		"--imsi", "123456",
		"--mcc", "204",
		"--mnc", "01",
		"--page-number", "0",
		"--page-size", "0",
		"--phone-number", "+12345678910",
		"--sim-card-id", "877f80a6-e5b2-4687-9a04-88076265720f",
		"--sim-group-id", "f05a189f-7c46-4531-ac56-1460dc465a42",
		"--sim-group-name", "sim name",
		"--sort", "created_at",
		"--start-date", "2021-05-01T00:00:00Z",
	)
}
