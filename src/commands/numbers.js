const { Command } = require('commander');
const inquirer = require('inquirer');
const ora = require('ora');
const Table = require('cli-table3');
const chalk = require('chalk');
const { 
  searchAvailableNumbers, 
  listPhoneNumbers, 
  orderPhoneNumber 
} = require('../api/client');
const { showSuccess, showError, showInfo, COLORS } = require('../ui/layout');

const primary = chalk.hex(COLORS.primary);
const gray = chalk.gray;
const yellow = chalk.yellow;

// ==================== NUMBER SEARCH ====================

const search = new Command('search')
  .description('Search for available phone numbers')
  .alias('s')  // Short alias like original CLI
  .option('-c, --country <code>', 'Country code (e.g., US, GB, CA)')
  .option('-l, --locality <city>', 'City/Locality (e.g., Chicago)')
  .option('-t, --type <type>', 'Number type: local, toll_free, mobile', 'local')
  .option('-n, --limit <number>', 'Number of results', '10')
  .option('-a, --area-code <code>', 'Area code/National destination code (NDC)')
  .option('-s, --starts-with <digits>', 'Numbers starting with these digits')
  .option('-e, --ends-with <digits>', 'Numbers ending with these digits')
  .option('--contains <digits>', 'Numbers containing these digits')
  .option('-f, --features <features>', 'Comma-separated features: sms,emergency,mms,fax')
  .option('--quickship', 'Show only quickship numbers (US toll-free, ready immediately)')
  .option('--best-effort', 'Return similar results if exact matches not found')
  .option('--state <state>', 'State/Province code (US/CA only, e.g., IL, CA)')
  .action(async (options) => {
    let { 
      country, locality, type, limit, areaCode, 
      startsWith, endsWith, contains, features, 
      quickship, bestEffort, state 
    } = options;
    
    // Interactive prompts for missing required fields
    const prompts = [];
    
    if (!country) {
      prompts.push({
        type: 'input',
        name: 'country',
        message: 'Country code (ISO 3166-1 alpha-2, e.g., US, GB, CA):',
        default: 'US',
        validate: (input) => {
          if (!input || input.trim().length < 2) {
            return 'Country code is required (2-letter ISO code)';
          }
          return true;
        }
      });
    }
    
    // For US/CA numbers, prompt for additional filters
    const countryCode = (country || 'US').toUpperCase();
    if (countryCode === 'US' || countryCode === 'CA') {
      if (!locality && !state && !areaCode) {
        prompts.push({
          type: 'list',
          name: 'searchBy',
          message: 'How would you like to search?',
          choices: [
            { name: 'By City/Locality', value: 'locality' },
            { name: 'By State/Province', value: 'state' },
            { name: 'By Area Code', value: 'areacode' },
            { name: 'No specific location', value: 'none' }
          ],
          default: 'locality'
        });
      }
    }
    
    if (prompts.length > 0) {
      const answers = await inquirer.prompt(prompts);
      country = country || answers.country;
      
      // Handle searchBy selection
      if (answers.searchBy) {
        if (answers.searchBy === 'locality') {
          const { loc } = await inquirer.prompt([{
            type: 'input',
            name: 'loc',
            message: 'City/Locality (e.g., Chicago):',
            validate: (input) => input.length > 0 || 'Locality is required'
          }]);
          locality = loc;
        } else if (answers.searchBy === 'state') {
          const { st } = await inquirer.prompt([{
            type: 'input',
            name: 'st',
            message: 'State/Province code (e.g., IL, CA):',
            validate: (input) => input.length === 2 || 'Please enter 2-letter state code'
          }]);
          state = st;
        } else if (answers.searchBy === 'areacode') {
          const { ac } = await inquirer.prompt([{
            type: 'input',
            name: 'ac',
            message: 'Area code (e.g., 312):',
            validate: (input) => input.length === 3 || 'Please enter 3-digit area code'
          }]);
          areaCode = ac;
        }
      }
    }
    
    // Ask for features if not provided
    if (!features) {
      const { enableFeatures } = await inquirer.prompt([{
        type: 'checkbox',
        name: 'enableFeatures',
        message: 'Select required features (optional):',
        choices: [
          { name: 'SMS', value: 'sms', checked: true },
          { name: 'MMS', value: 'mms' },
          { name: 'Emergency Calling', value: 'emergency' },
          { name: 'Fax', value: 'fax' },
          { name: 'Voice', value: 'voice' }
        ]
      }]);
      if (enableFeatures.length > 0) {
        features = enableFeatures.join(',');
      }
    }
    
    const spinner = ora({
      text: `Searching for ${type} numbers in ${locality || state || areaCode || country}...`,
      spinner: 'dots'
    }).start();
    
    try {
      const data = await searchAvailableNumbers({
        country: country.toUpperCase(),
        locality,
        numberType: type,
        limit: parseInt(limit),
        areaCode,
        startsWith,
        endsWith,
        contains,
        features,
        quickship,
        bestEffort,
        state
      });
      
      spinner.stop();
      
      if (!data.data || data.data.length === 0) {
        showInfo('No available numbers found matching your criteria.');
        showInfo('Try:');
        showInfo('  • Using --best-effort for similar results');
        showInfo('  • Broadening your location criteria');
        showInfo('  • Searching for different number types');
        return;
      }
      
      showSuccess(`Found ${data.data.length} available ${type} number(s)`);
      console.log('');
      
      // Create table with cost information
      const table = new Table({
        head: ['#', 'Phone Number', 'Location', 'Features', 'Monthly Cost'],
        colWidths: [4, 18, 20, 25, 15],
        style: {
          head: [COLORS.primary.replace('#', '')],
          border: ['gray']
        }
      });
      
      data.data.forEach((item, index) => {
        const num = item.data || item;
        const cost = num.cost_information?.upfront_cost || num.monthly_cost || 'N/A';
        const costStr = cost !== 'N/A' ? `$${cost}` : 'N/A';
        
        table.push([
          index + 1,
          primary(num.phone_number || num.e164_address || 'N/A'),
          num.locality || num.region || num.state || 'N/A',
          (num.features || []).join(', ') || 'Standard',
          yellow(costStr)
        ]);
      });
      
      console.log(table.toString());
      console.log('');
      showInfo('💡 Use "telnyx number order <phone_number>" to purchase');
      showInfo('   Note: Numbers must be purchased within 30 minutes of search');
      
    } catch (error) {
      spinner.stop();
      handleApiError(error, 'search');
    }
  });

