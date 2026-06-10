# frozen_string_literal: true

module Coordinates
  module_function

  def parse(value)
    case value.to_s.strip.downcase
    when 'nan'
      Float::NAN
    when 'inf', '+inf', 'infinity', '+infinity'
      Float::INFINITY
    when '-inf', '-infinity'
      -Float::INFINITY
    else
      value.to_f
    end
  end
end
