/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/team-telnyx/telnyx-cli/cmd/consul"
)

// consulCmd represents the consul command
var consulCmd = &cobra.Command{
	Use:   "consul",
	Short: "Utility commands for managing services in Consul",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("consul called")
	},
}

func init() {
	rootCmd.AddCommand(consulCmd)
	consulCmd.AddCommand(consul.ListCmd)
}
