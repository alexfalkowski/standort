# frozen_string_literal: true

When('I request a location with gRPC:') do |table|
  rows = table.rows_hash
  @request_id = SecureRandom.uuid
  metadata = { 'request-id' => @request_id }
  params = {}

  if rows['method'] == 'params'
    params[:ip] = rows['ip'] if rows['ip']
    if rows['latitude'] && rows['longitude']
      params[:point] = Standort::V2::Point.new(
        lat: Coordinates.parse(rows['latitude']),
        lng: Coordinates.parse(rows['longitude'])
      )
    end
  end

  if rows['method'] == 'metadata'
    metadata['x-forwarded-for'] = rows['ip'] if rows['ip']
    metadata['geolocation'] = "geo:#{rows['latitude']},#{rows['longitude']}" if rows['latitude'] && rows['longitude']
  end

  request = Standort::V2::GetLocationRequest.new(params)
  @operation = Standort::V2.grpc.get_location(request, Standort.grpc_options(metadata).merge(return_op: true))
  @response = @operation.execute
rescue StandardError => e
  @response = e
end

Then('I should receive a valid locations with gRPC:') do |table|
  expect(@response.meta['requestId']).to eq(@request_id)

  rows = table.rows_hash
  location = case rows['kind']
             when 'ip'
               expect(@response.geo).to be_nil
               @response.ip
             when 'geo'
               expect(@response.ip).to be_nil
               @response.geo
             else
               raise "unsupported location kind: #{rows['kind']}"
             end

  expect(location.country).to eq(rows['country'])
  expect(location.continent).to eq(rows['continent'])
end

Then('I should receive valid locations with gRPC:') do |table|
  expect(@response.meta['requestId']).to eq(@request_id)

  table.hashes.each do |row|
    location = case row['kind']
               when 'ip'
                 @response.ip
               when 'geo'
                 @response.geo
               else
                 raise "unsupported location kind: #{row['kind']}"
               end

    expect(location.country).to eq(row['country'])
    expect(location.continent).to eq(row['continent'])
  end
end

Then('I should receive a not found response with gRPC:') do |table|
  expect(@response).to be_a(GRPC::NotFound)

  rows = table.rows_hash

  expect(Array(@operation.trailing_metadata[rows['diagnostic']])).to include(rows['code'])
end

Then('I should receive a partial location with gRPC:') do |table|
  expect(@response.meta['requestId']).to eq(@request_id)

  rows = table.rows_hash
  location = case rows['kind']
             when 'ip'
               expect(@response.geo).to be_nil
               @response.ip
             when 'geo'
               expect(@response.ip).to be_nil
               @response.geo
             else
               raise "unsupported location kind: #{rows['kind']}"
             end

  expect(location.country).to eq(rows['country'])
  expect(location.continent).to eq(rows['continent'])
end
