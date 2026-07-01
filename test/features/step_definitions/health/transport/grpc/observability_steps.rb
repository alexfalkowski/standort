# frozen_string_literal: true

When('the system requests the {string} health status with gRPC') do |service|
  @response = Standort.health_grpc(service).check(deadline: Time.now + 10)
end

Then('the system should respond with a healthy status with gRPC') do
  expect(@response.status).to eq(:SERVING)
end
