// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestEmailDedicatedIPsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email:dedicated-ips", "create",
			"--idempotency-key", "8e03978e-40d5-43e8-bc93-6894a57f9326",
			"--ip-version", "ipv4",
			"--region", "ewr-nj1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"ip_version: ipv4\n" +
			"region: ewr-nj1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"email:dedicated-ips", "create",
			"--idempotency-key", "8e03978e-40d5-43e8-bc93-6894a57f9326",
		)
	})
}

func TestEmailDedicatedIPsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email:dedicated-ips", "retrieve",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestEmailDedicatedIPsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email:dedicated-ips", "update",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--drain=true",
			"--force=true",
			"--pool-id", "9b1c4d70-abcd-4e5f-8a9b-0c1d2e3f4a5b",
			"--warmup", "{paused: true}",
			"--warmup-paused=true",
			"--weight", "100",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(emailDedicatedIPsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email:dedicated-ips", "update",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--drain=true",
			"--force=true",
			"--pool-id", "9b1c4d70-abcd-4e5f-8a9b-0c1d2e3f4a5b",
			"--warmup.paused=true",
			"--warmup-paused=true",
			"--weight", "100",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"drain: true\n" +
			"force: true\n" +
			"pool_id: 9b1c4d70-abcd-4e5f-8a9b-0c1d2e3f4a5b\n" +
			"warmup:\n" +
			"  paused: true\n" +
			"warmup_paused: true\n" +
			"weight: 100\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"email:dedicated-ips", "update",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestEmailDedicatedIPsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email:dedicated-ips", "list",
			"--filter-pool-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--filter-region", "filter[region]",
			"--filter-status", "provisioning",
		)
	})
}

func TestEmailDedicatedIPsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email:dedicated-ips", "delete",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--drain=true",
			"--force=true",
			"--reason", "reason",
			"--drain=true",
			"--force=true",
			"--reason", "No longer needed",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"drain: true\n" +
			"force: true\n" +
			"reason: No longer needed\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"email:dedicated-ips", "delete",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--drain=true",
			"--force=true",
			"--reason", "reason",
		)
	})
}
