/*
Copyright © Telnyx LLC

*/
package cmd

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
	instanceCmd.Flags().StringP("env", "e", "dev", "The Consul environment to use")
	instanceCmd.Flags().StringP("datacenter", "d", "", "The Consul datacenter to use")

	rootCmd.AddCommand(instanceCmd)
}

// instanceCmd represents the instances command
var instanceCmd = &cobra.Command{
	Use:     "instances service",
	Aliases: []string{"ist"},
	Short:   "List the instances from a service in Consul",
	Args:    cobra.ExactArgs(1),
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

func printInstances(svcs []*api.CatalogService) {
	// Group by version
	groupedInstances := make(map[string][]*api.CatalogService)
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

func instanceVersion(svc *api.CatalogService) string {
	var version string
	if svc.ServiceMeta["version"] != "" {
		version = "v" + svc.ServiceMeta["version"]
	} else {
		version = svc.ServiceTags[len(svc.ServiceTags)-1]
	}

	return version
}

func printService(svc string) {
	color.Magenta("▶ %s", svc)
}

func printInstance(svc *api.CatalogService) {
	var version string
	if svc.ServiceMeta["version"] != "" {
		version = "v" + svc.ServiceMeta["version"]
	} else {
		version = svc.ServiceTags[len(svc.ServiceTags)-1]
	}

	healthStatus := svc.Checks.AggregatedStatus()
	healthColor := colorByHealthStatus(healthStatus).SprintFunc()

	fmt.Printf(
		"    » [%s][%s:%s][%s]%s[%s]\n",
		color.CyanString(svc.Node),
		color.CyanString(svc.ServiceAddress),
		color.CyanString(strconv.Itoa(svc.ServicePort)),
		color.CyanString(version),
		printTags(svc),
		healthColor(svc.Checks.AggregatedStatus()),
	)
}

func colorByHealthStatus(status string) *color.Color {
	var c *color.Color
	switch status {
	case api.HealthAny:
		c = color.New(color.FgWhite)
	case api.HealthWarning:
		c = color.New(color.FgWhite)
	case api.HealthCritical:
		c = color.New(color.FgRed)
	case api.HealthPassing:
		c = color.New(color.FgGreen)
	}

	return c
}

func printTags(svc *api.CatalogService) string {
	if isCanary(svc) {
		return fmt.Sprintf("[%s]", color.YellowString("canary"))
	}

	return ""
}

func isCanary(svc *api.CatalogService) bool {
	for _, tag := range svc.ServiceTags {
		if tag == "canary" {
			return true
		}
	}

	return false
}
