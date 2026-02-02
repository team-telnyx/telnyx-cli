const { Command } = require('commander');
const inquirer = require('inquirer');
const ora = require('ora');
const Table = require('cli-table3');
const chalk = require('chalk');
const { sendMessage, listMessages } = require('../api/client');
const { showSuccess, showError, showInfo, COLORS } = require('../ui/layout');

const primary = chalk.hex(COLORS.primary);
const gray = chalk.gray;
const yellow = chalk.yellow;

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

// ==================== MESSAGE SEND ====================

const send = new Command('send')
  .description('Send an SMS or MMS message')
  .alias('s')
  .option('-t, --to <number>', 'Recipient phone number (E.164 format, e.g., +13125551234)')
  .option('-f, --from <number>', 'Sender phone number (your Telnyx number)')
  .option('-m, --text <message>', 'Message text content')
  .option('--media <url>', 'Media URL for MMS (can be used multiple times)', collect, [])
  .option('--subject <text>', 'Subject for MMS messages')
  .option('--webhook <url>', 'Webhook URL for delivery status')
  .option('--no-confirm', 'Skip confirmation prompt')
  .option('--dry-run', 'Show preview without sending')
  .option('-j, --json', 'Output raw JSON')
  .option('-o, --output <format>', 'Output format: json, table, csv', 'table')
  .action(async (options) => {
    let { to, from, text, media, subject, webhook, confirm, dryRun } = options;
    const outputFormat = getOutputFormat(options);
    
    // Interactive prompts for missing fields
    const prompts = [];
    
    if (!to) {
      prompts.push({
        type: 'input',
        name: 'to',
        message: 'To (recipient phone number):',
        validate: (input) => {
          if (!input || !input.startsWith('+')) {
            return 'Phone number must be in E.164 format (e.g., +13125551234)';
          }
          if (input.length < 10) {
            return 'Phone number appears too short';
          }
          return true;
        }
      });
    }
    
    if (!from) {
      prompts.push({
        type: 'input',
        name: 'from',
        message: 'From (your Telnyx phone number):',
        validate: (input) => {
          if (!input || !input.startsWith('+')) {
            return 'Phone number must be in E.164 format (e.g., +13125551234)';
          }
          return true;
        }
      });
    }
    
    if (!text && media.length === 0) {
      prompts.push({
        type: 'editor',
        name: 'text',
        message: 'Message text (opens in default editor):',
        validate: (input) => {
          if (!input || input.trim().length === 0) {
            return 'Message text is required (or provide media URLs)';
          }
          return true;
        }
      });
    }
    
    if (prompts.length > 0) {
      const answers = await inquirer.prompt(prompts);
      to = to || answers.to;
      from = from || answers.from;
      text = text || answers.text;
    }
    
    // Validate required fields
    if (!to || !from) {
      showError('❌ Both --to and --from are required');
      process.exit(1);
    }
    
    if (!text && media.length === 0) {
      showError('❌ Either --text or --media is required');
      process.exit(1);
    }
    
    // Calculate estimated cost
    const segmentCount = text ? Math.ceil(text.length / 160) : 0;
    const estimatedCost = (segmentCount * 0.004) + (media.length * 0.02);
    
    // Show preview
    console.log('');
    console.log('┌─────────────────────────────────────────────────────────┐');
    console.log('│  📱 ' + primary('MESSAGE PREVIEW') + '                                     │');
    console.log('├─────────────────────────────────────────────────────────┤');
    console.log(`│  To:     ${to.substring(0, 45).padEnd(45)}│`);
    console.log(`│  From:   ${from.substring(0, 45).padEnd(45)}│`);
    
    if (text) {
      const previewText = text.length > 40 ? text.substring(0, 40) + '...' : text;
      const lines = previewText.match(/.{1,45}/g) || [previewText];
      lines.forEach((line, idx) => {
        if (idx === 0) {
          console.log(`│  Text:   ${line.padEnd(45)}│`);
        } else {
          console.log(`│          ${line.padEnd(45)}│`);
        }
      });
      console.log(`│  Length: ${text.length} chars (${segmentCount} segment${segmentCount > 1 ? 's' : ''})${''.padEnd(28)}│`);
    }
    
    if (media.length > 0) {
      console.log(`│  Media:  ${media.length} attachment(s)${''.padEnd(36)}│`);
      media.forEach((url, i) => {
        const shortUrl = url.length > 40 ? url.substring(0, 40) + '...' : url;
        console.log(`│         ${i + 1}. ${shortUrl.padEnd(41)}│`);
      });
    }
    
    if (subject) {
      console.log(`│  Subject: ${subject.substring(0, 44).padEnd(44)}│`);
    }
    
    console.log('├─────────────────────────────────────────────────────────┤');
    console.log(`│  💰 Estimated Cost: ~$${estimatedCost.toFixed(3)}${''.padEnd(37)}│`);
    console.log('└─────────────────────────────────────────────────────────┘');
    console.log('');
    
    if (dryRun) {
      showInfo('🔍 Dry run - message not sent.');
      showInfo('Remove --dry-run to actually send the message');
      return;
    }
    
    // Confirm before sending
    if (confirm !== false) {
      const { confirmed } = await inquirer.prompt([
        {
          type: 'confirm',
          name: 'confirmed',
          message: 'Send this message?',
          default: true
        }
      ]);
      
      if (!confirmed) {
        showInfo('Message cancelled.');
        return;
      }
    }
    
    const spinner = ora({
      text: 'Sending message...',
      spinner: 'dots'
    }).start();
    
    try {
      const messageOptions = {};
      if (media.length > 0) {
        messageOptions.media_urls = media;
      }
      if (subject) {
        messageOptions.subject = subject;
      }
      if (webhook) {
        messageOptions.webhook_url = webhook;
        messageOptions.use_webhooks = true;
      }
      
      const data = await sendMessage(to, from, text || '', messageOptions);
      
      spinner.stop();
      
      // Handle JSON/CSV output
      if (outputFormat !== 'table') {
        formatOutput(data.data || data, outputFormat);
        return;
      }
      
      showSuccess('✅ Message sent successfully!');
      console.log('');
      
      const msg = data.data;
      console.log(`${primary('Message ID:')}    ${msg.id}`);
      console.log(`${primary('To:')}           ${Array.isArray(msg.to) ? msg.to.join(', ') : msg.to || to}`);
      console.log(`${primary('From:')}         ${msg.from || from}`);
      console.log(`${primary('Direction:')}    ${msg.direction || 'outbound'}`);
      console.log(`${primary('Status:')}       ${msg.status || 'sent'}`);
      
      if (msg.text) {
        const preview = msg.text.substring(0, 50) + (msg.text.length > 50 ? '...' : '');
        console.log(`${primary('Text:')}         ${preview}`);
      }
      
      if (msg.cost) {
        console.log(`${primary('Cost:')}         ${msg.cost.amount} ${msg.cost.currency}`);
      }
      
      if (msg.status === 'queued' || msg.status === 'sending') {
        console.log('');
        showInfo('⏱️  Message is being processed. Check status with:');
        showInfo(`   telnyx message status ${msg.id}`);
      }
      
      console.log('');
      
    } catch (error) {
      spinner.stop();
      showError(error.message);
      process.exit(1);
    }
  });

