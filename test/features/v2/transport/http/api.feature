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
    Then I should receive a not found response with HTTP:
      | diagnostic | <diagnostic> |
      | code       | <code>       |

    Examples: With geoip2 parameters
      | source | method | ip        | diagnostic        | code      |
      | geoip2 | params | 0.0.0.0   | location-ip-error | not_found |
      | geoip2 | params | test      | location-ip-error | not_found |
      | geoip2 | params | <test>    | location-ip-error | not_found |
      | geoip2 | params | 154.6     | location-ip-error | not_found |
      | geoip2 | params | 192.0.2.1 | location-ip-error | not_found |

    Examples: With geoip2 headers
      | source | method  | ip        | diagnostic        | code      |
      | geoip2 | headers | 0.0.0.0   | location-ip-error | not_found |
      | geoip2 | headers | test      | location-ip-error | not_found |
      | geoip2 | headers | <test>    | location-ip-error | not_found |
      | geoip2 | headers | 154.6     | location-ip-error | not_found |
      | geoip2 | headers | 192.0.2.1 | location-ip-error | not_found |

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

  Scenario: Lookup locations in a batch with HTTP.
    When I lookup locations with HTTP:
      | kind | ip            | latitude  | longitude |
      | ip   | 95.91.246.242 |           |           |
      | geo  |               | 52.377956 |  4.897070 |
      | none | 192.0.2.1     |           |           |
    Then I should receive batch locations with HTTP:
      | index | kind | country | continent | code |
      | 0     | ip   | DE      | EU        |      |
      | 1     | geo  | NL      | EU        |      |
      | 2     | none |         |           | 5    |

  Scenario: Lookup too many locations with HTTP.
    When I lookup 101 locations with HTTP
    Then I should receive a bad request response with HTTP

  Scenario: Get lookup assets with HTTP.
    When I request lookup assets with HTTP
    Then I should receive lookup assets with HTTP:
      | name          | checksum_algorithm |
      | geoip2.mmdb   | sha256             |
      | earth.geojson | sha256             |

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

    Examples: With parameters
      | method | kind | ip            | latitude  | longitude | country | continent |
      | params | geo  |       0.0.0.0 | 52.377956 |  4.897070 | NL      | EU        |
      | params | ip   | 95.91.246.242 |        91 |        10 | DE      | EU        |

    Examples: With headers
      | method  | kind | ip            | latitude | longitude | country | continent |
      | headers | ip   | 95.91.246.242 | test     |       180 | DE      | EU        |

  Scenario Outline: Get location by a bad latitude and longitude.
    When I request a location with HTTP:
      | latitude  | <latitude>  |
      | longitude | <longitude> |
      | method    | <method>    |
    Then I should receive a not found response with HTTP:
      | diagnostic | <diagnostic> |
      | code       | <code>       |

    Examples: With parameters
      | method | latitude | longitude | diagnostic             | code          |
      | params |       91 |        10 | location-lat-lng-error | invalid_point |
      | params |       10 |       181 | location-lat-lng-error | invalid_point |

    Examples: With headers
      | method  | latitude | longitude | diagnostic             | code            |
      | headers |       91 |        10 | location-lat-lng-error | invalid_point   |
      | headers |       10 |       181 | location-lat-lng-error | invalid_point   |
      | headers | test     |       180 | location-point-error   | invalid_geo_uri |
      | headers |       90 | test      | location-point-error   | invalid_geo_uri |

  Scenario Outline: Get location by a not found latitude and longitude.
    When I request a location with HTTP:
      | latitude  | <latitude>  |
      | longitude | <longitude> |
      | method    | <method>    |
    Then I should receive a not found response with HTTP:
      | diagnostic | <diagnostic> |
      | code       | <code>       |

    Examples:
      | method  | latitude   | longitude | diagnostic             | code      |
      | params  |         90 |       180 | location-lat-lng-error | not_found |
      | headers |         90 |       180 | location-lat-lng-error | not_found |
      | params  | -49.303721 | 69.122136 | location-lat-lng-error | not_found |
      | headers | -49.303721 | 69.122136 | location-lat-lng-error | not_found |
