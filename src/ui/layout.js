const chalk = require('chalk');

// Telnyx Brand Colors
const COLORS = {
  primary: '#00E3AA',
  white: '#FFFFFF',
  gray: '#6B7280',
  dark: '#1F2937',
  yellow: '#F59E0B',
  red: '#EF4444'
};

// Create chalk instances with custom colors
const primary = chalk.hex(COLORS.primary);
const white = chalk.white;
const gray = chalk.gray;
const yellow = chalk.hex(COLORS.yellow);
const red = chalk.hex(COLORS.red);

// ASCII Art Banner
const ASCII_BANNER = `
    ████████╗███████╗██╗     ███╗   ██╗██╗   ██╗██╗  ██╗
    ╚══██╔══╝██╔════╝██║     ████╗  ██║╚██╗ ██╔╝╚██╗██╔╝
       ██║   █████╗  ██║     ██╔██╗ ██║ ╚████╔╝  ╚███╔╝ 
       ██║   ██╔══╝  ██║     ██║╚██╗██║  ╚██╔╝   ██╔██╗ 
       ██║   ███████╗███████╗██║ ╚████║   ██║   ██╔╝ ██╗
       ╚═╝   ╚══════╝╚══════╝╚═╝  ╚═══╝   ╚═╝   ╚═╝  ╚═╝
                                                            
       ${white('Command-line interface for Telnyx APIs')}
                                                            
       ${primary('Voice')} • ${primary('Messaging')} • ${primary('AI')} • ${primary('IoT')} • ${primary('Storage')}
`;

// Dynamic Menu Box - shows actual available commands
const MENU_BOX = `
┌──────────────────────────────────────────────────────────┐
│  ${white('QUICK COMMANDS')}                                          │
├──────────────────────────────────────────────────────────┤
│                                                          │
│  ${primary('Numbers')}                                                │
│  ${primary('  search')}    Search for available phone numbers        │
│  ${primary('  list')}      List your purchased numbers              │
│  ${primary('  order')}     Purchase a phone number                   │
│                                                          │
│  ${primary('Messaging')}                                              │
│  ${primary('  send')}      Send SMS/MMS messages                     │
│  ${primary('  list')}      View recent messages                      │
│                                                          │
│  ${primary('Voice')}                                                  │
│  ${primary('  dial')}      Make an outbound call                      │
│  ${primary('  hangup')}    End an active call                         │
│  ${primary('  list')}      List active calls                          │
│                                                          │
│  ${primary('Utilities')}                                              │
│  ${primary('  lookup')}    Phone number lookup & info                 │
│  ${primary('  sim list')}  IoT SIM cards                              │
│  ${primary('  fax send')}  Send a fax                                 │
│  ${primary('  verify')}   2FA verification codes                      │
│  ${primary('  debugger')}  Webhook debugging                          │
│                                                          │
│  ${primary('Account')}                                               │
│  ${primary('  auth')}      Login / Whoami                            │
│  ${primary('  billing')}   Check balance & billing                   │
│                                                          │
├──────────────────────────────────────────────────────────┤
│  ${gray('Examples:')}                                               │
│  ${gray('  telnyx number search --country US')}                     │
│  ${gray('  telnyx message send --to +1234567890')}                  │
│  ${gray('  telnyx call dial --from +1234 --to +5678')}              │
│  ${gray('  telnyx billing balance')}                               │
├──────────────────────────────────────────────────────────┤
│  Run ${primary('telnyx --help')} for all commands                      │
└──────────────────────────────────────────────────────────┘
`;

// Slash commands reference (for enhanced mode)
const SLASH_COMMANDS = {
  '/ns': 'number search --country US',
  '/nl': 'number list',
  '/no': 'number order',
  '/ms': 'message send',
  '/ml': 'message list',
  '/cd': 'call dial',
  '/cl': 'call list',
  '/lookup': 'lookup',
  '/sim': 'sim list',
  '/fax': 'fax send',
  '/verify': 'verify send',
  '/debug': 'debugger list',
  '/bb': 'billing balance',
  '/whoami': 'auth whoami',
  '/help': '--help'
};

function showBanner() {
  console.log(primary(ASCII_BANNER));
}

function showMenu() {
  console.log(MENU_BOX);
}

function showError(message) {
  console.log(red(`✖ ${message}`));
}

function showSuccess(message) {
  console.log(primary(`✔ ${message}`));
}

function showInfo(message) {
  console.log(gray(message));
}

function showWarning(message) {
  console.log(yellow(`⚠ ${message}`));
}

// Create a styled box for displaying information
function createBox(title, content, options = {}) {
  const width = options.width || 59;
  const titleStr = title ? ` ${title} ` : '';
  const titlePadding = Math.max(0, (width - titleStr.length) / 2);
  
  let box = '';
  
  // Top border with title
  if (title) {
    box += '┌' + '─'.repeat(Math.floor(titlePadding)) + titleStr + '─'.repeat(Math.ceil(titlePadding) - 1) + '┐\n';
  } else {
    box += '┌' + '─'.repeat(width) + '┐\n';
  }
  
  // Content lines
  const lines = Array.isArray(content) ? content : [content];
  lines.forEach(line => {
    const paddedLine = line.padEnd(width - 2);
    box += '│ ' + paddedLine + ' │\n';
  });
  
  // Bottom border
  box += '└' + '─'.repeat(width) + '┘';
  
  return box;
}

// Format a table header with Telnyx colors
function formatTableHeader(headers) {
  return headers.map(h => primary(h));
}

module.exports = {
  COLORS,
  SLASH_COMMANDS,
  showBanner,
  showMenu,
  showError,
  showSuccess,
  showInfo,
  showWarning,
  createBox,
  formatTableHeader,
  // Color helpers
  primary,
  gray,
  yellow,
  red
};
