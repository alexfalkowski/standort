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

When('I lookup locations with gRPC:') do |table|
  @request_id = SecureRandom.uuid
  metadata = { 'request-id' => @request_id }
  lookups = table.hashes.map do |row|
    params = {}
    params[:ip] = row['ip'] unless row['ip'].empty?
    if !row['latitude'].empty? && !row['longitude'].empty?
      params[:point] = Standort::V2::Point.new(
        lat: Coordinates.parse(row['latitude']),
        lng: Coordinates.parse(row['longitude'])
      )
    end

    Standort::V2::LocationLookup.new(params)
  end

  request = Standort::V2::LookupLocationsRequest.new(lookups:)
  @response = Standort::V2.grpc.lookup_locations(request, Standort.grpc_options(metadata))
rescue StandardError => e
  @response = e
end

When('I lookup {int} locations with gRPC') do |count|
  lookups = Array.new(count) do
    Standort::V2::LocationLookup.new(ip: '95.91.246.242')
  end

  request = Standort::V2::LookupLocationsRequest.new(lookups:)
  @response = Standort::V2.grpc.lookup_locations(request, Standort.grpc_options)
rescue StandardError => e
  @response = e
end

When('I request lookup assets with gRPC') do
  @request_id = SecureRandom.uuid
  metadata = { 'request-id' => @request_id }
  request = Standort::V2::GetLookupAssetsRequest.new

  @response = Standort::V2.grpc.get_lookup_assets(request, Standort.grpc_options(metadata))
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

Then('I should receive batch locations with gRPC:') do |table|
  expect(@response.meta['requestId']).to eq(@request_id)

  table.hashes.each do |row|
    lookup = @response.lookups.fetch(row['index'].to_i)
    location = case row['kind']
               when 'ip'
                 expect(lookup.geo).to be_nil
                 lookup.ip
               when 'geo'
                 expect(lookup.ip).to be_nil
                 lookup.geo
               when 'none'
                 expect(lookup.ip).to be_nil
                 expect(lookup.geo).to be_nil
                 expect(lookup.status.code).to eq(row['code'].to_i)
                 next
               else
                 raise "unsupported location kind: #{row['kind']}"
               end

    expect(lookup.status).to be_nil
    expect(location.country).to eq(row['country'])
    expect(location.continent).to eq(row['continent'])
  end
end

Then('I should receive an invalid argument response with gRPC') do
  expect(@response).to be_a(GRPC::InvalidArgument)
end

Then('I should receive lookup assets with gRPC:') do |table|
  expect(@response.meta['requestId']).to eq(@request_id)

  table.hashes.each do |row|
    asset = @response.assets.find { |lookup_asset| lookup_asset.name == row['name'] }

    expect(asset).not_to be_nil
    expect(asset.size_bytes).to be_positive
    expect(asset.checksum_algorithm).to eq(row['checksum_algorithm'])
    expect(asset.checksum).not_to be_empty
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
