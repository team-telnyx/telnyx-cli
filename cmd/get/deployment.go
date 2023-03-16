/*
Copyright © Telnyx LLC

*/
package get

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/team-telnyx/telnyx-cli/consul"
	"github.com/team-telnyx/telnyx-cli/metaservice"
)

func init() {
	defaultStartDate := time.Now().UTC().AddDate(0, 0, -7)
	defaultEndDate := time.Now().UTC()

	deploymentCmd.Flags().StringP("startDate", "s", formatTime(defaultStartDate), "The starting date of the deployments to check. Format: YYYY-MM-DD.")
	deploymentCmd.Flags().StringP("endDate", "e", formatTime(defaultEndDate), "The end date of the deployments to check. Format: YYYY-MM-DD.")
	deploymentCmd.Flags().Bool("last", false, "List the latest successful deployment from the given service.")

	getCmd.AddCommand(deploymentCmd)
}

// deploymentCmd represents the deployment command
var deploymentCmd = &cobra.Command{
	Use:     "deployment <service>",
	Aliases: []string{"d"},
	Args:    cobra.ExactArgs(1),
	Short:   "List successful production deployments from a service via Jenkins.",
	Long: `
Lists the deployments via Jenkins of the given service. Defaults to listing the
deployments from the last 7 days. You can define both the 'startDate' and 'endDate'
of the deployments, with the maximum window being 7 days.
Both 'startDate' and 'endDate' are optional, and when missing, the CLI will consider
a 7 days window based on the given parameter.

Examples:

# List the successful deployments from past 7 days for call-control service
telnyx-cli get deployment call-control

# List the successful deployments from call-control service during the first week of Jan/2023
telnyx-cli get deployment call-control --startDate 2023-01-01

# List the successful deployments from call-control service in the week ending on 2023-01-01
telnyx-cli get deployment call-control --endDate 2023-01-01
  `,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		svcs := consul.GetServicesByEnv("prod")

		var completions []string
		for s, _ := range svcs {
			if strings.HasPrefix(s, toComplete) {
				completions = append(completions, s)
			}
		}

		return completions, cobra.ShellCompDirectiveNoFileComp
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		s, _ := cmd.Flags().GetString("startDate")
		e, _ := cmd.Flags().GetString("endDate")

		startDate, err := time.Parse("2006-01-02", s)
		if err != nil {
			cobra.CheckErr(err)
		}

		endDate, err := time.Parse("2006-01-02", e)
		if err != nil {
			cobra.CheckErr(err)
		}

		if endDate.Before(startDate) {
			cobra.CheckErr("endDate happens before startDate")
		}

		difference := endDate.Sub(startDate)
		if difference.Hours() > 168 {
			cobra.CheckErr("cannot exceed 1 week")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		svc := args[0]
		last, _ := cmd.Flags().GetBool("last")
		startDate, _ := cmd.Flags().GetString("startDate")
		endDate, _ := cmd.Flags().GetString("endDate")

		if last {
			dp, err := getLastDeployment(svc)
			if err != nil {
				cobra.CheckErr(err)
			}
			printDeployment(*dp)
		}

		dps, err := metaservice.FetchDeployments(svc, startDate, endDate)
		if err != nil {
			cobra.CheckErr(err)
		}

		for _, dp := range dps {
			printDeployment(dp)
		}
	},
}

// getLastDeployment returns the latest deployment from the given service
// It will look for the latest deployment in the last 20 week, and return
// the latest deployment found
func getLastDeployment(svc string) (*metaservice.Deployment, error) {
	for i := 0; i < 20; i++ {
		fmt.Printf("Looking for deployments in the last %d weeks...\n", i+1)

		startDate := time.Now().UTC().AddDate(0, 0, -7*(i+1)).Format("2006-01-02")
		endDate := time.Now().UTC().AddDate(0, 0, -7*i).Format("2006-01-02")
		dps, err := metaservice.FetchDeployments(svc, startDate, endDate)
		if err != nil {
			return nil, err
		}
		if len(dps) > 0 {
			return &dps[0], nil
		}
	}

	return nil, fmt.Errorf("no deployment found for service %s in the last 20 weeks", svc)
}

func formatTime(t time.Time) string {
	return t.Format("2006-01-02")
}

func printDeployment(dp metaservice.Deployment) {
	fmt.Printf(
		"\n> Started at: %s\n"+
			"Version: %s\n"+
			"Started by: %s\n"+
			"Title: %s\n"+
			"Description: %s\n"+
			"Enable K8s single region cluster: %s\n"+
			"K8s cluster: %s\n"+
			"Host Group: %s\n"+
			"------------------------------------\n",
		color.CyanString(dp.StartTimeStamp),
		color.CyanString(dp.Version),
		color.CyanString(dp.StartedByUser),
		color.CyanString(dp.BuildParameters.ChangeTitle),
		color.CyanString(dp.BuildParameters.Description),
		color.CyanString(dp.BuildParameters.EnableSingleRegionKubernetesCluster),
		color.CyanString(dp.BuildParameters.SingleRegionKubernetesCluster),
		color.CyanString(dp.BuildParameters.HostGroup),
	)
}
