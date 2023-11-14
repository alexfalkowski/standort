@grpc
Feature: Server

  Server allows users to get locations by different means.

  @manual
  Scenario Outline: Get location by a valid IP address.
    Given I have "<source>" as the config file
    And I start the system
    When I request a location by IP address with gRPC:
      | ip | <ip> |
    Then I should receive a valid location by IP adress with gRPC:
      | country   | <country>   |
      | continent | <continent> |

    Examples: With ip2location
      | source      | ip             | country | continent |
      | ip2location | 95.91.246.242  | DE      | EU        |
      | ip2location | 45.128.199.236 | NL      | EU        |
      | ip2location | 154.6.22.65    | US      | NA        |

    Examples: With geoip2
      | source | ip             | country | continent |
      | geoip2 | 95.91.246.242  | DE      | EU        |
      | geoip2 | 45.128.199.236 | NL      | EU        |
      | geoip2 | 154.6.22.65    | US      | NA        |

  @manual
  Scenario Outline: Get location by a not found IP address.
    Given I have "<source>" as the config file
    And I start the system
    When I request a location by IP address with gRPC:
      | ip | <ip> |
    Then I should receive a not found response with gRPC

    Examples: With ip2location
      | source      | ip      |
      | ip2location | 0.0.0.0 |
      | ip2location | test    |
      | ip2location | <test>  |
      | ip2location | 154.6   |

    Examples: With geoip2
      | source | ip      |
      | geoip2 | 0.0.0.0 |
      | geoip2 | test    |
      | geoip2 | <test>  |
      | geoip2 | 154.6   |

  @startup
  Scenario Outline: Get location by a valid latitude and longitude.
    When I request a location by latitude and longitude with gRPC:
      | latitude  | <latitude>  |
      | longitude | <longitude> |
    Then I should receive a valid location by latitude and longitude with gRPC:
      | country   | <country>   |
      | continent | <continent> |

    Examples:
      | latitude  | longitude  | country | continent |
      | 52.520008 | 13.404954  | DE      | EU        |
      | 52.377956 | 4.897070   | NL      | EU        |
      | 43.000000 | -75.000000 | US      | NA        |

  @startup
  Scenario Outline: Get location by a not found latitude and longitude.
    When I request a location by latitude and longitude with gRPC:
      | latitude  | <latitude>  |
      | longitude | <longitude> |
    Then I should receive a not found response with gRPC

    Examples:
      | latitude | longitude |
      | 90       | 180       |
      | 91       | 10        |
      | 10       | 181       |
