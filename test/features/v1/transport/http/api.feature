@http
Feature: HTTP API
  These endpoints allows users to get locations by different means.

  Scenario Outline: Get location by a valid IP address.
    When I request a location by IP address with HTTP:
      | ip | <ip> |
    Then I should receive a valid location by IP adress with HTTP:
      | country   | <country>   |
      | continent | <continent> |

    Examples: With geoip2
      | source | ip             | country | continent |
      | geoip2 |  95.91.246.242 | DE      | EU        |
      | geoip2 | 45.128.199.236 | NL      | EU        |
      | geoip2 |    154.6.22.65 | US      | NA        |

  Scenario Outline: Get location by a not found IP address.
    When I request a location by IP address with HTTP:
      | ip | <ip> |
    Then I should receive a not found response with HTTP

    Examples: With geoip2
      | source | ip      |
      | geoip2 | 0.0.0.0 |
      | geoip2 | test    |
      | geoip2 | <test>  |
      | geoip2 |   154.6 |
      | geoip2 |         |

  Scenario Outline: Get location by a valid latitude and longitude.
    When I request a location by latitude and longitude with HTTP:
      | latitude  | <latitude>  |
      | longitude | <longitude> |
    Then I should receive a valid location by latitude and longitude with HTTP:
      | country   | <country>   |
      | continent | <continent> |

    Examples:
      | latitude  | longitude  | country | continent |
      | 52.520008 |  13.404954 | DE      | EU        |
      | 52.377956 |   4.897070 | NL      | EU        |
      | 43.000000 | -75.000000 | US      | NA        |

  Scenario Outline: Get location by a not found latitude and longitude.
    When I request a location by latitude and longitude with HTTP:
      | latitude  | <latitude>  |
      | longitude | <longitude> |
    Then I should receive a not found response with HTTP

    Examples:
      | latitude | longitude |
      |       90 |       180 |
      |       91 |        10 |
      |       10 |       181 |
