const { Command } = require('commander');
const ora = require('ora');
const chalk = require('chalk');
const { getBillingBalance } = require('../api/client');
const { showSuccess, showError, showInfo, COLORS } = require('../ui/layout');

const primary = chalk.hex(COLORS.primary);
const gray = chalk.gray;
const yellow = chalk.yellow;
const green = chalk.green;
const red = chalk.red;

// ==================== BILLING BALANCE ====================

const balance = new Command('balance')
  .description('Check your account balance and billing information')
  .alias('bal')
  .option('-j, --json', 'Output raw JSON')
  .option('--no-tips', 'Hide usage tips')
  .action(async (options) => {
    const spinner = ora({
      text: 'Fetching billing information...',
      spinner: 'dots'
    }).start();
    
    try {
      const data = await getBillingBalance();
      
      spinner.stop();
      
      if (options.json) {
        console.log(JSON.stringify(data, null, 2));
        return;
      }
      
      if (!data.data) {
        showInfo('💳 No billing information available.');
        return;
      }
      
      const bal = data.data;
      const currency = bal.currency || 'USD';
      
      // Format amounts
      const balanceCents = bal.balance_cents || 0;
      const availableCents = bal.available_cents || balanceCents;
      const pendingCents = balanceCents - availableCents;
      
      const balanceStr = formatCurrency(balanceCents, currency);
      const availableStr = formatCurrency(availableCents, currency);
      const pendingStr = pendingCents > 0 ? formatCurrency(pendingCents, currency) : null;
      
      // Determine balance status
      let statusIcon = '✅';
      let statusColor = green;
      if (balanceCents < 1000) { // Less than $10
        statusIcon = '⚠️';
        statusColor = yellow;
      }
      if (balanceCents < 100) { // Less than $1
        statusIcon = '❌';
        statusColor = red;
      }
      
      showSuccess(`${statusIcon} Billing information retrieved`);
      console.log('');
      
      // Main balance box
      console.log('┌─────────────────────────────────────────────────────────┐');
      console.log('│  💰 ' + primary('ACCOUNT BALANCE') + '                                      │');
      console.log('├─────────────────────────────────────────────────────────┤');
      console.log(`│                                                         │`);
      console.log(`│  Current Balance:      ${balanceStr.padEnd(35)}│`);
      console.log(`│  Available Balance:    ${availableStr.padEnd(35)}│`);
      if (pendingStr) {
        console.log(`│  Pending Charges:     -${pendingStr.padEnd(34)}│`);
      }
      console.log(`│                                                         │`);
      console.log(`│  Currency:             ${(currency).padEnd(35)}│`);
      console.log(`│  Billing Type:         ${(bal.billing_type || 'Prepaid').padEnd(35)}│`);
      console.log(`│                                                         │`);
      console.log('└─────────────────────────────────────────────────────────┘');
      
      // Auto-recharge info
      if (bal.auto_recharge_enabled) {
        console.log('');
        console.log('┌─────────────────────────────────────────────────────────┐');
        console.log('│  🔄 ' + primary('AUTO-RECHARGE SETTINGS') + '                             │');
        console.log('├─────────────────────────────────────────────────────────┤');
        
        if (bal.auto_recharge_threshold_cents) {
          const threshold = formatCurrency(bal.auto_recharge_threshold_cents, currency);
          console.log(`│  Trigger When Below:   ${threshold.padEnd(35)}│`);
        }
        
        if (bal.auto_recharge_amount_cents) {
          const amount = formatCurrency(bal.auto_recharge_amount_cents, currency);
          console.log(`│  Recharge Amount:      ${amount.padEnd(35)}│`);
        }
        
        if (bal.auto_recharge_payment_method) {
          console.log(`│  Payment Method:       ${bal.auto_recharge_payment_method.padEnd(35)}│`);
        }
        
        console.log('└─────────────────────────────────────────────────────────┘');
      }
      
      console.log('');
      
      // Warnings and tips
      if (options.tips !== false) {
        if (balanceCents < 1000) {
          showInfo('⚠️  Your balance is running low!');
          showInfo('   Add funds at: https://portal.telnyx.com/#/app/account/billing');
          console.log('');
        }
        
        if (!bal.auto_recharge_enabled && balanceCents > 5000) {
          showInfo('💡 Tip: Enable auto-recharge to avoid service interruptions');
          console.log('');
        }
      }
      
    } catch (error) {
      spinner.stop();
      handleApiError(error);
    }
  });

// ==================== HELPERS ====================

function formatCurrency(cents, currency = 'USD') {
  if (cents === undefined || cents === null) return 'N/A';
  
  const amount = (cents / 100).toFixed(2);
  const symbol = getCurrencySymbol(currency);
  
  // Add color based on amount
  if (cents < 100) return red(`${symbol}${amount}`);
  if (cents < 1000) return yellow(`${symbol}${amount}`);
  return green(`${symbol}${amount}`);
}

function getCurrencySymbol(currency) {
  const symbols = {
    'USD': '$',
    'EUR': '€',
    'GBP': '£',
    'CAD': 'C$',
    'AUD': 'A$',
    'JPY': '¥',
    'MXN': 'Mex$'
  };
  
  return symbols[currency] || currency + ' ';
}

function handleApiError(error) {
  if (error.response?.status === 401) {
    showError('🔐 Authentication failed. Run: telnyx auth login');
  } else if (error.response?.status === 403) {
    showError('🚫 You do not have permission to view billing information.');
    showInfo('   Contact your account administrator.');
  } else if (error.response?.status === 404) {
    showError('💳 Billing information not found.');
    showInfo('   Your account may not be fully set up yet.');
  } else {
    showError(`❌ Failed to fetch balance: ${error.message}`);
  }
  process.exit(1);
}

module.exports = {
  balance
};
