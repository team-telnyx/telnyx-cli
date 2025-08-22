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

	"github.com/spf13/cobra"
)

type PrometheusResponse struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string      `json:"resultType"`
		Result     interface{} `json:"result"`
	} `json:"data"`
	Error     string `json:"error,omitempty"`
	ErrorType string `json:"errorType,omitempty"`
}

func init() {
	prometheusCmd.Flags().StringP("query", "q", "", "PromQL query to execute")
	prometheusCmd.Flags().StringP("environment", "e", "prod", "Environment to query (dev or prod)")
	
	prometheusCmd.MarkFlagRequired("query")

	getCmd.AddCommand(prometheusCmd)
}

// prometheusCmd represents the prometheus command
var prometheusCmd = &cobra.Command{
	Use:   "prometheus",
	Short: "Query Prometheus using PromQL",
	Long: `Query Prometheus/Thanos using PromQL syntax. Examples:

# Query CPU usage
telnyx-cli get prometheus -q "cpu_usage_percent" -e prod

# Query with time range
telnyx-cli get prometheus -q "rate(http_requests_total[5m])" -e dev`,
	Run: func(cmd *cobra.Command, args []string) {
		query, _ := cmd.Flags().GetString("query")
		env, _ := cmd.Flags().GetString("environment")

		baseURL := getPrometheusURL(env)
		
		// Output info to stderr
		fmt.Fprintf(os.Stderr, "Query: %s\n", query)
		fmt.Fprintf(os.Stderr, "Environment: %s\n", env)
		
		// Construct the query URL
		queryURL := fmt.Sprintf("%s/api/v1/query", baseURL)
		
		// Create URL with query parameter
		u, err := url.Parse(queryURL)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing URL: %v\n", err)
			return
		}
		
		params := url.Values{}
		params.Add("query", query)
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
		var promResp PrometheusResponse
		err = json.Unmarshal(body, &promResp)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing JSON response: %v\n", err)
			return
		}
		
		// Display results
		if promResp.Status == "success" {
			// Pretty print the result to stdout
			resultJSON, err := json.MarshalIndent(promResp.Data.Result, "", "  ")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error formatting results: %v\n", err)
				return
			}
			fmt.Println(string(resultJSON))
		} else {
			fmt.Fprintf(os.Stderr, "Error: %s - %s\n", promResp.ErrorType, promResp.Error)
		}
	},
}

func getPrometheusURL(environment string) string {
	env := strings.ToLower(environment)
	switch env {
	case "dev":
		return "https://thanosdev.internal.telnyx.com"
	case "prod":
		return "https://thanos.internal.telnyx.com"
	default:
		fmt.Fprintf(os.Stderr, "Warning: Unknown environment '%s', using prod\n", environment)
		return "https://thanos.internal.telnyx.com"
	}
}