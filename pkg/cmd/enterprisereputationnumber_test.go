// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestEnterprisesReputationNumbersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"enterprises:reputation:numbers", "retrieve",
			"--enterprise-id", "4a6192a4-573d-446d-b3ce-aff9117272a6",
			"--phone-number", "+19493253498",
			"--fresh=true",
		)
	})
}

func TestEnterprisesReputationNumbersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"enterprises:reputation:numbers", "list",
			"--max-items", "10",
			"--enterprise-id", "4a6192a4-573d-446d-b3ce-aff9117272a6",
			"--filter-phone-number-contains", "+16035551234",
			"--filter-phone-number-eq", "+16035551234",
			"--page-number", "1",
			"--page-size", "10",
			"--phone-number", "+16035551234",
		)
	})
}

func TestEnterprisesReputationNumbersAssociate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"enterprises:reputation:numbers", "associate",
			"--enterprise-id", "4a6192a4-573d-446d-b3ce-aff9117272a6",
			"--phone-number", "+19493253498",
			"--phone-number", "+12134445566",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"phone_numbers:\n" +
			"  - '+19493253498'\n" +
			"  - '+12134445566'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"enterprises:reputation:numbers", "associate",
			"--enterprise-id", "4a6192a4-573d-446d-b3ce-aff9117272a6",
		)
	})
}

func TestEnterprisesReputationNumbersDisassociate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"enterprises:reputation:numbers", "disassociate",
			"--enterprise-id", "4a6192a4-573d-446d-b3ce-aff9117272a6",
			"--phone-number", "+19493253498",
		)
	})
}

func TestEnterprisesReputationNumbersRefresh(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"enterprises:reputation:numbers", "refresh",
			"--enterprise-id", "4a6192a4-573d-446d-b3ce-aff9117272a6",
			"--phone-number", "+19493253498",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"phone_numbers:\n" +
			"  - '+19493253498'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"enterprises:reputation:numbers", "refresh",
			"--enterprise-id", "4a6192a4-573d-446d-b3ce-aff9117272a6",
		)
	})
}
