# frozen_string_literal: true

module Standort
  module V1
    class HTTP < Nonnative::HTTPClient
      def get_location_by_ip(ip, opts = {})
        get("/v1/location/ip/#{URI::Parser.new.escape(ip)}", opts)
      end

      def get_location_by_lat_lng(lat, lng, opts = {})
        get("/v1/location/lat/#{lat}/lng/#{lng}", opts)
      end
    end
  end
end
