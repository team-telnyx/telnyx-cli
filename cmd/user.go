/*
Copyright © Telnyx LLC

*/
package cmd

// TODO
//  - Add connection_id to look for users
//  - Add aliases to find users by big customers: Avaya, Cisco, telapps.squad, Ringba, LivePerson, OneReach, others

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/team-telnyx/telnyx-cli/privateapi"
)

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Display one or more users",
	Run: func(cmd *cobra.Command, args []string) {
		env, _ := cmd.Flags().GetString("env")
		userId, _ := cmd.Flags().GetString("id")
		email, _ := cmd.Flags().GetString("email")
		orgId, _ := cmd.Flags().GetString("organization_id")
		long, _ := cmd.Flags().GetBool("long")
		query := &privateapi.UserQuery{Id: userId, Email: email, OrganizationId: orgId}

		users, err := privateapi.FetchUsers(env, query)
		if err != nil {
			cobra.CheckErr(err)
		}

		if len(users) == 0 {
			fmt.Println("Not found!")
			return
		}

		for _, user := range users {
			if long {
				printUser(user)
			} else {
				printUserShort(user)
			}
		}
	},
}

func printUserShort(user *privateapi.User) {
	u, err := json.MarshalIndent(user.ToShort(), "", "  ")
	if err != nil {
		cobra.CheckErr(err)
	}
	fmt.Printf("%s", string(u))
}

func printUser(user *privateapi.User) {
	u, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		cobra.CheckErr(err)
	}
	fmt.Printf("%s", string(u))
}

func init() {
	userCmd.Flags().StringP("env", "e", "prod", "The environment to use: [dev|prod]. Defaults to \"prod\"")
	userCmd.Flags().String("id", "", "The user ID to be searched")
	userCmd.Flags().String("email", "", "The user email to be searched")
	userCmd.Flags().String("organization_id", "", "The organization ID to be searched")
	userCmd.Flags().BoolP("long", "l", false, "Display the user's full information")
	getCmd.AddCommand(userCmd)
}
