/**
 * Tests for src/api/client.js
 * API client and error normalization
 */

// Mock dependencies
jest.mock('axios');
jest.mock('../config/store', () => ({
  getApiKey: jest.fn(() => 'test-api-key')
}));

const axios = require('axios');
const { getApiKey } = require('../config/store');
const { 
  normalizeError, 
  createClient, 
  setVerbose, 
  setTimeout,
  get,
  post
} = require('../api/client');

describe('API Client', () => {
  beforeEach(() => {
    jest.clearAllMocks();
    setVerbose(false);
    setTimeout(30000);
  });

  describe('createClient', () => {
    it('should throw error when no API key is configured', () => {
      getApiKey.mockReturnValueOnce('');
      
      expect(() => createClient()).toThrow('No API key configured. Run: telnyx auth login');
    });

    it('should create axios client with correct headers', () => {
      const mockCreate = jest.fn(() => ({}));
      axios.create = mockCreate;
      
      createClient();
      
      expect(mockCreate).toHaveBeenCalledWith({
        baseURL: 'https://api.telnyx.com/v2',
        headers: {
          'Authorization': 'Bearer test-api-key',
          'Content-Type': 'application/json',
          'Accept': 'application/json'
        },
        timeout: 30000
      });
    });

    it('should use custom timeout when set', () => {
      const mockCreate = jest.fn(() => ({}));
      axios.create = mockCreate;
      
      setTimeout(60000);
      createClient();
      
      expect(mockCreate).toHaveBeenCalledWith(
        expect.objectContaining({ timeout: 60000 })
      );
    });
  });

  describe('normalizeError', () => {
    describe('Network errors', () => {
      it('should handle timeout error', () => {
        const error = { code: 'ECONNABORTED', message: 'timeout' };
        const result = normalizeError(error);
        
        expect(result.message).toContain('Request timed out');
        expect(result.message).toContain('--timeout');
      });

      it('should handle connection refused', () => {
        const error = { code: 'ECONNREFUSED', message: 'connection refused' };
        const result = normalizeError(error);
        
        expect(result.message).toBe('Connection failed. Check your internet connection.');
      });

      it('should handle DNS resolution failure', () => {
        const error = { code: 'ENOTFOUND', message: 'getaddrinfo ENOTFOUND' };
        const result = normalizeError(error);
        
        expect(result.message).toBe('Connection failed. Check your internet connection.');
      });

      it('should handle generic network error', () => {
        const error = { message: 'Network Error' };
        const result = normalizeError(error);
        
        expect(result.message).toBe('Network error: Network Error');
      });
    });

    describe('HTTP errors', () => {
      it('should handle 401 Unauthorized', () => {
        const error = {
          response: { status: 401, data: {} }
        };
        const result = normalizeError(error);
        
        expect(result.message).toContain('Authentication failed');
        expect(result.message).toContain('telnyx auth login');
      });

      it('should handle 403 Forbidden', () => {
        const error = {
          response: { status: 403, data: {} }
        };
        const result = normalizeError(error);
        
        expect(result.message).toContain('Permission denied');
      });

      it('should handle 404 Not Found with resource extraction', () => {
        const error = {
          response: { status: 404, data: {} },
          config: { url: '/phone_numbers/123' }
        };
        const result = normalizeError(error);
        
        expect(result.message).toBe('Resource not found: phone numbers');
      });

      it('should handle 402 Payment Required', () => {
        const error = {
          response: { status: 402, data: {} }
        };
        const result = normalizeError(error);
        
        expect(result.message).toContain('Insufficient balance');
      });

      it('should handle 422 with error details', () => {
        const error = {
          response: {
            status: 422,
            data: {
              errors: [
                { title: 'Invalid phone number', detail: 'Number must be E.164 format' },
                { title: 'Missing field', detail: 'from is required' }
              ]
            }
          }
        };
        const result = normalizeError(error);
        
        expect(result.message).toContain('Validation failed');
        expect(result.message).toContain('E.164 format');
        expect(result.message).toContain('from is required');
      });

      it('should handle 422 without error details', () => {
        const error = {
          response: { status: 422, data: {} }
        };
        const result = normalizeError(error);
        
        expect(result.message).toBe('Validation failed: Invalid request parameters');
      });

      it('should handle 429 Rate Limited', () => {
        const error = {
          response: { status: 429, data: {} }
        };
        const result = normalizeError(error);
        
        expect(result.message).toContain('Rate limit exceeded');
      });

      it('should handle 5xx Server errors', () => {
        [500, 502, 503, 504].forEach(status => {
          const error = {
            response: { status, data: {} }
          };
          const result = normalizeError(error);
          
          expect(result.message).toBe('Telnyx API error. Please try again later.');
        });
      });

      it('should handle unknown status with error details', () => {
        const error = {
          response: {
            status: 418,
            data: {
              errors: [{ title: 'I am a teapot', detail: 'Cannot brew coffee' }]
            }
          }
        };
        const result = normalizeError(error);
        
        expect(result.message).toBe('Cannot brew coffee');
      });

      it('should handle unknown status without error details', () => {
        const error = {
          response: { status: 418, data: {} },
          message: 'Request failed with status 418'
        };
        const result = normalizeError(error);
        
        expect(result.message).toBe('Request failed with status 418');
      });
    });
  });
});

