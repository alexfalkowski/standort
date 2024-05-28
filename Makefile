include bin/build/make/grpc.mak
include bin/build/make/git.mak

# Diagrams generated from https://github.com/loov/goda.
diagrams: client-diagram location-diagram server-diagram

client-diagram:
	$(MAKE) package=client create-diagram

location-diagram:
	$(MAKE) package=location create-diagram

server-diagram:
	$(MAKE) package=server create-diagram
