/*
Copyright © Telnyx LLC
*/
package get

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/team-telnyx/telnyx-cli/consul"
	ghclient "github.com/team-telnyx/telnyx-cli/github"
	"github.com/team-telnyx/telnyx-cli/metaservice"
)

// commandOptions holds all the command flags
type commandOptions struct {
	service        string
	oldVersionFlag string
	newVersionFlag string
	oldEnv         string
	newEnv         string
	force          bool
	format         string
	githubToken    string
}

// fetchResult holds the results of parallel data fetching
type fetchResult struct {
	oldVersions *consul.VersionsByDatacenter
	newVersions *consul.VersionsByDatacenter
	repoName    string
	errors      []error
}

// DiffOutput represents the structured output for JSON format
type DiffOutput struct {
	Service      string                 `json:"service"`
	Repository   string                 `json:"repository"`
	OldVersion   VersionInfo            `json:"old_version"`
	NewVersion   VersionInfo            `json:"new_version"`
	Comparison   *ComparisonOutput      `json:"comparison,omitempty"`
	Error        string                 `json:"error,omitempty"`
}

type VersionInfo struct {
	Environment  string   `json:"environment"`
	Version      string   `json:"version"`
	Commit       string   `json:"commit"`
	Datacenters  int      `json:"datacenters"`
	Consistent   bool     `json:"consistent"`
	Conflicts    []string `json:"conflicts,omitempty"`
}

type ComparisonOutput struct {
	URL          string   `json:"url"`
	TotalCommits int      `json:"total_commits"`
	Commits      []string `json:"commits"`
	Authors      []string `json:"authors"`
	JiraTickets  []string `json:"jira_tickets,omitempty"`
}

func init() {
	serviceDiffCmd.Flags().StringP("old-version", "o", "", "Manual override for old version")
	serviceDiffCmd.Flags().StringP("new-version", "n", "", "Manual override for new version")
	serviceDiffCmd.Flags().String("old-env", "prod", "Old environment (default: prod)")
	serviceDiffCmd.Flags().String("new-env", "dev", "New environment (default: dev)")
	serviceDiffCmd.Flags().Bool("force", false, "Proceed despite version conflicts")
	serviceDiffCmd.Flags().StringP("format", "f", "default", "Output format: default, plain, json")

	getCmd.AddCommand(serviceDiffCmd)
}

// serviceDiffCmd represents the service-diff command
var serviceDiffCmd = &cobra.Command{
	Use:   "service-diff <service>",
	Short: "Compare service versions between environments",
	Long: `
Compares service versions between environments (default: prod vs dev) and displays
the GitHub diff including commits, authors, and Jira tickets.

Examples:

# Compare prod (old) vs dev (new) versions
telnyx-cli get service-diff call-control

# Compare specific versions manually
telnyx-cli get service-diff call-control --old-version 2025.09.24.10.55.48d3bae --new-version 2025.09.30.09.21.ab4b47e

# Compare dev vs staging
telnyx-cli get service-diff call-control --old-env dev --new-env staging

# Output in plain format (parsable)
telnyx-cli get service-diff call-control --format plain

# Output in JSON format (programmatic)
telnyx-cli get service-diff call-control --format json
`,
	Args: cobra.ExactArgs(1),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		svcs := consul.GetServicesByEnv("prod")

		var completions []string
		for s := range svcs {
			if strings.HasPrefix(s, toComplete) {
				completions = append(completions, s)
			}
		}

		return completions, cobra.ShellCompDirectiveNoFileComp
	},
	Run: func(cmd *cobra.Command, args []string) {
		opts := parseCommandOptions(cmd, args)

		if !validateOptions(opts) {
			return
		}

		if opts.format == "default" {
			printService(opts.service)
		}

		result := fetchServiceData(opts)
		if !handleFetchErrors(result, opts.format, opts.service) {
			return
		}

		if opts.format == "default" {
			fmt.Printf("%s %s\n", color.YellowString("▶▶ Repository:"), result.repoName)
			fmt.Println()
		}

		oldVersion, newVersion, shouldContinue := resolveVersions(opts, result)
		if !shouldContinue {
			return
		}

		if oldVersion == newVersion {
			handleIdenticalVersions(opts, result, oldVersion, newVersion)
			return
		}

		comparison, err := fetchGitHubComparison(opts, result, oldVersion, newVersion)
		if err != nil {
			fmt.Printf("%s %s\n", color.RedString("✗ ERROR:"), err)
			return
		}

		displayDiff(opts, result, oldVersion, newVersion, comparison)
	},
}

