@http
Feature: HTTP API
  These endpoints allows users to get locations by different types.

  Scenario Outline: Get location by a valid IP address.
    When I request a location with HTTP:
      | ip     | <ip>     |
      | method | <method> |
    Then I should receive a valid locations with HTTP:
      | kind      | <kind>      |
      | country   | <country>   |
      | continent | <continent> |

    Examples: With geoip2 parameters
      | source | method | kind | ip             | country | continent |
      | geoip2 | params | ip   |  95.91.246.242 | DE      | EU        |
      | geoip2 | params | ip   | 45.128.199.236 | NL      | EU        |
      | geoip2 | params | ip   |    154.6.22.65 | US      | NA        |
      | geoip2 | params | ip   |       1.40.0.0 | AU      | OC        |

    Examples: With geoip2 headers
      | source | method  | kind | ip             | country | continent |
      | geoip2 | headers | ip   |  95.91.246.242 | DE      | EU        |
      | geoip2 | headers | ip   | 45.128.199.236 | NL      | EU        |
      | geoip2 | headers | ip   |    154.6.22.65 | US      | NA        |
      | geoip2 | headers | ip   |       1.40.0.0 | AU      | OC        |

  Scenario Outline: Get location by an bad IP address.
    When I request a location with HTTP:
      | ip     | <ip>     |
      | method | <method> |
    Then I should receive a not found response with HTTP

    Examples: With geoip2 parameters
      | source | method | ip        |
      | geoip2 | params | 0.0.0.0   |
      | geoip2 | params | test      |
      | geoip2 | params | <test>    |
      | geoip2 | params | 154.6     |
      | geoip2 | params | 192.0.2.1 |

    Examples: With geoip2 headers
      | source | method  | ip        |
      | geoip2 | headers | 0.0.0.0   |
      | geoip2 | headers | test      |
      | geoip2 | headers | <test>    |
      | geoip2 | headers | 154.6     |
      | geoip2 | headers | 192.0.2.1 |

  Scenario Outline: Get location by a valid latitude and longitude.
    When I request a location with HTTP:
      | latitude  | <latitude>  |
      | longitude | <longitude> |
      | method    | <method>    |
    Then I should receive a valid locations with HTTP:
      | kind      | <kind>      |
      | country   | <country>   |
      | continent | <continent> |

    Examples: With parameters
      | method | kind | latitude  | longitude  | country | continent |
      | params | geo  | 52.520008 |  13.404954 | DE      | EU        |
      | params | geo  | 52.377956 |   4.897070 | NL      | EU        |
      | params | geo  | 43.000000 | -75.000000 | US      | NA        |

    Examples: With headers
      | method  | kind | latitude  | longitude  | country | continent |
      | headers | geo  | 52.520008 |  13.404954 | DE      | EU        |
      | headers | geo  | 52.377956 |   4.897070 | NL      | EU        |
      | headers | geo  | 43.000000 | -75.000000 | US      | NA        |

  Scenario Outline: Get location by a valid IP address and latitude and longitude.
    When I request a location with HTTP:
      | ip        | <ip>        |
      | latitude  | <latitude>  |
      | longitude | <longitude> |
      | method    | <method>    |
    Then I should receive valid locations with HTTP:
      | kind | country       | continent       |
      | ip   | <ip_country>  | <ip_continent>  |
      | geo  | <geo_country> | <geo_continent> |

    Examples: With parameters
      | method | ip            | latitude  | longitude | ip_country | ip_continent | geo_country | geo_continent |
      | params | 95.91.246.242 | 52.377956 |  4.897070 | DE         | EU           | NL          | EU            |

  Scenario Outline: Get location by partially valid inputs.
    When I request a location with HTTP:
      | ip        | <ip>        |
      | latitude  | <latitude>  |
      | longitude | <longitude> |
      | method    | <method>    |
    Then I should receive a partial location with HTTP:
      | kind      | <kind>      |
      | country   | <country>   |
      | continent | <continent> |
      | error     | <error>     |

    Examples: With parameters
      | method | kind | ip            | latitude  | longitude | country | continent | error               |
      | params | geo  |       0.0.0.0 | 52.377956 |  4.897070 | NL      | EU        | locationIpError     |
      | params | ip   | 95.91.246.242 |        91 |        10 | DE      | EU        | locationLatLngError |

    Examples: With headers
      | method  | kind | ip            | latitude | longitude | country | continent | error              |
      | headers | ip   | 95.91.246.242 | test     |       180 | DE      | EU        | locationPointError |

  Scenario Outline: Get location by a bad latitude and longitude.
    When I request a location with HTTP:
      | latitude  | <latitude>  |
      | longitude | <longitude> |
      | method    | <method>    |
    Then I should receive a not found response with HTTP

    Examples: With parameters
      | method | latitude | longitude |
      | params |       91 |        10 |
      | params |       10 |       181 |

    Examples: With headers
      | method  | latitude | longitude |
      | headers |       91 |        10 |
      | headers |       10 |       181 |
      | headers | test     |       180 |
      | headers |       90 | test      |

  Scenario Outline: Get location by a not found latitude and longitude.
    When I request a location with HTTP:
      | latitude  | <latitude>  |
      | longitude | <longitude> |
      | method    | <method>    |
    Then I should receive a not found response with HTTP

    Examples:
      | method  | latitude   | longitude |
      | params  |         90 |       180 |
      | headers |         90 |       180 |
      | params  | -49.303721 | 69.122136 |
      | headers | -49.303721 | 69.122136 |
