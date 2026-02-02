const { Command } = require('commander');
const inquirer = require('inquirer');
const ora = require('ora');
const Table = require('cli-table3');
const chalk = require('chalk');
const { 
  listEvents, 
  getEvent, 
  listWebhooks, 
  getWebhook, 
  resendWebhook 
} = require('../api/client');
const { showSuccess, showError, showInfo, COLORS } = require('../ui/layout');

const primary = chalk.hex(COLORS.primary);
const gray = chalk.gray;
const yellow = chalk.yellow;
const red = chalk.red;

// Helper to format output based on --json and --output flags
function formatOutput(data, format) {
  if (format === 'json') {
    console.log(JSON.stringify(data, null, 2));
    return;
  }
  
  if (format === 'csv') {
    if (Array.isArray(data) && data.length > 0) {
      const firstItem = data[0].data || data[0];
      const headers = Object.keys(firstItem);
      console.log(headers.join(','));
      
      data.forEach(item => {
        const row = item.data || item;
        const values = headers.map(h => {
          const val = row[h];
          if (val === null || val === undefined) return '';
          if (typeof val === 'object') return JSON.stringify(val).replace(/,/g, ';');
          return String(val).replace(/,/g, ';');
        });
        console.log(values.join(','));
      });
    } else {
      console.log('data');
      console.log(JSON.stringify(data));
    }
    return;
  }
  
  return false;
}

function getOutputFormat(options) {
  if (options.json) return 'json';
  if (options.output) return options.output;
  return 'table';
}

// ==================== DEBUGGER LIST ====================

const list = new Command('list')
  .description('List recent webhook events and deliveries for debugging')
  .alias('ls')
  .option('-t, --type <type>', 'Filter by event type (e.g., message.received, call.initiated)')
  .option('-s, --status <status>', 'Filter by status: pending, delivered, failed')
  .option('-n, --limit <number>', 'Number of results', '20')
  .option('--webhooks', 'Show webhook deliveries instead of events')
  .option('--failed', 'Show only failed deliveries')
  .option('--since <hours>', 'Show events from last N hours', '24')
  .option('-j, --json', 'Output raw JSON')
  .option('-o, --output <format>', 'Output format: json, table, csv', 'table')
  .action(async (options) => {
    const { type, status, limit, webhooks, failed, since } = options;
    const outputFormat = getOutputFormat(options);
    
    // Calculate since timestamp
    const sinceHours = parseInt(since);
    const sinceDate = new Date(Date.now() - (sinceHours * 60 * 60 * 1000));
    const sinceIso = sinceDate.toISOString();
    
    if (webhooks || failed) {
      await listWebhookDeliveries({ 
        status: failed ? 'failed' : status, 
        limit: parseInt(limit),
        since: sinceIso,
        outputFormat
      });
    } else {
      await listEventsList({ 
        type, 
        status, 
        limit: parseInt(limit),
        since: sinceIso,
        outputFormat
      });
    }
  });

async function listEventsList(filters) {
  const spinner = ora({
    text: 'Fetching events...',
    spinner: 'dots'
  }).start();
  
  try {
    const data = await listEvents({
      type: filters.type,
      status: filters.status,
      limit: filters.limit,
      since: filters.since
    });
    
    spinner.stop();
    
    if (!data.data || data.data.length === 0) {
      showInfo('📭 No events found.');
      showInfo(`   Try: --since 48 for events from last 48 hours`);
      return;
    }
    
    // Handle JSON/CSV output
    if (filters.outputFormat !== 'table') {
      formatOutput(data.data, filters.outputFormat);
      return;
    }
    
    showSuccess(`Found ${data.data.length} event(s)`);
    console.log('');
    
    const table = new Table({
      head: ['Event ID', 'Type', 'Status', 'Created'],
      colWidths: [26, 30, 12, 20],
      style: {
        head: [COLORS.primary.replace('#', '')],
        border: ['gray']
      }
    });
    
    data.data.forEach((item) => {
      const event = item.data || item;
      const time = event.created_at ? new Date(event.created_at).toLocaleString() : 'N/A';
      const eventType = event.event_type || 'unknown';
      
      table.push([
        (event.id || 'N/A').substring(0, 24),
        truncate(eventType, 28),
        getStatusIcon(event.status) + ' ' + (event.status || 'pending'),
        time
      ]);
    });
    
    console.log(table.toString());
    console.log('');
    showInfo('💡 Use "telnyx debugger inspect <event_id>" for details');
    
  } catch (error) {
    spinner.stop();
    showError(error.message);
    process.exit(1);
  }
}

