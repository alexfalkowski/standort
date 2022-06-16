Feature: Server

  Server allows users to get locations by different types.

  @manual
  Scenario Outline: Get location by a valid IP address.
    Given I have "<source>" as the config file
    And I start the system
    When I request a location with gRPC:
      | ip     | <ip>     |
      | method | <method> |
    Then I should receive a valid locations with gRPC:
      | kind      | <kind>      |
      | country   | <country>   |
      | continent | <continent> |

    Examples: With ip2location parameters
      | source      | method | ip             | country | continent | kind    |
      | ip2location | params | 95.91.246.242  | DE      | EU        | KIND_IP |
      | ip2location | params | 45.128.199.236 | NL      | EU        | KIND_IP |
      | ip2location | params | 154.6.22.65    | US      | NA        | KIND_IP |

    Examples: With ip2location metadata
      | source      | method   | ip             | country | continent | kind    |
      | ip2location | metadata | 95.91.246.242  | DE      | EU        | KIND_IP |
      | ip2location | metadata | 45.128.199.236 | NL      | EU        | KIND_IP |
      | ip2location | metadata | 154.6.22.65    | US      | NA        | KIND_IP |

    Examples: With geoip2 parameters
      | source | method | ip             | country | continent | kind    |
      | geoip2 | params | 95.91.246.242  | DE      | EU        | KIND_IP |
      | geoip2 | params | 45.128.199.236 | NL      | EU        | KIND_IP |
      | geoip2 | params | 154.6.22.65    | US      | NA        | KIND_IP |

    Examples: With geoip2 metadata
      | source | method   | ip             | country | continent | kind    |
      | geoip2 | metadata | 95.91.246.242  | DE      | EU        | KIND_IP |
      | geoip2 | metadata | 45.128.199.236 | NL      | EU        | KIND_IP |
      | geoip2 | metadata | 154.6.22.65    | US      | NA        | KIND_IP |

  @manual
  Scenario Outline: Get location by a not found IP address.
    Given I have "<source>" as the config file
    And I start the system
    When I request a location with gRPC:
      | ip     | <ip>     |
      | method | <method> |
    Then I should receive an empty response with gRPC

    Examples: With ip2location parameters
      | source      | method | ip      |
      | ip2location | params | 0.0.0.0 |
      | ip2location | params | test    |
      | ip2location | params | <test>  |
      | ip2location | params | 154.6   |

    Examples: With ip2location metadata
      | source      | method   | ip      |
      | ip2location | metadata | 0.0.0.0 |
      | ip2location | metadata | test    |
      | ip2location | metadata | <test>  |
      | ip2location | metadata | 154.6   |

    Examples: With geoip2 parameters
      | source | method | ip      |
      | geoip2 | params | 0.0.0.0 |
      | geoip2 | params | test    |
      | geoip2 | params | <test>  |
      | geoip2 | params | 154.6   |

    Examples: With geoip2 metadata
      | source | method   | ip      |
      | geoip2 | metadata | 0.0.0.0 |
      | geoip2 | metadata | test    |
      | geoip2 | metadata | <test>  |
      | geoip2 | metadata | 154.6   |

  @startup
  Scenario Outline: Get location by a valid latitude and longitude.
    When I request a location with gRPC:
      | latitude  | <latitude>  |
      | longitude | <longitude> |
      | method    | <method>    |
    Then I should receive a valid locations with gRPC:
      | kind      | <kind>      |
      | country   | <country>   |
      | continent | <continent> |

    Examples: With parameters
      | method | latitude  | longitude  | country | continent | kind     |
      | params | 52.520008 | 13.404954  | DE      | EU        | KIND_GEO |
      | params | 52.377956 | 4.897070   | NL      | EU        | KIND_GEO |
      | params | 43.000000 | -75.000000 | US      | NA        | KIND_GEO |

    Examples: With metadata
      | method   | latitude  | longitude  | country | continent | kind     |
      | metadata | 52.520008 | 13.404954  | DE      | EU        | KIND_GEO |
      | metadata | 52.377956 | 4.897070   | NL      | EU        | KIND_GEO |
      | metadata | 43.000000 | -75.000000 | US      | NA        | KIND_GEO |

  @startup
  Scenario Outline: Get location by a bad latitude and longitude.
    When I request a location with gRPC:
      | latitude  | <latitude>  |
      | longitude | <longitude> |
      | method    | <method>    |
    Then I should receive an empty response with gRPC

    Examples: With parameters
      | method | latitude | longitude |
      | params | 91       | 10        |
      | params | 10       | 181       |

    Examples: With metadata
      | method   | latitude | longitude |
      | metadata | 91       | 10        |
      | metadata | 10       | 181       |
      | metadata | test     | 180       |
      | metadata | 90       | test      |

  @startup
  Scenario Outline: Get location by a not found latitude and longitude.
    When I request a location with gRPC:
      | latitude  | <latitude>  |
      | longitude | <longitude> |
      | method    | <method>    |
    Then I should receive an empty response with gRPC

    Examples:
      | method   | latitude | longitude |
      | params   | 90       | 180       |
      | metadata | 90       | 180       |
