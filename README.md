# Telnyx CLI v2

A fully functional, production-ready command-line interface for Telnyx APIs.

## Features

### Phase 1 - Core Commands ✅
- **Auth**: Login, Whoami
- **Numbers**: Search, List, Order
- **Messages**: Send, List
- **Billing**: Balance

### Phase 2 - Advanced Commands ✅
- **Calls**: Dial, Hangup, List
- **Debugger**: List events, Inspect, Resend webhooks
- **Lookup**: Phone number lookup (caller name, carrier, portability)
- **SIMs**: List, Show, Update IoT SIM cards
- **Fax**: Send, List, Status, Cancel
- **Verify**: Send, Check, Status, List (2FA)

### Phase 3 - Generic API Support ✅
- **Generic CRUD**: List, Get, Create, Update, Delete for 17+ API resources
- **Resources Command**: View all available API endpoints

## Installation

```bash
cd telnyx-cli
npm install
```

## Usage

### Getting Started

```bash
# Show the welcome menu
telnyx

# Authenticate
telnyx auth login

# Check account info
telnyx auth whoami

# View all commands
telnyx --help
```

### Phone Numbers

```bash
# Search for available numbers
telnyx number search --country US --locality "Chicago"
telnyx number search --country US --area-code 312

# List your numbers
telnyx number list

# Purchase a number
telnyx number order +13125551234
```

### Messaging

```bash
# Send SMS (interactive)
telnyx message send

# Send SMS (direct)
telnyx message send --to +13125551234 --from +13125555678 --text "Hello!"

# List recent messages
telnyx message list
```

### Voice Calls

```bash
# Make a call (interactive)
telnyx call dial

# Make a call (direct)
telnyx call dial --from +13125555678 --to +13125551234 --connection <conn_id>

# List active calls
telnyx call list --active

# Hang up a call
telnyx call hangup <call_control_id>
```

### Number Lookup

```bash
# Lookup a number (interactive)
telnyx lookup

# Lookup with specific type
telnyx lookup +13125551234 --type caller-name
telnyx lookup +13125551234 --type carrier
telnyx lookup +13125551234 --type portability
```

### IoT SIMs

```bash
# List SIM cards
telnyx sim list
telnyx sim list --active

# Show SIM details
telnyx sim show <sim_id>

# Update SIM tags
telnyx sim update <sim_id> --tags "device1,iot,production"
```

### Fax

```bash
# Send a fax (interactive)
telnyx fax send

# Send a fax (direct)
telnyx fax send --to +13125551234 --from +13125555678 --media https://example.com/doc.pdf

# List faxes
telnyx fax list
telnyx fax list --outbound

# Check fax status
telnyx fax status <fax_id>

# Cancel a pending fax
telnyx fax cancel <fax_id>
```

### 2FA Verification

```bash
# Send verification code (interactive)
telnyx verify send

# Send verification code (direct)
telnyx verify send --number +13125551234 --type sms

# Check/submit code (interactive)
telnyx verify check

# Check/submit code (direct)
telnyx verify check --number +13125551234 --code 123456

# List verifications
telnyx verify list
```

### Webhook Debugging

```bash
# List recent events
telnyx debugger list
telnyx debugger list --since 48

# List webhook deliveries
telnyx debugger list --webhooks
telnyx debugger list --failed

# Inspect an event
telnyx debugger inspect <event_id>

# Resend a failed webhook
telnyx debugger resend <webhook_id>
```

### Billing

```bash
# Check balance
telnyx billing balance
telnyx billing balance --json
```

### Generic API Resources

```bash
# List all available resources
telnyx resources

# Messaging profiles
telnyx messaging-profile list
telnyx messaging-profile get <id>

# SIP connections
telnyx connection list
telnyx connection get <id>

# TeXML applications
telnyx texml-application list
telnyx texml-application create --data '{"name": "My App"}'

# Call control applications
telnyx call-control-application list

# Storage buckets
telnyx bucket list

# AI assistants
telnyx assistant list

# And more...
```

## Interactive Mode

Most commands support interactive prompts when arguments are omitted:

```bash
# This will prompt for all required fields
telnyx call dial
telnyx message send
telnyx verify send
```

## Environment Variables

```bash
# Set API key via environment (optional)
export TELNYX_API_KEY="your-api-key"
```

## Configuration

The CLI stores your API key securely using the `conf` package:
- macOS: `~/Library/Preferences/telnyx-cli`
- Linux: `~/.config/telnyx-cli`
- Windows: `%APPDATA%/telnyx-cli`

## API Reference

This CLI uses the [Telnyx API v2](https://developers.telnyx.com/docs/api/v2/overview).

## Project Structure

```
telnyx-cli/
├── bin/
│   └── telnyx              # Main entry point
├── src/
│   ├── api/
│   │   └── client.js       # API client with all methods
│   ├── commands/
│   │   ├── auth.js         # Authentication commands
│   │   ├── billing.js      # Billing commands
│   │   ├── calls.js        # Voice call commands
│   │   ├── debugger.js     # Webhook debugging commands
│   │   ├── fax.js          # Fax commands
│   │   ├── generic.js      # Generic API resource commands
│   │   ├── lookup.js       # Number lookup commands
│   │   ├── messages.js     # Messaging commands
│   │   ├── numbers.js      # Phone number commands
│   │   ├── sims.js         # IoT SIM commands
│   │   └── verify.js       # 2FA verification commands
│   ├── config/
│   │   └── store.js        # Configuration storage
│   └── ui/
│       └── layout.js       # UI components and styling
├── package.json
└── README.md
```

## Command Summary

| Category | Commands | Status |
|----------|----------|--------|
| Auth | login, whoami | ✅ |
| Numbers | search, list, order | ✅ |
| Messages | send, list | ✅ |
| Calls | dial, hangup, list | ✅ |
| Debugger | list, inspect, resend | ✅ |
| Lookup | lookup, batch | ✅ |
| SIMs | list, show, update | ✅ |
| Fax | send, list, status, cancel | ✅ |
| Verify | send, check, status, list | ✅ |
| Billing | balance | ✅ |
| Generic | 17+ resource types | ✅ |

## License

ISC
