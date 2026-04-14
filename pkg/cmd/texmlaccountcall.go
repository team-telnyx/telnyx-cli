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

var texmlAccountsCallsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns an individual call identified by its CallSid. This endpoint is\neventually consistent.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-sid",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "call-sid",
			Required: true,
		},
	},
	Action:          handleTexmlAccountsCallsRetrieve,
	HideHelpCommand: true,
}

var texmlAccountsCallsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update TeXML call. Please note that the keys present in the payload MUST BE\nformatted in CamelCase as specified in the example.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-sid",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "call-sid",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "fallback-method",
			Usage:    "HTTP request type used for `FallbackUrl`.",
			BodyPath: "FallbackMethod",
		},
		&requestflag.Flag[string]{
			Name:     "fallback-url",
			Usage:    "A failover URL for which Telnyx will retrieve the TeXML call instructions if the Url is not responding.",
			BodyPath: "FallbackUrl",
		},
		&requestflag.Flag[string]{
			Name:     "method",
			Usage:    "HTTP request type used for `Url`.",
			BodyPath: "Method",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    "The value to set the call status to. Setting the status to completed ends the call.",
			BodyPath: "Status",
		},
		&requestflag.Flag[string]{
			Name:     "status-callback",
			Usage:    "URL destination for Telnyx to send status callback events to for the call.",
			BodyPath: "StatusCallback",
		},
		&requestflag.Flag[string]{
			Name:     "status-callback-method",
			Usage:    "HTTP request type used for `StatusCallback`.",
			BodyPath: "StatusCallbackMethod",
		},
		&requestflag.Flag[string]{
			Name:     "texml",
			Usage:    "TeXML to replace the current one with.",
			BodyPath: "Texml",
		},
		&requestflag.Flag[string]{
			Name:     "url",
			Usage:    "The URL where TeXML will make a request to retrieve a new set of TeXML instructions to continue the call flow.",
			BodyPath: "Url",
		},
	},
	Action:          handleTexmlAccountsCallsUpdate,
	HideHelpCommand: true,
}

var texmlAccountsCallsCalls = cli.Command{
	Name:    "calls",
	Usage:   "Initiate an outbound TeXML call. Telnyx will request TeXML from the XML Request\nURL configured for the connection in the Mission Control Portal.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-sid",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "params",
			Required: true,
			BodyRoot: true,
		},
	},
	Action:          handleTexmlAccountsCallsCalls,
	HideHelpCommand: true,
}

var texmlAccountsCallsRetrieveCalls = cli.Command{
	Name:    "retrieve-calls",
	Usage:   "Returns multiple call resouces for an account. This endpoint is eventually\nconsistent.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-sid",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "end-time",
			Usage:     "Filters calls by their end date. Expected format is YYYY-MM-DD",
			QueryPath: "EndTime",
		},
		&requestflag.Flag[string]{
			Name:      "end-time-gt",
			Usage:     "Filters calls by their end date (after). Expected format is YYYY-MM-DD",
			QueryPath: "EndTime_gt",
		},
		&requestflag.Flag[string]{
			Name:      "end-time-lt",
			Usage:     "Filters calls by their end date (before). Expected format is YYYY-MM-DD",
			QueryPath: "EndTime_lt",
		},
		&requestflag.Flag[string]{
			Name:      "from",
			Usage:     "Filters calls by the from number.",
			QueryPath: "From",
		},
		&requestflag.Flag[int64]{
			Name:      "page",
			Usage:     "The number of the page to be displayed, zero-indexed, should be used in conjuction with PageToken.",
			QueryPath: "Page",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "The number of records to be displayed on a page",
			QueryPath: "PageSize",
		},
		&requestflag.Flag[string]{
			Name:      "page-token",
			Usage:     "Used to request the next page of results.",
			QueryPath: "PageToken",
		},
		&requestflag.Flag[string]{
			Name:      "start-time",
			Usage:     "Filters calls by their start date. Expected format is YYYY-MM-DD.",
			QueryPath: "StartTime",
		},
		&requestflag.Flag[string]{
			Name:      "start-time-gt",
			Usage:     "Filters calls by their start date (after). Expected format is YYYY-MM-DD",
			QueryPath: "StartTime_gt",
		},
		&requestflag.Flag[string]{
			Name:      "start-time-lt",
			Usage:     "Filters calls by their start date (before). Expected format is YYYY-MM-DD",
			QueryPath: "StartTime_lt",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Filters calls by status.",
			QueryPath: "Status",
		},
		&requestflag.Flag[string]{
			Name:      "to",
			Usage:     "Filters calls by the to number.",
			QueryPath: "To",
		},
	},
	Action:          handleTexmlAccountsCallsRetrieveCalls,
	HideHelpCommand: true,
}