// ==================== NUMBER LIST ====================

const list = new Command('list')
  .description('List your purchased phone numbers')
  .alias('ls')  // Common alias
  .option('-s, --status <status>', 'Filter by status: active, disabled, pending', 'active')
  .option('-t, --type <type>', 'Filter by type: local, toll_free, mobile')
  .option('-n, --limit <number>', 'Number of results', '20')
  .option('--no-features', 'Hide features column for compact view')
  .action(async (options) => {
    const { status, type, limit, features } = options;
    
    const spinner = ora({
      text: 'Fetching your phone numbers...',
      spinner: 'dots'
    }).start();
    
    try {
      const data = await listPhoneNumbers({
        status,
        numberType: type,
        limit: parseInt(limit)
      });
      
      spinner.stop();
      
      if (!data.data || data.data.length === 0) {
        showInfo('📭 You have no phone numbers.');
        console.log('');
        showInfo('Run "telnyx number search" to find and purchase numbers.');
        return;
      }
      
      showSuccess(`You have ${data.data.length} phone number(s)`);
      console.log('');
      
      // Create table
      const headers = features !== false 
        ? ['Phone Number', 'Type', 'Status', 'Messaging', 'Voice', 'Features']
        : ['Phone Number', 'Type', 'Status', 'Messaging', 'Voice'];
      
      const colWidths = features !== false
        ? [18, 12, 12, 12, 12, 20]
        : [18, 12, 12, 12, 12];
      
      const table = new Table({
        head: headers,
        colWidths: colWidths,
        style: {
          head: [COLORS.primary.replace('#', '')],
          border: ['gray']
        }
      });
      
      data.data.forEach((item) => {
        const num = item.data || item;
        const messagingEnabled = num.messaging_enabled || num.messaging?.enabled ? '✓' : '✗';
        const voiceEnabled = num.voice_enabled || num.voice?.enabled ? '✓' : '✗';
        const featuresList = (num.features || []).slice(0, 3).join(', ') || '-';
        
        const row = features !== false
          ? [
              primary(num.phone_number || num.e164_address || 'N/A'),
              num.number_type || num.type || 'N/A',
              num.status || 'N/A',
              messagingEnabled,
              voiceEnabled,
              featuresList
            ]
          : [
              primary(num.phone_number || num.e164_address || 'N/A'),
              num.number_type || num.type || 'N/A',
              num.status || 'N/A',
              messagingEnabled,
              voiceEnabled
            ];
        
        table.push(row);
      });
      
      console.log(table.toString());
      
      if (data.meta?.total_pages > 1) {
        console.log('');
        showInfo(`Page ${data.meta.page_number} of ${data.meta.total_pages} (${data.meta.total_results} total)`);
      }
      
    } catch (error) {
      spinner.stop();
      handleApiError(error, 'list');
    }
  });

// ==================== NUMBER ORDER ====================

