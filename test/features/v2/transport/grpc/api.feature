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
      | source | method | ip                                      | country | continent |
      | geoip2 | params |                           95.91.246.242 | DE      | EU        |
      | geoip2 | params |                          45.128.199.236 | NL      | EU        |
      | geoip2 | params |                             154.6.22.65 | US      | NA        |
      | geoip2 | params | 2a02:8109:9f2e:4600:861e:b845:8bd4:b047 | DE      | EU        |

    Examples: With geoip2 metadata
      | source | method   | ip             | country | continent |
      | geoip2 | metadata |  95.91.246.242 | DE      | EU        |
      | geoip2 | metadata | 45.128.199.236 | NL      | EU        |
      | geoip2 | metadata |    154.6.22.65 | US      | NA        |

  Scenario Outline: Get location by a not found IP address.
    When I request a location with gRPC:
      | ip     | <ip>     |
      | method | <method> |
    Then I should receive a not found response with gRPC

    Examples: With geoip2 parameters
      | source | method | ip      |
      | geoip2 | params | 0.0.0.0 |
      | geoip2 | params | test    |
      | geoip2 | params | <test>  |
      | geoip2 | params |   154.6 |

    Examples: With geoip2 metadata
      | source | method   | ip      |
      | geoip2 | metadata | 0.0.0.0 |
      | geoip2 | metadata | test    |
      | geoip2 | metadata | <test>  |
      | geoip2 | metadata |   154.6 |

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
      | method | latitude  | longitude  | country | continent |
      | params | 52.520008 |  13.404954 | DE      | EU        |
      | params | 52.377956 |   4.897070 | NL      | EU        |
      | params | 43.000000 | -75.000000 | US      | NA        |

    Examples: With metadata
      | method   | latitude  | longitude  | country | continent |
      | metadata | 52.520008 |  13.404954 | DE      | EU        |
      | metadata | 52.377956 |   4.897070 | NL      | EU        |
      | metadata | 43.000000 | -75.000000 | US      | NA        |

  Scenario Outline: Get location by a bad latitude and longitude.
    When I request a location with gRPC:
      | latitude  | <latitude>  |
      | longitude | <longitude> |
      | method    | <method>    |
    Then I should receive a not found response with gRPC

    Examples: With parameters
      | method | latitude | longitude |
      | params |       91 |        10 |
      | params |       10 |       181 |

    Examples: With metadata
      | method   | latitude | longitude |
      | metadata |       91 |        10 |
      | metadata |       10 |       181 |
      | metadata | test     |       180 |
      | metadata |       90 | test      |

  Scenario Outline: Get location by a not found latitude and longitude.
    When I request a location with gRPC:
      | latitude  | <latitude>  |
      | longitude | <longitude> |
      | method    | <method>    |
    Then I should receive a not found response with gRPC

    Examples:
      | method   | latitude | longitude |
      | params   |       90 |       180 |
      | metadata |       90 |       180 |
