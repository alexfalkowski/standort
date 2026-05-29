# frozen_string_literal: true

When('I request a location with HTTP which performs in {int} ms') do |time|
  @request_id = SecureRandom.uuid
  opts = {
    headers: {
      request_id: @request_id, user_agent: 'Standort-ruby-client/2.0 HTTP/1.0',
      content_type: :json, accept: :json
    },
    read_timeout: 10, open_timeout: 10
  }
  params = { ip: '95.91.246.242' }
  response = nil

  expect { response = Standort::V2.http.get_location(params, opts) }.to perform_under(time).ms

  body = JSON.parse(response.body)

  expect(response.code).to eq(200)
  expect(body.fetch('meta').fetch('requestId')).to eq(@request_id)
  expect(body.fetch('ip').fetch('country')).to eq('DE')
  expect(body.fetch('ip').fetch('continent')).to eq('EU')
  expect(body['geo']).to be_nil
end
