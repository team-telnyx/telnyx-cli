const { Command } = require('commander');
const inquirer = require('inquirer');
const ora = require('ora');
const Table = require('cli-table3');
const chalk = require('chalk');
const { 
  genericList, 
  genericGet, 
  genericCreate, 
  genericUpdate, 
  genericDelete 
} = require('../api/client');
const { showSuccess, showError, showInfo, COLORS } = require('../ui/layout');

const primary = chalk.hex(COLORS.primary);
const gray = chalk.gray;

// ==================== RESOURCE DEFINITIONS ====================
// Define common Telnyx resources for the "Long Tail"

const RESOURCES = {
  // Messaging
  'messaging-profiles': {
    name: 'messaging-profile',
    display: 'Messaging Profile',
    description: 'Manage messaging profiles',
    listable: true,
    hasId: true
  },
  'messaging-hosted-number-orders': {
    name: 'messaging-hosted-number-order',
    display: 'Hosted Number Order',
    description: 'Manage hosted number orders',
    listable: true,
    hasId: true
  },
  // Voice
  'connections': {
    name: 'connection',
    display: 'Connection',
    description: 'Manage SIP connections',
    listable: true,
    hasId: true
  },
  'credential-connections': {
    name: 'credential-connection',
    display: 'Credential Connection',
    description: 'Manage credential connections',
    listable: true,
    hasId: true
  },
  'texml-applications': {
    name: 'texml-application',
    display: 'TeXML Application',
    description: 'Manage TeXML applications',
    listable: true,
    hasId: true
  },
  'call-control-applications': {
    name: 'call-control-application',
    display: 'Call Control Application',
    description: 'Manage call control applications',
    listable: true,
    hasId: true
  },
  'sip-trunks': {
    name: 'sip-trunk',
    display: 'SIP Trunk',
    description: 'Manage SIP trunks',
    listable: true,
    hasId: true
  },
  // Storage
  'buckets': {
    name: 'bucket',
    display: 'Storage Bucket',
    description: 'Manage storage buckets',
    listable: true,
    hasId: true
  },
  'storage-files': {
    name: 'storage-file',
    display: 'Storage File',
    description: 'Manage storage files',
    listable: true,
    hasId: true
  },
  // Networking
  'networks': {
    name: 'network',
    display: 'Network',
    description: 'Manage networks',
    listable: true,
    hasId: true
  },
  'network-bridge': {
    name: 'network-bridge',
    display: 'Network Bridge',
    description: 'Manage network bridges',
    listable: true,
    hasId: true
  },
  // Portability
  'portability-checks': {
    name: 'portability-check',
    display: 'Portability Check',
    description: 'Check number portability',
    listable: true,
    hasId: true
  },
  'porting-orders': {
    name: 'porting-order',
    display: 'Porting Order',
    description: 'Manage porting orders',
    listable: true,
    hasId: true
  },
  // Conferences
  'conferences': {
    name: 'conference',
    display: 'Conference',
    description: 'Manage conferences',
    listable: true,
    hasId: true
  },
  // AI/Assistants
  'assistants': {
    name: 'assistant',
    display: 'AI Assistant',
    description: 'Manage AI assistants',
    listable: true,
    hasId: true
  },
  // Verification
  'verify-profiles': {
    name: 'verify-profile',
    display: 'Verify Profile',
    description: 'Manage verification profiles',
    listable: true,
    hasId: true
  },
  // IoT
  'iot-packages': {
    name: 'iot-package',
    display: 'IoT Package',
    description: 'Manage IoT packages',
    listable: true,
    hasId: true
  },
  'iot-coverage': {
    name: 'iot-coverage',
    display: 'IoT Coverage',
    description: 'View IoT coverage',
    listable: true,
    hasId: false
  }
};

// ==================== GENERIC LIST COMMAND ====================

function createListCommand(resourceKey, resourceDef) {
  return new Command('list')
    .description(`List ${resourceDef.display}s`)
    .alias('ls')
    .option('-n, --limit <number>', 'Number of results', '20')
    .option('-s, --status <status>', 'Filter by status')
    .option('-t, --type <type>', 'Filter by type')
    .option('-j, --json', 'Output raw JSON')
    .action(async (options) => {
      const { limit, status, type, json } = options;
      
      const spinner = ora({
        text: `Fetching ${resourceDef.display.toLowerCase()}s...`,
        spinner: 'dots'
      }).start();
      
      try {
        const data = await genericList(resourceKey, {
          limit: parseInt(limit),
          status,
          type
        });
        
        spinner.stop();
        
        if (json) {
          console.log(JSON.stringify(data, null, 2));
          return;
        }
        
        if (!data.data || data.data.length === 0) {
          showInfo(`📭 No ${resourceDef.display.toLowerCase()}s found.`);
          return;
        }
        
        showSuccess(`Found ${data.data.length} ${resourceDef.display.toLowerCase()}(s)`);
        console.log('');
        
        // Dynamic table generation based on available fields
        const table = generateTable(data.data, resourceKey);
        console.log(table.toString());
        console.log('');
        
        if (resourceDef.hasId) {
          showInfo(`💡 Use "telnyx ${resourceDef.name} get <id>" for details`);
        }
        
      } catch (error) {
        spinner.stop();
        handleApiError(error);
      }
    });
}

