/*
Copyright © Telnyx LLC

*/
package cmd

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

	rootCmd.AddCommand(deploymentCmd)
}

// deploymentCmd represents the deployments command
var deploymentCmd = &cobra.Command{
	Use:     "deployments",
	Aliases: []string{"dp"},
	Args:    cobra.ExactArgs(1),
	Short:   "List latest successful production deployments from a service via Jenkins.",
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
		startDate, _ := cmd.Flags().GetString("startDate")
		endDate, _ := cmd.Flags().GetString("endDate")

		dps, err := metaservice.FetchDeployments(svc, startDate, endDate)
		if err != nil {
			cobra.CheckErr(err)
		}

		for _, dp := range dps {
			printDeployment(dp)
		}
	},
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
