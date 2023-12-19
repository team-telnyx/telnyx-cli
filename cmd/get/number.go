/*
Copyright © Telnyx LLC

*/
package get

import (
	"encoding/json"
	"fmt"
	"slices"

	"github.com/spf13/cobra"
	"github.com/team-telnyx/telnyx-cli/privateapi"
)

// numCmd represents the number command
var numCmd = &cobra.Command{
	Use:     "number",
	Aliases: []string{"n"},
	Short:   "Display one or more numbers",
	Run: func(cmd *cobra.Command, args []string) {
		env, _ := cmd.Flags().GetString("env")
		all, _ := cmd.Flags().GetBool("all")
		tnE164, _ := cmd.Flags().GetString("tn")
		connId, _ := cmd.Flags().GetString("connection_id")
		userId, _ := cmd.Flags().GetString("user_id")
		fqdn, _ := cmd.Flags().GetString("fqdn")
		ip, _ := cmd.Flags().GetString("ip")
		username, _ := cmd.Flags().GetString("username")

		query := &privateapi.NumberQuery{
			TnE164:       tnE164,
			ConnectionId: connId,
			UserId:       userId,
			Fqdn:         fqdn,
			Ip:           ip,
			Username:     username,
		}

		nums, err := privateapi.FetchNumbers(env, query)
		if err != nil {
			cobra.CheckErr(err)
		}

		if len(nums) == 0 {
			fmt.Println("Not found!")
			return
		}

		if !all {
			nums = slices.DeleteFunc(nums, func(num *privateapi.Number) bool {
				return num.Status != "active"
			})
		}

		json, err := json.MarshalIndent(nums, "", "  ")
		if err != nil {
			cobra.CheckErr(err)
		}
		fmt.Printf("%s\n", string(json))

	},
}

func init() {
	numCmd.Flags().StringP("env", "e", "prod", "The environment to use: [dev|prod]. Defaults to \"prod\"")
	numCmd.Flags().BoolP("all", "a", false, "Return non-active and pending numbers")
	numCmd.Flags().String("tn", "", "The number in E164 format")
	numCmd.Flags().String("connection_id", "", "The number's connection_id")
	numCmd.Flags().String("user_id", "", "The number's user_id")
	numCmd.Flags().String("fqdn", "", "The number's FQDN")
	numCmd.Flags().String("ip", "", "The number's IP address")
	numCmd.Flags().String("username", "", "The number's SIP username")
	getCmd.AddCommand(numCmd)
}