// parseCommandOptions extracts all command flags into a struct
func parseCommandOptions(cmd *cobra.Command, args []string) commandOptions {
	oldVersionFlag, _ := cmd.Flags().GetString("old-version")
	newVersionFlag, _ := cmd.Flags().GetString("new-version")
	oldEnv, _ := cmd.Flags().GetString("old-env")
	newEnv, _ := cmd.Flags().GetString("new-env")
	force, _ := cmd.Flags().GetBool("force")
	format, _ := cmd.Flags().GetString("format")

	return commandOptions{
		service:        args[0],
		oldVersionFlag: oldVersionFlag,
		newVersionFlag: newVersionFlag,
		oldEnv:         oldEnv,
		newEnv:         newEnv,
		force:          force,
		format:         format,
		githubToken:    viper.GetString("github_token"),
	}
}

// validateOptions validates command options and returns false if invalid
func validateOptions(opts commandOptions) bool {
	if opts.format != "default" && opts.format != "plain" && opts.format != "json" {
		fmt.Printf("Invalid format: %s. Valid options: default, plain, json\n", opts.format)
		return false
	}

	if opts.githubToken == "" {
		printGitHubTokenError()
		return false
	}

	return true
}

// fetchServiceData fetches version and repository data in parallel
func fetchServiceData(opts commandOptions) fetchResult {
	result := fetchResult{errors: make([]error, 0)}
	var wg sync.WaitGroup
	var mu sync.Mutex

	// Fetch old version if not manually overridden
	if opts.oldVersionFlag == "" {
		wg.Add(1)
		go func() {
			defer wg.Done()
			v, err := consul.GetServiceVersions(opts.service, opts.oldEnv)
			mu.Lock()
			defer mu.Unlock()
			if err != nil {
				result.errors = append(result.errors, fmt.Errorf("old env (%s): %w", opts.oldEnv, err))
			} else {
				result.oldVersions = v
			}
		}()
	}

	// Fetch new version if not manually overridden
	if opts.newVersionFlag == "" {
		wg.Add(1)
		go func() {
			defer wg.Done()
			v, err := consul.GetServiceVersions(opts.service, opts.newEnv)
			mu.Lock()
			defer mu.Unlock()
			if err != nil {
				result.errors = append(result.errors, fmt.Errorf("new env (%s): %w", opts.newEnv, err))
			} else {
				result.newVersions = v
			}
		}()
	}

	// Fetch repository name
	wg.Add(1)
	go func() {
		defer wg.Done()
		r, err := metaservice.GetRepoName(opts.service)
		mu.Lock()
		defer mu.Unlock()
		if err != nil {
			result.errors = append(result.errors, err)
		} else {
			result.repoName = r
		}
	}()

	wg.Wait()
	return result
}

// handleFetchErrors handles errors from parallel fetching, returns false if errors occurred
func handleFetchErrors(result fetchResult, format, service string) bool {
	if len(result.errors) == 0 {
		return true
	}

	if format == "json" {
		errorOutput := DiffOutput{
			Service: service,
			Error:   fmt.Sprintf("%v", result.errors),
		}
		jsonData, _ := json.MarshalIndent(errorOutput, "", "  ")
		fmt.Println(string(jsonData))
	} else {
		for _, err := range result.errors {
			fmt.Printf("%s %s\n", color.RedString("✗ ERROR:"), err)
		}
	}
	return false
}

// resolveVersions determines which versions to compare, returns shouldContinue=false if conflicts found
func resolveVersions(opts commandOptions, result fetchResult) (oldVersion, newVersion string, shouldContinue bool) {
	// Resolve old version
	if opts.oldVersionFlag != "" {
		oldVersion = opts.oldVersionFlag
		if opts.format == "default" {
			fmt.Printf("%s %s: %s (manual override)\n",
				color.YellowString("▶▶ Version Analysis:"),
				color.CyanString("OLD"),
				oldVersion)
		}
	} else {
		if opts.format == "default" {
			printVersionAnalysis("OLD", opts.oldEnv, result.oldVersions, opts.force)
		}
		if result.oldVersions.HasConflict && !opts.force {
			if opts.format == "default" {
				printConflictSolutions(opts.oldEnv, result.oldVersions)
			}
			return "", "", false
		}
		oldVersion = result.oldVersions.UniqueVersion
	}

	// Resolve new version
	if opts.newVersionFlag != "" {
		newVersion = opts.newVersionFlag
		if opts.format == "default" {
			fmt.Printf("  %s: %s (manual override)\n",
				color.CyanString("NEW"),
				newVersion)
		}
	} else {
		if opts.format == "default" {
			printVersionAnalysis("NEW", opts.newEnv, result.newVersions, opts.force)
		}
		if result.newVersions.HasConflict && !opts.force {
			if opts.format == "default" {
				printConflictSolutions(opts.newEnv, result.newVersions)
			}
			return "", "", false
		}
		newVersion = result.newVersions.UniqueVersion
	}

	if opts.format == "default" {
		fmt.Println()
	}

	return oldVersion, newVersion, true
}

