# frozen_string_literal: true

When('I request a location by IP address with HTTP:') do |table|
  rows = table.rows_hash
  headers = {
    request_id: SecureRandom.uuid,
    user_agent: Standort.server_config['transport']['grpc']['user_agent']
  }

  @response = Standort.server_http.get_location_by_ip(rows['ip'].strip, headers)
end

When('I request a location by latitude and longitude with HTTP:') do |table|
  rows = table.rows_hash
  headers = {
    request_id: SecureRandom.uuid,
    user_agent: Standort.server_config['transport']['grpc']['user_agent']
  }

  @response = Standort.server_http.get_location_by_lat_lng(rows['latitude'], rows['longitude'], headers)
end

Then('I should receive a valid location by IP adress with HTTP:') do |table|
  expect(@response.code).to eq(200)

  resp = JSON.parse(@response.body)
  location = resp['location']
  rows = table.rows_hash

  expect(location['country']).to eq(rows['country'])
  expect(location['continent']).to eq(rows['continent'])
end

Then('I should receive a bad response with HTTP') do
  expect(@response.code).to eq(400)
end

Then('I should receive a not found response with HTTP') do
  expect(@response.code).to eq(404)
end

Then('I should receive a valid location by latitude and longitude with HTTP:') do |table|
  expect(@response.code).to eq(200)

  resp = JSON.parse(@response.body)
  location = resp['location']
  rows = table.rows_hash

  expect(location['country']).to eq(rows['country'])
  expect(location['continent']).to eq(rows['continent'])
end
