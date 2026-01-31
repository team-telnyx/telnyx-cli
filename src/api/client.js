const axios = require('axios');
const { getApiKey } = require('../config/store');

const BASE_URL = 'https://api.telnyx.com/v2';

function createClient() {
  const apiKey = getApiKey();
  
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
    timeout: 30000
  });
}

async function get(path, params = {}) {
  const client = createClient();
  const response = await client.get(path, { params });
  return response.data;
}

async function post(path, data) {
  const client = createClient();
  const response = await client.post(path, data);
  return response.data;
}

async function patch(path, data) {
  const client = createClient();
  const response = await client.patch(path, data);
  return response.data;
}

async function del(path) {
  const client = createClient();
  const response = await client.delete(path);
  return response.data;
}

// ==================== Numbers API ====================

async function searchAvailableNumbers(filters = {}) {
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
  
  return get('/available_phone_numbers', params);
}

async function listPhoneNumbers(filters = {}) {
  const params = {};
  
  if (filters.status) params['filter[status]'] = filters.status;
  if (filters.numberType) params['filter[phone_number_type]'] = filters.numberType;
  if (filters.limit) params['page[size]'] = filters.limit;
  if (filters.page) params['page[number]'] = filters.page;
  
  return get('/phone_numbers', params);
}

async function orderPhoneNumber(phoneNumber, options = {}) {
  const data = {
    phone_number: phoneNumber,
    ...options
  };
  
  return post('/phone_numbers', { data });
}

async function getPhoneNumber(numberId) {
  return get(`/phone_numbers/${numberId}`);
}

async function updatePhoneNumber(numberId, updates) {
  return patch(`/phone_numbers/${numberId}`, { data: updates });
}

async function deletePhoneNumber(numberId) {
  return del(`/phone_numbers/${numberId}`);
}

// ==================== Messaging API ====================

async function sendMessage(to, from, text, options = {}) {
  const data = {
    to: to,
    from: from,
    text: text,
    ...options
  };
  
  return post('/messages', { data });
}

async function getMessage(messageId) {
  return get(`/messages/${messageId}`);
}

async function listMessages(filters = {}) {
  const params = {};
  
  if (filters.from) params['filter[from]'] = filters.from;
  if (filters.to) params['filter[to]'] = filters.to;
  if (filters.status) params['filter[status]'] = filters.status;
  if (filters.limit) params['page[size]'] = filters.limit;
  
  return get('/messages', params);
}

// ==================== Billing API ====================

async function getBillingBalance() {
  return get('/balance');
}

// ==================== AI Assistants API ====================

async function listAssistants(filters = {}) {
  const params = {};
  
  if (filters.limit) params['limit'] = filters.limit;
  if (filters.cursor) params['cursor'] = filters.cursor;
  
  return get('/ai/assistants', params);
}

async function getAssistant(assistantId) {
  return get(`/ai/assistants/${assistantId}`);
}

// ==================== Voice/Calls API ====================

async function listCalls(filters = {}) {
  const params = {};
  
  if (filters.status) params['filter[status]'] = filters.status;
  if (filters.limit) params['page[size]'] = filters.limit;
  
  return get('/calls', params);
}

async function createCall(options) {
  return post('/calls', { data: options });
}

async function getCall(callId) {
  return get(`/calls/${callId}`);
}

async function hangupCall(callId) {
  return post(`/calls/${callId}/actions/hangup`, {});
}

async function answerCall(callId) {
  return post(`/calls/${callId}/actions/answer`, {});
}

async function transferCall(callId, to, options = {}) {
  const data = {
    to: to,
    ...options
  };
  return post(`/calls/${callId}/actions/transfer`, { data });
}

// ==================== Number Lookup API ====================

async function lookupNumber(phoneNumber, options = {}) {
  const params = {};
  
  if (options.type) params['type'] = options.type;
  if (options.fields) params['fields'] = options.fields;
  
  return get(`/number_lookup/${phoneNumber}`, params);
}

// ==================== IoT SIMs API ====================

async function listSims(filters = {}) {
  const params = {};
  
  if (filters.status) params['filter[status]'] = filters.status;
  if (filters.limit) params['page[size]'] = filters.limit;
  
  return get('/sim_cards', params);
}

async function getSim(simId) {
  return get(`/sim_cards/${simId}`);
}

async function updateSim(simId, updates) {
  return patch(`/sim_cards/${simId}`, { data: updates });
}

// ==================== Fax API ====================

async function sendFax(to, from, mediaUrl, options = {}) {
  const data = {
    to: to,
    from: from,
    media_url: mediaUrl,
    ...options
  };
  
  return post('/faxes', { data });
}

async function listFaxes(filters = {}) {
  const params = {};
  
  if (filters.status) params['filter[status]'] = filters.status;
  if (filters.limit) params['page[size]'] = filters.limit;
  
  return get('/faxes', params);
}

async function getFax(faxId) {
  return get(`/faxes/${faxId}`);
}

async function cancelFax(faxId) {
  return post(`/faxes/${faxId}/actions/cancel`, {});
}

// ==================== Verify API ====================

async function createVerification(phoneNumber, options = {}) {
  const data = {
    phone_number: phoneNumber,
    verify_profile_id: options.profileId,
    type: options.type || 'sms',
    timeout: options.timeout || 300,
    ...options
  };
  
  return post('/verifications', { data });
}

async function submitVerification(phoneNumber, code, options = {}) {
  const data = {
    code: code,
    verify_profile_id: options.profileId,
    ...options
  };
  
  return post(`/verifications/by_phone_number/${phoneNumber}/actions/verify`, { data });
}

async function getVerification(verificationId) {
  return get(`/verifications/${verificationId}`);
}

async function listVerifications(filters = {}) {
  const params = {};
  
  if (filters.status) params['filter[status]'] = filters.status;
  if (filters.limit) params['page[size]'] = filters.limit;
  
  return get('/verifications', params);
}

// ==================== Debugger/Events API ====================

async function listEvents(filters = {}) {
  const params = {};
  
  if (filters.type) params['filter[event_type]'] = filters.type;
  if (filters.status) params['filter[status]'] = filters.status;
  if (filters.limit) params['page[size]'] = filters.limit;
  if (filters.since) params['filter[created_at][gt]'] = filters.since;
  
  return get('/events', params);
}

async function getEvent(eventId) {
  return get(`/events/${eventId}`);
}

async function listWebhooks(filters = {}) {
  const params = {};
  
  if (filters.status) params['filter[status]'] = filters.status;
  if (filters.limit) params['page[size]'] = filters.limit;
  
  return get('/webhook_deliveries', params);
}

async function getWebhook(webhookId) {
  return get(`/webhook_deliveries/${webhookId}`);
}

async function resendWebhook(webhookId) {
  return post(`/webhook_deliveries/${webhookId}/actions/resend`, {});
}

// ==================== Generic API Methods (Long Tail Support) ====================

async function genericList(resource, filters = {}) {
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
  return get(endpoint, params);
}

async function genericGet(resource, id) {
  return get(`/${resource}/${id}`);
}

async function genericCreate(resource, data) {
  return post(`/${resource}`, { data });
}

async function genericUpdate(resource, id, data) {
  return patch(`/${resource}/${id}`, { data });
}

async function genericDelete(resource, id) {
  return del(`/${resource}/${id}`);
}

module.exports = {
  BASE_URL,
  get,
  post,
  patch,
  del,
  createClient,
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