var texmlAccountsCallsSiprecJson = cli.Command{
	Name:    "siprec-json",
	Usage:   "Starts siprec session with specified parameters for call idientified by\ncall_sid.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-sid",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "call-sid",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "connector-name",
			Usage:    "The name of the connector to use for the SIPREC session.",
			BodyPath: "ConnectorName",
		},
		&requestflag.Flag[bool]{
			Name:     "include-metadata-custom-headers",
			Usage:    "When set, custom parameters will be added as metadata (recording.session.ExtensionParameters). Otherwise, they’ll be added to sip headers.",
			BodyPath: "IncludeMetadataCustomHeaders",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the SIPREC session. May be used to stop the SIPREC session from TeXML instruction.",
			BodyPath: "Name",
		},
		&requestflag.Flag[bool]{
			Name:     "secure",
			Usage:    "Controls whether to encrypt media sent to your SRS using SRTP and TLS. When set you need to configure SRS port in your connector to 5061.",
			BodyPath: "Secure",
		},
		&requestflag.Flag[int64]{
			Name:     "session-timeout-secs",
			Usage:    "Sets `Session-Expires` header to the INVITE. A reinvite is sent every half the value set. Usefull for session keep alive. Minimum value is 90, set to 0 to disable.",
			Default:  1800,
			BodyPath: "SessionTimeoutSecs",
		},
		&requestflag.Flag[string]{
			Name:     "sip-transport",
			Usage:    "Specifies SIP transport protocol.",
			Default:  "udp",
			BodyPath: "SipTransport",
		},
		&requestflag.Flag[string]{
			Name:     "status-callback",
			Usage:    "URL destination for Telnyx to send status callback events to for the siprec session.",
			BodyPath: "StatusCallback",
		},
		&requestflag.Flag[string]{
			Name:     "status-callback-method",
			Usage:    "HTTP request type used for `StatusCallback`.",
			BodyPath: "StatusCallbackMethod",
		},
		&requestflag.Flag[string]{
			Name:     "track",
			Usage:    "The track to be used for siprec session. Can be `both_tracks`, `inbound_track` or `outbound_track`. Defaults to `both_tracks`.",
			BodyPath: "Track",
		},
	},
	Action:          handleTexmlAccountsCallsSiprecJson,
	HideHelpCommand: true,
}

var texmlAccountsCallsStreamsJson = cli.Command{
	Name:    "streams-json",
	Usage:   "Starts streaming media from a call to a specific WebSocket address.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-sid",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "call-sid",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "bidirectional-codec",
			Usage:    "Indicates codec for bidirectional streaming RTP payloads. Used only with stream_bidirectional_mode=rtp. Case sensitive.",
			Default:  "PCMU",
			BodyPath: "BidirectionalCodec",
		},
		&requestflag.Flag[string]{
			Name:     "bidirectional-mode",
			Usage:    "Configures method of bidirectional streaming (mp3, rtp).",
			Default:  "mp3",
			BodyPath: "BidirectionalMode",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "The user specified name of Stream.",
			BodyPath: "Name",
		},
		&requestflag.Flag[string]{
			Name:     "status-callback",
			Usage:    "Url where status callbacks will be sent.",
			BodyPath: "StatusCallback",
		},
		&requestflag.Flag[string]{
			Name:     "status-callback-method",
			Usage:    "HTTP method used to send status callbacks.",
			Default:  "POST",
			BodyPath: "StatusCallbackMethod",
		},
		&requestflag.Flag[string]{
			Name:     "track",
			Usage:    "Tracks to be included in the stream",
			Default:  "inbound_track",
			BodyPath: "Track",
		},
		&requestflag.Flag[string]{
			Name:     "url",
			Usage:    "The destination WebSocket address where the stream is going to be delivered.",
			BodyPath: "Url",
		},
	},
	Action:          handleTexmlAccountsCallsStreamsJson,
	HideHelpCommand: true,
}

func handleTexmlAccountsCallsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-sid") && len(unusedArgs) > 0 {
		cmd.Set("call-sid", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlAccountCallGetParams{
		AccountSid: cmd.Value("account-sid").(string),
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
	_, err = client.Texml.Accounts.Calls.Get(
		ctx,
		cmd.Value("call-sid").(string),
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
	return ShowJSON(os.Stdout, os.Stderr, "texml:accounts:calls retrieve", obj, format, explicitFormat, transform)
}

func handleTexmlAccountsCallsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-sid") && len(unusedArgs) > 0 {
		cmd.Set("call-sid", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlAccountCallUpdateParams{
		AccountSid: cmd.Value("account-sid").(string),
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Texml.Accounts.Calls.Update(
		ctx,
		cmd.Value("call-sid").(string),
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
	return ShowJSON(os.Stdout, os.Stderr, "texml:accounts:calls update", obj, format, explicitFormat, transform)
}

func handleTexmlAccountsCallsCalls(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-sid") && len(unusedArgs) > 0 {
		cmd.Set("account-sid", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlAccountCallCallsParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Texml.Accounts.Calls.Calls(
		ctx,
		cmd.Value("account-sid").(string),
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
	return ShowJSON(os.Stdout, os.Stderr, "texml:accounts:calls calls", obj, format, explicitFormat, transform)
}

func handleTexmlAccountsCallsRetrieveCalls(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-sid") && len(unusedArgs) > 0 {
		cmd.Set("account-sid", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlAccountCallGetCallsParams{}

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
	_, err = client.Texml.Accounts.Calls.GetCalls(
		ctx,
		cmd.Value("account-sid").(string),
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
	return ShowJSON(os.Stdout, os.Stderr, "texml:accounts:calls retrieve-calls", obj, format, explicitFormat, transform)
}

func handleTexmlAccountsCallsSiprecJson(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-sid") && len(unusedArgs) > 0 {
		cmd.Set("call-sid", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlAccountCallSiprecJsonParams{
		AccountSid: cmd.Value("account-sid").(string),
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Texml.Accounts.Calls.SiprecJson(
		ctx,
		cmd.Value("call-sid").(string),
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
	return ShowJSON(os.Stdout, os.Stderr, "texml:accounts:calls siprec-json", obj, format, explicitFormat, transform)
}

func handleTexmlAccountsCallsStreamsJson(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-sid") && len(unusedArgs) > 0 {
		cmd.Set("call-sid", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlAccountCallStreamsJsonParams{
		AccountSid: cmd.Value("account-sid").(string),
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Texml.Accounts.Calls.StreamsJson(
		ctx,
		cmd.Value("call-sid").(string),
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
	return ShowJSON(os.Stdout, os.Stderr, "texml:accounts:calls streams-json", obj, format, explicitFormat, transform)
}
