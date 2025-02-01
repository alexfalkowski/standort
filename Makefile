include bin/build/make/grpc.mak
include bin/build/make/git.mak

# Diagrams generated from https://github.com/loov/goda.
diagrams: location-diagram

location-diagram:
	$(MAKE) package=location create-diagram
