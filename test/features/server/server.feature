@manual @grpc
Feature: Server

  Server allows users to get locations by different types.

  Scenario: Starting server with an invalid ip2location path
    Given the server is configured with an invalid "ip2location" path
    When starting the system should raise an error
    Then I should see a log entry of "open /assets/ip2location.bin: no such file or directory" in the file "reports/server.log"
    And the server is configured with a valid configuration

  Scenario: Starting server with an invalid geoip2 path
    Given the server is configured with an invalid "geoip2" path
    When starting the system should raise an error
    Then I should see a log entry of "open /assets/geoip2.mmdb: no such file or directory" in the file "reports/server.log"
    And the server is configured with a valid configuration

  Scenario: Starting server with an invalid location path
    Given the server is configured with an invalid "location" path
    When starting the system should raise an error
    Then I should see a log entry of "open /assets/africa.geojson: no such file or directory" in the file "reports/server.log"
    And the server is configured with a valid configuration

  Scenario: Starting server with an invalid ip provider
    Given the server is configured with an invalid ip provider
    When starting the system should raise an error
    And I should see a log entry of "no provider configured" in the file "reports/server.log"
    And the server is configured with a valid configuration
