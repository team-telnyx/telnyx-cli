package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
	"github.com/team-telnyx/telnyx-cli/consul"
)

// maintCmd represents the maintenance command
var maintCmd = &cobra.Command{
	Use:     "maintenance [(-d|--disable)] service site",
	Aliases: []string{"maint"},
	Short:   "(BETA) Enable/disable services maintenance in Consul",
	Example: "maintenance --disable call-control-agent ch1-dev",
	Args:    cobra.ExactArgs(2),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		// Checks the current args so we know which one to auto-complete. Perhaps there is a better way to do this.
		switch len(args) {
		case 0:
			// We fetch only the services from PROD to speed up the lookup process
			svcs := consul.GetServicesByEnv("prod")

			var completions []string
			for s, _ := range svcs {
				if strings.HasPrefix(s, toComplete) {
					completions = append(completions, s)
				}
			}

			return completions, cobra.ShellCompDirectiveNoFileComp

		case 1:
			devDcs, err := consul.FetchDatacenters("dev")
			if err != nil {
				cobra.CheckErr(err)
			}

			prodDcs, err := consul.FetchDatacenters("prod")
			if err != nil {
				cobra.CheckErr(err)
			}

			dcs := append(devDcs, prodDcs...)

			var completions []string
			for _, dc := range dcs {
				if strings.HasPrefix(dc, toComplete) {
					completions = append(completions, dc)
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

		err := confirm(disable, svc, dc)
		if err != nil {
			return err
		}

		instances := consul.GetInstancesByDc(dc, svc)
		bar := progressbar.NewOptions(
			len(instances),
			progressbar.OptionEnableColorCodes(true),
			progressbar.OptionShowBytes(false),
			progressbar.OptionSetWidth(15),
			progressbar.OptionShowCount(),
			progressbar.OptionSetDescription("[yellow]Toogling maintenance in Consul...[reset]"),
			progressbar.OptionSetTheme(progressbar.Theme{
				Saucer:        "[green]=[reset]",
				SaucerHead:    "[green]>[reset]",
				SaucerPadding: " ",
				BarStart:      "[",
				BarEnd:        "]",
			}),
		)

		if disable {
			for _, inst := range instances {
				err := consul.DisableInstanceMaintenance(inst, dc)
				cobra.CheckErr(err)
				bar.Add(1)
			}
		} else {
			for _, inst := range instances {
				err := consul.EnableInstanceMaintenance(inst, dc)
				cobra.CheckErr(err)
				bar.Add(1)
			}
		}

		fmt.Println("\nDone!")
		return nil
	},
}

func confirm(disable bool, dc, site string) error {
	var operation string
	if disable {
		operation = "Disabling maintenance"
	} else {
		operation = "Starting maintenance"
	}

	fmt.Printf("%s service: %s, site: %s. Continue? [Y|n]: ", operation, dc, site)

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
	maintCmd.Flags().BoolP("disable", "d", false, "Disables the service maintenance in Consul")
	rootCmd.AddCommand(maintCmd)
}
