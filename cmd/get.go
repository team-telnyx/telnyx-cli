/*
Copyright © Telnyx LLC

The get command works as a placeholder where other subcommands can be nested upon, e.g. `get user`.
It does nothing besides providing some documentation. Check each nested subcommand file for its
actual implementation.

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Display one or many resources",
	Long: `
Prints a table of the most important information about the specified resources. Examples:

# List all users belonging to a given organization
telnyx-cli get user --organization-id <organization_id>

# Search for user with given email, displaying the user's full information
telnyx-cli get user --email <email> --long


Options:
	-e, --env=prod:
		The environment to use. Defaults to "prod".

	-l, --long=false:
		Display's the resource's full information. Defaults to "false".`,
}

func init() {
	rootCmd.AddCommand(getCmd)
}
