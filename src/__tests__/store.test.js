/**
 * Tests for src/config/store.js
 * Profile and API key management
 */

// Track mock store state that we can reset
let mockStore = {
  activeProfile: 'default',
  profiles: {}
};

// Mock conf before requiring store
jest.mock('conf', () => {
  return jest.fn().mockImplementation(() => {
    return {
      get: jest.fn((key) => mockStore[key]),
      set: jest.fn((key, value) => { mockStore[key] = value; }),
      delete: jest.fn((key) => { delete mockStore[key]; })
    };
  });
});

// Helper to reset store between tests
function resetMockStore() {
  mockStore = {
    activeProfile: 'default',
    profiles: {}
  };
}

// Re-require the module to get fresh functions that reference our mock
let store;
beforeEach(() => {
  resetMockStore();
  jest.resetModules();
  store = require('../config/store');
});

describe('Profile Management', () => {
  describe('getActiveProfile', () => {
    it('should return "default" when no profile is set', () => {
      const profile = store.getActiveProfile();
      expect(profile).toBe('default');
    });

    it('should return the active profile name', () => {
      mockStore.activeProfile = 'production';
      expect(store.getActiveProfile()).toBe('production');
    });
  });

  describe('setActiveProfile', () => {
    it('should set the active profile', () => {
      store.setActiveProfile('production');
      expect(mockStore.activeProfile).toBe('production');
    });
  });

  describe('profileExists', () => {
    it('should return false for non-existent profile', () => {
      expect(store.profileExists('nonexistent')).toBe(false);
    });

    it('should return true for existing profile', () => {
      mockStore.profiles = { 'test-profile': { apiKey: 'test-key' } };
      expect(store.profileExists('test-profile')).toBe(true);
    });
  });

  describe('setProfile / getProfile', () => {
    it('should store and retrieve profile data', () => {
      const profileData = { apiKey: 'KEY123', customField: 'value' };
      store.setProfile('myprofile', profileData);
      
      expect(mockStore.profiles['myprofile']).toEqual(profileData);
    });

    it('should return null for non-existent profile', () => {
      expect(store.getProfile('ghost')).toBeNull();
    });

    it('should retrieve existing profile', () => {
      mockStore.profiles = { existing: { apiKey: 'secret' } };
      expect(store.getProfile('existing')).toEqual({ apiKey: 'secret' });
    });
  });

  describe('deleteProfile', () => {
    it('should remove a profile', () => {
      mockStore.profiles = { 'to-delete': { apiKey: 'key' } };
      
      store.deleteProfile('to-delete');
      
      expect(mockStore.profiles['to-delete']).toBeUndefined();
    });

    it('should switch to default if active profile is deleted', () => {
      mockStore.profiles = { temp: { apiKey: 'key' } };
      mockStore.activeProfile = 'temp';
      
      store.deleteProfile('temp');
      
      expect(mockStore.activeProfile).toBe('default');
    });

    it('should not change active profile if different profile is deleted', () => {
      mockStore.profiles = { 
        keep: { apiKey: 'key1' },
        remove: { apiKey: 'key2' }
      };
      mockStore.activeProfile = 'keep';
      
      store.deleteProfile('remove');
      
      expect(mockStore.activeProfile).toBe('keep');
    });
  });

  describe('listProfiles', () => {
    it('should return empty array when no profiles exist', () => {
      expect(store.listProfiles()).toEqual([]);
    });

    it('should list all profiles with active status', () => {
      mockStore.profiles = {
        profile1: { apiKey: 'key1' },
        profile2: { apiKey: 'key2' }
      };
      mockStore.activeProfile = 'profile1';
      
      const profiles = store.listProfiles();
      
      expect(profiles).toHaveLength(2);
      expect(profiles.find(p => p.name === 'profile1')).toEqual({
        name: 'profile1',
        active: true,
        hasKey: true
      });
      expect(profiles.find(p => p.name === 'profile2')).toEqual({
        name: 'profile2',
        active: false,
        hasKey: true
      });
    });

    it('should correctly report hasKey status', () => {
      mockStore.profiles = {
        'with-key': { apiKey: 'secret' },
        'no-key': {}
      };
      
      const profiles = store.listProfiles();
      
      expect(profiles.find(p => p.name === 'with-key').hasKey).toBe(true);
      expect(profiles.find(p => p.name === 'no-key').hasKey).toBe(false);
    });
  });
});

describe('API Key Management', () => {
  describe('setApiKey / getApiKey', () => {
    it('should store API key for active profile by default', () => {
      mockStore.activeProfile = 'default';
      mockStore.profiles = {};
      
      store.setApiKey('my-api-key');
      
      expect(mockStore.profiles['default'].apiKey).toBe('my-api-key');
    });

    it('should store API key for specific profile', () => {
      mockStore.profiles = {};
      
      store.setApiKey('prod-key', 'production');
      store.setApiKey('staging-key', 'staging');
      
      expect(mockStore.profiles['production'].apiKey).toBe('prod-key');
      expect(mockStore.profiles['staging'].apiKey).toBe('staging-key');
    });

    it('should return empty string for profile without key', () => {
      expect(store.getApiKey('nonexistent')).toBe('');
    });

    it('should retrieve stored API key', () => {
      mockStore.profiles = { test: { apiKey: 'test-key-123' } };
      expect(store.getApiKey('test')).toBe('test-key-123');
    });
  });

  describe('hasApiKey', () => {
    it('should return true when profile has API key', () => {
      mockStore.profiles = { test: { apiKey: 'key' } };
      expect(store.hasApiKey('test')).toBe(true);
    });

    it('should return false when profile has no API key', () => {
      mockStore.profiles = { empty: {} };
      expect(store.hasApiKey('empty')).toBe(false);
    });

    it('should return false for non-existent profile', () => {
      expect(store.hasApiKey('ghost')).toBe(false);
    });

    it('should check active profile by default', () => {
      mockStore.activeProfile = 'default';
      mockStore.profiles = { default: { apiKey: 'default-key' } };
      
      expect(store.hasApiKey()).toBe(true);
    });
  });

  describe('clearApiKey', () => {
    it('should remove API key from profile', () => {
      mockStore.profiles = { myprofile: { apiKey: 'to-clear', otherData: 'keep' } };
      
      store.clearApiKey('myprofile');
      
      expect(mockStore.profiles['myprofile'].apiKey).toBeUndefined();
      expect(mockStore.profiles['myprofile'].otherData).toBe('keep');
    });

    it('should not throw if profile does not exist', () => {
      expect(() => store.clearApiKey('ghost')).not.toThrow();
    });
  });
});
