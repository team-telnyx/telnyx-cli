/*
Copyright © Telnyx LLC

The get command works as a placeholder where other subcommands can be nested upon, e.g. `get user`.
It does nothing besides providing some documentation. Check each nested subcommand file for the
actual implementation.
*/
package get

import (
	"github.com/spf13/cobra"
	"github.com/team-telnyx/telnyx-cli/cmd"
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
`}

func init() {
	cmd.RootCmd.AddCommand(getCmd)
}
