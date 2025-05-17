# frozen_string_literal: true

When('I request a location with HTTP which performs in {int} ms') do |time|
  @request_id = SecureRandom.uuid
  metadata = { 'request-id' => @request_id }
  params = { 'ip' => '95.91.246.242' }
  request = Standort::V2::GetLocationRequest.new(params)

  expect { Standort::V2.grpc.get_location(request, { metadata: }) }.to perform_under(time).ms
end
