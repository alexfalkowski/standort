# frozen_string_literal: true

When('I request a location with HTTP:') do |table|
  rows = table.rows_hash
  @request_id = SecureRandom.uuid
  opts = Standort.http_options(
    headers: {
      request_id: @request_id, user_agent: 'Standort-ruby-client/2.0 HTTP/1.0',
      content_type: :json, accept: :json
    }
  )

  params = {}

  if rows['method'] == 'params'
    params[:ip] = rows['ip'] if rows['ip']
    params[:point] = [rows['latitude'], rows['longitude']] if rows['latitude'] && rows['longitude']
  end

  if rows['method'] == 'headers'
    opts[:headers]['x-forwarded-for'] = rows['ip'] if rows['ip']
    opts[:headers]['geolocation'] = "geo:#{rows['latitude']},#{rows['longitude']}" if rows['latitude'] && rows['longitude']
  end

  @response = Standort::V2.http.get_location(params, opts)
end

When('I lookup locations with HTTP:') do |table|
  @request_id = SecureRandom.uuid
  opts = Standort.http_options(
    headers: {
      request_id: @request_id, user_agent: 'Standort-ruby-client/2.0 HTTP/1.0',
      content_type: 'application/pbjson', accept: 'application/pbjson'
    }
  )

  @lookup_positions = {}
  lookups = table.hashes.each_with_index.map do |row, lookup_position|
    lookup_label = row.fetch('lookup')
    raise "duplicate lookup label: #{lookup_label}" if @lookup_positions.key?(lookup_label)

    @lookup_positions[lookup_label] = lookup_position
    params = {}
    params[:ip] = row['ip'] unless row['ip'].empty?
    params[:point] = [row['latitude'], row['longitude']] unless row['latitude'].empty? || row['longitude'].empty?
    params
  end

  @response = Standort::V2.http.lookup_locations(lookups, opts)
end

When('I lookup a location using metadata with HTTP:') do |table|
  @request_id = SecureRandom.uuid
  rows = table.rows_hash
  opts = Standort.http_options(
    headers: {
      request_id: @request_id, user_agent: 'Standort-ruby-client/2.0 HTTP/1.0',
      content_type: 'application/pbjson', accept: 'application/pbjson',
      'x-forwarded-for': rows['ip'], geolocation: rows['geolocation']
    }
  )

  @lookup_positions = { rows.fetch('lookup') => 0 }
  @response = Standort::V2.http.lookup_locations([{}], opts)
end

When('I lookup {int} locations with HTTP') do |count|
  opts = Standort.http_options(
    headers: {
      user_agent: 'Standort-ruby-client/2.0 HTTP/1.0',
      content_type: :json, accept: :json
    }
  )
  lookups = Array.new(count) { { ip: '95.91.246.242' } }

  @response = Standort::V2.http.lookup_locations(lookups, opts)
end

When('I request lookup assets with HTTP') do
  @request_id = SecureRandom.uuid
  opts = Standort.http_options(
    headers: {
      request_id: @request_id, user_agent: 'Standort-ruby-client/2.0 HTTP/1.0',
      content_type: :json, accept: :json
    }
  )

  @response = Standort::V2.http.get_lookup_assets(opts)
end

Then('I should receive a valid locations with HTTP:') do |table|
  expect(@response.code).to eq(200)

  resp = JSON.parse(@response.body)

  expect(resp.fetch('meta').fetch('requestId')).to eq(@request_id)

  rows = table.rows_hash
  other = rows['kind'] == 'ip' ? 'geo' : 'ip'
  location = resp.fetch(rows['kind'])

  expect(resp[other]).to be_nil

  expect(location['country']).to eq(rows['country'])
  expect(location['continent']).to eq(rows['continent'])
end

Then('I should receive valid locations with HTTP:') do |table|
  expect(@response.code).to eq(200)

  resp = JSON.parse(@response.body)

  expect(resp.fetch('meta').fetch('requestId')).to eq(@request_id)

  table.hashes.each do |row|
    location = resp.fetch(row['kind'])

    expect(location['country']).to eq(row['country'])
    expect(location['continent']).to eq(row['continent'])
  end
