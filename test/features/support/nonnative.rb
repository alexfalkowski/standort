# frozen_string_literal: true

Nonnative.configure do |config|
  config.load_file('nonnative.yml')
end

Given('I have {string} as the config file') do |source|
  Nonnative.configuration.processes[0].environment['CONFIG_FILE'] = ".config/#{source}.server.config.yml"
end
