# frozen_string_literal: true

module Standort
  module V1
    class HTTP < Nonnative::HTTPClient
      def get_config(ip, headers = {})
        default_headers = {
          content_type: :json,
          accept: :json
        }

        default_headers.merge!(headers)

        get("/v1/location/ip/#{URI::Parser.new.escape(ip)}", headers, 10)
      end
    end
  end
end