// ==================== GENERIC GET COMMAND ====================

function createGetCommand(resourceKey, resourceDef) {
  return new Command('get')
    .description(`Get a specific ${resourceDef.display}`)
    .alias('show')
    .argument('<id>', `${resourceDef.display} ID`)
    .option('-j, --json', 'Output raw JSON')
    .action(async (id, options) => {
      const { json } = options;
      
      const spinner = ora({
        text: `Fetching ${resourceDef.display.toLowerCase()}...`,
        spinner: 'dots'
      }).start();
      
      try {
        const data = await genericGet(resourceKey, id);
        
        spinner.stop();
        
        if (json) {
          console.log(JSON.stringify(data, null, 2));
          return;
        }
        
        if (!data.data) {
          showError(`${resourceDef.display} not found`);
          return;
        }
        
        showSuccess(`${resourceDef.display} Details`);
        console.log('');
        
        // Pretty print the object
        printObject(data.data, resourceDef.display);
        console.log('');
        
      } catch (error) {
        spinner.stop();
        handleApiError(error);
      }
    });
}

// ==================== GENERIC CREATE COMMAND ====================

function createCreateCommand(resourceKey, resourceDef) {
  return new Command('create')
    .description(`Create a new ${resourceDef.display}`)
    .alias('new')
    .option('-d, --data <json>', 'JSON data for the resource')
    .option('-f, --file <path>', 'JSON file containing resource data')
    .option('-j, --json', 'Output raw JSON')
    .action(async (options) => {
      const { data: jsonData, file, json } = options;
      
      let resourceData = {};
      
      // Get data from file or command line
      if (file) {
        const fs = require('fs');
        try {
          const fileContent = fs.readFileSync(file, 'utf8');
          resourceData = JSON.parse(fileContent);
        } catch (err) {
          showError(`Failed to read file: ${err.message}`);
          process.exit(1);
        }
      } else if (jsonData) {
        try {
          resourceData = JSON.parse(jsonData);
        } catch (err) {
          showError(`Invalid JSON: ${err.message}`);
          process.exit(1);
        }
      } else {
        showError(`Please provide resource data via --data or --file`);
        process.exit(1);
      }
      
      const spinner = ora({
        text: `Creating ${resourceDef.display.toLowerCase()}...`,
        spinner: 'dots'
      }).start();
      
      try {
        const data = await genericCreate(resourceKey, resourceData);
        
        spinner.stop();
        
        if (json) {
          console.log(JSON.stringify(data, null, 2));
          return;
        }
        
        showSuccess(`✅ ${resourceDef.display} created successfully!`);
        console.log('');
        
        if (data.data && data.data.id) {
          console.log(`${primary('ID:')}  ${data.data.id}`);
        }
        
        console.log('');
        
      } catch (error) {
        spinner.stop();
        handleApiError(error);
      }
    });
}

// ==================== GENERIC UPDATE COMMAND ====================

function createUpdateCommand(resourceKey, resourceDef) {
  return new Command('update')
    .description(`Update a ${resourceDef.display}`)
    .alias('edit')
    .argument('<id>', `${resourceDef.display} ID`)
    .option('-d, --data <json>', 'JSON data with updates')
    .option('-f, --file <path>', 'JSON file containing updates')
    .option('-j, --json', 'Output raw JSON')
    .action(async (id, options) => {
      const { data: jsonData, file, json } = options;
      
      let updates = {};
      
      // Get data from file or command line
      if (file) {
        const fs = require('fs');
        try {
          const fileContent = fs.readFileSync(file, 'utf8');
          updates = JSON.parse(fileContent);
        } catch (err) {
          showError(`Failed to read file: ${err.message}`);
          process.exit(1);
        }
      } else if (jsonData) {
        try {
          updates = JSON.parse(jsonData);
        } catch (err) {
          showError(`Invalid JSON: ${err.message}`);
          process.exit(1);
        }
      } else {
        showError(`Please provide update data via --data or --file`);
        process.exit(1);
      }
      
      const spinner = ora({
        text: `Updating ${resourceDef.display.toLowerCase()}...`,
        spinner: 'dots'
      }).start();
      
      try {
        const data = await genericUpdate(resourceKey, id, updates);
        
        spinner.stop();
        
        if (json) {
          console.log(JSON.stringify(data, null, 2));
          return;
        }
        
        showSuccess(`✅ ${resourceDef.display} updated successfully!`);
        console.log('');
        
      } catch (error) {
        spinner.stop();
        handleApiError(error);
      }
    });
}

// ==================== GENERIC DELETE COMMAND ====================

