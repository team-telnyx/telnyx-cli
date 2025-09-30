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
	"github.com/team-telnyx/telnyx-cli/tailscale"
)

var deregisterCmd = &cobra.Command{
	Use:     "deregister [service] [datacenter] [(-e|--env) environment] [(-n|--node) node...]",
	Aliases: []string{"dereg"},
	Short:   "Deregister service instances or nodes from Consul catalog",
	Long: `
Deregister service instances or nodes from the Consul catalog. This command allows
you to remove specific service instances or entire nodes from Consul's service discovery.

By default, this command will deregister all instances of a service in the specified datacenter.
You can use the --node flag to specify particular nodes to deregister.

Note: This operation directly modifies the Consul catalog. Use with caution as
deregistered services will no longer be discoverable until re-registered.

Examples:

# Deregister all instances of call-control service in sv1-dev datacenter
telnyx-cli deregister call-control sv1-dev -e dev

# Deregister specific call-control instances on particular nodes
telnyx-cli deregister call-control sv1-dev -e dev -n ip-10-48-192-204.us-west-1.compute.internal -n ip-10-48-192-205.us-west-1.compute.internal

# Deregister an entire node (all services on that node)
telnyx-cli deregister sv1-dev -e dev --node-only -n ip-10-48-192-204.us-west-1.compute.internal
	`,
	Example: "deregister call-control sv1-dev -e dev --node ip-10-48-192-204.us-west-1.compute.internal",
	Args:    cobra.RangeArgs(1, 2),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		switch len(args) {
		case 0:
			svcs := consul.GetServicesByEnv("prod")
			var completions []string
			for s, _ := range svcs {
				if strings.HasPrefix(s, toComplete) {
					completions = append(completions, s)
				}
			}
			return completions, cobra.ShellCompDirectiveNoFileComp

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
		nodes, _ := cmd.Flags().GetStringSlice("node")
		nodes = removeDuplicates(nodes)
		nodeOnly, _ := cmd.Flags().GetBool("node-only")
		env, _ := cmd.Flags().GetString("env")

		if nodeOnly {
			if len(args) != 1 {
				return errors.New("when using --node-only, only specify the datacenter")
			}
			if len(nodes) == 0 {
				return errors.New("when using --node-only, you must specify at least one node with --node")
			}
			dc := args[0]
			return deregisterNodes(dc, env, nodes)
		}

		if len(args) != 2 {
			return errors.New("service and datacenter are required unless using --node-only")
		}

		svc := args[0]
		dc := args[1]

		return deregisterService(svc, dc, env, nodes)
	},
}

func deregisterService(svc, dc, env string, nodes []string) error {
	instances, err := consul.GetInstancesByDc(dc, svc, env)
	if err != nil {
		return fmt.Errorf("failed to get instances: %w", err)
	}

	if len(instances) == 0 {
		return fmt.Errorf("no instances found for service %s in %s", svc, dc)
	}

	var instancesToDeregister []*api.ServiceEntry
	if len(nodes) == 0 {
		instancesToDeregister = instances
	} else {
		for _, inst := range instances {
			for _, node := range nodes {
				if inst.Node.Node == node {
					instancesToDeregister = append(instancesToDeregister, inst)
					break
				}
			}
		}
		if len(instancesToDeregister) == 0 {
			return fmt.Errorf("no matching instances found for service %s in %s with specified nodes", svc, dc)
		}
	}

	nodeInfo := ""
	if len(nodes) > 0 {
		nodeInfo = fmt.Sprintf(" (nodes: %s)", strings.Join(nodes, ", "))
	}

	err = confirmDeregistration(fmt.Sprintf("service %s in %s%s", svc, dc, nodeInfo), len(instancesToDeregister))
	if err != nil {
		return err
	}

	client, err := api.NewClient(consul.ConsulConfigForDc(dc, env))
	if err != nil {
		return fmt.Errorf("failed to create consul client: %w", err)
	}

	bar := buildDeregisterProgressBar(instancesToDeregister)

	user := tailscale.GetTailscaleUser()
	for _, inst := range instancesToDeregister {
		deregistration := &api.CatalogDeregistration{
			Node:       inst.Node.Node,
			ServiceID:  inst.Service.ID,
			Datacenter: dc,
		}

		_, err := client.Catalog().Deregister(deregistration, nil)
		if err != nil {
			return fmt.Errorf("failed to deregister service %s on node %s: %w", inst.Service.ID, inst.Node.Node, err)
		}

		fmt.Printf("Deregistered by %s: %s on %s\n", user, inst.Service.ID, inst.Node.Node)
		bar.Add(1)
	}

	fmt.Println("\nDeregistration completed!")
	return nil
}

func deregisterNodes(dc, env string, nodes []string) error {
	err := confirmDeregistration(fmt.Sprintf("nodes %s in %s", strings.Join(nodes, ", "), dc), len(nodes))
	if err != nil {
		return err
	}

	client, err := api.NewClient(consul.ConsulConfigForDc(dc, env))
	if err != nil {
		return fmt.Errorf("failed to create consul client: %w", err)
	}

	bar := buildNodeDeregisterProgressBar(nodes)

	user := tailscale.GetTailscaleUser()
	for _, node := range nodes {
		deregistration := &api.CatalogDeregistration{
			Node:       node,
			Datacenter: dc,
		}

		_, err := client.Catalog().Deregister(deregistration, nil)
		if err != nil {
			return fmt.Errorf("failed to deregister node %s: %w", node, err)
		}

		fmt.Printf("Node deregistered by %s: %s\n", user, node)
		bar.Add(1)
	}

	fmt.Println("\nNode deregistration completed!")
	return nil
}

func buildDeregisterProgressBar(instances []*api.ServiceEntry) *progressbar.ProgressBar {
	return progressbar.NewOptions(
		len(instances),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(false),
		progressbar.OptionSetWidth(15),
		progressbar.OptionShowCount(),
		progressbar.OptionSetDescription("[red]Deregistering services from Consul...[reset]"),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[red]=[reset]",
			SaucerHead:    "[red]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
	)
}

func buildNodeDeregisterProgressBar(nodes []string) *progressbar.ProgressBar {
	return progressbar.NewOptions(
		len(nodes),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(false),
		progressbar.OptionSetWidth(15),
		progressbar.OptionShowCount(),
		progressbar.OptionSetDescription("[red]Deregistering nodes from Consul...[reset]"),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[red]=[reset]",
			SaucerHead:    "[red]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
	)
}

func confirmDeregistration(target string, count int) error {
	fmt.Printf(
		"DANGER: You are about to deregister %s. %d item(s) will be removed from Consul catalog. Continue?[y|n] ",
		target,
		count,
	)

	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		return fmt.Errorf("error getting user's input: %w", err)
	}
	if char == 'Y' || char == 'y' {
		return nil
	}
	return errors.New("operation aborted")
}

func init() {
	deregisterCmd.Flags().StringP("env", "e", "prod", "The Consul environment to use (dev or prod)")
	deregisterCmd.Flags().StringSliceP("node", "n", []string{}, "Specific nodes to deregister services from (can be specified multiple times)")
	deregisterCmd.Flags().Bool("node-only", false, "Deregister entire nodes instead of services")
	RootCmd.AddCommand(deregisterCmd)
}