// handleIdenticalVersions handles the case when old and new versions are the same
func handleIdenticalVersions(opts commandOptions, result fetchResult, oldVersion, newVersion string) {
	if opts.format == "json" {
		oldCommit, err := ghclient.ExtractCommitFromVersion(oldVersion)
		if err != nil {
			errorOutput := map[string]string{
				"error": fmt.Sprintf("Invalid old version format: %v", err),
			}
			jsonData, _ := json.MarshalIndent(errorOutput, "", "  ")
			fmt.Println(string(jsonData))
			return
		}

		newCommit, err := ghclient.ExtractCommitFromVersion(newVersion)
		if err != nil {
			errorOutput := map[string]string{
				"error": fmt.Sprintf("Invalid new version format: %v", err),
			}
			jsonData, _ := json.MarshalIndent(errorOutput, "", "  ")
			fmt.Println(string(jsonData))
			return
		}

		identicalOutput := DiffOutput{
			Service:    opts.service,
			Repository: result.repoName,
			OldVersion: VersionInfo{
				Environment: opts.oldEnv,
				Version:     oldVersion,
				Commit:      oldCommit,
			},
			NewVersion: VersionInfo{
				Environment: opts.newEnv,
				Version:     newVersion,
				Commit:      newCommit,
			},
			Error: "Versions are identical",
		}
		if result.oldVersions != nil {
			identicalOutput.OldVersion.Datacenters = len(result.oldVersions.Versions)
			identicalOutput.OldVersion.Consistent = !result.oldVersions.HasConflict
		}
		if result.newVersions != nil {
			identicalOutput.NewVersion.Datacenters = len(result.newVersions.Versions)
			identicalOutput.NewVersion.Consistent = !result.newVersions.HasConflict
		}
		jsonData, _ := json.MarshalIndent(identicalOutput, "", "  ")
		fmt.Println(string(jsonData))
	} else if opts.format == "plain" {
		fmt.Printf("SERVICE\t%s\n", opts.service)
		fmt.Printf("REPOSITORY\t%s\n", result.repoName)
		fmt.Printf("OLD_VERSION\t%s\n", oldVersion)
		fmt.Printf("NEW_VERSION\t%s\n", newVersion)
		fmt.Printf("STATUS\tidentical\n")
	} else {
		fmt.Printf("%s Versions are identical - no diff needed\n", color.GreenString("✓"))
	}
}

// fetchGitHubComparison fetches the GitHub commit comparison
func fetchGitHubComparison(opts commandOptions, result fetchResult, oldVersion, newVersion string) (*ghclient.ComparisonResult, error) {
	oldCommit, err := ghclient.ExtractCommitFromVersion(oldVersion)
	if err != nil {
		return nil, fmt.Errorf("invalid old version format: %w", err)
	}

	newCommit, err := ghclient.ExtractCommitFromVersion(newVersion)
	if err != nil {
		return nil, fmt.Errorf("invalid new version format: %w", err)
	}

	return ghclient.CompareCommits(opts.githubToken, result.repoName, oldCommit, newCommit)
}

// displayDiff displays the diff in the requested format
func displayDiff(opts commandOptions, result fetchResult, oldVersion, newVersion string, comparison *ghclient.ComparisonResult) {
	oldCommit, err := ghclient.ExtractCommitFromVersion(oldVersion)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return
	}

	newCommit, err := ghclient.ExtractCommitFromVersion(newVersion)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return
	}

	switch opts.format {
	case "plain":
		printPlainDiff(opts.service, result.repoName, oldVersion, newVersion, oldCommit, newCommit, comparison)
	case "json":
		printJSONDiff(opts.service, result.repoName, opts.oldEnv, opts.newEnv, oldVersion, newVersion,
			result.oldVersions, result.newVersions, oldCommit, newCommit, comparison)
	default:
		printDiff(oldCommit, newCommit, comparison)
	}
}

