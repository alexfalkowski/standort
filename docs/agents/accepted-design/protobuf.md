# Accepted design for protobuf generation

Part of the [accepted-design index](../accepted-design.md).

## Generated files from deleted or renamed schemas

When deleting or renaming a protobuf file, the change author removes orphaned
Go and Ruby generated files as part of the same change. The generation config
defines both output locations, and `make proto-stale` regenerates the outputs
and checks the Git diff.

Revisit automatic cleanup in reliability reviews if there is evidence that
the current repository workflow publishes or validates stale generated API
artifacts.
[Generation config](../../../api/buf.gen.yaml) · [CircleCI workflow](../../../.circleci/config.yml)
