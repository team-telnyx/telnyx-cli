package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/hashicorp/consul/api"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
	"github.com/team-telnyx/telnyx-cli/consul"
)

// maintCmd represents the maintenance command
var maintCmd = &cobra.Command{
	Use:     "maintenance [(-d|--disable)] service site [(-n|--node) node] [(-r|--reason) reason]",
	Aliases: []string{"maint"},
	Short:   "Toggle service maintenance in Consul",
	Long: `
Enable or disables maintenance in Consul for a given service. By default, it
tries to enable maintenance mode, unless the flag -d|--disable is provided.

Tip: the "--node" identifier can be found by listing the service instances via
"telnyx-cli get instance <service>". The node identifier is the value defined
inside the first [] characters.


Examples:

# Enable maintenance mode on all call-control instances present in sv1-dev with provided reason
telnyx-cli maintenance sv1-dev call-control --reason "Something wrong happened"

# Disable maintenance mode on all call-control instances present in sv1-dev
telnyx-cli maintenance --disable sv1-dev call-control

# Enable maintenance mode on the call-control instance present in node ip-10-48-192-204.us-west-1.compute.internal
telnyx-cli maintenance sv1-dev call-control --node ip-10-48-192-204.us-west-1.compute.internal
	`,
	Example: "maintenance --disable call-control-agent ch1-dev",
	Args:    cobra.ExactArgs(2),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		switch len(args) {
		// Auto-complete services
		case 0:
			// We fetch only the services from PROD to speed up the lookup process. This implies that services
			// not present in PROD might not show up for completion when dealing with dev environment.
			svcs := consul.GetServicesByEnv("prod")

			var completions []string
			for s, _ := range svcs {
				if strings.HasPrefix(s, toComplete) {
					completions = append(completions, s)
				}
			}

			return completions, cobra.ShellCompDirectiveNoFileComp

		// Auto-complete Consul datacenters
		case 1:
			var completions []string

			devDcs, err := consul.FetchDatacenters("dev")
			if err != nil {
				cobra.CheckErr(err)
			}

			prodDcs, err := consul.FetchDatacenters("prod")
			if err != nil {
				cobra.CheckErr(err)
			}

			dcs := append(devDcs, prodDcs...)

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
		node, _ := cmd.Flags().GetString("node")
		reason, _ := cmd.Flags().GetString("reason")

		if node == "" {
			svc := args[0]
			dc := args[1]
			instances := consul.GetInstancesByDc(dc, svc)

			if len(instances) == 0 {
				err := fmt.Errorf("service %s was not found on %s", svc, dc)
				cobra.CheckErr(err)
			}

			err := confirm(disable, svc, dc, instances)
			if err != nil {
				cobra.CheckErr(err)
			}

			bar := buildProgressBar(instances)

			if disable {
				for _, inst := range instances {
					err := consul.DisableInstanceMaintenance(inst, dc)
					cobra.CheckErr(err)
					bar.Add(1)
				}
			} else {
				for _, inst := range instances {
					err := consul.EnableInstanceMaintenance(inst, dc, reason)
					cobra.CheckErr(err)
					bar.Add(1)
				}
			}
		} else {
			svc := args[0]
			dc := args[1]
			node, _ := cmd.Flags().GetString("node")

			dcInstances := consul.GetInstancesByDc(dc, svc)
			var instances []*api.ServiceEntry

			for _, inst := range dcInstances {
				if inst.Node.Node == node {
					instances = append(instances, inst)
				}
			}

			if len(instances) == 0 {
				err := fmt.Errorf("service %s was not found on %s", svc, dc)
				cobra.CheckErr(err)
			}

			err := confirm(disable, svc, node, instances)
			if err != nil {
				cobra.CheckErr(err)
			}

			if disable {
				err := consul.DisableInstanceMaintenance(instances[0], dc)
				cobra.CheckErr(err)
			} else {
				err := consul.EnableInstanceMaintenance(instances[0], dc, reason)
				cobra.CheckErr(err)
			}
		}
		fmt.Println("\nDone!")
		return nil
	},
}

func buildProgressBar(instances []*api.ServiceEntry) *progressbar.ProgressBar {
	return progressbar.NewOptions(
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
}

func confirm(disable bool, svc, siteOrNode string, instances []*api.ServiceEntry) error {
	var operation string
	if disable {
		operation = "Disabling"
	} else {
		operation = "Enabling"
	}

	fmt.Printf(
		"%s maintenance for %s in %s. %d instance(s) will be affected. Continue?[y|n] ",
		operation,
		svc,
		siteOrNode,
		len(instances),
	)

	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		return errors.New("error getting user's input: " + err.Error())
	}
	if char == 'Y' || char == 'y' {
		return nil
	}
	return errors.New("operation aborted")
}

func init() {
	maintCmd.Flags().BoolP("disable", "d", false, "Disable the service maintenance in Consul")
	maintCmd.Flags().StringP("node", "n", "", "Defines the instance to be put/removed into maintenance in Consul")
	maintCmd.Flags().StringP(
		"reason",
		"r",
		"Maintenance enabled via telnyx-cli, no reason provided (default message)",
		"Reason for putting the service under maintenance",
	)
	RootCmd.AddCommand(maintCmd)
}
