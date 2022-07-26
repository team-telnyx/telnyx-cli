/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package consul

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/go-cleanhttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Env string

// ListCmd represents the list command
var ListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List resources from Consul",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := api.NewClient(consulConfig())
		if err != nil {
			cobra.CheckErr(err)
		}

		color.Cyan("Datacenters")

		var dcs []string
		if dcs, err = client.Catalog().Datacenters(); err != nil {
			cobra.CheckErr(err)
		}

		for _, dc := range dcs {
			fmt.Println(dc)
		}
	},
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

func init() {
	ListCmd.Flags().StringVarP(&Env, "env", "e", "dev", "The Consul environment to use")
}
