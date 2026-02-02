const axios = require('axios');
const { getApiKey } = require('../config/store');

const BASE_URL = 'https://api.telnyx.com/v2';

// Global settings (set by CLI flags)
let globalVerbose = false;
let globalTimeout = 30000;

function setVerbose(verbose) {
  globalVerbose = verbose;
}

function setTimeout(timeoutMs) {
  globalTimeout = timeoutMs;
}

function logRequest(method, url, params = null) {
  if (!globalVerbose) return;
  const timestamp = new Date().toISOString();
  console.error(`[${timestamp}] → ${method.toUpperCase()} ${url}`);
  if (params) {
    console.error(`[${timestamp}]   Params: ${JSON.stringify(params)}`);
  }
}

function logResponse(method, url, status, duration) {
  if (!globalVerbose) return;
  const timestamp = new Date().toISOString();
  console.error(`[${timestamp}] ← ${method.toUpperCase()} ${url} ${status} (${duration}ms)`);
}

function logError(method, url, error, duration) {
  if (!globalVerbose) return;
  const timestamp = new Date().toISOString();
  const status = error.response?.status || 'NETWORK_ERROR';
  console.error(`[${timestamp}] ✖ ${method.toUpperCase()} ${url} ${status} (${duration}ms)`);
  if (error.message) {
    console.error(`[${timestamp}]   Error: ${error.message}`);
  }
}

function createClient(profileName) {
  const apiKey = getApiKey(profileName);
  
  if (!apiKey) {
    throw new Error('No API key configured. Run: telnyx auth login');
  }

  return axios.create({
    baseURL: BASE_URL,
    headers: {
      'Authorization': `Bearer ${apiKey}`,
      'Content-Type': 'application/json',
      'Accept': 'application/json'
    },
    timeout: globalTimeout
  });
}

// Normalize API errors into friendly messages
function normalizeError(error) {
  const status = error.response?.status;
  const errorData = error.response?.data;
  
  // Network/connection errors
  if (!error.response) {
    if (error.code === 'ECONNABORTED') {
      const timeoutSec = Math.round(globalTimeout / 1000);
      return new Error(`Request timed out after ${timeoutSec}s. Use --timeout to increase.`);
    }
    if (error.code === 'ENOTFOUND' || error.code === 'ECONNREFUSED') {
      return new Error('Connection failed. Check your internet connection.');
    }
    return new Error(`Network error: ${error.message}`);
  }
  
  // HTTP errors
  switch (status) {
    case 401:
      return new Error('Authentication failed. Run \'telnyx auth login\' to re-authenticate.');
    
    case 404: {
      // Try to extract resource from URL
      const url = error.config?.url || '';
      const resourceMatch = url.match(/\/([^/]+)(?:\/[^/]+)?$/);
      const resource = resourceMatch ? resourceMatch[1].replace(/_/g, ' ') : 'resource';
      return new Error(`Resource not found: ${resource}`);
    }
    
    case 422: {
      // Parse Telnyx API error details
      const errors = errorData?.errors || [];
      if (errors.length > 0) {
        const details = errors.map(e => {
          const title = e.title || 'Error';
          const detail = e.detail || e.message || '';
          return detail ? `${title}: ${detail}` : title;
        });
        return new Error(`Validation failed: ${details.join('; ')}`);
      }
      return new Error('Validation failed: Invalid request parameters');
    }
    
    case 403:
      return new Error('Permission denied. Your account may not have access to this feature.');
    
    case 402:
      return new Error('Insufficient balance. Please add funds to your account.');
    
    case 429:
      return new Error('Rate limit exceeded. Please wait a moment and try again.');
    
    case 500:
    case 502:
    case 503:
    case 504:
      return new Error('Telnyx API error. Please try again later.');
    
    default:
      const message = errorData?.errors?.[0]?.detail || 
                      errorData?.errors?.[0]?.title || 
                      error.message || 
                      `HTTP ${status} error`;
      return new Error(message);
  }
}

