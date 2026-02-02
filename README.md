# Telnyx CLI

[![CI](https://github.com/team-telnyx/telnyx-cli/actions/workflows/ci.yml/badge.svg)](https://github.com/team-telnyx/telnyx-cli/actions/workflows/ci.yml)
[![npm version](https://badge.fury.io/js/@telnyx%2Fcli.svg)](https://www.npmjs.com/package/@telnyx/cli)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

A powerful command-line interface for the [Telnyx API](https://developers.telnyx.com). Manage phone numbers, send messages, make calls, and more — all from your terminal.

## Installation

```bash
npm install -g @telnyx/cli
```

Or with yarn:

```bash
yarn global add @telnyx/cli
```

**Requirements:** Node.js 18 or higher

## Quick Start

```bash
# Authenticate with your API key
telnyx auth login

# Check your account
telnyx auth whoami

# Search for phone numbers
telnyx number search --country US --locality "San Francisco"

# Send an SMS
telnyx message send --from +15551234567 --to +15559876543 --text "Hello from Telnyx CLI!"
```

## Features

- **Interactive Prompts** — Guided workflows when arguments are omitted
- **Multi-Profile Support** — Manage multiple API keys (sandbox, production, etc.)
- **Human-Friendly Output** — Colored tables, spinners, and clear formatting
- **Flexible Output Formats** — `--json`, `--output csv`, or default tables
- **Verbose Mode** — Debug HTTP requests with `--verbose`
- **Dry-Run Mode** — Preview destructive operations before executing

## Commands

### Authentication

```bash
telnyx auth login              # Authenticate with API key
telnyx auth logout             # Clear stored credentials
telnyx auth whoami             # Show current account info

# Multi-profile support
telnyx profile list            # List all profiles
telnyx profile add production  # Add a new profile
telnyx profile use production  # Switch active profile
```

### Phone Numbers

```bash
telnyx number search --country US --area-code 415
telnyx number list
telnyx number order +14155551234
telnyx number delete <number_id> --dry-run
```

### Messaging

```bash
telnyx message send --from +15551234567 --to +15559876543 --text "Hello!"
telnyx message list --limit 10
```

### Voice Calls

```bash
telnyx call dial --from +15551234567 --to +15559876543 --connection <id>
telnyx call list --active
telnyx call hangup <call_control_id>
```

### Additional Commands

| Category | Commands |
|----------|----------|
| **Billing** | `balance` |
| **Debugger** | `list`, `inspect`, `resend` |
| **Lookup** | `lookup` (carrier, caller-name, portability) |
| **SIMs** | `list`, `show`, `update` |
| **Fax** | `send`, `list`, `status`, `cancel` |
| **Verify** | `send`, `check`, `status`, `list` |

Run `telnyx --help` for the full command list, or `telnyx <command> --help` for details.

## Global Options

| Option | Description |
|--------|-------------|
| `-p, --profile <name>` | Use a specific profile |
| `-v, --verbose` | Show HTTP request/response details |
| `--timeout <seconds>` | Request timeout (default: 30) |
| `--json` | Output raw JSON |
| `-o, --output <format>` | Output format: `json`, `table`, `csv` |

## Configuration

Credentials are stored securely using your OS keychain:

| OS | Location |
|----|----------|
| macOS | `~/Library/Preferences/@telnyx/cli` |
| Linux | `~/.config/@telnyx/cli` |
| Windows | `%APPDATA%/@telnyx/cli` |

You can also set `TELNYX_API_KEY` as an environment variable.

## API Reference

This CLI wraps the [Telnyx API v2](https://developers.telnyx.com/docs/api/v2/overview). Get your API key from the [Telnyx Portal](https://portal.telnyx.com/#/app/api-keys).

## Contributing

Contributions are welcome! Please read our [contributing guidelines](CONTRIBUTING.md) before submitting a PR.

```bash
# Clone the repo
git clone https://github.com/team-telnyx/telnyx-cli.git
cd telnyx-cli

# Install dependencies
npm install

# Run tests
npm test

# Run in development
node bin/telnyx --help
```

## License

[MIT](LICENSE) © Telnyx
