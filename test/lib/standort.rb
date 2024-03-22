# frozen_string_literal: true

require 'securerandom'
require 'yaml'
require 'base64'

require 'grpc/health/v1/health_services_pb'

require 'standort/v1/http'
require 'standort/v1/service_pb'
require 'standort/v1/service_services_pb'
require 'standort/v2/http'
require 'standort/v2/service_pb'
require 'standort/v2/service_services_pb'

module Standort
  class << self
    def observability
      @observability ||= Nonnative::Observability.new('http://localhost:11000')
    end

    def server_config
      @server_config ||= Nonnative.configurations('.config/server.config.yml')
    end

    def health_grpc
      @health_grpc ||= Grpc::Health::V1::Health::Stub.new('localhost:12000', :this_channel_is_insecure, channel_args: Standort.user_agent)
    end

    def user_agent
      @user_agent ||= Nonnative::Header.grpc_user_agent('Standort-ruby-client/2.0 gRPC/1.0')
    end
  end

  module V1
    class << self
      def server_http
        @server_http ||= Standort::V1::HTTP.new('http://localhost:11000')
      end

      def server_grpc
        @server_grpc ||= Standort::V1::Service::Stub.new('localhost:12000', :this_channel_is_insecure, channel_args: Standort.user_agent)
      end
    end
  end

  module V2
    class << self
      def server_http
        @server_http ||= Standort::V2::HTTP.new('http://localhost:11000')
      end

      def server_grpc
        @server_grpc ||= Standort::V2::Service::Stub.new('localhost:12000', :this_channel_is_insecure, channel_args: Standort.user_agent)
      end
    end
  end
end
