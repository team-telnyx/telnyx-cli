# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Telnyx CLI is an internal operations toolkit for managing Telnyx infrastructure services. Built on Cobra (command framework) and Viper (configuration), it provides commands for querying Consul, managing service instances, accessing logs/metrics, and toggling maintenance modes.

## Build and Development Commands

```bash
# Build the project (creates ./telnyx-cli executable)
go build

# Install to $GOPATH/bin or $HOME/go/bin
go install

# Run tests
go test ./...

# Run the locally built binary
./telnyx-cli [command]

# Initialize CLI configuration (creates ~/.telnyx-cli/)
telnyx-cli init
```

## Architecture

### Command Structure

The CLI uses Cobra's hierarchical command pattern:

- **Root command** (`cmd/root.go`): Base command with global config initialization
- **Top-level commands** (`cmd/*.go`): `init`, `maintenance`, `deregister`
- **Nested commands** (`cmd/get/*.go`): All `get` subcommands like `get user`, `get logs`, `get prometheus`, `get service-diff`

Commands register themselves via `init()` functions that call `RootCmd.AddCommand()` or parent command's `AddCommand()`.

Main entry point (`main.go`) imports command packages with blank imports to trigger their `init()` functions.

### Configuration System

- **Config location**: `~/.telnyx-cli/config.yaml`
- **Cache location**: `~/.telnyx-cli/cache/`
- **Config struct** (`config/config.go`): Defines Consul and Private API URLs for dev/prod environments, Metaservice URL, and GitHub token
- **Viper integration**: Loads YAML config and supports environment variable overrides

### Service Clients

**Consul Client** (`consul/client.go`):
- Fetches datacenters, services, and instances across dev/prod environments
- Parallel fetching using goroutines and channels for multi-datacenter queries
- Caching layer for service lists (`consul/cache.go`)
- Instance version checking via HTTP `/version` endpoint
- Maintenance mode enable/disable operations

**Private API Client** (`privateapi/client.go`):
- User lookups, number queries, and organization data

**Metaservice Client** (`metaservice/client.go`):
- Deployment and connection information
- Repository name lookup for GitHub integration

**GitHub Client** (`github/client.go`):
- Compares commits between versions via GitHub API
- Extracts Jira ticket IDs from commit messages
- Parses version strings to extract commit SHAs

**Tailscale Integration** (`tailscale/user.go`):
- Retrieves current Tailscale user for audit trails

### Key Patterns

**Environment handling**: Commands typically accept `-e/--environment` flags for `dev` or `prod`, mapped to different Consul/API endpoints.

**Datacenter queries**: Most service operations involve querying multiple datacenters concurrently. The pattern is:
1. Fetch datacenter list for environment
2. Spawn goroutine per datacenter
3. Aggregate results via channels
4. Sort and present unified view

**Shell completion**: Commands implement `ValidArgsFunction` for autocomplete of services, datacenters, and other dynamic values.

**Progress indicators**: Long-running operations use `progressbar` package for visual feedback.

**User confirmation**: Destructive operations (like maintenance mode) require user confirmation via stdin.

## Common Development Workflows

**Adding a new `get` subcommand**:
1. Create file in `cmd/get/`
2. Define cobra.Command with Use, Short, Long, Run/RunE
3. Add to parent via `getCmd.AddCommand(yourCmd)` in `init()`
4. Implement flag handling and command logic
5. Add shell completion if appropriate

**Adding a new service client**:
1. Create package directory with `client.go`
2. Define HTTP client configuration
3. Implement API call functions
4. Use types for JSON marshaling/unmarshaling
5. Follow error handling pattern: return errors, let commands handle with `cobra.CheckErr()`

**Testing private repository access**: Requires GitHub personal access token in `~/.netrc` and `GOPRIVATE=github.com/team-telnyx` environment variable.

## Important Implementation Details

- **Datacenter filtering**: Only datacenters containing "dev" or "prod" in their names are considered valid
- **Consul port mapping**: Dev uses port 18500, prod uses 28500 for direct agent access
- **Service version endpoint**: Services expected to expose `/version` endpoint returning `{"version": "..."}`
- **Maintenance reasons**: Automatically prepend Tailscale user to maintenance reason for audit trails
- **Cache invalidation**: Service cache stored per environment, manually managed

## Testing Notes

- Mock external dependencies (Consul API, HTTP clients) in tests
- Test both positive and negative cases
- Ensure datacenter filtering logic is covered
- Test concurrent operations don't race

## Commit Style

- Present tense, imperative mood ("Add feature" not "Added feature")
- First line ≤72 characters
- Reference issues/PRs after first line