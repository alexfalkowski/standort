# frozen_string_literal: true

When('I request a location with HTTP:') do |table|
  rows = table.rows_hash
  @request_id = SecureRandom.uuid
  opts = {
    headers: {
      request_id: @request_id, user_agent: 'Standort-ruby-client/2.0 HTTP/1.0',
      content_type: :json, accept: :json
    },
    read_timeout: 10, open_timeout: 10
  }

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

Then('I should receive a partial location with HTTP:') do |table|
  expect(@response.code).to eq(200)

  resp = JSON.parse(@response.body)
  rows = table.rows_hash
  other = rows['kind'] == 'ip' ? 'geo' : 'ip'
  location = resp.fetch(rows['kind'])

  expect(resp[other]).to be_nil
  expect(resp.fetch('meta')).to include(rows['error'])
  expect(resp.fetch('meta').fetch('requestId')).to eq(@request_id)
  expect(location['country']).to eq(rows['country'])
  expect(location['continent']).to eq(rows['continent'])
end
