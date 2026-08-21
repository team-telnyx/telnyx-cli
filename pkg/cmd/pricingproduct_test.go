// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestPricingProductsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pricing:products", "retrieve",
			"--max-items", "10",
			"--slug", "slug",
			"--filter-country-iso", "SE",
			"--page-number", "1",
			"--page-size", "1",
		)
	})
}

func TestPricingProductsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pricing:products", "list",
			"--max-items", "10",
			"--page-number", "1",
			"--page-size", "1",
		)
	})
}
