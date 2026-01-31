const { Command } = require('commander');
const inquirer = require('inquirer');
const ora = require('ora');
const Table = require('cli-table3');
const chalk = require('chalk');
const fs = require('fs');
const path = require('path');
const { sendFax, listFaxes, getFax, cancelFax } = require('../api/client');
const { showSuccess, showError, showInfo, COLORS } = require('../ui/layout');

const primary = chalk.hex(COLORS.primary);
const gray = chalk.gray;
const yellow = chalk.yellow;

// ==================== FAX SEND ====================

const send = new Command('send')
  .description('Send a fax')
  .alias('transmit')
  .option('-t, --to <number>', 'Destination fax number (E.164 format)')
  .option('-f, --from <number>', 'Sender fax number (your Telnyx number)')
  .option('-m, --media <url>', 'Media URL of the document to fax')
  .option('-F, --file <path>', 'Local file path to upload and fax')
  .option('--connection <id>', 'Connection ID for fax routing')
  .option('--webhook <url>', 'Webhook URL for delivery status')
  .option('--cover-page', 'Include cover page', false)
  .option('--cover-sheet', 'Include cover sheet', false)
  .option('--no-confirm', 'Skip confirmation prompt')
  .action(async (options) => {
    let { to, from, media, file, connection, webhook, coverPage, coverSheet, confirm } = options;
    
    // Interactive prompts for missing fields
    const prompts = [];
    
    if (!to) {
      prompts.push({
        type: 'input',
        name: 'to',
        message: 'To (destination fax number):',
        validate: (input) => {
          if (!input || !input.startsWith('+')) {
            return 'Fax number must be in E.164 format (e.g., +13125551234)';
          }
          return true;
        }
      });
    }
    
    if (!from) {
      prompts.push({
        type: 'input',
        name: 'from',
        message: 'From (your Telnyx fax-enabled number):',
        validate: (input) => {
          if (!input || !input.startsWith('+')) {
            return 'Phone number must be in E.164 format (e.g., +13125551234)';
          }
          return true;
        }
      });
    }
    
    if (!media && !file) {
      prompts.push({
        type: 'list',
        name: 'sourceType',
        message: 'How would you like to provide the document?',
        choices: [
          { name: 'Media URL (already hosted)', value: 'url' },
          { name: 'Local file (will be uploaded)', value: 'file' }
        ],
        default: 'url'
      });
    }
    
    if (prompts.length > 0) {
      const answers = await inquirer.prompt(prompts);
      to = to || answers.to;
      from = from || answers.from;
      
      if (answers.sourceType === 'url' && !media) {
        const { url } = await inquirer.prompt([{
          type: 'input',
          name: 'url',
          message: 'Media URL (must be publicly accessible PDF/TIFF):',
          validate: (input) => input.length > 0 || 'Media URL is required'
        }]);
        media = url;
      } else if (answers.sourceType === 'file' && !file) {
        const { filePath } = await inquirer.prompt([{
          type: 'input',
          name: 'filePath',
          message: 'File path:',
          validate: (input) => {
            if (!input || input.length === 0) return 'File path is required';
            if (!fs.existsSync(input)) return 'File not found';
            return true;
          }
        }]);
        file = filePath;
      }
    }
    
    // Validate required fields
    if (!to || !from) {
      showError('❌ Both --to and --from are required');
      process.exit(1);
    }
    
    if (!media && !file) {
      showError('❌ Either --media or --file is required');
      process.exit(1);
    }
    
    // Handle file upload (simplified - in production would upload to storage)
    if (file) {
      showInfo('📁 Local file upload not yet implemented.');
      showInfo('   Please upload your file to a publicly accessible URL and use --media');
      process.exit(1);
    }
    
    // Show preview
    console.log('');
    console.log('┌─────────────────────────────────────────────────────────┐');
    console.log('│  📠 ' + primary('FAX PREVIEW') + '                                         │');
    console.log('├─────────────────────────────────────────────────────────┤');
    console.log(`│  To:           ${to.substring(0, 45).padEnd(45)}│`);
    console.log(`│  From:         ${from.substring(0, 45).padEnd(45)}│`);
    console.log(`│  Media URL:    ${(media || 'N/A').substring(0, 45).padEnd(45)}│`);
    if (connection) {
      console.log(`│  Connection:   ${connection.substring(0, 45).padEnd(45)}│`);
    }
    if (coverPage || coverSheet) {
      console.log(`│  Cover Page:   ${'Yes ✓'.padEnd(45)}│`);
    }
    console.log('├─────────────────────────────────────────────────────────┤');
    console.log('│  💰 Estimated Cost: ~$0.02-0.05 per page                 │');
    console.log('└─────────────────────────────────────────────────────────┘');
    console.log('');
    
    // Confirm before sending
    if (confirm !== false) {
      const { confirmed } = await inquirer.prompt([
        {
          type: 'confirm',
          name: 'confirmed',
          message: 'Send this fax?',
          default: true
        }
      ]);
      
      if (!confirmed) {
        showInfo('Fax cancelled.');
        return;
      }
    }
    
    const spinner = ora({
      text: 'Sending fax...',
      spinner: 'dots'
    }).start();
    
    try {
      const faxOptions = {};
      if (connection) faxOptions.connection_id = connection;
      if (webhook) faxOptions.webhook_url = webhook;
      if (coverPage || coverSheet) faxOptions.cover_page = true;
      
      const data = await sendFax(to, from, media, faxOptions);
      
      spinner.stop();
      
      showSuccess('✅ Fax queued successfully!');
      console.log('');
      
      const fax = data.data;
      console.log(`${primary('Fax ID:')}       ${fax.id}`);
      console.log(`${primary('To:')}           ${fax.to || to}`);
      console.log(`${primary('From:')}         ${fax.from || from}`);
      console.log(`${primary('Status:')}       ${fax.status || 'queued'}`);
      console.log(`${primary('Direction:')}    ${fax.direction || 'outbound'}`);
      
      if (fax.pages) {
        console.log(`${primary('Pages:')}        ${fax.pages}`);
      }
      
      console.log('');
      showInfo('⏱️  Fax is being processed. Check status with:');
      showInfo(`   telnyx fax status ${fax.id}`);
      
    } catch (error) {
      spinner.stop();
      handleApiError(error);
    }
  });

