const { Command } = require('commander');
const inquirer = require('inquirer');
const ora = require('ora');
const Table = require('cli-table3');
const chalk = require('chalk');
const { 
  listCalls, 
  createCall, 
  getCall, 
  hangupCall 
} = require('../api/client');
const { showSuccess, showError, showInfo, COLORS } = require('../ui/layout');

const primary = chalk.hex(COLORS.primary);
const gray = chalk.gray;
const yellow = chalk.yellow;
const green = chalk.green;
const red = chalk.red;

// ==================== CALL DIAL ====================

const dial = new Command('dial')
  .description('Make an outbound voice call')
  .alias('call')
  .option('-f, --from <number>', 'Caller ID (your Telnyx number)')
  .option('-t, --to <number>', 'Destination phone number (E.164 format)')
  .option('-c, --connection <id>', 'Connection ID for call routing')
  .option('--texml <url>', 'TeXML URL for call handling')
  .option('--webhook <url>', 'Webhook URL for call events')
  .option('--timeout <seconds>', 'Call timeout in seconds', '30')
  .option('--no-confirm', 'Skip confirmation prompt')
  .action(async (options) => {
    let { from, to, connection, texml, webhook, timeout, confirm } = options;
    
    // Interactive prompts for missing fields
    const prompts = [];
    
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
    
    if (!to) {
      prompts.push({
        type: 'input',
        name: 'to',
        message: 'To (destination phone number):',
        validate: (input) => {
          if (!input || !input.startsWith('+')) {
            return 'Phone number must be in E.164 format (e.g., +13125551234)';
          }
          return true;
        }
      });
    }
    
    if (!connection && !texml) {
      prompts.push({
        type: 'list',
        name: 'routingType',
        message: 'How should the call be handled?',
        choices: [
          { name: 'Connection ID (for SIP routing)', value: 'connection' },
          { name: 'TeXML URL (for programmable voice)', value: 'texml' }
        ],
        default: 'connection'
      });
    }
    
    if (prompts.length > 0) {
      const answers = await inquirer.prompt(prompts);
      from = from || answers.from;
      to = to || answers.to;
      
      if (answers.routingType === 'connection' && !connection) {
        const { connId } = await inquirer.prompt([{
          type: 'input',
          name: 'connId',
          message: 'Connection ID:',
          validate: (input) => input.length > 0 || 'Connection ID is required'
        }]);
        connection = connId;
      } else if (answers.routingType === 'texml' && !texml) {
        const { texmlUrl } = await inquirer.prompt([{
          type: 'input',
          name: 'texmlUrl',
          message: 'TeXML URL:',
          validate: (input) => input.length > 0 || 'TeXML URL is required'
        }]);
        texml = texmlUrl;
      }
    }
    
    // Validate required fields
    if (!from || !to) {
      showError('❌ Both --from and --to are required');
      process.exit(1);
    }
    
    if (!connection && !texml) {
      showError('❌ Either --connection or --texml is required');
      process.exit(1);
    }
    
    // Show preview
    console.log('');
    console.log('┌─────────────────────────────────────────────────────────┐');
    console.log('│  📞 ' + primary('CALL PREVIEW') + '                                          │');
    console.log('├─────────────────────────────────────────────────────────┤');
    console.log(`│  From:       ${from.substring(0, 45).padEnd(45)}│`);
    console.log(`│  To:         ${to.substring(0, 45).padEnd(45)}│`);
    
    if (connection) {
      console.log(`│  Connection: ${connection.substring(0, 45).padEnd(45)}│`);
    }
    if (texml) {
      console.log(`│  TeXML URL:  ${texml.substring(0, 45).padEnd(45)}│`);
    }
    if (webhook) {
      console.log(`│  Webhook:    ${webhook.substring(0, 45).padEnd(45)}│`);
    }
    
    console.log(`│  Timeout:    ${(timeout + ' seconds').padEnd(45)}│`);
    console.log('└─────────────────────────────────────────────────────────┘');
    console.log('');
    
    // Confirm before dialing
    if (confirm !== false) {
      const { confirmed } = await inquirer.prompt([
        {
          type: 'confirm',
          name: 'confirmed',
          message: 'Place this call?',
          default: true
        }
      ]);
      
      if (!confirmed) {
        showInfo('Call cancelled.');
        return;
      }
    }
    
    const spinner = ora({
      text: 'Initiating call...',
      spinner: 'dots'
    }).start();
    
    try {
      const callOptions = {
        to: to,
        from: from,
        timeout: parseInt(timeout)
      };
      
      if (connection) {
        callOptions.connection_id = connection;
      }
      if (texml) {
        callOptions.texml = texml;
      }
      if (webhook) {
        callOptions.webhook_url = webhook;
      }
      
      const data = await createCall(callOptions);
      
      spinner.stop();
      
      showSuccess('✅ Call initiated successfully!');
      console.log('');
      
      const call = data.data;
      console.log(`${primary('Call ID:')}      ${call.call_control_id || call.id}`);
      console.log(`${primary('From:')}         ${call.from || from}`);
      console.log(`${primary('To:')}           ${call.to || to}`);
      console.log(`${primary('Status:')}       ${call.call_leg_id ? 'initiated' : 'pending'}`);
      console.log(`${primary('Direction:')}    ${call.direction || 'outbound'}`);
      
      if (call.call_control_id) {
        console.log('');
        showInfo('💡 To hang up this call:');
        showInfo(`   telnyx call hangup ${call.call_control_id}`);
      }
      
      console.log('');
      
    } catch (error) {
      spinner.stop();
      handleApiError(error);
    }
  });

