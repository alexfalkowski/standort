# frozen_string_literal: true

When('I request a location with HTTP:') do |table|
  rows = table.rows_hash
  opts = {
    headers: {
      request_id: SecureRandom.uuid, user_agent: 'Standort-ruby-client/2.0 HTTP/1.0',
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

  expect(resp['meta'].length).to be > 0
  expect(resp['locations'].length).to eq(1)

  rows = table.rows_hash
  location = resp['locations'][0]
  kind = location['kind'] == 1 ? 'ip' : 'geo'

  expect(location['country']).to eq(rows['country'])
  expect(location['continent']).to eq(rows['continent'])
  expect(kind).to eq(rows['kind'])
end
