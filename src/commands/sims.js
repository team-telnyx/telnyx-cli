const { Command } = require('commander');
const inquirer = require('inquirer');
const ora = require('ora');
const Table = require('cli-table3');
const chalk = require('chalk');
const { listSims, getSim, updateSim } = require('../api/client');
const { showSuccess, showError, showInfo, COLORS } = require('../ui/layout');

const primary = chalk.hex(COLORS.primary);
const gray = chalk.gray;
const yellow = chalk.yellow;
const green = chalk.green;
const red = chalk.red;

// ==================== SIM LIST ====================

const list = new Command('list')
  .description('List your IoT SIM cards')
  .alias('ls')
  .option('-s, --status <status>', 'Filter by status: active, inactive, suspended, pending_activation')
  .option('-n, --limit <number>', 'Number of results', '20')
  .option('--active', 'Show only active SIMs')
  .action(async (options) => {
    const { status, limit, active } = options;
    
    let filterStatus = status;
    if (active) {
      filterStatus = 'active';
    }
    
    const spinner = ora({
      text: 'Fetching SIM cards...',
      spinner: 'dots'
    }).start();
    
    try {
      const data = await listSims({
        status: filterStatus,
        limit: parseInt(limit)
      });
      
      spinner.stop();
      
      if (!data.data || data.data.length === 0) {
        showInfo('📭 No SIM cards found.');
        showInfo('   Visit https://portal.telnyx.com to order IoT SIMs');
        return;
      }
      
      showSuccess(`Found ${data.data.length} SIM card(s)`);
      console.log('');
      
      const table = new Table({
        head: ['SIM ID', 'ICCID', 'Status', 'Network', 'Data Usage', 'Last Updated'],
        colWidths: [24, 22, 12, 12, 12, 20],
        style: {
          head: [COLORS.primary.replace('#', '')],
          border: ['gray']
        }
      });
      
      data.data.forEach((item) => {
        const sim = item.data || item;
        const lastUpdated = sim.updated_at 
          ? new Date(sim.updated_at).toLocaleDateString() 
          : 'N/A';
        
        const dataUsage = sim.data_usage 
          ? formatBytes(sim.data_usage.bytes_used || 0)
          : '-';
        
        table.push([
          (sim.id || 'N/A').substring(0, 22),
          (sim.iccid || 'N/A').substring(0, 20),
          getStatusIcon(sim.status) + ' ' + (sim.status || 'unknown'),
          sim.network?.name || sim.network_id || '-',
          dataUsage,
          lastUpdated
        ]);
      });
      
      console.log(table.toString());
      console.log('');
      showInfo('💡 Use "telnyx sim show <sim_id>" for detailed information');
      
    } catch (error) {
      spinner.stop();
      handleApiError(error);
    }
  });

// ==================== SIM SHOW ====================

