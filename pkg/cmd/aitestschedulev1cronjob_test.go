// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAITestsSchedulesV1CronJobsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tests:schedules:v1:cron-jobs", "retrieve",
			"--id", "0",
		)
	})
}

func TestAITestsSchedulesV1CronJobsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tests:schedules:v1:cron-jobs", "update",
			"--id", "0",
			"--agent", "0",
			"--crontab-expression", "0 4 * * *",
			"--execution-mode", "execution_mode",
			"--frequency", "0",
			"--is-active=true",
			"--name", "Nightly regression - updated",
			"--notify-on", "both",
			"--personality", "0",
			"--run-as-text=true",
			"--scenario", "0",
			"--tag", "string",
			"--timezone", "timezone",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"agent: 0\n" +
			"crontab_expression: 0 4 * * *\n" +
			"execution_mode: execution_mode\n" +
			"frequency: 0\n" +
			"is_active: true\n" +
			"name: Nightly regression - updated\n" +
			"notify_on: both\n" +
			"personalities:\n" +
			"  - 0\n" +
			"run_as_text: true\n" +
			"scenarios:\n" +
			"  - 0\n" +
			"tags:\n" +
			"  - string\n" +
			"timezone: timezone\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:tests:schedules:v1:cron-jobs", "update",
			"--id", "0",
		)
	})
}

func TestAITestsSchedulesV1CronJobsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tests:schedules:v1:cron-jobs", "delete",
			"--id", "0",
		)
	})
}

func TestAITestsSchedulesV1CronJobsCronJobs(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tests:schedules:v1:cron-jobs", "cron-jobs",
			"--crontab-expression", "0 3 * * *",
			"--agent", "1259",
			"--execution-mode", "voice",
			"--frequency", "0",
			"--is-active=true",
			"--name", "Nightly regression",
			"--notify-on", "failure",
			"--personality", "0",
			"--run-as-text=true",
			"--scenario", "9512",
			"--tag", "string",
			"--timezone", "America/Chicago",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"crontab_expression: 0 3 * * *\n" +
			"agent: 1259\n" +
			"execution_mode: voice\n" +
			"frequency: 0\n" +
			"is_active: true\n" +
			"name: Nightly regression\n" +
			"notify_on: failure\n" +
			"personalities:\n" +
			"  - 0\n" +
			"run_as_text: true\n" +
			"scenarios:\n" +
			"  - 9512\n" +
			"tags:\n" +
			"  - string\n" +
			"timezone: America/Chicago\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:tests:schedules:v1:cron-jobs", "cron-jobs",
		)
	})
}

func TestAITestsSchedulesV1CronJobsRetrieveCronJobs(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tests:schedules:v1:cron-jobs", "retrieve-cron-jobs",
			"--agent-id", "0",
		)
	})
}
