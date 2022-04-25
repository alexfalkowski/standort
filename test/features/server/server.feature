@manual
Feature: Server

  Server allows users to get locations by different types.

  Scenario: Starting server with an invalid ip path
    When the server is configured with an invalid ip path
    Then starting the system should raise an error
    And the server is configured with a valid configuration
