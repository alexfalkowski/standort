# frozen_string_literal: true

When('the server is configured with an invalid ip path') do
  Nonnative.configuration.processes[0].environment['CONFIG_FILE'] = '.config/invalid.ip.path.server.config.yml'
end

When('the server is configured with an invalid location path') do
  Nonnative.configuration.processes[0].environment['CONFIG_FILE'] = '.config/invalid.location.path.server.config.yml'
end

Then('the server is configured with a valid configuration') do
  Nonnative.configuration.processes[0].environment['CONFIG_FILE'] = '.config/server.config.yml'
end
