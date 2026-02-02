const { Command } = require('commander');
const inquirer = require('inquirer');
const ora = require('ora');
const Table = require('cli-table3');
const chalk = require('chalk');
const { 
  getApiKey, 
  setApiKey, 
  hasApiKey, 
  clearApiKey,
  getActiveProfile,
  setActiveProfile,
  listProfiles,
  profileExists,
  deleteProfile
} = require('../config/store');
const { get } = require('../api/client');
const { showSuccess, showError, showInfo, COLORS } = require('../ui/layout');

const primary = chalk.hex(COLORS.primary);
const gray = chalk.gray;

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
        const profile = getActiveProfile();
        showSuccess('Authenticated');
        console.log('');
        console.log(`${primary('Profile:')}    ${profile}`);
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
      showError(error.message);
      process.exit(1);
    }
  });

// login command
const login = new Command('login')
  .description('Authenticate with your Telnyx API key')
  .option('-p, --profile <name>', 'Profile to authenticate', 'default')
  .action(async (options) => {
    const profileName = options.profile;
    const currentKey = getApiKey(profileName);
    
    if (currentKey) {
      const { overwrite } = await inquirer.prompt([
        {
          type: 'confirm',
          name: 'overwrite',
          message: `An API key is already configured for profile "${profileName}". Overwrite?`,
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
      setApiKey(apiKey, profileName);
      setActiveProfile(profileName);
      showSuccess(`Successfully authenticated with Telnyx (profile: ${profileName})`);
      
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

// logout command
const logout = new Command('logout')
  .description('Clear stored credentials for the active profile')
  .option('-p, --profile <name>', 'Profile to logout (defaults to active profile)')
  .option('-a, --all', 'Logout of all profiles')
  .action(async (options) => {
    if (options.all) {
      const { confirm } = await inquirer.prompt([
        {
          type: 'confirm',
          name: 'confirm',
          message: 'Clear credentials for ALL profiles?',
          default: false
        }
      ]);
      
      if (!confirm) {
        showInfo('Logout cancelled.');
        return;
      }
      
      const profiles = listProfiles();
      for (const profile of profiles) {
        clearApiKey(profile.name);
      }
      setActiveProfile('default');
      showSuccess('Logged out of all profiles');
      return;
    }
    
    const profileName = options.profile || getActiveProfile();
    
    if (!hasApiKey(profileName)) {
      showInfo(`No credentials stored for profile "${profileName}"`);
      return;
    }
    
    const { confirm } = await inquirer.prompt([
      {
        type: 'confirm',
        name: 'confirm',
        message: `Clear credentials for profile "${profileName}"?`,
        default: false
      }
    ]);
    
    if (!confirm) {
      showInfo('Logout cancelled.');
      return;
    }
    
    clearApiKey(profileName);
    showSuccess(`Logged out of profile "${profileName}"`);
    
    if (getActiveProfile() === profileName) {
      showInfo('You are no longer authenticated. Run "telnyx auth login" to re-authenticate.');
    }
  });

// profile list command
const profileList = new Command('list')
  .description('List all profiles')
  .action(() => {
    const profiles = listProfiles();
    const active = getActiveProfile();
    
    if (profiles.length === 0) {
      showInfo('No profiles configured.');
      console.log('');
      showInfo('Create a profile with: telnyx profile add <name>');
      return;
    }
    
    showSuccess('Profiles');
    console.log('');
    
    const table = new Table({
      head: ['Name', 'Active', 'Authenticated'],
      colWidths: [30, 10, 15],
      style: {
        head: [COLORS.primary.replace('#', '')],
        border: ['gray']
      }
    });
    
    profiles.forEach(profile => {
      table.push([
        profile.name,
        profile.active ? '✓' : '',
        profile.hasKey ? '✓' : '✗'
      ]);
    });
    
    console.log(table.toString());
    console.log('');
    showInfo(`Active profile: ${active}`);
  });

// profile add command
const profileAdd = new Command('add')
  .description('Add a new profile')
  .argument('<name>', 'Profile name')
  .action(async (name) => {
    if (profileExists(name)) {
      showError(`Profile "${name}" already exists`);
      process.exit(1);
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
      // Test the API key
      const axios = require('axios');
      await axios.get('https://api.telnyx.com/v2/whoami', {
        headers: {
          'Authorization': `Bearer ${apiKey}`,
          'Content-Type': 'application/json'
        },
        timeout: 10000
      });
      
      spinner.stop();
      setApiKey(apiKey, name);
      showSuccess(`Profile "${name}" created successfully`);
      
      const { switchNow } = await inquirer.prompt([
        {
          type: 'confirm',
          name: 'switchNow',
          message: 'Switch to this profile now?',
          default: true
        }
      ]);
      
      if (switchNow) {
        setActiveProfile(name);
        showSuccess(`Switched to profile "${name}"`);
      }
      
    } catch (error) {
      spinner.stop();
      if (error.response && error.response.status === 401) {
        showError('Invalid API key. Please check your key and try again.');
      } else {
        showError(`Validation failed: ${error.message}`);
      }
      process.exit(1);
    }
  });

// profile use command
const profileUse = new Command('use')
  .description('Switch to a different profile')
  .argument('<name>', 'Profile name to activate')
  .action((name) => {
    if (!profileExists(name)) {
      showError(`Profile "${name}" does not exist`);
      showInfo('Create it with: telnyx profile add ' + name);
      process.exit(1);
    }
    
    setActiveProfile(name);
    showSuccess(`Switched to profile "${name}"`);
    
    if (!hasApiKey(name)) {
      showInfo('This profile has no API key configured.');
      showInfo('Run: telnyx auth login --profile ' + name);
    }
  });

// profile remove command
const profileRemove = new Command('remove')
  .description('Remove a profile')
  .argument('<name>', 'Profile name to remove')
  .option('-y, --yes', 'Skip confirmation prompt')
  .action(async (name, options) => {
    if (!profileExists(name)) {
      showError(`Profile "${name}" does not exist`);
      process.exit(1);
    }
    
    if (!options.yes) {
      const { confirm } = await inquirer.prompt([
        {
          type: 'confirm',
          name: 'confirm',
          message: `Remove profile "${name}"?`,
          default: false
        }
      ]);
      
      if (!confirm) {
        showInfo('Removal cancelled.');
        return;
      }
    }
    
    deleteProfile(name);
    showSuccess(`Profile "${name}" removed`);
  });

// profile command (parent)
const profile = new Command('profile')
  .description('Manage authentication profiles')
  .addCommand(profileList)
  .addCommand(profileAdd)
  .addCommand(profileUse)
  .addCommand(profileRemove);

module.exports = {
  whoami,
  login,
  logout,
  profile
};
