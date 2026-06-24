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
      | geoip2 | params | ip   |                           95.91.246.242 | DE      | EU        |
      | geoip2 | params | ip   |                          45.128.199.236 | NL      | EU        |
      | geoip2 | params | ip   |                             154.6.22.65 | US      | NA        |
      | geoip2 | params | ip   |                                1.40.0.0 | AU      | OC        |
      | geoip2 | params | ip   | 2a02:8109:9f2e:4600:861e:b845:8bd4:b047 | DE      | EU        |

    Examples: With geoip2 metadata
      | source | method   | kind | ip             | country | continent |
      | geoip2 | metadata | ip   |  95.91.246.242 | DE      | EU        |
      | geoip2 | metadata | ip   | 45.128.199.236 | NL      | EU        |
      | geoip2 | metadata | ip   |    154.6.22.65 | US      | NA        |

  Scenario Outline: Get location by a not found IP address.
    When I request a location with gRPC:
      | ip     | <ip>     |
      | method | <method> |
    Then I should receive a not found response with gRPC:
      | diagnostic | <diagnostic> |
      | code       | <code>       |

    Examples: With geoip2 parameters
      | source | method | ip        | diagnostic        | code      |
      | geoip2 | params | 0.0.0.0   | location-ip-error | not_found |
      | geoip2 | params | test      | location-ip-error | not_found |
      | geoip2 | params | <test>    | location-ip-error | not_found |
      | geoip2 | params | 154.6     | location-ip-error | not_found |
      | geoip2 | params | 192.0.2.1 | location-ip-error | not_found |

    Examples: With geoip2 metadata
      | source | method   | ip        | diagnostic        | code      |
      | geoip2 | metadata | 0.0.0.0   | location-ip-error | not_found |
      | geoip2 | metadata | test      | location-ip-error | not_found |
      | geoip2 | metadata | <test>    | location-ip-error | not_found |
      | geoip2 | metadata | 154.6     | location-ip-error | not_found |
      | geoip2 | metadata | 192.0.2.1 | location-ip-error | not_found |

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
      | params | geo  | 52.520008 |  13.404954 | DE      | EU        |
      | params | geo  | 52.377956 |   4.897070 | NL      | EU        |
      | params | geo  | 43.000000 | -75.000000 | US      | NA        |

    Examples: With metadata
      | method   | kind | latitude  | longitude  | country | continent |
      | metadata | geo  | 52.520008 |  13.404954 | DE      | EU        |
      | metadata | geo  | 52.377956 |   4.897070 | NL      | EU        |
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
      | params | 95.91.246.242 | 52.377956 |  4.897070 | DE         | EU           | NL          | EU            |

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
      | params | geo  | 0.0.0.0       | 52.377956 |  4.897070 | NL      | EU        |
      | params | ip   | 95.91.246.242 | 91        | 10        | DE      | EU        |

    Examples: With metadata
      | method   | kind | ip            | latitude | longitude | country | continent |
      | metadata | ip   | 95.91.246.242 | test     | 180       | DE      | EU        |

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
      | params |       91 |        10 | location-lat-lng-error | invalid_point |
      | params |       10 |       181 | location-lat-lng-error | invalid_point |
      | params |      nan |        10 | location-lat-lng-error | invalid_point |
      | params |       10 |       inf | location-lat-lng-error | invalid_point |

    Examples: With metadata
      | method   | latitude | longitude | diagnostic             | code            |
      | metadata |       91 |        10 | location-lat-lng-error | invalid_point   |
      | metadata |       10 |       181 | location-lat-lng-error | invalid_point   |
      | metadata | test     |       180 | location-point-error   | invalid_geo_uri |
      | metadata |       90 | test      | location-point-error   | invalid_geo_uri |

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