const show = new Command('show')
  .description('Show detailed information about a SIM card')
  .alias('info')
  .argument('[simId]', 'SIM card ID or ICCID')
  .option('-j, --json', 'Output raw JSON')
  .action(async (simId, options) => {
    let targetSimId = simId;
    
    // If no SIM ID provided, prompt for one
    if (!targetSimId) {
      const { id } = await inquirer.prompt([
        {
          type: 'input',
          name: 'id',
          message: 'SIM ID or ICCID:',
          validate: (input) => input.length > 0 || 'SIM ID is required'
        }
      ]);
      targetSimId = id;
    }
    
    const spinner = ora({
      text: 'Fetching SIM details...',
      spinner: 'dots'
    }).start();
    
    try {
      const data = await getSim(targetSimId);
      
      spinner.stop();
      
      if (!data.data) {
        showError('SIM card not found');
        return;
      }
      
      const sim = data.data;
      
      if (options.json) {
        console.log(JSON.stringify(data, null, 2));
        return;
      }
      
      showSuccess(`SIM Details: ${sim.iccid || targetSimId}`);
      console.log('');
      
      // Main info box
      console.log('┌─────────────────────────────────────────────────────────┐');
      console.log('│  📱 ' + primary('SIM CARD INFORMATION') + '                                │');
      console.log('├─────────────────────────────────────────────────────────┤');
      console.log(`│  SIM ID:          ${(sim.id || 'N/A').padEnd(40)}│`);
      console.log(`│  ICCID:           ${(sim.iccid || 'N/A').padEnd(40)}│`);
      console.log(`│  Status:          ${(sim.status || 'N/A').padEnd(40)}│`);
      if (sim.sim_type) {
        console.log(`│  SIM Type:        ${(sim.sim_type).padEnd(40)}│`);
      }
      console.log('└─────────────────────────────────────────────────────────┘');
      
      // Network section
      if (sim.network || sim.network_id) {
        console.log('');
        console.log('┌─────────────────────────────────────────────────────────┐');
        console.log('│  📡 ' + primary('NETWORK INFORMATION') + '                                 │');
        console.log('├─────────────────────────────────────────────────────────┤');
        console.log(`│  Network ID:      ${(sim.network_id || 'N/A').padEnd(40)}│`);
        if (sim.network?.name) {
          console.log(`│  Network Name:    ${(sim.network.name).padEnd(40)}│`);
        }
        if (sim.msisdn) {
          console.log(`│  MSISDN:          ${(sim.msisdn).padEnd(40)}│`);
        }
        if (sim.imsi) {
          console.log(`│  IMSI:            ${(sim.imsi).padEnd(40)}│`);
        }
        console.log('└─────────────────────────────────────────────────────────┘');
      }
      
      // Data usage section
      if (sim.data_usage) {
        console.log('');
        console.log('┌─────────────────────────────────────────────────────────┐');
        console.log('│  📊 ' + primary('DATA USAGE') + '                                            │');
        console.log('├─────────────────────────────────────────────────────────┤');
        console.log(`│  Data Used:       ${formatBytes(sim.data_usage.bytes_used || 0).padEnd(40)}│`);
        if (sim.data_usage.bytes_remaining !== undefined) {
          console.log(`│  Data Remaining:  ${formatBytes(sim.data_usage.bytes_remaining).padEnd(40)}│`);
        }
        if (sim.data_usage.last_reported) {
          console.log(`│  Last Reported:   ${new Date(sim.data_usage.last_reported).toLocaleString().padEnd(40)}│`);
        }
        console.log('└─────────────────────────────────────────────────────────┘');
      }
      
      // Tags section
      if (sim.tags && sim.tags.length > 0) {
        console.log('');
        console.log('┌─────────────────────────────────────────────────────────┐');
        console.log('│  🏷️  ' + primary('TAGS') + '                                                 │');
        console.log('├─────────────────────────────────────────────────────────┤');
        sim.tags.forEach(tag => {
          console.log(`│  • ${tag.padEnd(51)}│`);
        });
        console.log('└─────────────────────────────────────────────────────────┘');
      }
      
      // Timestamps
      console.log('');
      console.log('┌─────────────────────────────────────────────────────────┐');
      console.log('│  📅 ' + primary('TIMESTAMPS') + '                                          │');
      console.log('├─────────────────────────────────────────────────────────┤');
      if (sim.created_at) {
        console.log(`│  Created:         ${new Date(sim.created_at).toLocaleString().padEnd(40)}│`);
      }
      if (sim.updated_at) {
        console.log(`│  Updated:         ${new Date(sim.updated_at).toLocaleString().padEnd(40)}│`);
      }
      if (sim.activated_at) {
        console.log(`│  Activated:       ${new Date(sim.activated_at).toLocaleString().padEnd(40)}│`);
      }
      console.log('└─────────────────────────────────────────────────────────┘');
      
      console.log('');
      
    } catch (error) {
      spinner.stop();
      handleApiError(error);
    }
  });

// ==================== SIM UPDATE ====================

