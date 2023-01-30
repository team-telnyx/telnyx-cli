/*
Copyright © Telnyx LLC

*/
package get

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/team-telnyx/telnyx-cli/consul"
)

func init() {
	serviceCmd.Flags().StringP("datacenter", "d", "", "The Consul datacenter to use")
	serviceCmd.Flags().StringP("filter", "f", "", "Filter by service name starting with the given string")

	serviceCmd.MarkFlagRequired("datacenter")

	getCmd.AddCommand(serviceCmd)
}

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:     "service",
	Aliases: []string{"s"},
	Short:   "List services registered in Consul",
	Run: func(cmd *cobra.Command, args []string) {
		dc, _ := cmd.Flags().GetString("datacenter")
		filter, _ := cmd.Flags().GetString("filter")

		printDatacenter(dc)

		svcs := consul.GetServicesByDc(dc)

		for svc, tags := range svcs {
			if filter == "" || strings.HasPrefix(svc, filter) {
				color.Magenta("  ▶ %s", svc)
				fmt.Printf("    Tags: %s \n", tags)
			}
		}
	},
}
