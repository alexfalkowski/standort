@manual
Feature: Server

  Server allows users to get locations by different means.

  Scenario Outline: Get location by a valid IP address.
    When I request a location by IP address with gRPC:
      | ip | <ip> |
    Then I should receive a valid location by IP adress with gRPC:
      | country   | <country>   |
      | continent | <continent> |
    And the process 'server' should consume less than '40mb' of memory

    Examples:
      | ip             | country | continent |
      | 95.91.246.242  | DE      | EU        |
      | 45.128.199.236 | NL      | EU        |
      | 154.6.22.65    | US      | NA        |

  Scenario Outline: Get location by an bad IP address.
    When I request a location by IP address with gRPC:
      | ip | <ip> |
    Then I should receive a bad response with gRPC
    And the process 'server' should consume less than '40mb' of memory

    Examples:
      | ip     |
      | test   |
      | <test> |
      | 154.6  |
      |        |

  Scenario Outline: Get location by a not found IP address.
    When I request a location by IP address with gRPC:
      | ip | <ip> |
    Then I should receive a not found response with gRPC
    And the process 'server' should consume less than '40mb' of memory

    Examples:
      | ip      |
      | 0.0.0.0 |

  Scenario Outline: Get location by a valid latitude and longitude.
    When I request a location by latitude and longitude with gRPC:
      | latitude  | <latitude>  |
      | longitude | <longitude> |
    Then I should receive a valid location by latitude and longitude with gRPC:
      | country   | <country>   |
      | continent | <continent> |
    And the process 'server' should consume less than '40mb' of memory

    Examples:
      | latitude  | longitude  | country | continent |
      | 52.520008 | 13.404954  | DE      | EU        |
      | 52.377956 | 4.897070   | NL      | EU        |
      | 43.000000 | -75.000000 | US      | NA        |

  Scenario Outline: Get location by a bad latitude and longitude.
    When I request a location by latitude and longitude with gRPC:
      | latitude  | <latitude>  |
      | longitude | <longitude> |
    Then I should receive a bad response with gRPC

    Examples:
      | latitude | longitude |
      | 91       | 10        |
      | 10       | 181       |

  Scenario Outline: Get location by a not found latitude and longitude.
    When I request a location by latitude and longitude with gRPC:
      | latitude  | <latitude>  |
      | longitude | <longitude> |
    Then I should receive a not found response with gRPC
    And the process 'server' should consume less than '40mb' of memory

    Examples:
      | latitude | longitude |
      | 90       | 180       |
