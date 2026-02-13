# frozen_string_literal: true

module Standort
  module V2
    ##
    # HTTP transport client for Standort API v2.
    #
    # This client talks to the service's HTTP-to-gRPC mapping by POSTing JSON
    # payloads to:
    #
    # - `/standort.v2.Service/GetLocation`
    #
    # It inherits request/response behavior (headers, timeouts, etc.) from
    # `Nonnative::HTTPClient`.
    #
    class HTTP < Nonnative::HTTPClient
      ##
      # Looks up location information.
      #
      # You can provide an IP, a point (lat/lng), or both. When present, values are
      # coerced into the JSON shape expected by the gateway.
      #
      # @param params [Hash] Request parameters.
      # @option params [String] :ip An IPv4 or IPv6 address (for example `"8.8.8.8"`).
      # @option params [Array<(Numeric,String)>] :point A 2-item array `[lat, lng]`.
      #   Latitude/longitude values are coerced to floats.
      # @param opts [Hash] Optional request options forwarded to `Nonnative::HTTPClient#post`
      #   (for example custom headers).
      #
      # @return [Object] Whatever `Nonnative::HTTPClient#post` returns (typically a parsed HTTP response).
      #
      # @example Lookup by IP
      #   client = Standort::V2::HTTP.new("http://localhost:11000")
      #   client.get_location({ ip: "8.8.8.8" })
      #
      # @example Lookup by point
      #   client = Standort::V2::HTTP.new("http://localhost:11000")
      #   client.get_location({ point: [52.52, 13.405] })
      #
      # @example Lookup by IP and point
      #   client = Standort::V2::HTTP.new("http://localhost:11000")
      #   client.get_location({ ip: "8.8.8.8", point: [52.52, 13.405] })
      #
      def get_location(params, opts = {})
        req = {}
        req[:ip] = params[:ip] if params[:ip]
        point = params[:point] || []
        req[:point] = { lat: point[0].to_f, lng: point[1].to_f } if point.length.positive?

        post('/standort.v2.Service/GetLocation', req.to_json, opts)
      end
    end
  end
end
