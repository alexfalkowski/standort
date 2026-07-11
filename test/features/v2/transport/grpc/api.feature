@grpc
Feature: gRPC API
  These endpoints allows users to get locations by different types.

  Scenario Outline: Get location by a valid IP address.
    When I request a location with gRPC:
      | ip     | <ip>     |
      | method | <method> |
    Then I should receive a valid locations with gRPC:
      | kind      | <kind>      |
      | country   | <country>   |
      | continent | <continent> |

    Examples: With geoip2 parameters
      | source | method | kind | ip                                      | country | continent |
      | geoip2 | params | ip   | 95.91.246.242                           | DE      | EU        |
      | geoip2 | params | ip   | 45.128.199.236                          | NL      | EU        |
      | geoip2 | params | ip   | 154.6.22.65                             | US      | NA        |
      | geoip2 | params | ip   | 1.40.0.0                                | AU      | OC        |
      | geoip2 | params | ip   | 2a02:8109:9f2e:4600:861e:b845:8bd4:b047 | DE      | EU        |

    Examples: With geoip2 metadata
      | source | method   | kind | ip             | country | continent |
      | geoip2 | metadata | ip   | 95.91.246.242  | DE      | EU        |
      | geoip2 | metadata | ip   | 45.128.199.236 | NL      | EU        |
      | geoip2 | metadata | ip   | 154.6.22.65    | US      | NA        |

  Scenario Outline: Get location by a not found IP address.
    When I request a location with gRPC:
      | ip     | <ip>     |
      | method | <method> |
    Then I should receive a not found response with gRPC:
      | diagnostic | <diagnostic> |
      | code       | <code>       |

    Examples: With geoip2 parameters
      | source | method | ip           | diagnostic        | code       |
      | geoip2 | params | 0.0.0.0      | location-ip-error | not_found  |
      | geoip2 | params | test         | location-ip-error | invalid_ip |
      | geoip2 | params | <test>       | location-ip-error | invalid_ip |
      | geoip2 | params | 154.6        | location-ip-error | invalid_ip |
      | geoip2 | params | 2001:db8::zz | location-ip-error | invalid_ip |
      | geoip2 | params | 192.0.2.1    | location-ip-error | not_found  |

    Examples: With geoip2 metadata
      | source | method   | ip           | diagnostic        | code       |
      | geoip2 | metadata | 0.0.0.0      | location-ip-error | not_found  |
      | geoip2 | metadata | test         | location-ip-error | invalid_ip |
      | geoip2 | metadata | <test>       | location-ip-error | invalid_ip |
      | geoip2 | metadata | 154.6        | location-ip-error | invalid_ip |
      | geoip2 | metadata | 2001:db8::zz | location-ip-error | invalid_ip |
      | geoip2 | metadata | 192.0.2.1    | location-ip-error | not_found  |

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
      | method | kind | latitude  | longitude  | country | continent |
      | params | geo  | 52.520008 | 13.404954  | DE      | EU        |
      | params | geo  | 52.377956 | 4.897070   | NL      | EU        |
      | params | geo  | 43.000000 | -75.000000 | US      | NA        |

    Examples: With metadata
      | method   | kind | latitude  | longitude  | country | continent |
      | metadata | geo  | 52.520008 | 13.404954  | DE      | EU        |
      | metadata | geo  | 52.377956 | 4.897070   | NL      | EU        |
      | metadata | geo  | 43.000000 | -75.000000 | US      | NA        |

  Scenario Outline: Get location by a valid IP address and latitude and longitude.
    When I request a location with gRPC:
      | ip        | <ip>        |
      | latitude  | <latitude>  |
      | longitude | <longitude> |
      | method    | <method>    |
    Then I should receive valid locations with gRPC:
      | kind | country       | continent       |
      | ip   | <ip_country>  | <ip_continent>  |
      | geo  | <geo_country> | <geo_continent> |

    Examples: With parameters
      | method | ip            | latitude  | longitude | ip_country | ip_continent | geo_country | geo_continent |
      | params | 95.91.246.242 | 52.377956 | 4.897070  | DE         | EU           | NL          | EU            |

  Scenario: Lookup locations in a batch with gRPC.
    When I lookup locations with gRPC:
      | lookup             | kind | ip            | latitude  | longitude |
      | German IP address  | ip   | 95.91.246.242 |           |           |
      | Netherlands point  | geo  |               | 52.377956 | 4.897070  |
      | unknown IP address | none | 192.0.2.1     |           |           |
    Then I should receive batch locations with gRPC:
      | lookup             | kind | country | continent | code |
      | German IP address  | ip   | DE      | EU        |      |
      | Netherlands point  | geo  | NL      | EU        |      |
      | unknown IP address | none |         |           | 5    |

  Scenario: Lookup location failure diagnostics in a batch with gRPC.
    When I lookup locations with gRPC:
      | lookup             | kind | ip        | latitude | longitude |
      | unknown IP address | none | 192.0.2.1 |          |           |
      | invalid point      | none | 154.6     | 91       | 10        |
    Then I should receive batch diagnostics with gRPC:
      | lookup             | diagnostic             | code          |
      | unknown IP address | location-ip-error      | not_found     |
      | invalid point      | location-ip-error      | invalid_ip    |
      | invalid point      | location-lat-lng-error | invalid_point |

  Scenario: Lookup location metadata failure diagnostics in a batch with gRPC.
    When I lookup a location using metadata with gRPC:
      | lookup      | metadata lookup |
      | ip          | 192.0.2.1       |
      | geolocation | geo:test,180    |
    Then I should receive batch diagnostics with gRPC:
      | lookup          | diagnostic           | code            |
      | metadata lookup | location-ip-error    | not_found       |
      | metadata lookup | location-point-error | invalid_geo_uri |

  Scenario: Lookup too many locations with gRPC.
    When I lookup 101 locations with gRPC
    Then I should receive an invalid argument response with gRPC

  Scenario: Get lookup assets with gRPC.
    When I request lookup assets with gRPC
    Then I should receive lookup assets with gRPC:
      | name          | checksum_algorithm |
      | geoip2.mmdb   | sha256             |
      | earth.geojson | sha256             |

  Scenario Outline: Get location by partially valid inputs.
    When I request a location with gRPC:
      | ip        | <ip>        |
      | latitude  | <latitude>  |
      | longitude | <longitude> |
      | method    | <method>    |
    Then I should receive a partial location with gRPC:
      | kind      | <kind>      |
      | country   | <country>   |
      | continent | <continent> |

    Examples: With parameters
      | method | kind | ip            | latitude  | longitude | country | continent |
      | params | geo  | 0.0.0.0       | 52.377956 | 4.897070  | NL      | EU        |
      | params | geo  | test          | 52.377956 | 4.897070  | NL      | EU        |
      | params | ip   | 95.91.246.242 | 91        | 10        | DE      | EU        |

    Examples: With metadata
      | method   | kind | ip            | latitude  | longitude | country | continent |
      | metadata | geo  | test          | 52.377956 | 4.897070  | NL      | EU        |
      | metadata | ip   | 95.91.246.242 | test      | 180       | DE      | EU        |

  Scenario Outline: Get location by a bad latitude and longitude.
    When I request a location with gRPC:
      | latitude  | <latitude>  |
      | longitude | <longitude> |
      | method    | <method>    |
    Then I should receive a not found response with gRPC:
      | diagnostic | <diagnostic> |
      | code       | <code>       |

    Examples: With parameters
      | method | latitude | longitude | diagnostic             | code          |
      | params | 91       | 10        | location-lat-lng-error | invalid_point |
      | params | 10       | 181       | location-lat-lng-error | invalid_point |
      | params | nan      | 10        | location-lat-lng-error | invalid_point |
      | params | 10       | inf       | location-lat-lng-error | invalid_point |

    Examples: With metadata
      | method   | latitude | longitude | diagnostic             | code            |
      | metadata | 91       | 10        | location-lat-lng-error | invalid_point   |
      | metadata | 10       | 181       | location-lat-lng-error | invalid_point   |
      | metadata | test     | 180       | location-point-error   | invalid_geo_uri |
      | metadata | 90       | test      | location-point-error   | invalid_geo_uri |

  Scenario Outline: Get location by a not found latitude and longitude.
    When I request a location with gRPC:
      | latitude  | <latitude>  |
      | longitude | <longitude> |
      | method    | <method>    |
    Then I should receive a not found response with gRPC:
      | diagnostic | <diagnostic> |
      | code       | <code>       |

    Examples:
      | method   | latitude   | longitude | diagnostic             | code      |
      | params   | 90         | 180       | location-lat-lng-error | not_found |
      | metadata | 90         | 180       | location-lat-lng-error | not_found |
      | params   | -49.303721 | 69.122136 | location-lat-lng-error | not_found |
      | metadata | -49.303721 | 69.122136 | location-lat-lng-error | not_found |
