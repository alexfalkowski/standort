@grpc
Feature: Observability

  Observability is a measure of how well internal states of a system can be inferred by knowledge of its external outputs.

  Scenario Outline: Health with gRPC
    When the system requests the "<service>" health status with gRPC
    Then the system should respond with a healthy status with gRPC

    Examples:
      | service             |
      | standort.v1.Service |
      | standort.v2.Service |
