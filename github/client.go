package github

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/google/go-github/v58/github"
	"github.com/spf13/cobra"
)

// ComparisonResult contains structured information about a GitHub commit comparison.
type ComparisonResult struct {
	HTMLURL      string
	Commits      []CommitInfo
	TotalCommits int
}

// CommitInfo represents a single commit with its metadata.
type CommitInfo struct {
	SHA     string
	Message string
	Author  string
}

var jiraTicketRegex = regexp.MustCompile(`([A-Z]+-\d+)`)

// CompareCommits fetches the GitHub comparison between two commits and returns
// structured information about the diff including commit messages and authors.
// Returns an error if the token is invalid, repository is not found, or network fails.
func CompareCommits(token, repo, oldCommit, newCommit string) (*ComparisonResult, error) {
	if token == "" {
		return nil, fmt.Errorf("GitHub token not configured")
	}

	ctx := context.Background()
	client := github.NewClient(nil).WithAuthToken(token)

	// Compare commits
	comparison, resp, err := client.Repositories.CompareCommits(
		ctx,
		"team-telnyx",
		repo,
		oldCommit,
		newCommit,
		nil,
	)

	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusUnauthorized {
			return nil, fmt.Errorf("invalid GitHub token - regenerate at https://github.com/settings/tokens")
		}
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			return nil, fmt.Errorf("repository '%s' or commits not found on GitHub", repo)
		}
		return nil, fmt.Errorf("failed to compare commits: %w", err)
	}

	result := &ComparisonResult{
		HTMLURL:      comparison.GetHTMLURL(),
		TotalCommits: comparison.GetTotalCommits(),
		Commits:      make([]CommitInfo, 0, len(comparison.Commits)),
	}

	for _, commit := range comparison.Commits {
		info := CommitInfo{
			SHA:     commit.GetSHA(),
			Message: extractFirstLine(commit.GetCommit().GetMessage()),
			Author:  commit.GetCommit().GetAuthor().GetName(),
		}
		result.Commits = append(result.Commits, info)
	}

	return result, nil
}

// ExtractJiraTickets extracts unique Jira ticket IDs (e.g., INFRA-123) from commit messages
// using regex pattern matching. Returns a deduplicated list of ticket IDs.
func ExtractJiraTickets(messages []string) []string {
	ticketSet := make(map[string]bool)
	tickets := []string{}

	for _, msg := range messages {
		matches := jiraTicketRegex.FindAllString(msg, -1)
		for _, ticket := range matches {
			if !ticketSet[ticket] {
				ticketSet[ticket] = true
				tickets = append(tickets, ticket)
			}
		}
	}

	return tickets
}

// ExtractCommitFromVersion extracts the commit SHA from a version string.
// Version format: YYYY.MM.DD.HH.MM.COMMIT_SHA (e.g., 2025.09.30.09.21.ab4b47e)
// Returns the last segment after the final dot.
func ExtractCommitFromVersion(version string) string {
	// Version format: YYYY.MM.DD.HH.MM.COMMIT_SHA
	// Extract last segment after final dot
	parts := strings.Split(version, ".")
	if len(parts) == 0 {
		cobra.CheckErr(fmt.Errorf("invalid version format: %s", version))
	}
	return parts[len(parts)-1]
}

func extractFirstLine(message string) string {
	lines := strings.Split(message, "\n")
	if len(lines) > 0 {
		return strings.TrimSpace(lines[0])
	}
	return message
}
