const { Command } = require('commander');
const inquirer = require('inquirer');
const ora = require('ora');
const Table = require('cli-table3');
const chalk = require('chalk');
const { 
  createVerification, 
  submitVerification, 
  getVerification, 
  listVerifications 
} = require('../api/client');
const { showSuccess, showError, showInfo, COLORS } = require('../ui/layout');

const primary = chalk.hex(COLORS.primary);
const gray = chalk.gray;
const yellow = chalk.yellow;
const green = chalk.green;
const red = chalk.red;

// ==================== VERIFY SEND ====================

const send = new Command('send')
  .description('Send a verification code (2FA)')
  .alias('create')
  .option('-n, --number <number>', 'Phone number to verify (E.164 format)')
  .option('-t, --type <type>', 'Verification type: sms, call, flashcall', 'sms')
  .option('-p, --profile <id>', 'Verify profile ID')
  .option('--timeout <seconds>', 'Code expiration timeout', '300')
  .option('--code-length <length>', 'Length of verification code', '6')
  .option('--locale <locale>', 'Locale for voice messages', 'en-US')
  .option('--no-confirm', 'Skip confirmation prompt')
  .action(async (options) => {
    let { number, type, profile, timeout, codeLength, locale, confirm } = options;
    
    // Interactive prompts for missing fields
    const prompts = [];
    
    if (!number) {
      prompts.push({
        type: 'input',
        name: 'number',
        message: 'Phone number to verify (E.164 format):',
        validate: (input) => {
          if (!input || !input.startsWith('+')) {
            return 'Phone number must be in E.164 format (e.g., +13125551234)';
          }
          return true;
        }
      });
    }
    
    if (prompts.length > 0) {
      const answers = await inquirer.prompt(prompts);
      number = number || answers.number;
    }
    
    // Validate required fields
    if (!number) {
      showError('❌ Phone number is required');
      process.exit(1);
    }
    
    // Show preview
    console.log('');
    console.log('┌─────────────────────────────────────────────────────────┐');
    console.log('│  🔐 ' + primary('VERIFICATION PREVIEW') + '                                │');
    console.log('├─────────────────────────────────────────────────────────┤');
    console.log(`│  Phone Number:    ${number.substring(0, 45).padEnd(45)}│`);
    console.log(`│  Type:            ${type.padEnd(45)}│`);
    console.log(`│  Timeout:         ${(timeout + ' seconds').padEnd(45)}│`);
    if (profile) {
      console.log(`│  Profile ID:      ${profile.substring(0, 45).padEnd(45)}│`);
    }
    console.log('└─────────────────────────────────────────────────────────┘');
    console.log('');
    
    // Confirm before sending
    if (confirm !== false) {
      const { confirmed } = await inquirer.prompt([
        {
          type: 'confirm',
          name: 'confirmed',
          message: 'Send verification code?',
          default: true
        }
      ]);
      
      if (!confirmed) {
        showInfo('Verification cancelled.');
        return;
      }
    }
    
    const spinner = ora({
      text: 'Sending verification code...',
      spinner: 'dots'
    }).start();
    
    try {
      const verifyOptions = {
        type,
        timeout: parseInt(timeout)
      };
      
      if (profile) verifyOptions.profileId = profile;
      if (codeLength) verifyOptions.code_length = parseInt(codeLength);
      if (locale) verifyOptions.locale = locale;
      
      const data = await createVerification(number, verifyOptions);
      
      spinner.stop();
      
      showSuccess('✅ Verification code sent!');
      console.log('');
      
      const verification = data.data;
      console.log(`${primary('Verification ID:')}  ${verification.id}`);
      console.log(`${primary('Phone Number:')}    ${verification.phone_number || number}`);
      console.log(`${primary('Type:')}            ${verification.type || type}`);
      console.log(`${primary('Status:')}          ${verification.status || 'pending'}`);
      
      if (verification.timeout_secs) {
        console.log(`${primary('Expires In:')}      ${verification.timeout_secs} seconds`);
      }
      
      console.log('');
      showInfo('💡 To verify the code:');
      showInfo(`   telnyx verify check ${number}`);
      
      // If in test mode, show the code
      if (verification.code) {
        console.log('');
        console.log('┌─────────────────────────────────────────────────────────┐');
        console.log('│  🧪 ' + primary('TEST MODE - Verification Code') + '                        │');
        console.log('├─────────────────────────────────────────────────────────┤');
        console.log(`│                                                         │`);
        console.log(`│  Code: ${(verification.code).padEnd(48)}│`);
        console.log(`│                                                         │`);
        console.log('└─────────────────────────────────────────────────────────┘');
      }
      
      console.log('');
      
    } catch (error) {
      spinner.stop();
      handleApiError(error);
    }
  });

