/*
Copyright © Telnyx LLC

*/
package get

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
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

type LokiStream struct {
	Stream map[string]string `json:"stream"`
	Values [][]string        `json:"values"`
}

type LokiTailResponse struct {
	Streams []LokiStream `json:"streams"`
}

func init() {
	logsCmd.Flags().StringP("service", "s", "", "Service name to query logs for (uses service_name label)")
	logsCmd.Flags().StringP("query", "q", "", "LogQL query to execute")
	logsCmd.Flags().StringP("environment", "e", "prod", "Environment to query (dev or prod)")
	logsCmd.Flags().StringP("region", "r", "global", "Region to query (global or eu)")
	logsCmd.Flags().StringP("since", "", "1h", "Time range to query (e.g., 1h, 30m, 1d)")
	logsCmd.Flags().IntP("limit", "l", 100, "Maximum number of log entries to return")
	logsCmd.Flags().BoolP("follow", "f", false, "Follow log stream in real-time")

	getCmd.AddCommand(logsCmd)
}

// logsCmd represents the logs command
var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Query Loki for logs using LogQL",
	Long: `Query Loki for logs using LogQL syntax. Examples:

# Query logs for a specific service
telnyx-cli get logs -s "my-service" -e prod

# Follow logs in real-time (like tail -f)
telnyx-cli get logs -s "my-service" -f -e prod

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
		follow, _ := cmd.Flags().GetBool("follow")

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
		
		if follow {
			fmt.Fprintf(os.Stderr, "Following logs... (Press Ctrl+C to stop)\n")
			err := followLogs(baseURL, logQuery, limit)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error following logs: %v\n", err)
			}
		} else {
			fmt.Fprintf(os.Stderr, "Time range: %s\n", since)
			err := queryHistoricalLogs(baseURL, logQuery, since, limit)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error querying logs: %v\n", err)
			}
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

func queryHistoricalLogs(baseURL, logQuery, since string, limit int) error {
	// Construct the query URL
	queryURL := fmt.Sprintf("%s/loki/api/v1/query_range", baseURL)
	
	// Create URL with query parameters
	u, err := url.Parse(queryURL)
	if err != nil {
		return fmt.Errorf("parsing URL: %v", err)
	}
	
	// Calculate time range
	now := time.Now()
	duration, err := time.ParseDuration(since)
	if err != nil {
		return fmt.Errorf("parsing time range '%s': %v", since, err)
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
		return fmt.Errorf("making request: %v", err)
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading response: %v", err)
	}
	
	// Parse JSON response
	var lokiResp LokiResponse
	err = json.Unmarshal(body, &lokiResp)
	if err != nil {
		return fmt.Errorf("parsing JSON response: %v", err)
	}
	
	// Display results
	if lokiResp.Status == "success" {
		// Pretty print the result to stdout
		resultJSON, err := json.MarshalIndent(lokiResp.Data.Result, "", "  ")
		if err != nil {
			return fmt.Errorf("formatting results: %v", err)
		}
		fmt.Println(string(resultJSON))
	} else {
		return fmt.Errorf("%s - %s", lokiResp.ErrorType, lokiResp.Error)
	}
	
	return nil
}

func followLogs(baseURL, logQuery string, limit int) error {
	// Set up signal handling for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	
	go func() {
		<-signalCh
		cancel()
	}()
	
	// Use polling approach since /tail endpoint might not be available
	// Keep track of the last timestamp we've seen to avoid duplicates
	lastTimestamp := time.Now().UnixNano()
	pollInterval := 1 * time.Second
	
	// Set of seen log entries to avoid duplicates
	seenEntries := make(map[string]bool)
	
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}
		
		// Query for recent logs
		now := time.Now()
		start := time.Unix(0, lastTimestamp)
		
		queryURL := fmt.Sprintf("%s/loki/api/v1/query_range", baseURL)
		u, err := url.Parse(queryURL)
		if err != nil {
			return fmt.Errorf("parsing URL: %v", err)
		}
		
		params := url.Values{}
		params.Add("query", logQuery)
		params.Add("start", fmt.Sprintf("%d", start.UnixNano()))
		params.Add("end", fmt.Sprintf("%d", now.UnixNano()))
		params.Add("limit", fmt.Sprintf("%d", limit))
		params.Add("direction", "forward")
		u.RawQuery = params.Encode()
		
		// Create HTTP request with context
		req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
		if err != nil {
			return fmt.Errorf("creating request: %v", err)
		}
		
		client := &http.Client{
			Timeout: 10 * time.Second,
		}
		resp, err := client.Do(req)
		if err != nil {
			// If we can't reach the server, wait and retry
			time.Sleep(pollInterval)
			continue
		}
		
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			time.Sleep(pollInterval)
			continue
		}
		
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			time.Sleep(pollInterval)
			continue
		}
		
		// Parse JSON response
		var lokiResp LokiResponse
		err = json.Unmarshal(body, &lokiResp)
		if err != nil {
			time.Sleep(pollInterval)
			continue
		}
		
		// Process results if successful
		if lokiResp.Status == "success" {
			// Parse the result to extract streams
			resultBytes, _ := json.Marshal(lokiResp.Data.Result)
			var streams []LokiStream
			if err := json.Unmarshal(resultBytes, &streams); err == nil {
				// Track the newest timestamp
				newestTimestamp := lastTimestamp
				
				// Process streams and print new entries
				for _, stream := range streams {
					for _, value := range stream.Values {
						if len(value) >= 2 {
							timestamp := value[0]
							logLine := value[1]
							
							// Create a unique key for this log entry
							entryKey := timestamp + ":" + logLine
							if seenEntries[entryKey] {
								continue // Skip duplicates
							}
							seenEntries[entryKey] = true
							
							// Convert nanosecond timestamp to readable format and track newest
							if ts, err := strconv.ParseInt(timestamp, 10, 64); err == nil {
								if ts > newestTimestamp {
									newestTimestamp = ts
								}
								t := time.Unix(0, ts)
								fmt.Printf("%s %s\n", t.Format("2006-01-02T15:04:05.000Z"), logLine)
							} else {
								fmt.Printf("%s %s\n", timestamp, logLine)
							}
						}
					}
				}
				
				// Update lastTimestamp if we found newer entries
				if newestTimestamp > lastTimestamp {
					lastTimestamp = newestTimestamp + 1 // Add 1ns to avoid getting the same entry again
				}
				
				// Clean up old entries from seen map to prevent memory growth
				if len(seenEntries) > 1000 {
					// Keep only the last 500 entries
					newSeen := make(map[string]bool)
					count := 0
					for k, v := range seenEntries {
						if count < 500 {
							newSeen[k] = v
							count++
						}
					}
					seenEntries = newSeen
				}
			}
		}
		
		// Wait before next poll
		time.Sleep(pollInterval)
	}
}