// ==================== CALL HANGUP ====================

const hangup = new Command('hangup')
  .description('Hang up an active call')
  .alias('end')
  .argument('[callId]', 'Call Control ID to hang up')
  .option('-y, --yes', 'Skip confirmation prompt')
  .action(async (callId, options) => {
    let targetCallId = callId;
    
    // If no call ID provided, prompt for one
    if (!targetCallId) {
      const { id } = await inquirer.prompt([
        {
          type: 'input',
          name: 'id',
          message: 'Call Control ID:',
          validate: (input) => input.length > 0 || 'Call ID is required'
        }
      ]);
      targetCallId = id;
    }
    
    // Confirm before hanging up
    if (!options.yes) {
      const { confirm } = await inquirer.prompt([
        {
          type: 'confirm',
          name: 'confirm',
          message: `Hang up call ${targetCallId}?`,
          default: true
        }
      ]);
      
      if (!confirm) {
        showInfo('Operation cancelled.');
        return;
      }
    }
    
    const spinner = ora({
      text: 'Hanging up call...',
      spinner: 'dots'
    }).start();
    
    try {
      await hangupCall(targetCallId);
      
      spinner.stop();
      showSuccess(`✅ Call ${targetCallId} hung up successfully`);
      
    } catch (error) {
      spinner.stop();
      handleApiError(error);
    }
  });

// ==================== CALL LIST ====================

const list = new Command('list')
  .description('List active and recent calls')
  .alias('ls')
  .option('-s, --status <status>', 'Filter by status: initiated, ringing, answered, hangup')
  .option('-n, --limit <number>', 'Number of results', '20')
  .option('--active', 'Show only active calls (initiated, ringing, answered)')
  .action(async (options) => {
    const { status, limit, active } = options;
    
    let filterStatus = status;
    if (active) {
      filterStatus = undefined; // We'll filter client-side for active calls
    }
    
    const spinner = ora({
      text: 'Fetching calls...',
      spinner: 'dots'
    }).start();
    
    try {
      const data = await listCalls({ 
        status: filterStatus, 
        limit: parseInt(limit) 
      });
      
      spinner.stop();
      
      if (!data.data || data.data.length === 0) {
        showInfo('📭 No calls found.');
        return;
      }
      
      // Filter for active calls if requested
      let calls = data.data;
      if (active) {
        const activeStatuses = ['initiated', 'ringing', 'answered', 'bridged'];
        calls = calls.filter(item => {
          const call = item.data || item;
          return activeStatuses.includes(call.status) || activeStatuses.includes(call.call_state);
        });
      }
      
      if (calls.length === 0) {
        showInfo('📭 No active calls found.');
        return;
      }
      
      showSuccess(`Found ${calls.length} call(s)`);
      console.log('');
      
      const table = new Table({
        head: ['Call ID', 'From', 'To', 'Status', 'Duration', 'Time'],
        colWidths: [24, 18, 18, 12, 10, 20],
        style: {
          head: [COLORS.primary.replace('#', '')],
          border: ['gray']
        }
      });
      
      calls.forEach((item) => {
        const call = item.data || item;
        const callStatus = call.status || call.call_state || 'unknown';
        const time = call.created_at ? new Date(call.created_at).toLocaleString() : 'N/A';
        const duration = call.duration || call.call_duration || '-';
        
        table.push([
          (call.call_control_id || call.id || 'N/A').substring(0, 22),
          call.from || call.from_display_name || 'N/A',
          call.to || 'N/A',
          getStatusIcon(callStatus) + ' ' + callStatus,
          typeof duration === 'number' ? formatDuration(duration) : duration,
          time
        ]);
      });
      
      console.log(table.toString());
      console.log('');
      showInfo('💡 Use "telnyx call hangup <call_id>" to end a call');
      
    } catch (error) {
      spinner.stop();
      handleApiError(error);
    }
  });

// ==================== HELPERS ====================

function getStatusIcon(status) {
  const icons = {
    'initiated': '📞',
    'ringing': '🔔',
    'answered': '✅',
    'bridged': '🔗',
    'hangup': '❌',
    'completed': '✓',
    'failed': '❌'
  };
  return icons[status?.toLowerCase()] || '•';
}

function formatDuration(seconds) {
  if (!seconds || seconds <= 0) return '-';
  const mins = Math.floor(seconds / 60);
  const secs = seconds % 60;
  return `${mins}:${secs.toString().padStart(2, '0')}`;
}

function handleApiError(error) {
  if (error.response?.status === 401) {
    showError('🔐 Authentication failed. Run: telnyx auth login');
  } else if (error.response?.status === 403) {
    showError('🚫 Permission denied. Your account may not have voice calling enabled.');
    showInfo('   Contact Telnyx support to enable voice on your account.');
  } else if (error.response?.status === 404) {
    showError('❌ Call not found. The call may have already ended or the ID is incorrect.');
  } else if (error.response?.status === 422) {
    const detail = error.response.data?.errors?.[0]?.detail || 
                   error.response.data?.errors?.[0]?.title || 
                   'Invalid request';
    showError(`❌ ${detail}`);
  } else if (error.response?.status === 429) {
    showError('⏱️  Rate limit exceeded. Please wait a moment and try again.');
  } else {
    showError(`❌ ${error.message}`);
  }
  process.exit(1);
}

module.exports = {
  dial,
  hangup,
  list
};
