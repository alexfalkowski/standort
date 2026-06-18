# frozen_string_literal: true

When('I request a location by IP address with HTTP:') do |table|
  rows = table.rows_hash
  @request_id = SecureRandom.uuid
  opts = Standort.http_options(
    headers: {
      request_id: @request_id, user_agent: Standort.config.transport.http.user_agent,
      content_type: :json, accept: :json
    }
  )

  @response = Standort::V1.http.get_location_by_ip(rows['ip'].strip, opts)
end

When('I request a location by latitude and longitude with HTTP:') do |table|
  rows = table.rows_hash
  @request_id = SecureRandom.uuid
  opts = Standort.http_options(
    headers: {
      request_id: @request_id, user_agent: Standort.config.transport.http.user_agent,
      content_type: :json, accept: :json
    }
  )

  @response = Standort::V1.http.get_location_by_lat_lng(rows['latitude'], rows['longitude'], opts)
end

Then('I should receive a valid location by IP adress with HTTP:') do |table|
  expect(@response.code).to eq(200)

  resp = JSON.parse(@response.body)
  location = resp['location']
  rows = table.rows_hash

  expect(resp.fetch('meta').fetch('requestId')).to eq(@request_id)
  expect(location['country']).to eq(rows['country'])
  expect(location['continent']).to eq(rows['continent'])
end

Then('I should receive a not found response with HTTP') do
  expect(@response.code).to eq(404)
end

Then('I should receive a valid location by latitude and longitude with HTTP:') do |table|
  expect(@response.code).to eq(200)

  resp = JSON.parse(@response.body)
  location = resp['location']
  rows = table.rows_hash

  expect(resp.fetch('meta').fetch('requestId')).to eq(@request_id)
  expect(location['country']).to eq(rows['country'])
  expect(location['continent']).to eq(rows['continent'])
end