async function get(path, params = {}, profileName) {
  const client = createClient(profileName);
  const startTime = Date.now();
  const fullUrl = `${BASE_URL}${path}`;
  
  logRequest('GET', fullUrl, params);
  
  try {
    const response = await client.get(path, { params });
    logResponse('GET', fullUrl, response.status, Date.now() - startTime);
    return response.data;
  } catch (error) {
    logError('GET', fullUrl, error, Date.now() - startTime);
    throw normalizeError(error);
  }
}

async function post(path, data, profileName) {
  const client = createClient(profileName);
  const startTime = Date.now();
  const fullUrl = `${BASE_URL}${path}`;
  
  logRequest('POST', fullUrl, data);
  
  try {
    const response = await client.post(path, data);
    logResponse('POST', fullUrl, response.status, Date.now() - startTime);
    return response.data;
  } catch (error) {
    logError('POST', fullUrl, error, Date.now() - startTime);
    throw normalizeError(error);
  }
}

async function patch(path, data, profileName) {
  const client = createClient(profileName);
  const startTime = Date.now();
  const fullUrl = `${BASE_URL}${path}`;
  
  logRequest('PATCH', fullUrl, data);
  
  try {
    const response = await client.patch(path, data);
    logResponse('PATCH', fullUrl, response.status, Date.now() - startTime);
    return response.data;
  } catch (error) {
    logError('PATCH', fullUrl, error, Date.now() - startTime);
    throw normalizeError(error);
  }
}

async function del(path, profileName) {
  const client = createClient(profileName);
  const startTime = Date.now();
  const fullUrl = `${BASE_URL}${path}`;
  
  logRequest('DELETE', fullUrl);
  
  try {
    const response = await client.delete(path);
    logResponse('DELETE', fullUrl, response.status, Date.now() - startTime);
    return response.data;
  } catch (error) {
    logError('DELETE', fullUrl, error, Date.now() - startTime);
    throw normalizeError(error);
  }
}

// ==================== Numbers API ====================

async function searchAvailableNumbers(filters = {}, profileName) {
  const params = {};
  
  // Required
  if (filters.country) params['filter[country_code]'] = filters.country;
  
  // Standard filters
  if (filters.locality) params['filter[locality]'] = filters.locality;
  if (filters.numberType) params['filter[phone_number_type]'] = filters.numberType;
  if (filters.limit) params['filter[limit]'] = filters.limit;
  if (filters.areaCode) params['filter[national_destination_code]'] = filters.areaCode;
  if (filters.state) params['filter[administrative_area]'] = filters.state;
  if (filters.features) params['filter[features]'] = filters.features;
  
  // Advanced filters
  if (filters.startsWith) params['filter[starts_with]'] = filters.startsWith;
  if (filters.endsWith) params['filter[ends_with]'] = filters.endsWith;
  if (filters.contains) params['filter[contains]'] = filters.contains;
  if (filters.quickship) params['filter[quickship]'] = 'true';
  if (filters.bestEffort) params['filter[best_effort]'] = 'true';
  if (filters.reservable) params['filter[reservable]'] = 'true';
  if (filters.excludeHeld) params['filter[exclude_held_numbers]'] = 'true';
  
  return get('/available_phone_numbers', params, profileName);
}

async function listPhoneNumbers(filters = {}, profileName) {
  const params = {};
  
  if (filters.status) params['filter[status]'] = filters.status;
  if (filters.numberType) params['filter[phone_number_type]'] = filters.numberType;
  if (filters.limit) params['page[size]'] = filters.limit;
  if (filters.page) params['page[number]'] = filters.page;
  
  return get('/phone_numbers', params, profileName);
}

async function orderPhoneNumber(phoneNumber, options = {}, profileName) {
  const data = {
    phone_number: phoneNumber,
    ...options
  };
  
  return post('/phone_numbers', { data }, profileName);
}

async function getPhoneNumber(numberId, profileName) {
  return get(`/phone_numbers/${numberId}`, {}, profileName);
}

async function updatePhoneNumber(numberId, updates, profileName) {
  return patch(`/phone_numbers/${numberId}`, { data: updates }, profileName);
}

