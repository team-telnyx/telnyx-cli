// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestEnterprisesDirCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"enterprises:dir", "create",
			"--enterprise-id", "4a6192a4-573d-446d-b3ce-aff9117272a6",
			"--authorizer-email", "sam@acmeplumbing.example.com",
			"--authorizer-name", "Sam Owner",
			"--certify-brand-is-accurate=true",
			"--certify-ip-ownership=true",
			"--certify-no-shaft-content=true",
			"--display-name", "Acme Plumbing",
			"--call-reason", "Appointment reminders",
			"--call-reason", "Billing inquiries",
			"--document", "{document_id: 2a7e8337-e803-4057-a4ae-26c40eb0bc6c, document_type: business_registration, description: Certificate of incorporation.}",
			"--logo-url", "https://acmeplumbing.example.com/logo-256.bmp",
			"--reselling=false",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(enterprisesDirCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"enterprises:dir", "create",
			"--enterprise-id", "4a6192a4-573d-446d-b3ce-aff9117272a6",
			"--authorizer-email", "sam@acmeplumbing.example.com",
			"--authorizer-name", "Sam Owner",
			"--certify-brand-is-accurate=true",
			"--certify-ip-ownership=true",
			"--certify-no-shaft-content=true",
			"--display-name", "Acme Plumbing",
			"--call-reason", "Appointment reminders",
			"--call-reason", "Billing inquiries",
			"--document.document-id", "2a7e8337-e803-4057-a4ae-26c40eb0bc6c",
			"--document.document-type", "business_registration",
			"--document.description", "Certificate of incorporation.",
			"--logo-url", "https://acmeplumbing.example.com/logo-256.bmp",
			"--reselling=false",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"authorizer_email: sam@acmeplumbing.example.com\n" +
			"authorizer_name: Sam Owner\n" +
			"certify_brand_is_accurate: true\n" +
			"certify_ip_ownership: true\n" +
			"certify_no_shaft_content: true\n" +
			"display_name: Acme Plumbing\n" +
			"call_reasons:\n" +
			"  - Appointment reminders\n" +
			"  - Billing inquiries\n" +
			"documents:\n" +
			"  - document_id: 2a7e8337-e803-4057-a4ae-26c40eb0bc6c\n" +
			"    document_type: business_registration\n" +
			"    description: Certificate of incorporation.\n" +
			"logo_url: https://acmeplumbing.example.com/logo-256.bmp\n" +
			"reselling: false\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"enterprises:dir", "create",
			"--enterprise-id", "4a6192a4-573d-446d-b3ce-aff9117272a6",
		)
	})
}

func TestEnterprisesDirList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"enterprises:dir", "list",
			"--max-items", "10",
			"--enterprise-id", "4a6192a4-573d-446d-b3ce-aff9117272a6",
			"--filter-call-reason-contains", "filter[call_reason][contains]",
			"--filter-display-name-contains", "filter[display_name][contains]",
			"--filter-expiring-at-gte", "'2019-12-27T18:11:19.117Z'",
			"--filter-expiring-at-lte", "'2019-12-27T18:11:19.117Z'",
			"--filter-expiring-within-days", "1",
			"--filter-status", "draft",
			"--page-number", "1",
			"--page-size", "20",
			"--search", "search",
			"--sort", "created_at",
			"--status", "draft",
		)
	})
}