end

Then('I should receive batch locations with HTTP:') do |table|
  expect(@response.code).to eq(200)

  resp = JSON.parse(@response.body)

  expect(resp.fetch('meta').fetch('requestId')).to eq(@request_id)

  table.hashes.each do |row|
    lookup = resp.fetch('lookups').fetch(@lookup_positions.fetch(row['lookup']))
    location = case row['kind']
               when 'ip'
                 expect(lookup['status']).to be_nil
                 expect(lookup.fetch('locations')['geo']).to be_nil
                 lookup.fetch('locations').fetch('ip')
               when 'geo'
                 expect(lookup['status']).to be_nil
                 expect(lookup.fetch('locations')['ip']).to be_nil
                 lookup.fetch('locations').fetch('geo')
               when 'none'
                 expect(lookup['locations']).to be_nil
                 expect(lookup.fetch('status').fetch('code')).to eq(row['code'].to_i)
                 next
               else
                 raise "unsupported location kind: #{row['kind']}"
               end

    expect(location['country']).to eq(row['country'])
    expect(location['continent']).to eq(row['continent'])
  end
end

Then('I should receive batch diagnostics with HTTP:') do |table|
  expect(@response.code).to eq(200)

  resp = JSON.parse(@response.body)

  expect(resp.fetch('meta').fetch('requestId')).to eq(@request_id)

  table.hashes.group_by { |row| row['lookup'] }.each do |lookup_label, rows|
    lookup = resp.fetch('lookups').fetch(@lookup_positions.fetch(lookup_label))
    status = lookup.fetch('status')
    detail = status.fetch('details').fetch(0)
    expected_metadata = rows.to_h { |row| [row['diagnostic'], row['code']] }

    expect(lookup['locations']).to be_nil
    expect(status.fetch('code')).to eq(5)
    expect(status.fetch('message')).to eq('not found')
    expect(status.fetch('details').length).to eq(1)
    expect(detail.fetch('@type')).to eq('type.googleapis.com/google.rpc.ErrorInfo')
    expect(detail.fetch('reason')).to eq('LOCATION_LOOKUP_FAILED')
    expect(detail.fetch('domain')).to eq('standort.v2')
    expect(detail.fetch('metadata')).to eq(expected_metadata)
  end
end

Then('I should receive a bad request response with HTTP') do
  expect(@response.code).to eq(400)
end

Then('I should receive lookup assets with HTTP:') do |table|
  expect(@response.code).to eq(200)

  resp = JSON.parse(@response.body)

  expect(resp.fetch('meta').fetch('requestId')).to eq(@request_id)

  table.hashes.each do |row|
    asset = resp.fetch('assets').find { |lookup_asset| lookup_asset.fetch('name') == row['name'] }

    expect(asset).not_to be_nil
    expect(asset.fetch('size_bytes').to_i).to be_positive
    expect(asset.fetch('checksum_algorithm')).to eq(row['checksum_algorithm'])
    expect(asset.fetch('checksum')).not_to be_empty
  end
end

Then('I should receive a not found response with HTTP:') do |table|
  expect(@response.code).to eq(404)

  rows = table.rows_hash

  expect(response_header(@response, rows['diagnostic'])).to eq(rows['code'])
end

Then('I should receive a partial location with HTTP:') do |table|
  expect(@response.code).to eq(200)

  resp = JSON.parse(@response.body)
  rows = table.rows_hash
  other = rows['kind'] == 'ip' ? 'geo' : 'ip'
  location = resp.fetch(rows['kind'])

  expect(resp[other]).to be_nil
  expect(resp.fetch('meta').fetch('requestId')).to eq(@request_id)
  expect(location['country']).to eq(rows['country'])
  expect(location['continent']).to eq(rows['continent'])
end

def response_header(response, key)
  response.headers[key.tr('-', '_').to_sym] || response.headers[key]
end