func printVersionAnalysis(label, env string, versions *consul.VersionsByDatacenter, force bool) {
	if label == "OLD" {
		fmt.Printf("%s\n", color.YellowString("▶▶ Version Analysis:"))
	}

	if versions.HasConflict {
		if force {
			fmt.Printf("  %s (%s): %s (proceeding with --force)\n",
				color.CyanString(label),
				env,
				color.YellowString("⚠ CONFLICT"))
		} else {
			fmt.Printf("  %s (%s): %s\n",
				color.CyanString(label),
				env,
				color.RedString("✗ CONFLICT DETECTED"))
		}
	} else {
		fmt.Printf("  %s (%s): %s\n",
			color.CyanString(label),
			env,
			color.GreenString("✓ Consistent"))
		fmt.Printf("    Version: %s (%d datacenters)\n",
			versions.UniqueVersion,
			len(versions.Versions))
		commit, err := ghclient.ExtractCommitFromVersion(versions.UniqueVersion)
		if err != nil {
			fmt.Printf("    Commit:  %s\n", color.RedString("invalid version format"))
		} else {
			fmt.Printf("    Commit:  %s\n", commit)
		}
	}
}

func printConflictSolutions(env string, versions *consul.VersionsByDatacenter) {
	fmt.Println()

	// Group datacenters by version
	versionToDcs := make(map[string][]string)
	for dc, version := range versions.Versions {
		versionToDcs[version] = append(versionToDcs[version], dc)
	}

	// Sort versions by datacenter count (descending)
	type versionCount struct {
		version string
		dcs     []string
	}
	var versionCounts []versionCount
	for v, dcs := range versionToDcs {
		sort.Strings(dcs)
		versionCounts = append(versionCounts, versionCount{version: v, dcs: dcs})
	}
	sort.Slice(versionCounts, func(i, j int) bool {
		return len(versionCounts[i].dcs) > len(versionCounts[j].dcs)
	})

	// Print breakdown
	for _, vc := range versionCounts {
		dcList := strings.Join(vc.dcs, ", ")
		if len(dcList) > 80 {
			// Truncate long lists
			dcList = strings.Join(vc.dcs[:5], ", ") + fmt.Sprintf(", ... [+%d more]", len(vc.dcs)-5)
		}
		fmt.Printf("    Version %s (%d DCs):\n      %s\n\n",
			vc.version,
			len(vc.dcs),
			dcList)
	}

	fmt.Println(strings.Repeat("─", 60))
	fmt.Printf("%s Version conflict - incomplete deployment detected\n\n",
		color.RedString("✗ ERROR:"))
	fmt.Println("Solutions:")
	fmt.Printf("  1. Complete deployment across all datacenters in %s\n", env)
	fmt.Printf("  2. Specify version manually: --old-version %s\n", versionCounts[0].version)
	fmt.Println("  3. Force (not recommended): --force")
	fmt.Println(strings.Repeat("─", 60))
}

func printDiff(oldCommit, newCommit string, comparison *ghclient.ComparisonResult) {
	fmt.Println(strings.Repeat("═", 60))
	fmt.Printf("GitHub: %s\n", color.CyanString(comparison.HTMLURL))
	fmt.Println(strings.Repeat("═", 60))
	fmt.Println()

	fmt.Printf("Commits (%s → %s):\n", oldCommit[:7], newCommit[:7])

	// Extract unique authors and messages for Jira tickets
	authorSet := make(map[string]bool)
	messages := make([]string, 0, len(comparison.Commits))

	for _, commit := range comparison.Commits {
		fmt.Printf("  %s %s\n", commit.SHA[:7], commit.Message)
		authorSet[commit.Author] = true
		messages = append(messages, commit.Message)
	}

	fmt.Println()
	fmt.Println(strings.Repeat("─", 60))

	// Print authors
	authors := make([]string, 0, len(authorSet))
	for author := range authorSet {
		authors = append(authors, author)
	}
	sort.Strings(authors)
	fmt.Printf("Authors: %s\n", strings.Join(authors, ", "))

	// Print Jira tickets
	tickets := ghclient.ExtractJiraTickets(messages)
	if len(tickets) > 0 {
		fmt.Printf("Jira: %s\n", strings.Join(tickets, ", "))
	}

	fmt.Println(strings.Repeat("─", 60))
}

func printGitHubTokenError() {
	fmt.Printf("%s GitHub token not configured\n\n", color.RedString("✗ ERROR:"))
	fmt.Println("Configure your GitHub token:")
	fmt.Println("  1. Create token: https://github.com/settings/tokens (scope: repo)")
	fmt.Println("  2. Edit ~/.telnyx-cli/config.yaml:")
	fmt.Println("     github_token: your_token_here")
	fmt.Println()
	fmt.Println("See README for detailed instructions.")
}

