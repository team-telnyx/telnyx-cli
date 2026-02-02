# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [2.0.0] - 2026-02-02

### Added
- Complete rewrite in Node.js for improved maintainability
- Interactive CLI with guided prompts using Inquirer
- Multi-profile support for managing multiple API keys
- Human-friendly colored output with tables and spinners
- Verbose mode (`--verbose`) for debugging HTTP requests
- Configurable timeout (`--timeout`)
- Dry-run mode for destructive operations
- Consistent output formats (`--json`, `--output csv|table`)

### Commands
- **Auth**: `login`, `logout`, `whoami`
- **Profile**: `list`, `add`, `use`, `remove`
- **Numbers**: `search`, `list`, `order`, `delete`
- **Messages**: `send`, `list`
- **Calls**: `dial`, `hangup`, `list`
- **Billing**: `balance`
- **Debugger**: `list`, `inspect`, `resend`
- **Lookup**: number lookup with carrier/caller-name info
- **SIMs**: `list`, `show`, `update`
- **Fax**: `send`, `list`, `status`, `cancel`
- **Verify**: `send`, `check`, `status`, `list`
- **Generic**: CRUD for 17+ API resource types

### Changed
- Migrated from auto-generated Go CLI to hand-crafted Node.js
- Improved error messages with actionable suggestions
- Secure credential storage using `conf` package

### Security
- API keys stored securely in OS-specific config directories
- No credentials logged in verbose mode

## [1.x] - Legacy

Previous versions were auto-generated from the OpenAPI spec in Go.
See the [v1 branch](https://github.com/team-telnyx/telnyx-cli/tree/v1) for historical releases.
