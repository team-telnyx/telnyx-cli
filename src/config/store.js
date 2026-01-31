const Conf = require('conf');

const config = new Conf({
  projectName: 'telnyx-cli',
  schema: {
    apiKey: {
      type: 'string',
      default: ''
    }
  }
});

module.exports = {
  getApiKey: () => config.get('apiKey'),
  setApiKey: (apiKey) => config.set('apiKey', apiKey),
  clearApiKey: () => config.delete('apiKey'),
  hasApiKey: () => !!config.get('apiKey')
};
