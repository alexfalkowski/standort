Feature: Server

  Server allows users to get locations by different types.

  @manual
  Scenario Outline: Get location by a valid IP address.
    Given I have "<source>" as the config file
    And I start the system
    When I request a location with HTTP:
      | ip     | <ip>     |
      | method | <method> |
    Then I should receive a valid locations with HTTP:
      | type      | <type>      |
      | country   | <country>   |
      | continent | <continent> |
    And the process 'server' should consume less than '40mb' of memory

    Examples: With ip2location parameters
      | source      | method | ip             | country | continent | type    |
      | ip2location | params | 95.91.246.242  | DE      | EU        | TYPE_IP |
      | ip2location | params | 45.128.199.236 | NL      | EU        | TYPE_IP |
      | ip2location | params | 154.6.22.65    | US      | NA        | TYPE_IP |

    Examples: With ip2location headers
      | source      | method  | ip             | country | continent | type    |
      | ip2location | headers | 95.91.246.242  | DE      | EU        | TYPE_IP |
      | ip2location | headers | 45.128.199.236 | NL      | EU        | TYPE_IP |
      | ip2location | headers | 154.6.22.65    | US      | NA        | TYPE_IP |

    Examples: With geoip2 parameters
      | source | method | ip             | country | continent | type    |
      | geoip2 | params | 95.91.246.242  | DE      | EU        | TYPE_IP |
      | geoip2 | params | 45.128.199.236 | NL      | EU        | TYPE_IP |
      | geoip2 | params | 154.6.22.65    | US      | NA        | TYPE_IP |

    Examples: With geoip2 headers
      | source | method  | ip             | country | continent | type    |
      | geoip2 | headers | 95.91.246.242  | DE      | EU        | TYPE_IP |
      | geoip2 | headers | 45.128.199.236 | NL      | EU        | TYPE_IP |
      | geoip2 | headers | 154.6.22.65    | US      | NA        | TYPE_IP |

  @manual
  Scenario Outline: Get location by an bad IP address.
    Given I have "<source>" as the config file
    And I start the system
    When I request a location with HTTP:
      | ip     | <ip>     |
      | method | <method> |
    Then I should receive an empty response with HTTP
    And the process 'server' should consume less than '40mb' of memory

    Examples: With ip2location parameters
      | source      | method | ip      |
      | ip2location | params | 0.0.0.0 |
      | ip2location | params | test    |
      | ip2location | params | <test>  |
      | ip2location | params | 154.6   |

    Examples: With ip2location headers
      | source      | method  | ip      |
      | ip2location | headers | 0.0.0.0 |
      | ip2location | headers | test    |
      | ip2location | headers | <test>  |
      | ip2location | headers | 154.6   |

    Examples: With geoip2 parameters
      | source | method | ip      |
      | geoip2 | params | 0.0.0.0 |
      | geoip2 | params | test    |
      | geoip2 | params | <test>  |
      | geoip2 | params | 154.6   |

    Examples: With ip2location headers
      | source | method  | ip      |
      | geoip2 | headers | 0.0.0.0 |
      | geoip2 | headers | test    |
      | geoip2 | headers | <test>  |
      | geoip2 | headers | 154.6   |

  @startup
  Scenario Outline: Get location by a valid latitude and longitude.
    When I request a location with HTTP:
      | latitude  | <latitude>  |
      | longitude | <longitude> |
      | method    | <method>    |
    Then I should receive a valid locations with HTTP:
      | type      | <type>      |
      | country   | <country>   |
      | continent | <continent> |
    And the process 'server' should consume less than '40mb' of memory

    Examples: With parameters
      | method | latitude  | longitude  | country | continent | type     |
      | params | 52.520008 | 13.404954  | DE      | EU        | TYPE_GEO |
      | params | 52.377956 | 4.897070   | NL      | EU        | TYPE_GEO |
      | params | 43.000000 | -75.000000 | US      | NA        | TYPE_GEO |

    Examples: With headers
      | method  | latitude  | longitude  | country | continent | type     |
      | headers | 52.520008 | 13.404954  | DE      | EU        | TYPE_GEO |
      | headers | 52.377956 | 4.897070   | NL      | EU        | TYPE_GEO |
      | headers | 43.000000 | -75.000000 | US      | NA        | TYPE_GEO |

  @startup
  Scenario Outline: Get location by a bad latitude and longitude.
    When I request a location with HTTP:
      | latitude  | <latitude>  |
      | longitude | <longitude> |
      | method    | <method>    |
    Then I should receive an empty response with HTTP
    And the process 'server' should consume less than '40mb' of memory

    Examples: With parameters
      | method | latitude | longitude |
      | params | 91       | 10        |
      | params | 10       | 181       |

    Examples: With headers
      | method  | latitude | longitude |
      | headers | 91       | 10        |
      | headers | 10       | 181       |
      | headers | test     | 180       |
      | headers | 90       | test      |

  @startup
  Scenario Outline: Get location by a not found latitude and longitude.
    When I request a location with HTTP:
      | latitude  | <latitude>  |
      | longitude | <longitude> |
      | method    | <method>    |
    Then I should receive an empty response with HTTP
    And the process 'server' should consume less than '40mb' of memory

    Examples:
      | method  | latitude | longitude |
      | params  | 90       | 180       |
      | headers | 90       | 180       |
