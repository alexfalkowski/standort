# Accepted design for deployment

Part of the [accepted-design index](../accepted-design.md).

## Rollout ownership and serialization

The version-controlled `area/apps` path in `alexfalkowski/infraops` owns the
deployed state, including runtime artifact references, startup and readiness
probes, serialized update jobs, and targeted release verification. Standort's
CircleCI deploy job hands off the release after publishing the image manifest.

Revisit Standort-side deploy serialization in reliability reviews if evidence
shows the infraops path has lost control of the deployed state or can roll it
back out of order.
[Standort CircleCI workflow](../../../.circleci/config.yml)
