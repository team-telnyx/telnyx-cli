const Conf = require('conf');

const config = new Conf({
  projectName: 'telnyx-cli',
  schema: {
    activeProfile: {
      type: 'string',
      default: 'default'
    },
    profiles: {
      type: 'object',
      default: {}
    }
  }
});

// Profile management
function getActiveProfile() {
  return config.get('activeProfile') || 'default';
}

function setActiveProfile(profileName) {
  config.set('activeProfile', profileName);
}

function getProfiles() {
  return config.get('profiles') || {};
}

function profileExists(profileName) {
  const profiles = getProfiles();
  return !!profiles[profileName];
}

function getProfile(profileName) {
  const profiles = getProfiles();
  return profiles[profileName] || null;
}

function setProfile(profileName, profileData) {
  const profiles = getProfiles();
  profiles[profileName] = profileData;
  config.set('profiles', profiles);
}

function deleteProfile(profileName) {
  const profiles = getProfiles();
  delete profiles[profileName];
  config.set('profiles', profiles);
  
  // If we deleted the active profile, switch to default
  if (getActiveProfile() === profileName) {
    setActiveProfile('default');
  }
}

function listProfiles() {
  const profiles = getProfiles();
  const active = getActiveProfile();
  
  return Object.keys(profiles).map(name => ({
    name,
    active: name === active,
    hasKey: !!profiles[name]?.apiKey
  }));
}

// API Key management (profile-aware)
function getApiKey(profileName) {
  const targetProfile = profileName || getActiveProfile();
  const profiles = getProfiles();
  return profiles[targetProfile]?.apiKey || '';
}

function setApiKey(apiKey, profileName) {
  const targetProfile = profileName || getActiveProfile();
  const profiles = getProfiles();
  
  if (!profiles[targetProfile]) {
    profiles[targetProfile] = {};
  }
  
  profiles[targetProfile].apiKey = apiKey;
  config.set('profiles', profiles);
}

function clearApiKey(profileName) {
  const targetProfile = profileName || getActiveProfile();
  const profiles = getProfiles();
  
  if (profiles[targetProfile]) {
    delete profiles[targetProfile].apiKey;
    config.set('profiles', profiles);
  }
}

function hasApiKey(profileName) {
  const targetProfile = profileName || getActiveProfile();
  const profiles = getProfiles();
  return !!profiles[targetProfile]?.apiKey;
}

// Backward compatibility - migrate old single-key storage
function migrateOldConfig() {
  const oldKey = config.get('apiKey');
  if (oldKey) {
    const profiles = getProfiles();
    if (!profiles.default) {
      profiles.default = { apiKey: oldKey };
      config.set('profiles', profiles);
      config.delete('apiKey');
    }
  }
}

// Run migration on load
migrateOldConfig();

module.exports = {
  // Profile management
  getActiveProfile,
  setActiveProfile,
  getProfiles,
  profileExists,
  getProfile,
  setProfile,
  deleteProfile,
  listProfiles,
  
  // API Key management
  getApiKey,
  setApiKey,
  clearApiKey,
  hasApiKey
};
