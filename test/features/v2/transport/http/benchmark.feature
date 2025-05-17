Feature: Benchmark HTTP API
  Make sure these endpoints perform at their best.

  Scenario: Get location in a good time frame and memory.
    When I request a location with HTTP which performs in 15 ms
    And the process 'server' should consume less than '150mb' of memory
