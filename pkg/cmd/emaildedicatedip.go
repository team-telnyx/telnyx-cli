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

var emailDedicatedIPsCreate = cli.Command{
	Name:    "create",
	Usage:   "Purchase a dedicated sending IP. Requires a verified, enterprise, or legacy\naccount, and an `Idempotency-Key` header: a request without the header returns\n`400` (code `10015`) and allocates nothing. Telnyx allocates the next available\nIP from stock. The optional `region` and `ip_version` fields each restrict which\nstock IP is allocated; they never change the allocated IP's metadata. An omitted\nfield applies no restriction for that field, and omitting the body applies\nneither restriction. An unsupported selector value returns `422`; a supported\nvalue that matches no available IP returns `429` and records the request on the\nwaitlist. Billing: $14/month/IP, charged as a monthly recurring charge (MRC)\nfrom purchase. The IP is returned in `provisioning` status with a Telnyx-managed\nPTR hostname and carries no mail until you call `start_warmup` and add it to a\npool that routes mail.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:       "idempotency-key",
			Required:   true,
			HeaderPath: "Idempotency-Key",
		},
		&requestflag.Flag[string]{
			Name:     "ip-version",
			Usage:    "Allocate an IP of this address family. Supported value: `ipv4`. Any other value, including `ipv6` or `null`, returns `422`.",
			BodyPath: "ip_version",
		},
		&requestflag.Flag[string]{
			Name:     "region",
			Usage:    "Allocate an IP in this region. Supported value: `ewr-nj1`. Any other value, including an empty string or `null`, returns `422`.",
			BodyPath: "region",
		},
	},
	Action:          handleEmailDedicatedIPsCreate,
	HideHelpCommand: true,
}

var emailDedicatedIPsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a single dedicated IP by its identifier, including warm-up state,\nbilling state, and the DNS/PTR provisioning checklist.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleEmailDedicatedIPsRetrieve,
	HideHelpCommand: true,
}

var emailDedicatedIPsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Move the IP to another pool, or pause/resume its warm-up ramp. One change\napplies per request. `pool_id` takes precedence over warm-up changes. Otherwise,\nuse `warmup_paused` or `warmup.paused`; a Boolean `warmup_paused` takes\nprecedence. Moving the last IP out of a pool requires `force: true` and\n`drain: true`. The PTR hostname is managed by Telnyx: a request that includes\n`ptr_hostname` returns `422` unless `pool_id` or a warm-up field is present, in\nwhich case `ptr_hostname` is ignored.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[bool]{
			Name:     "drain",
			Usage:    "Set together with `force: true` to move the last IP out of its current pool. Acknowledges that mail routed to the emptied pool is held. Used only with `pool_id`.",
			Default:  false,
			BodyPath: "drain",
		},
		&requestflag.Flag[bool]{
			Name:     "force",
			Usage:    "Set together with `drain: true` to move the last IP out of its current pool. Used only with `pool_id`.",
			Default:  false,
			BodyPath: "force",
		},
		&requestflag.Flag[string]{
			Name:     "pool-id",
			Usage:    "Move the IP to this pool, replacing its current membership. Use the IP's current pool to change only its weight. To take an IP out of a pool, use `DELETE /email/ip_pools/{id}/ips/{ip_id}`.",
			BodyPath: "pool_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "warmup",
			Usage:    "Alias for `warmup_paused`.",
			BodyPath: "warmup",
		},
		&requestflag.Flag[bool]{
			Name:     "warmup-paused",
			Usage:    "Pause (`true`) or resume (`false`) the warm-up ramp; equivalent to the `pause_warmup` and `resume_warmup` actions. A Boolean value takes precedence over `warmup.paused`. When a pause field is present and `pool_id` is absent, supplying neither pause field as a Boolean invokes resume.",
			BodyPath: "warmup_paused",
		},
		&requestflag.Flag[int64]{
			Name:     "weight",
			Usage:    "Routing weight in the target pool. Used only with `pool_id`.",
			Default:  100,
			BodyPath: "weight",
		},
	},
	Action:          handleEmailDedicatedIPsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"warmup": {
		&requestflag.InnerFlag[bool]{
			Name:       "warmup.paused",
			Usage:      "Pause (`true`) or resume (`false`) the warm-up ramp. Used only when `warmup_paused` is not a Boolean.",
			InnerField: "paused",
		},
	},
})

