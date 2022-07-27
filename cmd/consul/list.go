/*
Copyright © Telnyx LLC

*/
package consul

import (
	"fmt"
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
		}
	},
}

func datacenters() {
	client, err := api.NewClient(consulConfig())
	if err != nil {
		cobra.CheckErr(err)
	}

	var dcs []string
	if dcs, err = client.Catalog().Datacenters(); err != nil {
		cobra.CheckErr(err)
	}

	for _, dc := range dcs {
		color.Cyan("▶ %s", dc)
	}
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

	color.Cyan("▶▶ %s", Dc)

	for svc, tags := range svcs {
		if Filter == "" || strings.Contains(svc, Filter) {
			color.Magenta("  ▶ %s", svc)
			fmt.Printf("    Tags: %s \n", tags)
		}
	}
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

	if strings.Contains(Dc, "dev") {
		consulUrl = viper.GetString("consul_dev_url")
	} else if strings.Contains(Dc, "prod") {
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
}
