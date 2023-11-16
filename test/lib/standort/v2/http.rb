# frozen_string_literal: true

module Standort
  module V2
    class HTTP < Nonnative::HTTPClient
      def get_location(params, opts = {})
        ip = params[:ip]
        point = params[:point]
        uri = if ip || point
                params_uri(ip, point)
              else
                '/v2/location'
              end

        get(uri, opts)
      end

      private

      def params_uri(ip, point)
        point ||= []
        point.compact!

        params = {}
        params[:ip] = ip if ip

        if point.length.positive?
          params['point.lat'] = point[0]
          params['point.lng'] = point[1]
        end

        "/v2/location?#{URI.encode_www_form(params)}"
      end
    end
  end
end
