# frozen_string_literal: true

module Standort
  module V2
    class HTTP < Nonnative::HTTPClient
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
