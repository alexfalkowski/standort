@manual
Feature: Server

  Server allows users to get locations by different types.

  Scenario: Starting server with an invalid ip2location path
    When the server is configured with an invalid "ip2location" path
    Then starting the system should raise an error
    And the server is configured with a valid configuration

  Scenario: Starting server with an invalid geoip2 path
    When the server is configured with an invalid "geoip2" path
    Then starting the system should raise an error
    And the server is configured with a valid configuration

  Scenario: Starting server with an invalid location path
    When the server is configured with an invalid "location" path
    Then starting the system should raise an error
    And the server is configured with a valid configuration

  Scenario: Starting server with an invalid ip provider
    When the server is configured with an invalid ip provider
    Then starting the system should raise an error
    And the server is configured with a valid configuration
