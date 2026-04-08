// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestEnterprisesReputationRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"enterprises:reputation", "retrieve",
			"--enterprise-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestEnterprisesReputationDisable(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"enterprises:reputation", "disable",
			"--enterprise-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestEnterprisesReputationEnable(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"enterprises:reputation", "enable",
			"--enterprise-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--loa-document-id", "doc_01HXYZ1234ABCDEF",
			"--check-frequency", "business_daily",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"loa_document_id: doc_01HXYZ1234ABCDEF\n" +
			"check_frequency: business_daily\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"enterprises:reputation", "enable",
			"--enterprise-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestEnterprisesReputationUpdateFrequency(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"enterprises:reputation", "update-frequency",
			"--enterprise-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--check-frequency", "business_daily",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("check_frequency: business_daily")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"enterprises:reputation", "update-frequency",
			"--enterprise-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}