async function deletePhoneNumber(numberId, profileName) {
  return del(`/phone_numbers/${numberId}`, profileName);
}

// ==================== Messaging API ====================

async function sendMessage(to, from, text, options = {}, profileName) {
  const data = {
    to: to,
    from: from,
    text: text,
    ...options
  };
  
  return post('/messages', { data }, profileName);
}

async function getMessage(messageId, profileName) {
  return get(`/messages/${messageId}`, {}, profileName);
}

async function listMessages(filters = {}, profileName) {
  const params = {};
  
  if (filters.from) params['filter[from]'] = filters.from;
  if (filters.to) params['filter[to]'] = filters.to;
  if (filters.status) params['filter[status]'] = filters.status;
  if (filters.limit) params['page[size]'] = filters.limit;
  
  return get('/messages', params, profileName);
}

// ==================== Billing API ====================

async function getBillingBalance(profileName) {
  return get('/balance', {}, profileName);
}

// ==================== AI Assistants API ====================

async function listAssistants(filters = {}, profileName) {
  const params = {};
  
  if (filters.limit) params['limit'] = filters.limit;
  if (filters.cursor) params['cursor'] = filters.cursor;
  
  return get('/ai/assistants', params, profileName);
}

async function getAssistant(assistantId, profileName) {
  return get(`/ai/assistants/${assistantId}`, {}, profileName);
}

// ==================== Voice/Calls API ====================

async function listCalls(filters = {}, profileName) {
  const params = {};
  
  if (filters.status) params['filter[status]'] = filters.status;
  if (filters.limit) params['page[size]'] = filters.limit;
  
  return get('/calls', params, profileName);
}

async function createCall(options, profileName) {
  return post('/calls', { data: options }, profileName);
}

async function getCall(callId, profileName) {
  return get(`/calls/${callId}`, {}, profileName);
}

async function hangupCall(callId, profileName) {
  return post(`/calls/${callId}/actions/hangup`, {}, profileName);
}

async function answerCall(callId, profileName) {
  return post(`/calls/${callId}/actions/answer`, {}, profileName);
}

async function transferCall(callId, to, options = {}, profileName) {
  const data = {
    to: to,
    ...options
  };
  return post(`/calls/${callId}/actions/transfer`, { data }, profileName);
}

// ==================== Number Lookup API ====================

async function lookupNumber(phoneNumber, options = {}, profileName) {
  const params = {};
  
  if (options.type) params['type'] = options.type;
  if (options.fields) params['fields'] = options.fields;
  
  return get(`/number_lookup/${phoneNumber}`, params, profileName);
}

// ==================== IoT SIMs API ====================

async function listSims(filters = {}, profileName) {
  const params = {};
  
  if (filters.status) params['filter[status]'] = filters.status;
  if (filters.limit) params['page[size]'] = filters.limit;
  
  return get('/sim_cards', params, profileName);
}

async function getSim(simId, profileName) {
  return get(`/sim_cards/${simId}`, {}, profileName);
}

async function updateSim(simId, updates, profileName) {
  return patch(`/sim_cards/${simId}`, { data: updates }, profileName);
}

// ==================== Fax API ====================

async function sendFax(to, from, mediaUrl, options = {}, profileName) {
  const data = {
    to: to,
    from: from,
    media_url: mediaUrl,
    ...options
  };
  
  return post('/faxes', { data }, profileName);
}

async function listFaxes(filters = {}, profileName) {
  const params = {};
  
  if (filters.status) params['filter[status]'] = filters.status;
  if (filters.limit) params['page[size]'] = filters.limit;
  
  return get('/faxes', params, profileName);
}

async function getFax(faxId, profileName) {
  return get(`/faxes/${faxId}`, {}, profileName);
}

async function cancelFax(faxId, profileName) {
  return post(`/faxes/${faxId}/actions/cancel`, {}, profileName);
}

// ==================== Verify API ====================

async function createVerification(phoneNumber, options = {}, profileName) {
  const data = {
    phone_number: phoneNumber,
    verify_profile_id: options.profileId,
    type: options.type || 'sms',
    timeout: options.timeout || 300,
    ...options
  };
  
  return post('/verifications', { data }, profileName);
}

