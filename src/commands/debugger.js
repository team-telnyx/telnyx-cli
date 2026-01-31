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
  .action(async (options) => {
    const { type, status, limit, webhooks, failed, since } = options;
    
    // Calculate since timestamp
    const sinceHours = parseInt(since);
    const sinceDate = new Date(Date.now() - (sinceHours * 60 * 60 * 1000));
    const sinceIso = sinceDate.toISOString();
    
    if (webhooks || failed) {
      await listWebhookDeliveries({ 
        status: failed ? 'failed' : status, 
        limit: parseInt(limit),
        since: sinceIso
      });
    } else {
      await listEventsList({ 
        type, 
        status, 
        limit: parseInt(limit),
        since: sinceIso
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
    handleApiError(error);
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
    handleApiError(error);
  }
}

// ==================== DEBUGGER INSPECT ====================

const inspect = new Command('inspect')
  .description('Inspect a specific event in detail')
  .alias('view')
  .argument('<eventId>', 'Event ID to inspect')
  .option('-j, --json', 'Output raw JSON')
  .action(async (eventId, options) => {
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
      
      if (options.json) {
        console.log(JSON.stringify(data, null, 2));
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
      handleApiError(error);
    }
  });

// ==================== DEBUGGER RESEND ====================

const resend = new Command('resend')
  .description('Resend a failed webhook delivery')
  .alias('retry')
  .argument('<webhookId>', 'Webhook delivery ID to resend')
  .option('-y, --yes', 'Skip confirmation prompt')
  .action(async (webhookId, options) => {
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
      await resendWebhook(webhookId);
      
      spinner.stop();
      showSuccess(`✅ Webhook ${webhookId} resent successfully`);
      
    } catch (error) {
      spinner.stop();
      handleApiError(error);
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

function handleApiError(error) {
  if (error.response?.status === 401) {
    showError('🔐 Authentication failed. Run: telnyx auth login');
  } else if (error.response?.status === 403) {
    showError('🚫 Permission denied. Your account may not have access to events.');
  } else if (error.response?.status === 404) {
    showError('❌ Event or webhook not found.');
  } else if (error.response?.status === 422) {
    const detail = error.response.data?.errors?.[0]?.detail || 'Invalid request';
    showError(`❌ ${detail}`);
  } else if (error.response?.status === 429) {
    showError('⏱️  Rate limit exceeded. Please wait a moment and try again.');
  } else {
    showError(`❌ ${error.message}`);
  }
  process.exit(1);
}

module.exports = {
  list,
  inspect,
  resend
};
