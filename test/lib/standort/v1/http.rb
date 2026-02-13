# frozen_string_literal: true

module Standort
  module V1
    ##
    # HTTP transport client for Standort API v1.
    #
    # This client talks to the service's HTTP-to-gRPC mapping by POSTing JSON
    # payloads to routes shaped like:
    #
    # - `/standort.v1.Service/GetLocationByIP`
    # - `/standort.v1.Service/GetLocationByLatLng`
    #
    # It inherits request/response behavior (headers, timeouts, etc.) from
    # `Nonnative::HTTPClient`.
    #
    class HTTP < Nonnative::HTTPClient
      ##
      # Looks up location information by IP address.
      #
      # The IP is URI-escaped before being JSON encoded and sent to the gateway.
      #
      # @param ip [String] An IPv4 or IPv6 address (for example `"8.8.8.8"`).
      # @param opts [Hash] Optional request options forwarded to `Nonnative::HTTPClient#post`
      #   (for example custom headers).
      #
      # @return [Object] Whatever `Nonnative::HTTPClient#post` returns (typically a parsed HTTP response).
      #
      # @example
      #   client = Standort::V1::HTTP.new("http://localhost:11000")
      #   client.get_location_by_ip("8.8.8.8")
      #
      def get_location_by_ip(ip, opts = {})
        post('/standort.v1.Service/GetLocationByIP', { ip: URI::Parser.new.escape(ip) }.to_json, opts)
      end

      ##
      # Looks up location information by latitude and longitude.
      #
      # Inputs are coerced to floats before being JSON encoded and sent to the gateway.
      #
      # @param lat [Numeric, String] Latitude.
      # @param lng [Numeric, String] Longitude.
      # @param opts [Hash] Optional request options forwarded to `Nonnative::HTTPClient#post`
      #   (for example custom headers).
      #
      # @return [Object] Whatever `Nonnative::HTTPClient#post` returns (typically a parsed HTTP response).
      #
      # @example
      #   client = Standort::V1::HTTP.new("http://localhost:11000")
      #   client.get_location_by_lat_lng(52.52, 13.405)
      #
      def get_location_by_lat_lng(lat, lng, opts = {})
        post('/standort.v1.Service/GetLocationByLatLng', { lat: lat.to_f, lng: lng.to_f }.to_json, opts)
      end
    end
  end
end
