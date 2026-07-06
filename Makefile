include bin/build/make/help.mak
include bin/build/make/grpc.mak
include bin/build/make/git.mak
include bin/build/make/claude.mak

# Download and check the embedded lookup assets.
update-lookup-assets:
	@scripts/update-lookup-assets