const order = new Command('order')
  .description('Purchase a phone number')
  .alias('buy')  // More intuitive alias
  .argument('[phoneNumber]', 'Phone number to purchase (E.164 format, e.g., +13125551234)')
  .option('-m, --messaging', 'Enable messaging', true)
  .option('-v, --voice', 'Enable voice', true)
  .option('-y, --yes', 'Skip confirmation prompt')
  .option('--profile <id>', 'Messaging profile ID to assign')
  .option('--app <id>', 'TeXML application ID for voice')
  .action(async (phoneNumber, options) => {
    let targetNumber = phoneNumber;
    
    // If no number provided, prompt for one
    if (!targetNumber) {
      const { number } = await inquirer.prompt([
        {
          type: 'input',
          name: 'number',
          message: 'Enter the phone number to purchase (E.164 format, e.g., +13125551234):',
          validate: (input) => {
            if (!input || !input.startsWith('+')) {
              return 'Phone number must be in E.164 format (e.g., +13125551234)';
            }
            if (input.length < 10) {
              return 'Phone number appears to be too short';
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
    
    // Show order summary and confirm
    if (!options.yes) {
      console.log('');
      console.log('┌─────────────────────────────────────────────────────────┐');
      console.log('│  📦 ' + primary('ORDER SUMMARY') + '                                      │');
      console.log('├─────────────────────────────────────────────────────────┤');
      console.log(`│  Phone Number:  ${targetNumber.padEnd(40)}│`);
      console.log(`│  Messaging:     ${(options.messaging ? '✓ Enabled' : '✗ Disabled').padEnd(40)}│`);
      console.log(`│  Voice:         ${(options.voice ? '✓ Enabled' : '✗ Disabled').padEnd(40)}│`);
      if (options.profile) {
        console.log(`│  Msg Profile:   ${options.profile.substring(0, 40).padEnd(40)}│`);
      }
      if (options.app) {
        console.log(`│  Voice App:     ${options.app.substring(0, 40).padEnd(40)}│`);
      }
      console.log('├─────────────────────────────────────────────────────────┤');
      console.log('│  ⚠️  You will be charged for this number immediately    │');
      console.log('└─────────────────────────────────────────────────────────┘');
      console.log('');
      
      const { confirm } = await inquirer.prompt([
        {
          type: 'confirm',
          name: 'confirm',
          message: 'Confirm purchase?',
          default: false
        }
      ]);
      
      if (!confirm) {
        showInfo('Order cancelled.');
        return;
      }
    }
    
    const spinner = ora({
      text: `Processing purchase for ${targetNumber}...`,
      spinner: 'dots'
    }).start();
    
    try {
      const orderOptions = {};
      if (options.messaging) orderOptions.messaging_enabled = true;
      if (options.voice) orderOptions.voice_enabled = true;
      if (options.profile) orderOptions.messaging_profile_id = options.profile;
      if (options.app) orderOptions.texml_application_id = options.app;
      
      const data = await orderPhoneNumber(targetNumber, orderOptions);
      
      spinner.stop();
      
      showSuccess(`🎉 Successfully purchased ${targetNumber}`);
      console.log('');
      
      const num = data.data;
      console.log(`${primary('Phone Number:')}  ${num.phone_number || num.e164_address}`);
      console.log(`${primary('ID:')}            ${num.id}`);
      console.log(`${primary('Status:')}        ${num.status || 'active'}`);
      console.log(`${primary('Messaging:')}     ${num.messaging_enabled ? '✓ Enabled' : '✗ Disabled'}`);
      console.log(`${primary('Voice:')}         ${num.voice_enabled ? '✓ Enabled' : '✗ Disabled'}`);
      
      if (num.regulatory_requirements?.length > 0) {
        console.log('');
        showInfo('⚠️  This number has regulatory requirements that must be fulfilled.');
        showInfo('   Run: telnyx number requirements ' + num.id);
      }
      
      console.log('');
      showSuccess('Your number is ready to use! 🚀');
      
    } catch (error) {
      spinner.stop();
      handleApiError(error, 'order');
    }
  });

// ==================== ERROR HANDLER ====================

function handleApiError(error, context) {
  if (error.response?.status === 401) {
    showError('🔐 Authentication failed. Run: telnyx auth login');
  } else if (error.response?.status === 403) {
    showError('🚫 Permission denied. Your account may not have access to this feature.');
  } else if (error.response?.status === 422) {
    const detail = error.response.data?.errors?.[0]?.detail || 
                   error.response.data?.errors?.[0]?.title || 
                   'Invalid request parameters';
    
    if (context === 'search') {
      showError(`🔍 Search failed: ${detail}`);
      showInfo('   Tip: Use --best-effort flag for more flexible results');
    } else if (context === 'order') {
      showError(`📦 Order failed: ${detail}`);
      if (detail.includes('unavailable') || detail.includes('taken')) {
        showInfo('   This number may have been purchased by someone else.');
        showInfo('   Run "telnyx number search" to find similar numbers.');
      }
    } else {
      showError(`❌ ${detail}`);
    }
  } else if (error.response?.status === 402) {
    showError('💳 Insufficient balance. Please add funds to your account.');
    showInfo('   Visit: https://portal.telnyx.com/#/app/account/billing');
  } else if (error.response?.status === 429) {
    showError('⏱️  Rate limit exceeded. Please wait a moment and try again.');
  } else if (error.code === 'ECONNABORTED') {
    showError('⏱️  Request timed out. Please try again.');
  } else {
    showError(`❌ ${error.message}`);
  }
  process.exit(1);
}

module.exports = {
  search,
  list,
  order
};
