require 'rack/test'
require 'rspec'
require 'multiply_server'

ENV['RACK_ENV'] = 'test'

module RSpecMixin
  include Rack::Test::Methods
  def app() MultiplyServer end
end

RSpec.configure { |c| c.include RSpecMixin }
