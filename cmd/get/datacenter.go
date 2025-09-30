/*
Copyright © Telnyx LLC
*/
package get

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/team-telnyx/telnyx-cli/consul"
)

func init() {
	datacenterCmd.Flags().StringP("env", "e", "prod", "The Consul environment to use")

	getCmd.AddCommand(datacenterCmd)
}

// datacenterCmd represents the datacenter command
var datacenterCmd = &cobra.Command{
	Use:     "datacenter",
	Aliases: []string{"dc"},
	Short:   "List Consul datacenters",
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
