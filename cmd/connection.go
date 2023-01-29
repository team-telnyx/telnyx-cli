/*
Copyright © Telnyx LLC

*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/team-telnyx/telnyx-cli/privateapi"
)

// connCmd represents the connection command
var connCmd = &cobra.Command{
	Use:     "connection",
	Aliases: []string{"conn"},
	Short:   "Display one or more connections",
	Run: func(cmd *cobra.Command, args []string) {
		env, _ := cmd.Flags().GetString("env")
		connId, _ := cmd.Flags().GetString("id")
		userId, _ := cmd.Flags().GetString("user_id")
		long, _ := cmd.Flags().GetBool("long")
		query := &privateapi.ConnectionQuery{Id: connId, UserId: userId}

		conns, err := privateapi.FetchConnections(env, query)
		if err != nil {
			cobra.CheckErr(err)
		}

		if len(conns) == 0 {
			fmt.Println("Not found!")
			return
		}

		for _, conn := range conns {
			if long {
				printConnection(conn)
			} else {
				printConnectionShort(conn)
			}
		}
	},
}

func printConnectionShort(conn *privateapi.Connection) {
	u, err := json.MarshalIndent(conn.ToShort(), "", "  ")
	if err != nil {
		cobra.CheckErr(err)
	}
	fmt.Printf("%s", string(u))
}

func printConnection(conn *privateapi.Connection) {
	u, err := json.MarshalIndent(conn, "", "  ")
	if err != nil {
		cobra.CheckErr(err)
	}
	fmt.Printf("%s", string(u))
}

func init() {
	connCmd.Flags().StringP("env", "e", "prod", "The environment to use: [dev|prod]. Defaults to \"prod\"")
	connCmd.Flags().String("id", "", "The connection ID to be searched")
	connCmd.Flags().String("user_id", "", "The organization ID to be searched")
	connCmd.Flags().BoolP("long", "l", false, "Display the connection's full information")
	getCmd.AddCommand(connCmd)
}
