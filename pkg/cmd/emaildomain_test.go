// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestEmailDomainsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-domains", "create",
			"--domain", "example.com",
			"--dmarc-policy", "{p: none, pct: 0, rua: rua, sp: none}",
			"--inbound-enabled=true",
			"--tracking", "{click_tracking: true, open_tracking: true, unsubscribe_tracking: false}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(emailDomainsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-domains", "create",
			"--domain", "example.com",
			"--dmarc-policy.p", "none",
			"--dmarc-policy.pct", "0",
			"--dmarc-policy.rua", "rua",
			"--dmarc-policy.sp", "none",
			"--inbound-enabled=true",
			"--tracking.click-tracking=true",
			"--tracking.open-tracking=true",
			"--tracking.unsubscribe-tracking=false",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"domain: example.com\n" +
			"dmarc_policy:\n" +
			"  p: none\n" +
			"  pct: 0\n" +
			"  rua: rua\n" +
			"  sp: none\n" +
			"inbound_enabled: true\n" +
			"tracking:\n" +
			"  click_tracking: true\n" +
			"  open_tracking: true\n" +
			"  unsubscribe_tracking: false\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"email-domains", "create",
		)
	})
}

func TestEmailDomainsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-domains", "retrieve",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestEmailDomainsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-domains", "update",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--dmarc-policy", "{p: none, pct: 0, rua: rua, sp: none}",
			"--inbound-enabled=true",
			"--tracking", "{click_tracking: true, open_tracking: false, unsubscribe_tracking: true}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(emailDomainsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-domains", "update",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--dmarc-policy.p", "none",
			"--dmarc-policy.pct", "0",
			"--dmarc-policy.rua", "rua",
			"--dmarc-policy.sp", "none",
			"--inbound-enabled=true",
			"--tracking.click-tracking=true",
			"--tracking.open-tracking=false",
			"--tracking.unsubscribe-tracking=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"dmarc_policy:\n" +
			"  p: none\n" +
			"  pct: 0\n" +
			"  rua: rua\n" +
			"  sp: none\n" +
			"inbound_enabled: true\n" +
			"tracking:\n" +
			"  click_tracking: true\n" +
			"  open_tracking: false\n" +
			"  unsubscribe_tracking: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"email-domains", "update",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestEmailDomainsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-domains", "list",
			"--max-items", "10",
			"--filter-domain", "filter[domain]",
			"--filter-profile-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--filter-status", "pending",
			"--filter-type", "custom",
			"--filter-usable-for-inbound=true",
			"--filter-usable-for-sending=true",
			"--page-after", "page[after]",
			"--page-before", "page[before]",
			"--page-number", "1",
			"--page-size", "1",
			"--sort", "created_at",
		)
	})
}

func TestEmailDomainsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-domains", "delete",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--force=true",
		)
	})
}

func TestEmailDomainsRetrieveDNSRecords(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-domains", "retrieve-dns-records",
			"--domain-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestEmailDomainsRetrieveHealth(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-domains", "retrieve-health",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestEmailDomainsVerify(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email-domains", "verify",
			"--domain-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
