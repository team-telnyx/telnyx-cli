// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestDirReferencesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"dir:references", "update",
			"--dir-id", "16635d38-75a6-4481-82e8-69af60e05011",
			"--ref-type", "business",
			"--slot", "0",
			"--email", "dana.reyes@example.com",
			"--full-name", "Dana Reyes",
			"--job-title", "VP of Operations",
			"--organization", "Acme Logistics",
			"--phone-e164", "+14155550123",
			"--relationship-to-registrant", "Supplier",
			"--timezone", "America/New_York",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"email: dana.reyes@example.com\n" +
			"full_name: Dana Reyes\n" +
			"job_title: VP of Operations\n" +
			"organization: Acme Logistics\n" +
			"phone_e164: '+14155550123'\n" +
			"relationship_to_registrant: Supplier\n" +
			"timezone: America/New_York\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"dir:references", "update",
			"--dir-id", "16635d38-75a6-4481-82e8-69af60e05011",
			"--ref-type", "business",
			"--slot", "0",
		)
	})
}

func TestDirReferencesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"dir:references", "list",
			"--dir-id", "16635d38-75a6-4481-82e8-69af60e05011",
		)
	})
}

func TestDirReferencesSubmit(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"dir:references", "submit",
			"--dir-id", "16635d38-75a6-4481-82e8-69af60e05011",
			"--business-reference", "{email: dana.reyes@example.com, full_name: Dana Reyes, phone_e164: '+14155550123', timezone: America/New_York, job_title: VP of Operations, organization: Acme Logistics, relationship_to_registrant: Supplier}",
			"--business-reference", "{email: dana.reyes@example.com, full_name: Dana Reyes, phone_e164: '+14155550123', timezone: America/New_York, job_title: VP of Operations, organization: Acme Logistics, relationship_to_registrant: Supplier}",
			"--financial-reference", "{email: dana.reyes@example.com, full_name: Dana Reyes, phone_e164: '+14155550123', timezone: America/New_York, job_title: VP of Operations, organization: Acme Logistics, relationship_to_registrant: Supplier}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(dirReferencesSubmit)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"dir:references", "submit",
			"--dir-id", "16635d38-75a6-4481-82e8-69af60e05011",
			"--business-reference.email", "dana.reyes@example.com",
			"--business-reference.full-name", "Dana Reyes",
			"--business-reference.phone-e164", "+14155550123",
			"--business-reference.timezone", "America/New_York",
			"--business-reference.job-title", "VP of Operations",
			"--business-reference.organization", "Acme Logistics",
			"--business-reference.relationship-to-registrant", "Supplier",
			"--business-reference.email", "dana.reyes@example.com",
			"--business-reference.full-name", "Dana Reyes",
			"--business-reference.phone-e164", "+14155550123",
			"--business-reference.timezone", "America/New_York",
			"--business-reference.job-title", "VP of Operations",
			"--business-reference.organization", "Acme Logistics",
			"--business-reference.relationship-to-registrant", "Supplier",
			"--financial-reference.email", "dana.reyes@example.com",
			"--financial-reference.full-name", "Dana Reyes",
			"--financial-reference.phone-e164", "+14155550123",
			"--financial-reference.timezone", "America/New_York",
			"--financial-reference.job-title", "VP of Operations",
			"--financial-reference.organization", "Acme Logistics",
			"--financial-reference.relationship-to-registrant", "Supplier",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"business_references:\n" +
			"  - email: dana.reyes@example.com\n" +
			"    full_name: Dana Reyes\n" +
			"    phone_e164: '+14155550123'\n" +
			"    timezone: America/New_York\n" +
			"    job_title: VP of Operations\n" +
			"    organization: Acme Logistics\n" +
			"    relationship_to_registrant: Supplier\n" +
			"  - email: dana.reyes@example.com\n" +
			"    full_name: Dana Reyes\n" +
			"    phone_e164: '+14155550123'\n" +
			"    timezone: America/New_York\n" +
			"    job_title: VP of Operations\n" +
			"    organization: Acme Logistics\n" +
			"    relationship_to_registrant: Supplier\n" +
			"financial_reference:\n" +
			"  email: dana.reyes@example.com\n" +
			"  full_name: Dana Reyes\n" +
			"  phone_e164: '+14155550123'\n" +
			"  timezone: America/New_York\n" +
			"  job_title: VP of Operations\n" +
			"  organization: Acme Logistics\n" +
			"  relationship_to_registrant: Supplier\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"dir:references", "submit",
			"--dir-id", "16635d38-75a6-4481-82e8-69af60e05011",
		)
	})
}