// ==================== MESSAGE LIST ====================

const list = new Command('list')
  .description('List recent messages')
  .alias('ls')
  .option('-f, --from <number>', 'Filter by sender')
  .option('-t, --to <number>', 'Filter by recipient')
  .option('-s, --status <status>', 'Filter by status: queued, sending, sent, delivered, failed')
  .option('-n, --limit <number>', 'Number of results', '20')
  .option('-j, --json', 'Output raw JSON')
  .option('-o, --output <format>', 'Output format: json, table, csv', 'table')
  .action(async (options) => {
    const { from, to, status, limit } = options;
    const outputFormat = getOutputFormat(options);
    
    const spinner = ora({
      text: 'Fetching messages...',
      spinner: 'dots'
    }).start();
    
    try {
      const data = await listMessages({ from, to, status, limit: parseInt(limit) });
      
      spinner.stop();
      
      if (!data.data || data.data.length === 0) {
        showInfo('📭 No messages found.');
        return;
      }
      
      // Handle JSON/CSV output
      if (outputFormat !== 'table') {
        formatOutput(data.data, outputFormat);
        return;
      }
      
      showSuccess(`Found ${data.data.length} message(s)`);
      console.log('');
      
      const table = new Table({
        head: ['ID', 'To', 'From', 'Status', 'Direction', 'Time'],
        colWidths: [24, 18, 18, 12, 12, 20],
        style: {
          head: [COLORS.primary.replace('#', '')],
          border: ['gray']
        }
      });
      
      data.data.forEach((item) => {
        const msg = item.data || item;
        const time = msg.created_at ? new Date(msg.created_at).toLocaleString() : 'N/A';
        
        table.push([
          msg.id.substring(0, 22),
          msg.to?.[0] || msg.to || 'N/A',
          msg.from || 'N/A',
          getStatusIcon(msg.status) + ' ' + msg.status,
          msg.direction || 'N/A',
          time
        ]);
      });
      
      console.log(table.toString());
      
    } catch (error) {
      spinner.stop();
      showError(error.message);
      process.exit(1);
    }
  });

// ==================== HELPERS ====================

function collect(value, previous) {
  return previous.concat([value]);
}

function getStatusIcon(status) {
  const icons = {
    'delivered': '✅',
    'sent': '✓',
    'sending': '⏳',
    'queued': '⏸️',
    'failed': '❌',
    'received': '📥'
  };
  return icons[status] || '•';
}

module.exports = {
  send,
  list
};