function createDeleteCommand(resourceKey, resourceDef) {
  return new Command('delete')
    .description(`Delete a ${resourceDef.display}`)
    .alias('remove')
    .argument('<id>', `${resourceDef.display} ID`)
    .option('-y, --yes', 'Skip confirmation prompt')
    .option('-j, --json', 'Output raw JSON')
    .action(async (id, options) => {
      const { yes, json } = options;
      
      // Confirm before deleting
      if (!yes) {
        const { confirm } = await inquirer.prompt([
          {
            type: 'confirm',
            name: 'confirm',
            message: `Delete ${resourceDef.display.toLowerCase()} ${id}?`,
            default: false
          }
        ]);
        
        if (!confirm) {
          showInfo('Operation cancelled.');
          return;
        }
      }
      
      const spinner = ora({
        text: `Deleting ${resourceDef.display.toLowerCase()}...`,
        spinner: 'dots'
      }).start();
      
      try {
        const data = await genericDelete(resourceKey, id);
        
        spinner.stop();
        
        if (json) {
          console.log(JSON.stringify(data, null, 2));
          return;
        }
        
        showSuccess(`✅ ${resourceDef.display} deleted successfully`);
        
      } catch (error) {
        spinner.stop();
        handleApiError(error);
      }
    });
}

// ==================== RESOURCE COMMAND BUILDER ====================

function createResourceCommand(resourceKey, resourceDef) {
  const command = new Command(resourceDef.name)
    .description(resourceDef.description);
  
  // Add list command if resource is listable
  if (resourceDef.listable) {
    command.addCommand(createListCommand(resourceKey, resourceDef));
  }
  
  // Add get command if resource has IDs
  if (resourceDef.hasId) {
    command.addCommand(createGetCommand(resourceKey, resourceDef));
    command.addCommand(createCreateCommand(resourceKey, resourceDef));
    command.addCommand(createUpdateCommand(resourceKey, resourceDef));
    command.addCommand(createDeleteCommand(resourceKey, resourceDef));
  }
  
  return command;
}

// ==================== RESOURCES COMMAND (LIST ALL) ====================

const resources = new Command('resources')
  .description('List all available API resources')
  .alias('api')
  .action(() => {
    showSuccess('Available Telnyx API Resources');
    console.log('');
    
    const table = new Table({
      head: ['Command', 'Description', 'Operations'],
      colWidths: [30, 35, 25],
      style: {
        head: [COLORS.primary.replace('#', '')],
        border: ['gray']
      }
    });
    
    Object.entries(RESOURCES).forEach(([key, def]) => {
      const ops = [];
      if (def.listable) ops.push('list');
      if (def.hasId) ops.push('get', 'create', 'update', 'delete');
      
      table.push([
        `telnyx ${def.name}`,
        def.description,
        ops.join(', ')
      ]);
    });
    
    console.log(table.toString());
    console.log('');
    showInfo('💡 Use "telnyx <resource> --help" for command options');
  });

// ==================== HELPERS ====================

function generateTable(items, resourceKey) {
  // Extract common fields
  const firstItem = items[0].data || items[0];
  const fields = Object.keys(firstItem).slice(0, 5); // Limit to 5 columns
  
  const headers = fields.map(f => f.replace(/_/g, ' ').replace(/\b\w/g, l => l.toUpperCase()));
  const colWidths = fields.map(() => 20);
  
  const table = new Table({
    head: headers,
    colWidths: colWidths,
    style: {
      head: [COLORS.primary.replace('#', '')],
      border: ['gray']
    }
  });
  
  items.forEach((item) => {
    const data = item.data || item;
    const row = fields.map(field => {
      const value = data[field];
      if (value === null || value === undefined) return '-';
      if (typeof value === 'object') return JSON.stringify(value).substring(0, 18);
      return String(value).substring(0, 18);
    });
    table.push(row);
  });
  
  return table;
}

function printObject(obj, title) {
  console.log('┌─────────────────────────────────────────────────────────┐');
  console.log(`│  📋 ${primary(title.toUpperCase())}` + ' '.repeat(52 - title.length) + '│');
  console.log('├─────────────────────────────────────────────────────────┤');
  
  Object.entries(obj).forEach(([key, value]) => {
    let displayValue = value;
    
    if (value === null || value === undefined) {
      displayValue = 'null';
    } else if (typeof value === 'object') {
      displayValue = JSON.stringify(value).substring(0, 40);
    } else {
      displayValue = String(value).substring(0, 40);
    }
    
    const keyStr = key.replace(/_/g, ' ').replace(/\b\w/g, l => l.toUpperCase());
    console.log(`│  ${keyStr.padEnd(15)}: ${displayValue.padEnd(35)}│`);
  });
  
  console.log('└─────────────────────────────────────────────────────────┘');
}

function handleApiError(error) {
  if (error.response?.status === 401) {
    showError('🔐 Authentication failed. Run: telnyx auth login');
  } else if (error.response?.status === 403) {
    showError('🚫 Permission denied. This feature may not be enabled on your account.');
  } else if (error.response?.status === 404) {
    showError('❌ Resource not found.');
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

// Export all resource commands
const resourceCommands = {};
Object.entries(RESOURCES).forEach(([key, def]) => {
  resourceCommands[def.name.replace(/-/g, '_')] = createResourceCommand(key, def);
});

module.exports = {
  resources,
  ...resourceCommands
};