// ==================== VERIFY CHECK ====================

const check = new Command('check')
  .description('Check/Submit a verification code')
  .alias('submit')
  .option('-n, --number <number>', 'Phone number being verified')
  .option('-c, --code <code>', 'Verification code received')
  .option('-p, --profile <id>', 'Verify profile ID')
  .action(async (options) => {
    let { number, code, profile } = options;
    
    // Interactive prompts for missing fields
    const prompts = [];
    
    if (!number) {
      prompts.push({
        type: 'input',
        name: 'number',
        message: 'Phone number being verified:',
        validate: (input) => {
          if (!input || !input.startsWith('+')) {
            return 'Phone number must be in E.164 format (e.g., +13125551234)';
          }
          return true;
        }
      });
    }
    
    if (!code) {
      prompts.push({
        type: 'input',
        name: 'code',
        message: 'Verification code:',
        validate: (input) => {
          if (!input || input.length < 4) {
            return 'Please enter the verification code (usually 6 digits)';
          }
          return true;
        }
      });
    }
    
    if (prompts.length > 0) {
      const answers = await inquirer.prompt(prompts);
      number = number || answers.number;
      code = code || answers.code;
    }
    
    // Validate required fields
    if (!number || !code) {
      showError('❌ Both phone number and code are required');
      process.exit(1);
    }
    
    const spinner = ora({
      text: 'Verifying code...',
      spinner: 'dots'
    }).start();
    
    try {
      const checkOptions = {};
      if (profile) checkOptions.profileId = profile;
      
      const data = await submitVerification(number, code, checkOptions);
      
      spinner.stop();
      
      const result = data.data;
      
      if (result.verified || result.status === 'verified') {
        showSuccess('✅ Verification successful!');
        console.log('');
        console.log(`${primary('Phone Number:')}    ${number}`);
        console.log(`${primary('Status:')}          Verified ✓`);
        
        if (result.verified_at) {
          console.log(`${primary('Verified At:')}     ${new Date(result.verified_at).toLocaleString()}`);
        }
      } else {
        showError('❌ Verification failed');
        console.log('');
        console.log(`${primary('Phone Number:')}    ${number}`);
        console.log(`${primary('Status:')}          ${result.status || 'failed'}`);
        
        if (result.remaining_attempts !== undefined) {
          console.log(`${primary('Attempts Left:')}   ${result.remaining_attempts}`);
        }
        
        if (result.error_message) {
          showInfo(`Error: ${result.error_message}`);
        }
      }
      
      console.log('');
      
    } catch (error) {
      spinner.stop();
      handleApiError(error);
    }
  });

// ==================== VERIFY STATUS ====================