describe('API Methods', () => {
  let mockAxiosInstance;

  beforeEach(() => {
    jest.clearAllMocks();
    setVerbose(false);
    
    mockAxiosInstance = {
      get: jest.fn(),
      post: jest.fn(),
      patch: jest.fn(),
      delete: jest.fn()
    };
    axios.create = jest.fn(() => mockAxiosInstance);
  });

  describe('get', () => {
    it('should make GET request and return data', async () => {
      const mockData = { data: { id: '123', name: 'test' } };
      mockAxiosInstance.get.mockResolvedValue({ status: 200, data: mockData });
      
      const result = await get('/phone_numbers');
      
      expect(mockAxiosInstance.get).toHaveBeenCalledWith('/phone_numbers', { params: {} });
      expect(result).toEqual(mockData);
    });

    it('should pass query parameters', async () => {
      mockAxiosInstance.get.mockResolvedValue({ status: 200, data: {} });
      
      await get('/phone_numbers', { 'filter[status]': 'active' });
      
      expect(mockAxiosInstance.get).toHaveBeenCalledWith(
        '/phone_numbers',
        { params: { 'filter[status]': 'active' } }
      );
    });

    it('should normalize errors on failure', async () => {
      mockAxiosInstance.get.mockRejectedValue({
        response: { status: 401, data: {} }
      });
      
      await expect(get('/whoami')).rejects.toThrow('Authentication failed');
    });
  });

  describe('post', () => {
    it('should make POST request and return data', async () => {
      const mockData = { data: { id: 'msg_123' } };
      mockAxiosInstance.post.mockResolvedValue({ status: 201, data: mockData });
      
      const result = await post('/messages', { to: '+1234567890', text: 'Hello' });
      
      expect(mockAxiosInstance.post).toHaveBeenCalledWith(
        '/messages',
        { to: '+1234567890', text: 'Hello' }
      );
      expect(result).toEqual(mockData);
    });

    it('should normalize errors on failure', async () => {
      mockAxiosInstance.post.mockRejectedValue({
        response: {
          status: 422,
          data: {
            errors: [{ title: 'Invalid', detail: 'to is required' }]
          }
        }
      });
      
      await expect(post('/messages', {})).rejects.toThrow('to is required');
    });
  });
});

describe('Verbose Mode', () => {
  let consoleErrorSpy;

  beforeEach(() => {
    consoleErrorSpy = jest.spyOn(console, 'error').mockImplementation(() => {});
  });

  afterEach(() => {
    consoleErrorSpy.mockRestore();
  });

  it('should log requests when verbose is enabled', async () => {
    setVerbose(true);
    
    const mockAxiosInstance = {
      get: jest.fn().mockResolvedValue({ status: 200, data: {} })
    };
    axios.create = jest.fn(() => mockAxiosInstance);
    
    await get('/test');
    
    expect(consoleErrorSpy).toHaveBeenCalled();
    const loggedOutput = consoleErrorSpy.mock.calls.map(call => call[0]).join('\n');
    expect(loggedOutput).toContain('GET');
    expect(loggedOutput).toContain('/test');
  });

  it('should not log when verbose is disabled', async () => {
    setVerbose(false);
    
    const mockAxiosInstance = {
      get: jest.fn().mockResolvedValue({ status: 200, data: {} })
    };
    axios.create = jest.fn(() => mockAxiosInstance);
    
    await get('/test');
    
    expect(consoleErrorSpy).not.toHaveBeenCalled();
  });
});
