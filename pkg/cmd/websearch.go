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

var webSearchCreate = cli.Command{
	Name:    "create",
	Usage:   "Performs a real-time web search and returns structured, LLM-ready JSON results\nwith titles, URLs, descriptions, and snippets. Supports filtering by domain,\ncountry, safe search, freshness, and live crawl.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "query",
			Usage:    "The search query text.",
			Required: true,
			BodyPath: "query",
		},
		&requestflag.Flag[int64]{
			Name:     "count",
			Usage:    "Number of results to return (1-100).",
			BodyPath: "count",
		},
		&requestflag.Flag[string]{
			Name:     "country",
			Usage:    "Two-letter country code (ISO 3166-1 alpha-2) to bias results.",
			BodyPath: "country",
		},
		&requestflag.Flag[[]string]{
			Name:     "exclude-domain",
			Usage:    "Exclude results from these domains (bare hostnames, e.g. `pinterest.com`).",
			BodyPath: "exclude_domains",
		},
		&requestflag.Flag[string]{
			Name:     "freshness",
			Usage:    "Time-based filter for results. Common values: `day`, `week`, `month`, `year`.",
			BodyPath: "freshness",
		},
		&requestflag.Flag[[]string]{
			Name:     "include-domain",
			Usage:    "Restrict results to these domains (bare hostnames, e.g. `arxiv.org`).",
			BodyPath: "include_domains",
		},
		&requestflag.Flag[bool]{
			Name:     "livecrawl",
			Usage:    "When true, the provider crawls pages in real-time for fresh content. The boolean is translated to the provider's internal enum internally; callers always pass `true` or `false`.",
			BodyPath: "livecrawl",
		},
		&requestflag.Flag[string]{
			Name:     "safesearch",
			Usage:    "Safe search filter level.",
			BodyPath: "safesearch",
		},
	},
	Action:          handleWebSearchCreate,
	HideHelpCommand: true,
}

var webSearchContents = cli.Command{
	Name:    "contents",
	Usage:   "Retrieves clean HTML or Markdown content from a list of URLs. Supports up to 20\nURLs per request (public API limit). Specify which formats to return: `html`,\n`markdown`, `metadata`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "url",
			Usage:    "List of URLs to retrieve content from (max 20 for public API).",
			Required: true,
			BodyPath: "urls",
		},
		&requestflag.Flag[int64]{
			Name:     "crawl-timeout",
			Usage:    "Timeout for crawling each URL, in seconds (1-60).",
			BodyPath: "crawl_timeout",
		},
		&requestflag.Flag[[]string]{
			Name:     "format",
			Usage:    "Content formats to return. If omitted, `html` and `metadata` are returned by default. Retrieval is best-effort per URL: a format field appears only when that content could be produced, and a freshly crawled page may also include `html` even when not requested.",
			BodyPath: "formats",
		},
		&requestflag.Flag[*int64]{
			Name:     "max-age",
			Usage:    "Maximum age of cached content in seconds. `null` means no limit.",
			BodyPath: "max_age",
		},
	},
	Action:          handleWebSearchContents,
	HideHelpCommand: true,
}

func handleWebSearchCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.WebSearchNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.WebSearch.New(ctx, params, options...)
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
		Title:          "web-search create",
		Transform:      transform,
	})
}

func handleWebSearchContents(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.WebSearchContentsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.WebSearch.Contents(ctx, params, options...)
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
		Title:          "web-search contents",
		Transform:      transform,
	})
}
