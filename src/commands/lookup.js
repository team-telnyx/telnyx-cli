const { Command } = require('commander');
const inquirer = require('inquirer');
const ora = require('ora');
const Table = require('cli-table3');
const chalk = require('chalk');
const { lookupNumber } = require('../api/client');
const { showSuccess, showError, showInfo, COLORS } = require('../ui/layout');

const primary = chalk.hex(COLORS.primary);
const gray = chalk.gray;
const yellow = chalk.yellow;
const green = chalk.green;
const red = chalk.red;

// Helper to format output based on --json and --output flags
function formatOutput(data, format) {
  if (format === 'json') {
    console.log(JSON.stringify(data, null, 2));
    return;
  }
  
  if (format === 'csv') {
    if (data.data) {
      const item = data.data;
      const headers = Object.keys(item);
      console.log(headers.join(','));
      const values = headers.map(h => {
        const val = item[h];
        if (val === null || val === undefined) return '';
        if (typeof val === 'object') return JSON.stringify(val).replace(/,/g, ';');
        return String(val).replace(/,/g, ';');
      });
      console.log(values.join(','));
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

// ==================== LOOKUP COMMAND ====================

const lookup = new Command('lookup')
  .description('Look up information about a phone number')
  .alias('whois')
  .argument('[phoneNumber]', 'Phone number to look up (E.164 format)')
  .option('-t, --type <type>', 'Lookup type: caller-name, carrier, portability', 'caller-name')
  .option('-f, --fields <fields>', 'Comma-separated fields to include')
  .option('-j, --json', 'Output raw JSON')
  .option('-o, --output <format>', 'Output format: json, table, csv', 'table')
  .action(async (phoneNumber, options) => {
    let targetNumber = phoneNumber;
    const { type, fields } = options;
    const outputFormat = getOutputFormat(options);
    
    // Interactive prompt for missing phone number
    if (!targetNumber) {
      const { number } = await inquirer.prompt([
        {
          type: 'input',
          name: 'number',
          message: 'Phone number to look up (E.164 format):',
          validate: (input) => {
            if (!input || !input.startsWith('+')) {
              return 'Phone number must be in E.164 format (e.g., +13125551234)';
            }
            return true;
          }
        }
      ]);
      targetNumber = number;
    }
    
    // Validate E.164 format
    if (!targetNumber.startsWith('+')) {
      showError('❌ Phone number must be in E.164 format (e.g., +13125551234)');
      showInfo('   E.164 format: +[country code][number]');
      process.exit(1);
    }
    
    const spinner = ora({
      text: `Looking up ${targetNumber}...`,
      spinner: 'dots'
    }).start();
    
    try {
      const lookupOptions = { type };
      if (fields) {
        lookupOptions.fields = fields;
      }
      
      const data = await lookupNumber(targetNumber, lookupOptions);
      
      spinner.stop();
      
      if (!data.data) {
        showError('No lookup data available for this number');
        return;
      }
      
      const result = data.data;
      
      // Handle JSON/CSV output
      if (outputFormat !== 'table') {
        formatOutput(result, outputFormat);
        return;
      }
      
      showSuccess(`Lookup results for ${targetNumber}`);
      console.log('');
      
      // Main info box
      console.log('┌─────────────────────────────────────────────────────────┐');
      console.log('│  🔍 ' + primary('NUMBER LOOKUP RESULTS') + '                             │');
      console.log('├─────────────────────────────────────────────────────────┤');
      console.log(`│  Phone Number:    ${(result.phone_number || targetNumber).padEnd(40)}│`);
      console.log(`│  Number Type:     ${(result.number_type || 'Unknown').padEnd(40)}│`);
      console.log(`│  Country Code:    ${(result.country_code || 'Unknown').padEnd(40)}│`);
      console.log('└─────────────────────────────────────────────────────────┘');
      
      // Caller Name section
      if (result.caller_name) {
        console.log('');
        console.log('┌─────────────────────────────────────────────────────────┐');
        console.log('│  👤 ' + primary('CALLER NAME') + '                                         │');
        console.log('├─────────────────────────────────────────────────────────┤');
        console.log(`│  Name:            ${(result.caller_name.caller_name || 'Unavailable').padEnd(40)}│`);
        if (result.caller_name.error_code) {
          console.log(`│  Status:          ${(result.caller_name.error_code).padEnd(40)}│`);
        }
        console.log('└─────────────────────────────────────────────────────────┘');
      }
      
      // Carrier info section
      if (result.carrier) {
        console.log('');
        console.log('┌─────────────────────────────────────────────────────────┐');
        console.log('│  📡 ' + primary('CARRIER INFORMATION') + '                                 │');
        console.log('├─────────────────────────────────────────────────────────┤');
        console.log(`│  Name:            ${(result.carrier.name || 'Unknown').padEnd(40)}│`);
        console.log(`│  Type:            ${(result.carrier.type || 'Unknown').padEnd(40)}│`);
        console.log(`│  Mobile Country:  ${(result.carrier.mobile_country_code || 'N/A').padEnd(40)}│`);
        console.log(`│  Mobile Network:  ${(result.carrier.mobile_network_code || 'N/A').padEnd(40)}│`);
        console.log('└─────────────────────────────────────────────────────────┘');
      }
      
      // Portability section
      if (result.portability) {
        console.log('');
        console.log('┌─────────────────────────────────────────────────────────┐');
        console.log('│  🔄 ' + primary('PORTABILITY') + '                                         │');
        console.log('├─────────────────────────────────────────────────────────┤');
        console.log(`│  Portable:        ${(result.portability.portable ? 'Yes ✓' : 'No ✗').padEnd(40)}│`);
        if (result.portability.ported_status) {
          console.log(`│  Ported Status:   ${(result.portability.ported_status).padEnd(40)}│`);
        }
        if (result.portability.ported_date) {
          console.log(`│  Ported Date:     ${(result.portability.ported_date).padEnd(40)}│`);
        }
        if (result.portability.spid) {
          console.log(`│  SPID:            ${(result.portability.spid).padEnd(40)}│`);
        }
        if (result.portability.spid_name) {
          console.log(`│  SPID Name:       ${(result.portability.spid_name).padEnd(40)}│`);
        }
        console.log('└─────────────────────────────────────────────────────────┘');
      }
      
      // Fraud info section
      if (result.fraud) {
        console.log('');
        console.log('┌─────────────────────────────────────────────────────────┐');
        console.log('│  🛡️  ' + primary('FRAUD PROTECTION') + '                                    │');
        console.log('├─────────────────────────────────────────────────────────┤');
        console.log(`│  Risk Score:      ${(result.fraud.risk_score !== undefined ? result.fraud.risk_score : 'N/A').padEnd(40)}│`);
        console.log(`│  Risk Level:      ${(result.fraud.risk_level || 'Unknown').padEnd(40)}│`);
        if (result.fraud.recommendation) {
          console.log(`│  Recommendation:  ${(result.fraud.recommendation).padEnd(40)}│`);
        }
        console.log('└─────────────────────────────────────────────────────────┘');
      }
      
      console.log('');
      showInfo('💡 Lookup types available: caller-name, carrier, portability');
      showInfo('   Use --type to specify a different lookup type');
      
    } catch (error) {
      spinner.stop();
      showError(error.message);
      process.exit(1);
    }
  });

// ==================== LOOKUP BATCH (for future use) ====================

const batch = new Command('batch')
  .description('Look up multiple phone numbers (from file)')
  .argument('<file>', 'Path to file containing phone numbers (one per line)')
  .option('-t, --type <type>', 'Lookup type', 'caller-name')
  .option('-o, --output <file>', 'Output file for results')
  .option('-j, --json', 'Output raw JSON')
  .action(async (file, options) => {
    showInfo('Batch lookup feature coming soon!');
    showInfo('For now, use multiple individual lookup commands.');
  });

module.exports = {
  lookup,
  batch
};
