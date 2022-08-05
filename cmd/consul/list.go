/*
Copyright © Telnyx LLC

*/
package consul

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/go-cleanhttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Env string
var Dc string
var Filter string
var Serv string

// ListCmd represents the list command
var ListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List resources from Consul",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := args[0]

		switch arg {
		case "datacenters":
			datacenters()
		case "services":
			services()
		case "instances":
			instances()
		default:
			fmt.Println("invalid resource")
			os.Exit(1)
		}
	},
}

func datacenters() {
	dcs, err := fetchDatacenters()
	if err != nil {
		cobra.CheckErr(err)
	}

	for _, dc := range dcs {
		printDatacenter(dc)
	}
}

func fetchDatacenters() ([]string, error) {
	client, err := api.NewClient(consulConfig())
	if err != nil {
		cobra.CheckErr(err)
	}

	return client.Catalog().Datacenters()
}

func services() {
	client, err := api.NewClient(consulConfigForDc(Dc))
	if err != nil {
		cobra.CheckErr(err)
	}

	q := &api.QueryOptions{
		Datacenter: Dc,
	}
	var svcs map[string][]string
	if svcs, _, err = client.Catalog().Services(q); err != nil {
		cobra.CheckErr(err)
	}

	printDatacenter(Dc)

	for svc, tags := range svcs {
		if Filter == "" || strings.Contains(svc, Filter) {
			color.Magenta("  ▶ %s", svc)
			fmt.Printf("    Tags: %s \n", tags)
		}
	}
}

func instances() {
	printService(Serv)

	if Dc == "all" {
		listAllInstances()
	} else {
		listInstancesByDc(Dc)
	}
}

func listAllInstances() {
	dcs, err := fetchDatacenters()
	if err != nil {
		cobra.CheckErr(err)
	}

	for _, dc := range dcs {
		listInstancesByDc(dc)
	}
}

func listInstancesByDc(dc string) {
	printDatacenter(dc)

	instances, err := fetchInstancesByDc(dc)
	if err != nil {
		cobra.CheckErr(err)
	}

	printInstances(instances)
}

func fetchInstancesByDc(dc string) ([]*api.CatalogService, error) {
	client, err := api.NewClient(consulConfigForDc(dc))
	if err != nil {
		cobra.CheckErr(err)
	}

	q := &api.QueryOptions{
		Datacenter: dc,
	}

	svcs, _, err := client.Catalog().Service(Serv, "", q)

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

func printDatacenter(dc string) {
	color.Yellow("▶▶ %s", dc)
}

func printService(s string) {
	color.Magenta("▶ %s", Serv)
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

func consulConfig() *api.Config {
	var consulUrl string

	if Env == "dev" {
		consulUrl = viper.GetString("consul_dev_url")
	} else {
		consulUrl = viper.GetString("consul_prod_url")
	}

	return &api.Config{
		Address:   consulUrl,
		Scheme:    "http",
		Transport: cleanhttp.DefaultTransport(),
	}
}

func consulConfigForDc(dc string) *api.Config {
	var consulUrl string

	if strings.Contains(dc, "dev") {
		consulUrl = viper.GetString("consul_dev_url")
	} else if strings.Contains(dc, "prod") {
		consulUrl = viper.GetString("consul_prod_url")
	} else {
		err := "Invalid datacenter"
		cobra.CheckErr(err)
	}

	return &api.Config{
		Address:   consulUrl,
		Scheme:    "http",
		Transport: cleanhttp.DefaultTransport(),
	}
}

func init() {
	ListCmd.Flags().StringVarP(&Env, "env", "e", "dev", "The Consul environment to use")
	ListCmd.Flags().StringVarP(&Dc, "datacenter", "d", "ch1-dev", "The Consul datacenter to use")
	ListCmd.Flags().StringVarP(&Filter, "filter", "f", "", "Filter service by name containing the given string")
	ListCmd.Flags().StringVarP(&Serv, "service", "s", "", "The Consul service to search")
}
