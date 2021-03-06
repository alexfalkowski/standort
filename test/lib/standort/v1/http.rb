# frozen_string_literal: true

module Standort
  module V1
    class HTTP < Nonnative::HTTPClient
      def get_location_by_ip(ip, headers = {})
        default_headers = {
          content_type: :json,
          accept: :json
        }

        default_headers.merge!(headers)

        get("/v1/location/ip/#{URI::Parser.new.escape(ip)}", headers, 10)
      end

      def get_location_by_lat_lng(lat, lng, headers = {})
        default_headers = {
          content_type: :json,
          accept: :json
        }

        default_headers.merge!(headers)

        get("/v1/location/lat/#{lat}/lng/#{lng}", headers, 10)
      end
    end
  end
end
