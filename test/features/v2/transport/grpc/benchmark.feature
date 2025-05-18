@benchmark
Feature: Benchmark gRPC API
  Make sure these endpoints perform at their best.

  Scenario: Get location in a good time frame and memory.
    When I request a location with gRPC which performs in 15 ms
    And the process 'server' should consume less than '65mb' of memory
