# frozen_string_literal: true

When('I request a location with gRPC which performs in {int} ms') do |time|
  @request_id = SecureRandom.uuid
  opts = {
    headers: {
      request_id: SecureRandom.uuid, user_agent: 'Standort-ruby-client/2.0 HTTP/1.0',
      content_type: :json, accept: :json
    },
    read_timeout: 10, open_timeout: 10
  }
  params = { 'ip' => '95.91.246.242' }

  expect { Standort::V2.http.get_location(params, opts) }.to perform_under(time).ms
end