async function submitVerification(phoneNumber, code, options = {}, profileName) {
  const data = {
    code: code,
    verify_profile_id: options.profileId,
    ...options
  };
  
  return post(`/verifications/by_phone_number/${phoneNumber}/actions/verify`, { data }, profileName);
}

async function getVerification(verificationId, profileName) {
  return get(`/verifications/${verificationId}`, {}, profileName);
}

async function listVerifications(filters = {}, profileName) {
  const params = {};
  
  if (filters.status) params['filter[status]'] = filters.status;
  if (filters.limit) params['page[size]'] = filters.limit;
  
  return get('/verifications', params, profileName);
}

// ==================== Debugger/Events API ====================

async function listEvents(filters = {}, profileName) {
  const params = {};
  
  if (filters.type) params['filter[event_type]'] = filters.type;
  if (filters.status) params['filter[status]'] = filters.status;
  if (filters.limit) params['page[size]'] = filters.limit;
  if (filters.since) params['filter[created_at][gt]'] = filters.since;
  
  return get('/events', params, profileName);
}

async function getEvent(eventId, profileName) {
  return get(`/events/${eventId}`, {}, profileName);
}

async function listWebhooks(filters = {}, profileName) {
  const params = {};
  
  if (filters.status) params['filter[status]'] = filters.status;
  if (filters.limit) params['page[size]'] = filters.limit;
  
  return get('/webhook_deliveries', params, profileName);
}

async function getWebhook(webhookId, profileName) {
  return get(`/webhook_deliveries/${webhookId}`, {}, profileName);
}

async function resendWebhook(webhookId, profileName) {
  return post(`/webhook_deliveries/${webhookId}/actions/resend`, {}, profileName);
}

// ==================== Generic API Methods (Long Tail Support) ====================

async function genericList(resource, filters = {}, profileName) {
  const params = {};
  
  // Common filter mappings
  if (filters.limit) params['page[size]'] = filters.limit;
  if (filters.page) params['page[number]'] = filters.page;
  if (filters.status) params['filter[status]'] = filters.status;
  if (filters.type) params['filter[type]'] = filters.type;
  
  // Map resource names to correct API endpoints
  const endpointMap = {
    'assistants': '/ai/assistants'
  };
  
  const endpoint = endpointMap[resource] || `/${resource}`;
  return get(endpoint, params, profileName);
}

async function genericGet(resource, id, profileName) {
  return get(`/${resource}/${id}`, {}, profileName);
}

async function genericCreate(resource, data, profileName) {
  return post(`/${resource}`, { data }, profileName);
}

async function genericUpdate(resource, id, data, profileName) {
  return patch(`/${resource}/${id}`, { data }, profileName);
}

async function genericDelete(resource, id, profileName) {
  return del(`/${resource}/${id}`, profileName);
}

module.exports = {
  BASE_URL,
  get,
  post,
  patch,
  del,
  createClient,
  setVerbose,
  setTimeout,
  normalizeError,
  // Numbers
  searchAvailableNumbers,
  listPhoneNumbers,
  orderPhoneNumber,
  getPhoneNumber,
  updatePhoneNumber,
  deletePhoneNumber,
  // Messaging
  sendMessage,
  getMessage,
  listMessages,
  // Billing
  getBillingBalance,
  // AI
  listAssistants,
  getAssistant,
  // Voice/Calls
  listCalls,
  createCall,
  getCall,
  hangupCall,
  answerCall,
  transferCall,
  // Number Lookup
  lookupNumber,
  // IoT SIMs
  listSims,
  getSim,
  updateSim,
  // Fax
  sendFax,
  listFaxes,
  getFax,
  cancelFax,
  // Verify
  createVerification,
  submitVerification,
  getVerification,
  listVerifications,
  // Debugger/Events
  listEvents,
  getEvent,
  listWebhooks,
  getWebhook,
  resendWebhook,
  // Generic (Long Tail Support)
  genericList,
  genericGet,
  genericCreate,
  genericUpdate,
  genericDelete
};
