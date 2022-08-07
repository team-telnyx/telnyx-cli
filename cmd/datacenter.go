/*
Copyright © Telnyx LLC

*/
package cmd

import (
	"github.com/fatih/color"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/go-cleanhttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	datacenterCmd.Flags().StringP("env", "e", "dev", "The Consul environment to use")

	rootCmd.AddCommand(datacenterCmd)
}

// datacenterCmd represents the datacenters command
var datacenterCmd = &cobra.Command{
	Use:     "datacenters",
	Aliases: []string{"dcs"},
	Short:   "List datacenters from Consul",
	Run: func(cmd *cobra.Command, args []string) {
		env, _ := cmd.Flags().GetString("env")

		dcs, err := fetchDatacenters(env)
		if err != nil {
			cobra.CheckErr(err)
		}

		for _, dc := range dcs {
			printDatacenter(dc)
		}
	},
}

func fetchDatacenters(env string) ([]string, error) {
	client, err := api.NewClient(consulConfig(env))
	if err != nil {
		cobra.CheckErr(err)
	}

	return client.Catalog().Datacenters()
}

func consulConfig(env string) *api.Config {
	var consulUrl string

	if env == "dev" {
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

func printDatacenter(dc string) {
	color.Yellow("▶▶ %s", dc)
}
