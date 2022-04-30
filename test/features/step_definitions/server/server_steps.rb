# frozen_string_literal: true

When('the server is configured with an invalid {string} path') do |type|
  Nonnative.configuration.processes[0].environment['CONFIG_FILE'] = ".config/invalid.#{type}.path.server.config.yml"
end

When('the server is configured with an invalid ip provider') do
  Nonnative.configuration.processes[0].environment['CONFIG_FILE'] = '.config/invalid.ip.type.server.config.yml'
end

Then('the server is configured with a valid configuration') do
  Nonnative.configuration.processes[0].environment['CONFIG_FILE'] = '.config/server.config.yml'
end
