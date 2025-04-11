# frozen_string_literal: true

module Standort
  module V1
    class HTTP < Nonnative::HTTPClient
      def get_location_by_ip(ip, opts = {})
        post('/standort.v1.Service/GetLocationByIP', { ip: URI::Parser.new.escape(ip) }.to_json, opts)
      end

      def get_location_by_lat_lng(lat, lng, opts = {})
        post('/standort.v1.Service/GetLocationByLatLng', { lat: lat.to_f, lng: lng.to_f }.to_json, opts)
      end
    end
  end
end