func printPlainDiff(service, repo, oldVersion, newVersion, oldCommit, newCommit string, comparison *ghclient.ComparisonResult) {
	// Plain format: tab-separated values, easy to parse with awk/cut
	fmt.Printf("SERVICE\t%s\n", service)
	fmt.Printf("REPOSITORY\t%s\n", repo)
	fmt.Printf("OLD_VERSION\t%s\n", oldVersion)
	fmt.Printf("OLD_COMMIT\t%s\n", oldCommit)
	fmt.Printf("NEW_VERSION\t%s\n", newVersion)
	fmt.Printf("NEW_COMMIT\t%s\n", newCommit)
	fmt.Printf("GITHUB_URL\t%s\n", comparison.HTMLURL)
	fmt.Printf("TOTAL_COMMITS\t%d\n", comparison.TotalCommits)

	// Extract unique authors and messages for Jira tickets
	authorSet := make(map[string]bool)
	messages := make([]string, 0, len(comparison.Commits))

	for _, commit := range comparison.Commits {
		fmt.Printf("COMMIT\t%s\t%s\n", commit.SHA[:7], commit.Message)
		authorSet[commit.Author] = true
		messages = append(messages, commit.Message)
	}

	// Print authors
	authors := make([]string, 0, len(authorSet))
	for author := range authorSet {
		authors = append(authors, author)
	}
	sort.Strings(authors)
	for _, author := range authors {
		fmt.Printf("AUTHOR\t%s\n", author)
	}

	// Print Jira tickets
	tickets := ghclient.ExtractJiraTickets(messages)
	for _, ticket := range tickets {
		fmt.Printf("JIRA\t%s\n", ticket)
	}
}

func printJSONDiff(service, repo, oldEnv, newEnv, oldVersion, newVersion string,
	oldVersions, newVersions *consul.VersionsByDatacenter,
	oldCommit, newCommit string, comparison *ghclient.ComparisonResult) {

	output := DiffOutput{
		Service:    service,
		Repository: repo,
		OldVersion: VersionInfo{
			Environment: oldEnv,
			Version:     oldVersion,
			Commit:      oldCommit,
			Consistent:  true,
		},
		NewVersion: VersionInfo{
			Environment: newEnv,
			Version:     newVersion,
			Commit:      newCommit,
			Consistent:  true,
		},
	}

	// Add datacenter info if available (not manual override)
	if oldVersions != nil {
		output.OldVersion.Datacenters = len(oldVersions.Versions)
		output.OldVersion.Consistent = !oldVersions.HasConflict
		if oldVersions.HasConflict {
			conflicts := make([]string, 0)
			for dc, ver := range oldVersions.Versions {
				conflicts = append(conflicts, fmt.Sprintf("%s:%s", dc, ver))
			}
			output.OldVersion.Conflicts = conflicts
		}
	}

	if newVersions != nil {
		output.NewVersion.Datacenters = len(newVersions.Versions)
		output.NewVersion.Consistent = !newVersions.HasConflict
		if newVersions.HasConflict {
			conflicts := make([]string, 0)
			for dc, ver := range newVersions.Versions {
				conflicts = append(conflicts, fmt.Sprintf("%s:%s", dc, ver))
			}
			output.NewVersion.Conflicts = conflicts
		}
	}

	// Add comparison data
	if comparison != nil {
		authorSet := make(map[string]bool)
		commits := make([]string, 0, len(comparison.Commits))
		messages := make([]string, 0, len(comparison.Commits))

		for _, commit := range comparison.Commits {
			commits = append(commits, fmt.Sprintf("%s %s", commit.SHA[:7], commit.Message))
			authorSet[commit.Author] = true
			messages = append(messages, commit.Message)
		}

		authors := make([]string, 0, len(authorSet))
		for author := range authorSet {
			authors = append(authors, author)
		}
		sort.Strings(authors)

		output.Comparison = &ComparisonOutput{
			URL:          comparison.HTMLURL,
			TotalCommits: comparison.TotalCommits,
			Commits:      commits,
			Authors:      authors,
			JiraTickets:  ghclient.ExtractJiraTickets(messages),
		}
	}

	jsonData, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		fmt.Printf("Error generating JSON: %v\n", err)
		return
	}

	fmt.Println(string(jsonData))
}