const update = new Command('update')
  .description('Update SIM card settings (tags, status)')
  .alias('set')
  .argument('[simId]', 'SIM card ID')
  .option('-t, --tags <tags>', 'Comma-separated tags')
  .option('-s, --status <status>', 'Update status (if supported)')
  .option('-y, --yes', 'Skip confirmation prompt')
  .action(async (simId, options) => {
    let targetSimId = simId;
    const { tags, status, yes } = options;
    
    // If no SIM ID provided, prompt for one
    if (!targetSimId) {
      const { id } = await inquirer.prompt([
        {
          type: 'input',
          name: 'id',
          message: 'SIM ID:',
          validate: (input) => input.length > 0 || 'SIM ID is required'
        }
      ]);
      targetSimId = id;
    }
    
    // If no updates specified, prompt for them
    let updates = {};
    
    if (!tags && !status) {
      const { updateType } = await inquirer.prompt([
        {
          type: 'list',
          name: 'updateType',
          message: 'What would you like to update?',
          choices: [
            { name: 'Add/Update Tags', value: 'tags' },
            { name: 'Change Status', value: 'status' }
          ]
        }
      ]);
      
      if (updateType === 'tags') {
        const { newTags } = await inquirer.prompt([
          {
            type: 'input',
            name: 'newTags',
            message: 'Enter tags (comma-separated):',
            validate: (input) => input.length > 0 || 'At least one tag is required'
          }
        ]);
        updates.tags = newTags.split(',').map(t => t.trim());
      } else if (updateType === 'status') {
        const { newStatus } = await inquirer.prompt([
          {
            type: 'list',
            name: 'newStatus',
            message: 'Select new status:',
            choices: [
              { name: 'Active', value: 'active' },
              { name: 'Inactive', value: 'inactive' },
              { name: 'Suspended', value: 'suspended' }
            ]
          }
        ]);
        updates.status = newStatus;
      }
    } else {
      if (tags) {
        updates.tags = tags.split(',').map(t => t.trim());
      }
      if (status) {
        updates.status = status;
      }
    }
    
    // Confirm before updating
    if (!yes) {
      console.log('');
      console.log('┌─────────────────────────────────────────────────────────┐');
      console.log('│  📝 ' + primary('UPDATE SUMMARY') + '                                      │');
      console.log('├─────────────────────────────────────────────────────────┤');
      console.log(`│  SIM ID:  ${targetSimId.padEnd(45)}│`);
      if (updates.tags) {
        console.log(`│  Tags:    ${updates.tags.join(', ').padEnd(45)}│`);
      }
      if (updates.status) {
        console.log(`│  Status:  ${updates.status.padEnd(45)}│`);
      }
      console.log('└─────────────────────────────────────────────────────────┘');
      console.log('');
      
      const { confirm } = await inquirer.prompt([
        {
          type: 'confirm',
          name: 'confirm',
          message: 'Apply these changes?',
          default: true
        }
      ]);
      
      if (!confirm) {
        showInfo('Update cancelled.');
        return;
      }
    }
    
    const spinner = ora({
      text: 'Updating SIM card...',
      spinner: 'dots'
    }).start();
    
    try {
      await updateSim(targetSimId, updates);
      
      spinner.stop();
      showSuccess(`✅ SIM ${targetSimId} updated successfully`);
      
    } catch (error) {
      spinner.stop();
      handleApiError(error);
    }
  });

// ==================== HELPERS ====================

function getStatusIcon(status) {
  const icons = {
    'active': '🟢',
    'inactive': '⚪',
    'suspended': '🟡',
    'pending_activation': '⏳',
    'pending': '⏳'
  };
  return icons[status?.toLowerCase()] || '•';
}

function formatBytes(bytes) {
  if (bytes === 0) return '0 B';
  
  const k = 1024;
  const sizes = ['B', 'KB', 'MB', 'GB', 'TB'];
  const i = Math.floor(Math.log(bytes) / Math.log(k));
  
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i];
}

function handleApiError(error) {
  if (error.response?.status === 401) {
    showError('🔐 Authentication failed. Run: telnyx auth login');
  } else if (error.response?.status === 403) {
    showError('🚫 Permission denied. IoT SIM access may not be enabled on your account.');
    showInfo('   Contact Telnyx support to enable IoT services.');
  } else if (error.response?.status === 404) {
    showError('❌ SIM card not found. Check the ID and try again.');
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
  list,
  show,
  update
};
