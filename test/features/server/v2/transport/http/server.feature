@http
Feature: Server
  Server allows users to get locations by different types.

  Scenario Outline: Get location by a valid IP address.
    When I request a location with HTTP:
      | ip     | <ip>     |
      | method | <method> |
    Then I should receive a valid locations with HTTP:
      | kind      | <kind>      |
      | country   | <country>   |
      | continent | <continent> |

    Examples: With geoip2 parameters
      | source | method | ip             | country | continent | kind    |
      | geoip2 | params |  95.91.246.242 | DE      | EU        | KIND_IP |
      | geoip2 | params | 45.128.199.236 | NL      | EU        | KIND_IP |
      | geoip2 | params |    154.6.22.65 | US      | NA        | KIND_IP |

    Examples: With geoip2 headers
      | source | method  | ip             | country | continent | kind    |
      | geoip2 | headers |  95.91.246.242 | DE      | EU        | KIND_IP |
      | geoip2 | headers | 45.128.199.236 | NL      | EU        | KIND_IP |
      | geoip2 | headers |    154.6.22.65 | US      | NA        | KIND_IP |

  Scenario Outline: Get location by an bad IP address.
    When I request a location with HTTP:
      | ip     | <ip>     |
      | method | <method> |
    Then I should receive an empty response with HTTP

    Examples: With geoip2 parameters
      | source | method | ip      |
      | geoip2 | params | 0.0.0.0 |
      | geoip2 | params | test    |
      | geoip2 | params | <test>  |
      | geoip2 | params |   154.6 |

    Examples: With ip2location headers
      | source | method  | ip      |
      | geoip2 | headers | 0.0.0.0 |
      | geoip2 | headers | test    |
      | geoip2 | headers | <test>  |
      | geoip2 | headers |   154.6 |

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
      | method | latitude  | longitude  | country | continent | kind     |
      | params | 52.520008 |  13.404954 | DE      | EU        | KIND_GEO |
      | params | 52.377956 |   4.897070 | NL      | EU        | KIND_GEO |
      | params | 43.000000 | -75.000000 | US      | NA        | KIND_GEO |

    Examples: With headers
      | method  | latitude  | longitude  | country | continent | kind     |
      | headers | 52.520008 |  13.404954 | DE      | EU        | KIND_GEO |
      | headers | 52.377956 |   4.897070 | NL      | EU        | KIND_GEO |
      | headers | 43.000000 | -75.000000 | US      | NA        | KIND_GEO |

  Scenario Outline: Get location by a bad latitude and longitude.
    When I request a location with HTTP:
      | latitude  | <latitude>  |
      | longitude | <longitude> |
      | method    | <method>    |
    Then I should receive an empty response with HTTP

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
    Then I should receive an empty response with HTTP

    Examples:
      | method  | latitude | longitude |
      | params  |       90 |       180 |
      | headers |       90 |       180 |
