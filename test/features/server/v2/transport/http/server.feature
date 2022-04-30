@startup
Feature: Server

  Server allows users to get locations by different types.

  Scenario Outline: Get location by a valid IP address.
    When I request a location with HTTP:
      | ip     | <ip>     |
      | method | <method> |
    Then I should receive a valid locations with HTTP:
      | type      | <type>      |
      | country   | <country>   |
      | continent | <continent> |
    And the process 'server' should consume less than '40mb' of memory

    Examples: With parameters
      | method | ip             | country | continent | type    |
      | params | 95.91.246.242  | DE      | EU        | TYPE_IP |
      | params | 45.128.199.236 | NL      | EU        | TYPE_IP |
      | params | 154.6.22.65    | US      | NA        | TYPE_IP |

    Examples: With headers
      | method  | ip             | country | continent | type    |
      | headers | 95.91.246.242  | DE      | EU        | TYPE_IP |
      | headers | 45.128.199.236 | NL      | EU        | TYPE_IP |
      | headers | 154.6.22.65    | US      | NA        | TYPE_IP |

  Scenario Outline: Get location by an bad IP address.
    When I request a location with HTTP:
      | ip     | <ip>     |
      | method | <method> |
    Then I should receive an empty response with HTTP
    And the process 'server' should consume less than '40mb' of memory

    Examples: With parameters
      | method | ip      |
      | params | 0.0.0.0 |
      | params | test    |
      | params | <test>  |
      | params | 154.6   |

    Examples: With headers
      | method  | ip      |
      | headers | 0.0.0.0 |
      | headers | test    |
      | headers | <test>  |
      | headers | 154.6   |

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
