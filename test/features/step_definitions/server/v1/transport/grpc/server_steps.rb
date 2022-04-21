# frozen_string_literal: true

When('I request a location by IP address with gRPC:') do |table|
  rows = table.rows_hash
  @request_id = SecureRandom.uuid
  metadata = {
    'request-id' => @request_id,
    'ua' => Standort.server_config['transport']['grpc']['user_agent']
  }

  request = Standort::V1::GetLocationByIPRequest.new(ip: rows['ip'])
  @response = Standort.server_grpc.get_location_by_ip(request, { metadata: metadata })
rescue StandardError => e
  @response = e
end

Then('I should receive a valid location by IP adress with gRPC:') do |table|
  rows = table.rows_hash

  expect(@response.location.country).to eq(rows['country'])
  expect(@response.location.continent).to eq(rows['continent'])
end

Then('I should receive a bad response with gRPC') do
  expect(@response).to be_a(GRPC::InvalidArgument)
end

Then('I should receive a not found response with gRPC') do
  expect(@response).to be_a(GRPC::NotFound)
end
