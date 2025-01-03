# frozen_string_literal: true

When('I request a location by IP address with gRPC:') do |table|
  rows = table.rows_hash
  @request_id = SecureRandom.uuid
  metadata = { 'request-id' => @request_id }

  request = Standort::V1::GetLocationByIPRequest.new(ip: rows['ip'].strip)
  @response = Standort::V1.grpc.get_location_by_ip(request, { metadata: })
rescue StandardError => e
  @response = e
end

When('I request a location by latitude and longitude with gRPC:') do |table|
  rows = table.rows_hash
  @request_id = SecureRandom.uuid
  metadata = { 'request-id' => @request_id }

  request = Standort::V1::GetLocationByLatLngRequest.new(lat: rows['latitude'].to_f, lng: rows['longitude'].to_f)
  @response = Standort::V1.grpc.get_location_by_lat_lng(request, { metadata: })
rescue StandardError => e
  @response = e
end

Then('I should receive a valid location by IP adress with gRPC:') do |table|
  rows = table.rows_hash

  expect(@response.meta.length).to be > 0
  expect(@response.location.country).to eq(rows['country'])
  expect(@response.location.continent).to eq(rows['continent'])
end

Then('I should receive a not found response with gRPC') do
  expect(@response).to be_a(GRPC::NotFound)
end

Then('I should receive a valid location by latitude and longitude with gRPC:') do |table|
  rows = table.rows_hash

  expect(@response.meta.length).to be > 0
  expect(@response.location.country).to eq(rows['country'])
  expect(@response.location.continent).to eq(rows['continent'])
end
