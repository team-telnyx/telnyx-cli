/*
Copyright © Telnyx LLC

*/
package get

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/hashicorp/consul/api"
	"github.com/spf13/cobra"
	"github.com/team-telnyx/telnyx-cli/consul"
)

func init() {
	instanceCmd.Flags().StringP("env", "e", "prod", "The Consul environment to use")
	instanceCmd.Flags().StringP("datacenter", "d", "", "The Consul datacenter to use")

	getCmd.AddCommand(instanceCmd)
}

// instanceCmd represents the instances command
var instanceCmd = &cobra.Command{
	Use:     "instance <service>",
	Aliases: []string{"i"},
	Short:   "List the running instances of a given service",
	Long: `
Lists the running instances of a given service. Instances will be grouped by datacenter,
and will be printed following the format:

	[node-name][ipv4][version][canary|empty][health status]
	`,
	Args: cobra.ExactArgs(1),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		dc, _ := cmd.Flags().GetString("datacenter")
		env, _ := cmd.Flags().GetString("env")

		var svcs map[string][]string

		if dc != "" {
			svcs = consul.GetServicesByDc(dc)
		} else {
			svcs = consul.GetServicesByEnv(env)
		}

		var completions []string
		for s, _ := range svcs {
			if strings.HasPrefix(s, toComplete) {
				completions = append(completions, s)
			}
		}

		return completions, cobra.ShellCompDirectiveNoFileComp
	},
	Run: func(cmd *cobra.Command, args []string) {
		svc := args[0]
		env, _ := cmd.Flags().GetString("env")
		dc, _ := cmd.Flags().GetString("datacenter")

		printService(svc)

		if dc == "" {
			ists := consul.GetInstancesByEnv(env, svc)
			for _, ist := range ists {
				printDatacenter(ist.Dc)
				printInstances(ist.Instances)
			}
		} else {
			ists := consul.GetInstancesByDc(dc, svc)
			printDatacenter(dc)
			printInstances(ists)
		}
	},
}

func printInstances(svcs []*api.ServiceEntry) {
	// Group by version
	groupedInstances := make(map[string][]*api.ServiceEntry)
	for _, s := range svcs {
		grp := groupedInstances[instanceVersion(s)]
		groupedInstances[instanceVersion(s)] = append(grp, s)
	}

	for version, svcs := range groupedInstances {
		fmt.Printf("  • %d instances running version %s\n", len(svcs), version)
		for _, svc := range svcs {
			printInstance(svc)
		}
	}
}

func instanceVersion(svc *api.ServiceEntry) string {
	var version string

	if svc.Service.Meta["version"] != "" {
		version = svc.Service.Meta["version"]
	} else {
		version = consul.GetInstanceVersion(svc)
	}

	return version
}

func printService(svc string) {
	color.Magenta("▶ %s", svc)
}

func printInstance(svc *api.ServiceEntry) {
	version := instanceVersion(svc)
	healthStatus := svc.Checks.AggregatedStatus()
	healthColorFunc := colorByHealthStatus(healthStatus).SprintFunc()

	fmt.Printf(
		"    » [%s][%s:%s][%s]%s[%s]\n",
		color.CyanString(svc.Node.Node),
		color.CyanString(svc.Service.Address),
		color.CyanString(strconv.Itoa(svc.Service.Port)),
		color.CyanString(version),
		printTags(svc),
		healthColorFunc(svc.Checks.AggregatedStatus()),
	)
}

func colorByHealthStatus(status string) *color.Color {
	var c *color.Color
	switch status {
	case api.HealthAny:
		c = color.New(color.FgWhite)
	case api.HealthWarning:
		c = color.New(color.FgHiMagenta)
	case api.HealthCritical:
		c = color.New(color.FgRed)
	case api.HealthMaint:
		c = color.New(color.FgYellow)
	case api.HealthPassing:
		c = color.New(color.FgGreen)
	default:
		c = color.New(color.FgWhite)
	}

	return c
}

func printTags(svc *api.ServiceEntry) string {
	if isCanary(svc) {
		return fmt.Sprintf("[%s]", color.YellowString("canary"))
	}

	return ""
}

func isCanary(svc *api.ServiceEntry) bool {
	for _, tag := range svc.Service.Tags {
		if tag == "canary" {
			return true
		}
	}

	return false
}
