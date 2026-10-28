// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/team-telnyx/telnyx-cli/internal/apiquery"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var tunnelsCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a tunnel and returns its `tunnel_token`. The token is the only\ncredential the tunnel process needs, and this response and\n`POST /tunnels/{id}/actions/rotate_token` are the only ones that return it. A\nnew tunnel is `inactive` until a tunnel process connects with the token.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "A user specified name for the tunnel.",
			Required: true,
			BodyPath: "name",
		},
	},
	Action:          handleTunnelsCreate,
	HideHelpCommand: true,
}

var tunnelsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves a tunnel with its status and the tunnel processes connected to it. The\nresponse never includes `tunnel_token`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleTunnelsRetrieve,
	HideHelpCommand: true,
}

var tunnelsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Lists the tunnels on the account with their status and connected tunnel\nprocesses. Responses never include `tunnel_token`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[name], filter[status]",
			QueryPath: "filter",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "page",
			Usage:     "Consolidated page parameter (deepObject style). Originally: page[after], page[limit]",
			QueryPath: "page",
		},
		&requestflag.Flag[string]{
			Name:      "sort",
			Usage:     "Sorts the results by a field, ascending, or descending with a leading `-`. The default is `-created_at`.",
			QueryPath: "sort",
		},
	},
	Action:          handleTunnelsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.name",
			Usage:      "The name of the tunnel.",
			InnerField: "name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.status",
			Usage:      "The status of the tunnel.",
			InnerField: "status",
		},
	},
	"page": {
		&requestflag.InnerFlag[string]{
			Name:       "page.after",
			Usage:      "The cursor from `meta.cursors.after` of the previous page.",
			InnerField: "after",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "page.limit",
			Usage:      "The number of results per page.",
			InnerField: "limit",
		},
	},
})

var tunnelsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes a tunnel. A tunnel that a route names as its next hop cannot be deleted;\ndelete those routes first. Connected tunnel processes are disconnected and\ncannot reconnect with the tunnel's token.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleTunnelsDelete,
	HideHelpCommand: true,
}

func handleTunnelsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := telnyx.TunnelNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Tunnels.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "tunnels create",
		Transform:      transform,
	})
}

func handleTunnelsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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
	_, err = client.Tunnels.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "tunnels retrieve",
		Transform:      transform,
	})
}

func handleTunnelsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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

	params := telnyx.TunnelListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Tunnels.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "tunnels list",
		Transform:      transform,
	})
}

func handleTunnelsDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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

	return client.Tunnels.Delete(ctx, cmd.Value("id").(string), options...)
}
