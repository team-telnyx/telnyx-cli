/*
Copyright © Telnyx LLC

*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
	"github.com/hashicorp/consul/api"
	"github.com/spf13/cobra"
)

func init() {
	instanceCmd.Flags().StringP("env", "e", "dev", "The Consul environment to use")
	instanceCmd.Flags().StringP("datacenter", "d", "", "The Consul datacenter to use")

	rootCmd.AddCommand(instanceCmd)
}

// instanceCmd represents the instances command
var instanceCmd = &cobra.Command{
	Use:     "instances",
	Aliases: []string{"ist"},
	Short:   "List instances of a service from Consul",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		svc := args[0]
		env, _ := cmd.Flags().GetString("env")
		dc, _ := cmd.Flags().GetString("datacenter")

		printService(svc)

		if dc == "" {
			listAllInstancesByEnv(env, svc)
		} else {
			listInstancesByDc(dc, svc)
		}
	},
}

func listAllInstancesByEnv(env string, svc string) {
	dcs, err := fetchDatacenters(env)
	if err != nil {
		cobra.CheckErr(err)
	}

	for _, dc := range dcs {
		listInstancesByDc(dc, svc)
	}
}

func listInstancesByDc(dc string, svc string) {
	printDatacenter(dc)

	instances, err := fetchInstancesByDc(dc, svc)
	if err != nil {
		cobra.CheckErr(err)
	}

	printInstances(instances)
}

func fetchInstancesByDc(dc string, svc string) ([]*api.CatalogService, error) {
	client, err := api.NewClient(consulConfigForDc(dc))
	if err != nil {
		cobra.CheckErr(err)
	}

	q := &api.QueryOptions{
		Datacenter: dc,
	}

	svcs, _, err := client.Catalog().Service(svc, "", q)

	return svcs, err
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
		"    » [%s][%s:%s][%s][%s]\n",
		color.CyanString(svc.Node),
		color.CyanString(svc.ServiceAddress),
		color.CyanString(strconv.Itoa(svc.ServicePort)),
		color.CyanString(version),
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
