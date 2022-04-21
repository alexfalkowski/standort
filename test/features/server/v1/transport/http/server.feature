@manual
Feature: Server

  Server allows users to get locations by different means.

  Scenario Outline: Get location by a valid IP address.
    When I request a location by IP address with HTTP:
      | ip | <ip> |
    Then I should receive a valid location by IP adress with HTTP:
      | country   | <country>   |
      | continent | <continent> |

    Examples:
      | ip             | country | continent |
      | 95.91.246.242  | DE      | EU        |
      | 45.128.199.236 | NL      | EU        |
      | 154.6.22.65    | US      | NA        |

  Scenario Outline: Get location by an bad IP address.
    When I request a location by IP address with HTTP:
      | ip | <ip> |
    Then I should receive a bad response with HTTP

    Examples:
      | ip     |
      | test   |
      | <test> |
      | 154.6  |
      |        |

  Scenario Outline: Get location by a not found IP address.
    When I request a location by IP address with HTTP:
      | ip | <ip> |
    Then I should receive a not found response with HTTP

    Examples:
      | ip      |
      | 0.0.0.0 |
