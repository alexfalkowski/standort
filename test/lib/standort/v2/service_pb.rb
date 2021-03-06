# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: standort/v2/service.proto

require 'google/protobuf'

require 'google/api/annotations_pb'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("standort/v2/service.proto", :syntax => :proto3) do
    add_message "standort.v2.Location" do
      optional :kind, :enum, 1, "standort.v2.Kind", json_name: "kind"
      optional :country, :string, 2, json_name: "country"
      optional :continent, :string, 3, json_name: "continent"
    end
    add_message "standort.v2.Point" do
      optional :lat, :double, 1, json_name: "lat"
      optional :lng, :double, 2, json_name: "lng"
    end
    add_message "standort.v2.GetLocationRequest" do
      optional :ip, :string, 1, json_name: "ip"
      optional :point, :message, 2, "standort.v2.Point", json_name: "point"
    end
    add_message "standort.v2.GetLocationResponse" do
      repeated :locations, :message, 1, "standort.v2.Location", json_name: "locations"
    end
    add_enum "standort.v2.Kind" do
      value :KIND_UNSPECIFIED, 0
      value :KIND_IP, 1
      value :KIND_GEO, 2
    end
  end
end

module Standort
  module V2
    Location = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("standort.v2.Location").msgclass
    Point = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("standort.v2.Point").msgclass
    GetLocationRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("standort.v2.GetLocationRequest").msgclass
    GetLocationResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("standort.v2.GetLocationResponse").msgclass
    Kind = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("standort.v2.Kind").enummodule
  end
end
