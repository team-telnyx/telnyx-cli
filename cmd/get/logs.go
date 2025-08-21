/*
Copyright © Telnyx LLC

*/
package get

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

type LokiResponse struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string      `json:"resultType"`
		Result     interface{} `json:"result"`
	} `json:"data"`
	Error     string `json:"error,omitempty"`
	ErrorType string `json:"errorType,omitempty"`
}

func init() {
	logsCmd.Flags().StringP("service", "s", "", "Service name to query logs for (uses service_name label)")
	logsCmd.Flags().StringP("query", "q", "", "LogQL query to execute")
	logsCmd.Flags().StringP("environment", "e", "prod", "Environment to query (dev or prod)")
	logsCmd.Flags().StringP("region", "r", "global", "Region to query (global or eu)")
	logsCmd.Flags().StringP("since", "", "1h", "Time range to query (e.g., 1h, 30m, 1d)")
	logsCmd.Flags().IntP("limit", "l", 100, "Maximum number of log entries to return")

	getCmd.AddCommand(logsCmd)
}

// logsCmd represents the logs command
var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Query Loki for logs using LogQL",
	Long: `Query Loki for logs using LogQL syntax. Examples:

# Query logs for a specific service
telnyx-cli get logs -s "my-service" -e prod

# Query with custom LogQL
telnyx-cli get logs -q '{service_name="my-service"} |= "error"' -e dev

# Query EU region logs
telnyx-cli get logs -s "my-service" -r eu -e prod`,
	Run: func(cmd *cobra.Command, args []string) {
		service, _ := cmd.Flags().GetString("service")
		query, _ := cmd.Flags().GetString("query")
		env, _ := cmd.Flags().GetString("environment")
		region, _ := cmd.Flags().GetString("region")
		since, _ := cmd.Flags().GetString("since")
		limit, _ := cmd.Flags().GetInt("limit")

		// Validate input
		if service == "" && query == "" {
			fmt.Fprintf(os.Stderr, "Error: Either --service or --query must be provided\n")
			return
		}

		// Build the query
		var logQuery string
		if query != "" {
			logQuery = query
		} else {
			logQuery = fmt.Sprintf(`{service_name="%s"}`, service)
		}

		baseURL := getLokiURL(env, region)
		
		// Output info to stderr
		fmt.Fprintf(os.Stderr, "Query: %s\n", logQuery)
		fmt.Fprintf(os.Stderr, "Environment: %s\n", env)
		fmt.Fprintf(os.Stderr, "Region: %s\n", region)
		fmt.Fprintf(os.Stderr, "Time range: %s\n", since)
		
		// Construct the query URL
		queryURL := fmt.Sprintf("%s/loki/api/v1/query_range", baseURL)
		
		// Create URL with query parameters
		u, err := url.Parse(queryURL)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing URL: %v\n", err)
			return
		}
		
		// Calculate time range
		now := time.Now()
		duration, err := time.ParseDuration(since)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing time range '%s': %v\n", since, err)
			return
		}
		start := now.Add(-duration)
		
		params := url.Values{}
		params.Add("query", logQuery)
		params.Add("start", fmt.Sprintf("%d", start.UnixNano()))
		params.Add("end", fmt.Sprintf("%d", now.UnixNano()))
		params.Add("limit", fmt.Sprintf("%d", limit))
		u.RawQuery = params.Encode()
		
		// Make HTTP request
		resp, err := http.Get(u.String())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error making request: %v\n", err)
			return
		}
		defer resp.Body.Close()
		
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading response: %v\n", err)
			return
		}
		
		// Parse JSON response
		var lokiResp LokiResponse
		err = json.Unmarshal(body, &lokiResp)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing JSON response: %v\n", err)
			return
		}
		
		// Display results
		if lokiResp.Status == "success" {
			// Pretty print the result to stdout
			resultJSON, err := json.MarshalIndent(lokiResp.Data.Result, "", "  ")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error formatting results: %v\n", err)
				return
			}
			fmt.Println(string(resultJSON))
		} else {
			fmt.Fprintf(os.Stderr, "Error: %s - %s\n", lokiResp.ErrorType, lokiResp.Error)
		}
	},
}

func getLokiURL(environment, region string) string {
	env := strings.ToLower(environment)
	reg := strings.ToLower(region)
	
	switch env {
	case "dev":
		switch reg {
		case "global":
			return "http://loki-query-frontend.query.dev.telnyx.io:3100"
		case "eu":
			return "http://loki-eu.query.dev.telnyx.io:3100"
		default:
			fmt.Fprintf(os.Stderr, "Warning: Unknown region '%s', using global\n", region)
			return "http://loki-query-frontend.query.dev.telnyx.io:3100"
		}
	case "prod":
		switch reg {
		case "global":
			return "http://loki-query-frontend.query.prod.telnyx.io:3100"
		case "eu":
			return "http://loki-eu.query.prod.telnyx.io:3100"
		default:
			fmt.Fprintf(os.Stderr, "Warning: Unknown region '%s', using global\n", region)
			return "http://loki-query-frontend.query.prod.telnyx.io:3100"
		}
	default:
		fmt.Fprintf(os.Stderr, "Warning: Unknown environment '%s', using prod global\n", environment)
		return "http://loki-query-frontend.query.prod.telnyx.io:3100"
	}
}