// ==================== FAX LIST ====================

const list = new Command('list')
  .description('List fax transmissions')
  .alias('ls')
  .option('-s, --status <status>', 'Filter by status: queued, processing, delivered, failed')
  .option('-n, --limit <number>', 'Number of results', '20')
  .option('--inbound', 'Show only inbound faxes')
  .option('--outbound', 'Show only outbound faxes')
  .action(async (options) => {
    const { status, limit, inbound, outbound } = options;
    
    const spinner = ora({
      text: 'Fetching faxes...',
      spinner: 'dots'
    }).start();
    
    try {
      const data = await listFaxes({
        status,
        limit: parseInt(limit)
      });
      
      spinner.stop();
      
      if (!data.data || data.data.length === 0) {
        showInfo('📭 No faxes found.');
        return;
      }
      
      // Filter by direction if specified
      let faxes = data.data;
      if (inbound) {
        faxes = faxes.filter(item => {
          const fax = item.data || item;
          return fax.direction === 'inbound';
        });
      } else if (outbound) {
        faxes = faxes.filter(item => {
          const fax = item.data || item;
          return fax.direction === 'outbound';
        });
      }
      
      if (faxes.length === 0) {
        showInfo(`📭 No ${inbound ? 'inbound' : outbound ? 'outbound' : ''} faxes found.`);
        return;
      }
      
      showSuccess(`Found ${faxes.length} fax(es)`);
      console.log('');
      
      const table = new Table({
        head: ['Fax ID', 'To', 'From', 'Status', 'Pages', 'Time'],
        colWidths: [24, 18, 18, 12, 8, 20],
        style: {
          head: [COLORS.primary.replace('#', '')],
          border: ['gray']
        }
      });
      
      faxes.forEach((item) => {
        const fax = item.data || item;
        const time = fax.created_at ? new Date(fax.created_at).toLocaleString() : 'N/A';
        
        table.push([
          (fax.id || 'N/A').substring(0, 22),
          fax.to || 'N/A',
          fax.from || 'N/A',
          getStatusIcon(fax.status) + ' ' + (fax.status || 'unknown'),
          fax.pages || '-',
          time
        ]);
      });
      
      console.log(table.toString());
      console.log('');
      showInfo('💡 Use "telnyx fax status <fax_id>" for detailed information');
      
    } catch (error) {
      spinner.stop();
      handleApiError(error);
    }
  });

// ==================== FAX STATUS ====================

