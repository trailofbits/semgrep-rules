require 'timeout'
require 'faraday'

url = 'http://example.com'

begin
  # ruleid: global-timeout
  Timeout::timeout(5) do
    response = Faraday.get(url)
    puts response.body
  end
rescue Timeout::Error
  puts "The request to the URL has taken more time than the allotted 5 seconds."
end

# ok: global-timeout
connection = Faraday.new(url) do |req|
  req.options.timeout = 5
end

begin
  response = connection.get('/')
  puts response.body
rescue Faraday::TimeoutError
  puts "The request took more than 5 seconds to complete."
end
