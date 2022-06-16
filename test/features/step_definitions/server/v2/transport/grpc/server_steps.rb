# frozen_string_literal: true

When('I request a location with gRPC:') do |table|
  rows = table.rows_hash
  @request_id = SecureRandom.uuid
  metadata = { 'request-id' => @request_id, 'ua' => Standort.server_config['transport']['grpc']['user_agent'] }
  params = {}

  if rows['method'] == 'params'
    params[:ip] = rows['ip'] if rows['ip']
    params[:point] = Standort::V2::Point.new(lat: rows['latitude'].to_f, lng: rows['longitude'].to_f) if rows['latitude'] && rows['longitude']
  end

  if rows['method'] == 'metadata'
    metadata['forwarded-for'] = rows['ip'] if rows['ip']
    metadata['geolocation'] = "geo:#{rows['latitude']},#{rows['longitude']}" if rows['latitude'] && rows['longitude']
  end

  request = Standort::V2::GetLocationRequest.new(params)
  @response = Standort::V2.server_grpc.get_location(request, { metadata: metadata })
rescue StandardError => e
  @response = e
end

Then('I should receive a valid locations with gRPC:') do |table|
  rows = table.rows_hash

  expect(@response.locations.length).to eq(1)

  location = @response.locations[0]
  expect(location.country).to eq(rows['country'])
  expect(location.continent).to eq(rows['continent'])
  expect(location.kind).to eq(rows['kind'].to_sym)
end

Then('I should receive an empty response with gRPC') do
  expect(@response.locations.length).to eq(0)
end
