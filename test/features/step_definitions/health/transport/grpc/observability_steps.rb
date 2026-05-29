# frozen_string_literal: true

When('the system requests the {string} health status with gRPC') do |service|
  request = Grpc::Health::V1::HealthCheckRequest.new(service:)
  @response = Standort.health_grpc.check(request)
end

Then('the system should respond with a healthy status with gRPC') do
  expect(@response.status).to eq(:SERVING)
end
