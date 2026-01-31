const { Command } = require('commander');
const inquirer = require('inquirer');
const ora = require('ora');
const { getApiKey, setApiKey, hasApiKey } = require('../config/store');
const { get } = require('../api/client');
const { showSuccess, showError, showInfo, COLORS } = require('../ui/layout');
const chalk = require('chalk');

const primary = chalk.hex(COLORS.primary);

// whoami command
const whoami = new Command('whoami')
  .description('Display current account information')
  .action(async () => {
    if (!hasApiKey()) {
      showError('Not authenticated. Run: telnyx auth login');
      process.exit(1);
    }

    const spinner = ora('Fetching account info...').start();

    try {
      const data = await get('/whoami');
      spinner.stop();

      if (data && data.data) {
        const user = data.data;
        showSuccess('Authenticated');
        console.log('');
        console.log(`${primary('User ID:')}    ${user.user_id || user.id || 'N/A'}`);
        console.log(`${primary('Email:')}      ${user.email || 'N/A'}`);
        console.log(`${primary('Org Name:')}   ${user.organization_name || user.org_name || 'N/A'}`);
        console.log('');
      } else {
        spinner.stop();
        showInfo('Account info retrieved but no data available.');
      }
    } catch (error) {
      spinner.stop();
      if (error.response && error.response.status === 401) {
        showError('Invalid API key. Please run: telnyx auth login');
      } else {
        showError(`Failed to fetch account info: ${error.message}`);
      }
      process.exit(1);
    }
  });

// login command
const login = new Command('login')
  .description('Authenticate with your Telnyx API key')
  .action(async () => {
    const currentKey = getApiKey();
    
    if (currentKey) {
      const { overwrite } = await inquirer.prompt([
        {
          type: 'confirm',
          name: 'overwrite',
          message: 'An API key is already configured. Overwrite?',
          default: false
        }
      ]);
      
      if (!overwrite) {
        showInfo('Login cancelled.');
        return;
      }
    }

    const { apiKey } = await inquirer.prompt([
      {
        type: 'password',
        name: 'apiKey',
        message: 'Enter your Telnyx API Key:',
        mask: '*',
        validate: (input) => {
          if (!input || input.trim().length === 0) {
            return 'API key is required';
          }
          return true;
        }
      }
    ]);

    const spinner = ora('Validating API key...').start();

    try {
      // Test the API key by making a whoami request
      const axios = require('axios');
      const response = await axios.get('https://api.telnyx.com/v2/whoami', {
        headers: {
          'Authorization': `Bearer ${apiKey}`,
          'Content-Type': 'application/json'
        },
        timeout: 10000
      });

      spinner.stop();
      setApiKey(apiKey);
      showSuccess('Successfully authenticated with Telnyx');
      
      if (response.data && response.data.data) {
        const user = response.data.data;
        console.log('');
        console.log(`${primary('Email:')}    ${user.email || 'N/A'}`);
        console.log(`${primary('Org:')}      ${user.organization_name || user.org_name || 'N/A'}`);
      }
    } catch (error) {
      spinner.stop();
      if (error.response && error.response.status === 401) {
        showError('Invalid API key. Please check your key and try again.');
      } else {
        showError(`Authentication failed: ${error.message}`);
      }
      process.exit(1);
    }
  });

module.exports = {
  whoami,
  login
};
