# frozen_string_literal: true

require 'securerandom'
require 'yaml'
require 'base64'

require 'google/rpc/error_details_pb'

require 'standort/v1/http'
require 'standort/v1/service_pb'
require 'standort/v1/service_services_pb'
require 'standort/v2/http'
require 'standort/v2/service_pb'
require 'standort/v2/service_services_pb'

##
# Public entrypoint for the Ruby *test client* used by the feature/benchmark harness.
#
# This module provides:
#
# - A small configuration helper via {.config}
# - Pre-wired gRPC stubs for the service (v1/v2) and gRPC Health checking
# - Convenience constructors for the HTTP transport clients (v1/v2)
#
# The clients are intentionally configured for local development defaults:
#
# - HTTP: `http://localhost:11000`
# - gRPC: `localhost:12000`
#
# These values match the dev config in `test/.config/server.yml`.
#
# @example Load configuration used by the harness
#   Standort.config
#
# @example Call the v2 HTTP endpoint
#   response = Standort::V2.http.get_location({ ip: "8.8.8.8" })
#
# @example Call the v1 gRPC endpoint
#   stub = Standort::V1.grpc
#   stub.get_location_by_ip(Standort::V1::GetLocationByIPRequest.new(ip: "8.8.8.8"))
#
module Standort
  class << self
    ##
    # Loads and memoizes server configuration used by the Ruby harness.
    #
    # The underlying configuration loader is provided by `Nonnative` and reads
    # the YAML config file at `.config/server.yml` *relative to the `test/` directory*.
    #
    # @return [Object] The configuration object returned by `Nonnative::ConfigurationFile.load`.
    #
    def config
      @config ||= Nonnative::ConfigurationFile.load('.config/server.yml')
    end

    ##
    # Returns call options for unary gRPC requests made by the local harness.
    #
    # @param metadata [Hash] request metadata to attach to the call.
    # @return [Hash] gRPC call options including metadata and a client-side deadline.
    #
    def grpc_options(metadata = {})
      {
        metadata:,
        deadline: Time.now + 10
      }
    end

    ##
    # Returns bounded per-call options for HTTP feature-harness requests.
    #
    # Each call includes a generated `request_id` header. Caller-provided
    # headers are merged afterward, so scenarios can override that value or add
    # transport-specific headers such as content type and user agent.
    #
    # @param headers [Hash] HTTP headers merged after the generated request id.
    # @param read_timeout [Integer] read timeout in seconds.
    # @param open_timeout [Integer] connection open timeout in seconds.
    # @return [Hash] Options compatible with `Nonnative::HTTPClient` calls.
    #
    def http_options(headers: {}, read_timeout: 10, open_timeout: 10)
      {
        headers: { request_id: SecureRandom.uuid }.merge(headers),
        read_timeout:,
        open_timeout:
      }
    end

    ##
    # Returns a memoized gRPC Health Checking stub.
    #
    # This is used by health/observability feature steps to query the server's
    # gRPC health endpoint.
    #
    # The channel is created as insecure because this harness targets local
    # development endpoints.
    #
    # @param service [String] gRPC health service name.
    # @return [Nonnative::GRPCHealth]
    #
    def health_grpc(service)
      @health_grpc ||= {}
      @health_grpc[service] ||= Nonnative.grpc_health(host: 'localhost', port: 12_000, service:)
    end

    ##
    # Channel args used to set the gRPC user agent for all stubs created by this module.
    #
    # @return [Hash] gRPC channel args as returned by `Nonnative::Header.grpc_user_agent`.
    #
    def user_agent
      @user_agent ||= Nonnative::Header.grpc_user_agent('Standort-ruby-client/2.0 gRPC/1.0')
    end
  end

  ##
  # Convenience accessors for v1 clients.
  #
  # - {.http} returns an HTTP client that uses the service's HTTP mapping
  # - {.grpc} returns a generated gRPC stub for direct RPC calls
  #
  module V1
    class << self
      ##
      # Returns a memoized v1 HTTP client.
      #
      # This client posts JSON to the HTTP->gRPC gateway routes for v1.
      #
      # @return [Standort::V1::HTTP]
      #
      def http
        @http ||= Standort::V1::HTTP.new('http://localhost:11000')
      end

      ##
      # Returns a memoized v1 gRPC stub.
      #
      # The stub is configured with the same user agent as other gRPC clients in this harness.
      #
      # @return [Standort::V1::Service::Stub]
      #
      def grpc
        @grpc ||= Standort::V1::Service::Stub.new(
          'localhost:12000',
          :this_channel_is_insecure,
          channel_args: Standort.user_agent
        )
      end
    end
  end

  ##
  # Convenience accessors for v2 clients.
  #
  # - {.http} returns an HTTP client that uses the service's HTTP mapping
  # - {.grpc} returns a generated gRPC stub for direct RPC calls
  #
  module V2
    class << self
      ##
      # Returns a memoized v2 HTTP client.
      #
      # This client posts JSON to the HTTP->gRPC gateway routes for v2.
      #
      # @return [Standort::V2::HTTP]
      #
      def http
        @http ||= Standort::V2::HTTP.new('http://localhost:11000')
      end

      ##
      # Returns a memoized v2 gRPC stub.
      #
      # The stub is configured with the same user agent as other gRPC clients in this harness.
      #
      # @return [Standort::V2::Service::Stub]
      #
      def grpc
        @grpc ||= Standort::V2::Service::Stub.new(
          'localhost:12000',
          :this_channel_is_insecure,
          channel_args: Standort.user_agent
        )
      end
    end
  end
end