async function listWebhookDeliveries(filters) {
  const spinner = ora({
    text: 'Fetching webhook deliveries...',
    spinner: 'dots'
  }).start();
  
  try {
    const data = await listWebhooks({
      status: filters.status,
      limit: filters.limit
    });
    
    spinner.stop();
    
    if (!data.data || data.data.length === 0) {
      showInfo('📭 No webhook deliveries found.');
      return;
    }
    
    // Handle JSON/CSV output
    if (filters.outputFormat !== 'table') {
      formatOutput(data.data, filters.outputFormat);
      return;
    }
    
    showSuccess(`Found ${data.data.length} webhook delivery(s)`);
    console.log('');
    
    const table = new Table({
      head: ['Webhook ID', 'Event Type', 'Status', 'Attempts', 'Time'],
      colWidths: [26, 24, 12, 10, 20],
      style: {
        head: [COLORS.primary.replace('#', '')],
        border: ['gray']
      }
    });
    
    data.data.forEach((item) => {
      const webhook = item.data || item;
      const time = webhook.created_at ? new Date(webhook.created_at).toLocaleString() : 'N/A';
      const attempts = webhook.attempts || webhook.retry_count || 0;
      
      table.push([
        (webhook.id || 'N/A').substring(0, 24),
        truncate(webhook.event_type || 'unknown', 22),
        getDeliveryStatusIcon(webhook.status) + ' ' + (webhook.status || 'unknown'),
        attempts.toString(),
        time
      ]);
    });
    
    console.log(table.toString());
    console.log('');
    
    // Show summary of failed webhooks
    const failedCount = data.data.filter(item => {
      const webhook = item.data || item;
      return webhook.status === 'failed';
    }).length;
    
    if (failedCount > 0) {
      showInfo(`⚠️  ${failedCount} failed delivery(s) detected`);
      showInfo('   Use --failed to see only failed deliveries');
      showInfo('   Use "telnyx debugger resend <webhook_id>" to retry');
    }
    
  } catch (error) {
    spinner.stop();
    showError(error.message);
    process.exit(1);
  }
}

// ==================== DEBUGGER INSPECT ====================

const inspect = new Command('inspect')
  .description('Inspect a specific event in detail')
  .alias('view')
  .argument('<eventId>', 'Event ID to inspect')
  .option('-j, --json', 'Output raw JSON')
  .option('-o, --output <format>', 'Output format: json, table, csv', 'table')
  .action(async (eventId, options) => {
    const outputFormat = getOutputFormat(options);
    
    const spinner = ora({
      text: 'Fetching event details...',
      spinner: 'dots'
    }).start();
    
    try {
      const data = await getEvent(eventId);
      
      spinner.stop();
      
      if (!data.data) {
        showError('Event not found');
        return;
      }
      
      const event = data.data;
      
      // Handle JSON/CSV output
      if (outputFormat !== 'table') {
        formatOutput(event, outputFormat);
        return;
      }
      
      showSuccess('Event Details');
      console.log('');
      
      console.log('┌─────────────────────────────────────────────────────────┐');
      console.log('│  🔍 ' + primary('EVENT DETAILS') + '                                       │');
      console.log('├─────────────────────────────────────────────────────────┤');
      console.log(`│  Event ID:    ${(event.id || 'N/A').padEnd(45)}│`);
      console.log(`│  Type:        ${(event.event_type || 'N/A').padEnd(45)}│`);
      console.log(`│  Status:      ${(event.status || 'N/A').padEnd(45)}│`);
      console.log(`│  Created:     ${(event.created_at ? new Date(event.created_at).toLocaleString() : 'N/A').padEnd(45)}│`);
      console.log('└─────────────────────────────────────────────────────────┘');
      
      // Payload section
      if (event.payload) {
        console.log('');
        console.log('┌─────────────────────────────────────────────────────────┐');
        console.log('│  📦 ' + primary('PAYLOAD') + '                                           │');
        console.log('├─────────────────────────────────────────────────────────┤');
        
        const payload = typeof event.payload === 'string' 
          ? JSON.parse(event.payload) 
          : event.payload;
        
        const payloadStr = JSON.stringify(payload, null, 2);
        const lines = payloadStr.split('\n').slice(0, 20); // Limit to 20 lines
        
        lines.forEach(line => {
          const truncated = line.substring(0, 53).padEnd(53);
          console.log(`│  ${truncated}│`);
        });
        
        if (payloadStr.split('\n').length > 20) {
          console.log(`│  ... (${payloadStr.split('\n').length - 20} more lines)${''.padEnd(33)}│`);
        }
        
        console.log('└─────────────────────────────────────────────────────────┘');
      }
      
      console.log('');
      
    } catch (error) {
      spinner.stop();
      showError(error.message);
      process.exit(1);
    }
  });

