/*
Copyright © Telnyx LLC

*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/go-cleanhttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	serviceCmd.Flags().StringP("datacenter", "d", "", "The Consul datacenter to use")
	serviceCmd.Flags().StringP("filter", "f", "", "Filter by service name containing the given string")

	serviceCmd.MarkFlagRequired("datacenter")

	rootCmd.AddCommand(serviceCmd)
}

// serviceCmd represents the services command
var serviceCmd = &cobra.Command{
	Use:     "services",
	Aliases: []string{"svcs"},
	Short:   "List services from Consul",
	Run: func(cmd *cobra.Command, args []string) {
		dc, _ := cmd.Flags().GetString("datacenter")
		filter, _ := cmd.Flags().GetString("filter")

		printDatacenter(dc)

		svcs := getServicesByDc(dc)

		for svc, tags := range svcs {
			if filter == "" || strings.Contains(svc, filter) {
				color.Magenta("  ▶ %s", svc)
				fmt.Printf("    Tags: %s \n", tags)
			}
		}
	},
}

// TODO: move Map into its own type to improve readability
func getServicesByDc(dc string) map[string][]string {
	client, err := api.NewClient(consulConfigForDc(dc))
	if err != nil {
		cobra.CheckErr(err)
	}

	q := &api.QueryOptions{
		Datacenter: dc,
	}

	var svcs map[string][]string
	if svcs, _, err = client.Catalog().Services(q); err != nil {
		cobra.CheckErr(err)
	}

	return svcs
}

func consulConfigForDc(dc string) *api.Config {
	var consulUrl string

	if strings.Contains(dc, "dev") {
		consulUrl = viper.GetString("consul_dev_url")
	} else if strings.Contains(dc, "prod") {
		consulUrl = viper.GetString("consul_prod_url")
	} else {
		err := "invalid datacenter"
		cobra.CheckErr(err)
	}

	return &api.Config{
		Address:   consulUrl,
		Scheme:    "http",
		Transport: cleanhttp.DefaultTransport(),
	}
}
