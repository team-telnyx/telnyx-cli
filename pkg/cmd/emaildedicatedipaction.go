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

var emailDedicatedIPsActionsCompleteWarmup = cli.Command{
	Name:    "complete-warmup",
	Usage:   "Manually complete the warm-up ramp (`warming_up` to `active`), bypassing the\nschedule. Use only for an IP already warmed elsewhere. Completion sets\n`warmup.progress` to `100`, clears `warmup.daily_cap`, and applies the\nramp-complete allowance to the current UTC day: `warmup.today.cap` becomes\n`100000000`, `warmup.today.granted` keeps the sends already released today, and\n`warmup.today.remaining` becomes `cap` minus `granted`. Once completion applies\nthis allowance, an exhausted stage allowance does not block the next eligible\nfirst-attempt send. The response reports these values. On an `active` IP,\n`warmup.progress` and `warmup.daily_cap` are informational: they can keep\nearlier values after a suspension is lifted and do not limit sending.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleEmailDedicatedIPsActionsCompleteWarmup,
	HideHelpCommand: true,
}

var emailDedicatedIPsActionsPauseWarmup = cli.Command{
	Name:    "pause-warmup",
	Usage:   "Pause the warm-up ramp at its current stage. While paused, the IP releases no\nfirst-attempt sends from hold and the ramp does not advance. Only a `warming_up`\nIP can be paused; pausing an already paused IP refreshes `paused_at`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleEmailDedicatedIPsActionsPauseWarmup,
	HideHelpCommand: true,
}

var emailDedicatedIPsActionsResumeWarmup = cli.Command{
	Name:    "resume-warmup",
	Usage:   "Resume a paused warm-up ramp from its held stage. Only a `warming_up` IP can be\nresumed; resuming a ramp that is not paused returns the IP unchanged.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleEmailDedicatedIPsActionsResumeWarmup,
	HideHelpCommand: true,
}

var emailDedicatedIPsActionsStartWarmup = cli.Command{
	Name:    "start-warmup",
	Usage:   "Start the warm-up ramp (`provisioning` to `warming_up`). The ramp never starts\nautomatically. Only a `provisioning` IP can be started; a `suspended` IP returns\n`409`. Starting does not check DNS: the IP carries mail only while\n`dns_confirmed_at` is set and the IP is in a pool that routes mail.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleEmailDedicatedIPsActionsStartWarmup,
	HideHelpCommand: true,
}

func handleEmailDedicatedIPsActionsCompleteWarmup(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Email.DedicatedIPs.Actions.CompleteWarmup(ctx, cmd.Value("id").(string), options...)
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
		Title:          "email:dedicated-ips:actions complete-warmup",
		Transform:      transform,
	})
}

func handleEmailDedicatedIPsActionsPauseWarmup(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Email.DedicatedIPs.Actions.PauseWarmup(ctx, cmd.Value("id").(string), options...)
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
		Title:          "email:dedicated-ips:actions pause-warmup",
		Transform:      transform,
	})
}

func handleEmailDedicatedIPsActionsResumeWarmup(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Email.DedicatedIPs.Actions.ResumeWarmup(ctx, cmd.Value("id").(string), options...)
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
		Title:          "email:dedicated-ips:actions resume-warmup",
		Transform:      transform,
	})
}

func handleEmailDedicatedIPsActionsStartWarmup(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Email.DedicatedIPs.Actions.StartWarmup(ctx, cmd.Value("id").(string), options...)
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
		Title:          "email:dedicated-ips:actions start-warmup",
		Transform:      transform,
	})
}