const status = new Command('status')
  .description('Check the status of a fax transmission')
  .alias('info')
  .argument('[faxId]', 'Fax ID to check')
  .option('-j, --json', 'Output raw JSON')
  .action(async (faxId, options) => {
    let targetFaxId = faxId;
    
    // If no fax ID provided, prompt for one
    if (!targetFaxId) {
      const { id } = await inquirer.prompt([
        {
          type: 'input',
          name: 'id',
          message: 'Fax ID:',
          validate: (input) => input.length > 0 || 'Fax ID is required'
        }
      ]);
      targetFaxId = id;
    }
    
    const spinner = ora({
      text: 'Fetching fax status...',
      spinner: 'dots'
    }).start();
    
    try {
      const data = await getFax(targetFaxId);
      
      spinner.stop();
      
      if (!data.data) {
        showError('Fax not found');
        return;
      }
      
      const fax = data.data;
      
      if (options.json) {
        console.log(JSON.stringify(data, null, 2));
        return;
      }
      
      showSuccess(`Fax Status: ${targetFaxId}`);
      console.log('');
      
      console.log('┌─────────────────────────────────────────────────────────┐');
      console.log('│  📠 ' + primary('FAX DETAILS') + '                                         │');
      console.log('├─────────────────────────────────────────────────────────┤');
      console.log(`│  Fax ID:        ${(fax.id || 'N/A').padEnd(45)}│`);
      console.log(`│  Status:        ${(fax.status || 'N/A').padEnd(45)}│`);
      console.log(`│  Direction:     ${(fax.direction || 'N/A').padEnd(45)}│`);
      console.log(`│  To:            ${(fax.to || 'N/A').padEnd(45)}│`);
      console.log(`│  From:          ${(fax.from || 'N/A').padEnd(45)}│`);
      if (fax.pages) {
        console.log(`│  Pages:         ${(fax.pages.toString()).padEnd(45)}│`);
      }
      if (fax.media_url) {
        console.log(`│  Media URL:     ${(fax.media_url.substring(0, 45)).padEnd(45)}│`);
      }
      console.log('├─────────────────────────────────────────────────────────┤');
      console.log(`│  Created:       ${(fax.created_at ? new Date(fax.created_at).toLocaleString() : 'N/A').padEnd(45)}│`);
      if (fax.updated_at) {
        console.log(`│  Updated:       ${(fax.updated_at ? new Date(fax.updated_at).toLocaleString() : 'N/A').padEnd(45)}│`);
      }
      if (fax.completed_at) {
        console.log(`│  Completed:     ${(fax.completed_at ? new Date(fax.completed_at).toLocaleString() : 'N/A').padEnd(45)}│`);
      }
      console.log('└─────────────────────────────────────────────────────────┘');
      
      // Error details if failed
      if (fax.status === 'failed' && fax.errors) {
        console.log('');
        console.log('┌─────────────────────────────────────────────────────────┐');
        console.log('│  ❌ ' + primary('ERROR DETAILS') + '                                       │');
        console.log('├─────────────────────────────────────────────────────────┤');
        fax.errors.forEach(err => {
          console.log(`│  Code:    ${(err.code || 'N/A').padEnd(49)}│`);
          console.log(`│  Detail:  ${(err.detail || err.title || 'Unknown error').padEnd(49)}│`);
        });
        console.log('└─────────────────────────────────────────────────────────┘');
      }
      
      console.log('');
      
    } catch (error) {
      spinner.stop();
      handleApiError(error);
    }
  });

// ==================== FAX CANCEL ====================

const cancel = new Command('cancel')
  .description('Cancel a pending fax transmission')
  .alias('stop')
  .argument('[faxId]', 'Fax ID to cancel')
  .option('-y, --yes', 'Skip confirmation prompt')
  .action(async (faxId, options) => {
    let targetFaxId = faxId;
    
    // If no fax ID provided, prompt for one
    if (!targetFaxId) {
      const { id } = await inquirer.prompt([
        {
          type: 'input',
          name: 'id',
          message: 'Fax ID to cancel:',
          validate: (input) => input.length > 0 || 'Fax ID is required'
        }
      ]);
      targetFaxId = id;
    }
    
    // Confirm before cancelling
    if (!options.yes) {
      const { confirm } = await inquirer.prompt([
        {
          type: 'confirm',
          name: 'confirm',
          message: `Cancel fax ${targetFaxId}?`,
          default: true
        }
      ]);
      
      if (!confirm) {
        showInfo('Operation cancelled.');
        return;
      }
    }
    
    const spinner = ora({
      text: 'Cancelling fax...',
      spinner: 'dots'
    }).start();
    
    try {
      await cancelFax(targetFaxId);
      
      spinner.stop();
      showSuccess(`✅ Fax ${targetFaxId} cancelled successfully`);
      
    } catch (error) {
      spinner.stop();
      handleApiError(error);
    }
  });

// ==================== HELPERS ====================

function getStatusIcon(status) {
  const icons = {
    'queued': '⏳',
    'processing': '🔄',
    'delivered': '✅',
    'received': '✅',
    'failed': '❌',
    'cancelled': '❌'
  };
  return icons[status?.toLowerCase()] || '•';
}

function handleApiError(error) {
  if (error.response?.status === 401) {
    showError('🔐 Authentication failed. Run: telnyx auth login');
  } else if (error.response?.status === 403) {
    showError('🚫 Permission denied. Fax may not be enabled on your account.');
    showInfo('   Contact Telnyx support to enable fax services.');
  } else if (error.response?.status === 404) {
    showError('❌ Fax not found. The fax may have already completed or the ID is incorrect.');
  } else if (error.response?.status === 422) {
    const detail = error.response.data?.errors?.[0]?.detail || 
                   error.response.data?.errors?.[0]?.title || 
                   'Invalid request';
    showError(`❌ ${detail}`);
    
    if (detail.includes('media') || detail.includes('url')) {
      showInfo('   Tip: Media URLs must be publicly accessible and in PDF/TIFF format');
    }
  } else if (error.response?.status === 429) {
    showError('⏱️  Rate limit exceeded. Please wait a moment and try again.');
  } else {
    showError(`❌ ${error.message}`);
  }
  process.exit(1);
}

module.exports = {
  send,
  list,
  status,
  cancel
};