const status = new Command('status')
  .description('Check the status of a verification')
  .alias('info')
  .argument('[verificationId]', 'Verification ID to check')
  .option('-j, --json', 'Output raw JSON')
  .action(async (verificationId, options) => {
    let targetId = verificationId;
    
    // If no ID provided, prompt for one
    if (!targetId) {
      const { id } = await inquirer.prompt([
        {
          type: 'input',
          name: 'id',
          message: 'Verification ID:',
          validate: (input) => input.length > 0 || 'Verification ID is required'
        }
      ]);
      targetId = id;
    }
    
    const spinner = ora({
      text: 'Fetching verification status...',
      spinner: 'dots'
    }).start();
    
    try {
      const data = await getVerification(targetId);
      
      spinner.stop();
      
      if (!data.data) {
        showError('Verification not found');
        return;
      }
      
      const verification = data.data;
      
      if (options.json) {
        console.log(JSON.stringify(data, null, 2));
        return;
      }
      
      showSuccess(`Verification Status: ${targetId}`);
      console.log('');
      
      console.log('┌─────────────────────────────────────────────────────────┐');
      console.log('│  🔐 ' + primary('VERIFICATION DETAILS') + '                                │');
      console.log('├─────────────────────────────────────────────────────────┤');
      console.log(`│  Verification ID: ${(verification.id || 'N/A').padEnd(40)}│`);
      console.log(`│  Phone Number:    ${(verification.phone_number || 'N/A').padEnd(40)}│`);
      console.log(`│  Type:            ${(verification.type || 'N/A').padEnd(40)}│`);
      console.log(`│  Status:          ${(verification.status || 'N/A').padEnd(40)}│`);
      
      if (verification.verified !== undefined) {
        console.log(`│  Verified:        ${(verification.verified ? 'Yes ✓' : 'No ✗').padEnd(40)}│`);
      }
      
      console.log('├─────────────────────────────────────────────────────────┤');
      console.log(`│  Created:         ${(verification.created_at ? new Date(verification.created_at).toLocaleString() : 'N/A').padEnd(40)}│`);
      
      if (verification.verified_at) {
        console.log(`│  Verified At:     ${(new Date(verification.verified_at).toLocaleString()).padEnd(40)}│`);
      }
      
      if (verification.expires_at) {
        console.log(`│  Expires At:      ${(new Date(verification.expires_at).toLocaleString()).padEnd(40)}│`);
      }
      
      console.log('└─────────────────────────────────────────────────────────┘');
      
      if (verification.remaining_attempts !== undefined && verification.remaining_attempts < 3) {
        console.log('');
        showInfo(`⚠️  Only ${verification.remaining_attempts} attempt(s) remaining`);
      }
      
      console.log('');
      
    } catch (error) {
      spinner.stop();
      handleApiError(error);
    }
  });

// ==================== VERIFY LIST ====================

const list = new Command('list')
  .description('List recent verifications')
  .alias('ls')
  .option('-s, --status <status>', 'Filter by status: pending, verified, expired, failed')
  .option('-n, --limit <number>', 'Number of results', '20')
  .action(async (options) => {
    const { status, limit } = options;
    
    const spinner = ora({
      text: 'Fetching verifications...',
      spinner: 'dots'
    }).start();
    
    try {
      const data = await listVerifications({
        status,
        limit: parseInt(limit)
      });
      
      spinner.stop();
      
      if (!data.data || data.data.length === 0) {
        showInfo('📭 No verifications found.');
        return;
      }
      
      showSuccess(`Found ${data.data.length} verification(s)`);
      console.log('');
      
      const table = new Table({
        head: ['Verification ID', 'Phone Number', 'Type', 'Status', 'Created'],
        colWidths: [26, 18, 10, 12, 20],
        style: {
          head: [COLORS.primary.replace('#', '')],
          border: ['gray']
        }
      });
      
      data.data.forEach((item) => {
        const verification = item.data || item;
        const time = verification.created_at 
          ? new Date(verification.created_at).toLocaleString() 
          : 'N/A';
        
        table.push([
          (verification.id || 'N/A').substring(0, 24),
          verification.phone_number || 'N/A',
          verification.type || 'N/A',
          getStatusIcon(verification.status) + ' ' + (verification.status || 'unknown'),
          time
        ]);
      });
      
      console.log(table.toString());
      console.log('');
      
    } catch (error) {
      spinner.stop();
      handleApiError(error);
    }
  });

// ==================== HELPERS ====================

function getStatusIcon(status) {
  const icons = {
    'pending': '⏳',
    'sent': '📤',
    'verified': '✅',
    'expired': '⏰',
    'failed': '❌',
    'max_attempts_reached': '❌'
  };
  return icons[status?.toLowerCase()] || '•';
}

function handleApiError(error) {
  if (error.response?.status === 401) {
    showError('🔐 Authentication failed. Run: telnyx auth login');
  } else if (error.response?.status === 403) {
    showError('🚫 Permission denied. Verify API may not be enabled on your account.');
    showInfo('   Contact Telnyx support to enable verification services.');
  } else if (error.response?.status === 404) {
    showError('❌ Verification not found. Check the ID and try again.');
  } else if (error.response?.status === 422) {
    const detail = error.response.data?.errors?.[0]?.detail || 
                   error.response.data?.errors?.[0]?.title || 
                   'Invalid request';
    showError(`❌ ${detail}`);
    
    if (detail.includes('code') || detail.includes('expired')) {
      showInfo('   Tip: Verification codes expire after the timeout period');
      showInfo('   Use "telnyx verify send" to generate a new code');
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
  check,
  status,
  list
};
