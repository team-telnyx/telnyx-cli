// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/team-telnyx/telnyx-cli/internal/apiquery"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var organizationsUsersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns a user in your organization.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[bool]{
			Name:      "include-groups",
			Usage:     "When set to true, includes the groups array for each user in the response. The groups array contains objects with id and name for each group the user belongs to.",
			Default:   false,
			QueryPath: "include_groups",
		},
	},
	Action:          handleOrganizationsUsersRetrieve,
	HideHelpCommand: true,
}

var organizationsUsersList = cli.Command{
	Name:    "list",
	Usage:   "Returns a list of the users in your organization.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "filter-email",
			Usage:     "Filter by email address (partial match)",
			QueryPath: "filter[email]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-user-status",
			Usage:     "Filter by user status",
			QueryPath: "filter[user_status]",
		},
		&requestflag.Flag[bool]{
			Name:      "include-groups",
			Usage:     "When set to true, includes the groups array for each user in the response. The groups array contains objects with id and name for each group the user belongs to.",
			Default:   false,
			QueryPath: "include_groups",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Usage:     "The page number to load",
			Default:   1,
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "The size of the page",
			Default:   250,
			QueryPath: "page[size]",
		},
	},
	Action:          handleOrganizationsUsersList,
	HideHelpCommand: true,
}

var organizationsUsersGetGroupsReport = cli.Command{
	Name:    "get-groups-report",
	Usage:   "Returns a report of all users in your organization with their group memberships.\nThis endpoint returns all users without pagination and always includes group\ninformation. The report can be retrieved in JSON or CSV format by sending\nspecific content-type headers.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:       "accept",
			Default:    "application/json",
			HeaderPath: "Accept",
		},
	},
	Action:          handleOrganizationsUsersGetGroupsReport,
	HideHelpCommand: true,
}

func handleOrganizationsUsersRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.OrganizationUserGetParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Organizations.Users.Get(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "organizations:users retrieve", obj, format, transform)
}

func handleOrganizationsUsersList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.OrganizationUserListParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Organizations.Users.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "organizations:users list", obj, format, transform)
	} else {
		iter := client.Organizations.Users.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "organizations:users list", iter, format, transform)
	}
}

func handleOrganizationsUsersGetGroupsReport(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.OrganizationUserGetGroupsReportParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Organizations.Users.GetGroupsReport(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "organizations:users get-groups-report", obj, format, transform)
}
