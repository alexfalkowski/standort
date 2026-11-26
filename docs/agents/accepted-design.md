# Accepted design

This index records Standort's owner-accepted design choices and their scope.
The area files carry the same weight as [AGENTS.md](../../AGENTS.md).

Read the relevant area files before proposing a change or recording a finding
or audit entry. Revisit a decision when the owner asks or new evidence changes
the documented contract or tradeoff.

- [Location lookups](accepted-design/location.md) covers v2 metadata fallback
  and embedded lookup asset freshness.
- [Protobuf generation](accepted-design/protobuf.md) covers cleanup of generated
  Go and Ruby files after schema deletes and renames.
- [Deployment](accepted-design/deployment.md) covers rollout ownership and
  serialization.
