/*
Copyright © Telnyx LLC

*/
package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/team-telnyx/telnyx-cli/consul"
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

		dcs, err := consul.FetchDatacenters(env)
		if err != nil {
			cobra.CheckErr(err)
		}

		for _, dc := range dcs {
			printDatacenter(dc)
		}
	},
}

func printDatacenter(dc string) {
	color.Yellow("▶▶ %s", dc)
}