var emailDedicatedIPsList = cli.Command{
	Name:    "list",
	Usage:   "List the dedicated sending IPs owned by the account, newest first. Narrow the\nlist with `filter[status]`, `filter[pool_id]`, and `filter[region]`; supplied\nfilters combine, so an IP must match every one. An empty value, or a nested or\narray value such as `filter[status][]=active`, applies no filter, and other\n`filter[...]` keys are ignored. A value that no IP can match (an unknown status,\na `pool_id` that is not a UUID, or a value containing a NUL byte) returns an\nempty list, not an error.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "filter-pool-id",
			Usage:     "Return only IPs in this pool (exact match on `pool_id`). A value that is not a UUID returns an empty list. An empty value applies no filter.",
			QueryPath: "filter[pool_id]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-region",
			Usage:     "Return only IPs in this region (exact, case-sensitive match on `region`). An empty value applies no filter.",
			QueryPath: "filter[region]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-status",
			Usage:     "Return only IPs in this lifecycle status (exact match). Any other value returns an empty list. An empty value applies no filter.",
			QueryPath: "filter[status]",
		},
	},
	Action:          handleEmailDedicatedIPsList,
	HideHelpCommand: true,
}

var emailDedicatedIPsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Release the dedicated IP to Telnyx. The IP stops carrying mail immediately,\nleaves its pool, and moves to `released`. Telnyx sets `quarantine_until` to 30\ndays after release. Release requests that the MRC stop asynchronously. Check\n`billing`: `stopping` means the stop is pending, `stopped` means it completed,\nand `stop_failed` means the charge remains. If the charge was never created or\nits creation has not completed, `billing` keeps showing the creation status\n(`not_started`, `pending`, `in_progress`, or `failed`). Any status except\n`released` can be released. Releasing the last IP in a pool requires\n`force=true` and `drain=true`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[bool]{
			Name:      "drain",
			Usage:     "Set together with `force=true` to release the last IP in a pool. Acknowledges that mail routed to the emptied pool is held until the pool has a sendable IP again. Also accepted in the JSON body.",
			Default:   false,
			QueryPath: "drain",
		},
		&requestflag.Flag[bool]{
			Name:      "force",
			Usage:     "Set together with `drain=true` to release the last IP in a pool. Either flag alone has no effect. Also accepted in the JSON body.",
			Default:   false,
			QueryPath: "force",
		},
		&requestflag.Flag[string]{
			Name:      "reason",
			Usage:     "Free-text reason recorded with the release. Also accepted in the JSON body.",
			QueryPath: "reason",
		},
		&requestflag.Flag[bool]{
			Name:     "drain",
			Usage:    "Set together with `force: true` to release the last IP in a pool. Acknowledges that mail routed to the emptied pool is held until the pool has a sendable IP again.",
			Default:  false,
			BodyPath: "drain",
		},
		&requestflag.Flag[bool]{
			Name:     "force",
			Usage:    "Set together with `drain: true` to release the last IP in a pool. Either flag alone has no effect.",
			Default:  false,
			BodyPath: "force",
		},
		&requestflag.Flag[string]{
			Name:     "reason",
			Usage:    "Free-text reason recorded with the release.",
			BodyPath: "reason",
		},
	},
	Action:          handleEmailDedicatedIPsDelete,
	HideHelpCommand: true,
}

func handleEmailDedicatedIPsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EmailDedicatedIPNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Email.DedicatedIPs.New(ctx, params, options...)
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
		Title:          "email:dedicated-ips create",
		Transform:      transform,
	})
}

func handleEmailDedicatedIPsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Email.DedicatedIPs.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "email:dedicated-ips retrieve",
		Transform:      transform,
	})
}

func handleEmailDedicatedIPsUpdate(ctx context.Context, cmd *cli.Command) error {
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
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := telnyx.EmailDedicatedIPUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Email.DedicatedIPs.Update(
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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "email:dedicated-ips update",
		Transform:      transform,
	})
}

func handleEmailDedicatedIPsList(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EmailDedicatedIPListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Email.DedicatedIPs.List(ctx, params, options...)
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
		Title:          "email:dedicated-ips list",
		Transform:      transform,
	})
}

func handleEmailDedicatedIPsDelete(ctx context.Context, cmd *cli.Command) error {
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
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := telnyx.EmailDedicatedIPDeleteParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Email.DedicatedIPs.Delete(
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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "email:dedicated-ips delete",
		Transform:      transform,
	})
}
