package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/team-telnyx/telnyx-cli/consul"
)

// maintCmd represents the maintenance command
var maintCmd = &cobra.Command{
	Use:     "maintenance (SERVICE) (SITE)",
	Aliases: []string{"maint"},
	Short:   "Puts all service instances with given site into maintenance in Consul",
	Args:    cobra.MinimumNArgs(1),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		switch len(args) {
		case 0:
			dcs, err := consul.FetchDatacenters("dev")
			if err != nil {
				cobra.CheckErr(err)
			}

			var completions []string
			for _, dc := range dcs {
				if strings.HasPrefix(dc, toComplete) {
					completions = append(completions, dc)
				}
			}

			return completions, cobra.ShellCompDirectiveNoFileComp

		case 1:
			svcs := consul.GetServicesByEnv("dev")

			var completions []string
			for s, _ := range svcs {
				if strings.HasPrefix(s, toComplete) {
					completions = append(completions, s)
				}
			}

			return completions, cobra.ShellCompDirectiveNoFileComp

		default:
			return []string{}, cobra.ShellCompDirectiveNoFileComp
		}
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		disable, _ := cmd.Flags().GetBool("disable")
		svc := args[0]
		dc := args[1]

		err := confirm("maintenance", svc, dc)
		if err != nil {
			return err
		}

		instances := consul.GetInstancesByDc(dc, svc)

		if disable {
			for _, inst := range instances {
				err = consul.DisableInstanceMaintenance(inst, dc)
				cobra.CheckErr(err)
			}
		} else {
			for _, inst := range instances {
				err = consul.EnableInstanceMaintenance(inst, dc)
				cobra.CheckErr(err)
			}
		}

		fmt.Println("Done!")
		return nil
	},
}

func confirm(operation, dc, site string) error {
	fmt.Printf("Starting %s service: %s, site: %s. Continue? [Y|n]: ", operation, dc, site)

	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		return errors.New("Error getting user's input: " + err.Error())
	}
	if char == 'Y' || char == 'y' {
		return nil
	}
	return errors.New("Aborted")
}

func init() {
	maintCmd.Flags().BoolP("disable", "d", false, "Disables maintenance in Consul")
	rootCmd.AddCommand(maintCmd)
}