// ==================== DEBUGGER RESEND ====================

const resend = new Command('resend')
  .description('Resend a failed webhook delivery')
  .alias('retry')
  .argument('<webhookId>', 'Webhook delivery ID to resend')
  .option('-y, --yes', 'Skip confirmation prompt')
  .option('--dry-run', 'Show what would happen without resending')
  .option('-j, --json', 'Output raw JSON')
  .option('-o, --output <format>', 'Output format: json, table, csv', 'table')
  .action(async (webhookId, options) => {
    const outputFormat = getOutputFormat(options);
    
    // Dry run mode
    if (options.dryRun) {
      console.log('');
      console.log('┌─────────────────────────────────────────────────────────┐');
      console.log('│  🔍 ' + primary('DRY RUN - Webhook will not be resent') + '                 │');
      console.log('├─────────────────────────────────────────────────────────┤');
      console.log(`│  Webhook ID:  ${webhookId.substring(0, 45).padEnd(45)}│`);
      console.log('│  Would call: POST /v2/webhook_deliveries/{id}/resend    │');
      console.log('└─────────────────────────────────────────────────────────┘');
      console.log('');
      showInfo('Remove --dry-run to actually resend the webhook');
      return;
    }
    
    // Confirm before resending
    if (!options.yes) {
      const { confirm } = await inquirer.prompt([
        {
          type: 'confirm',
          name: 'confirm',
          message: `Resend webhook ${webhookId}?`,
          default: true
        }
      ]);
      
      if (!confirm) {
        showInfo('Operation cancelled.');
        return;
      }
    }
    
    const spinner = ora({
      text: 'Resending webhook...',
      spinner: 'dots'
    }).start();
    
    try {
      const data = await resendWebhook(webhookId);
      
      spinner.stop();
      
      // Handle JSON/CSV output
      if (outputFormat !== 'table') {
        formatOutput(data.data || data, outputFormat);
        return;
      }
      
      showSuccess(`✅ Webhook ${webhookId} resent successfully`);
      
    } catch (error) {
      spinner.stop();
      showError(error.message);
      process.exit(1);
    }
  });

// ==================== HELPERS ====================

function getStatusIcon(status) {
  const icons = {
    'delivered': '✅',
    'pending': '⏳',
    'processing': '⏳',
    'failed': '❌'
  };
  return icons[status?.toLowerCase()] || '•';
}

function getDeliveryStatusIcon(status) {
  const icons = {
    'delivered': '✅',
    'pending': '⏳',
    'failed': '❌',
    'retrying': '🔄'
  };
  return icons[status?.toLowerCase()] || '•';
}

function truncate(str, maxLength) {
  if (!str) return 'N/A';
  if (str.length <= maxLength) return str;
  return str.substring(0, maxLength - 3) + '...';
}

module.exports = {
  list,
  inspect,
  resend
};
