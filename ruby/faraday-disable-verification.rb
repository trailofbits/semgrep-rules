require 'faraday'

# ruleid: faraday-disable-verification
conn = Faraday.new(ssl: { verify: false }) do |faraday|
  # ...
end

# ok: faraday-disable-verification
conn = Faraday.new(ssl: { ca_file: '/path/to/ca_file', min_version: :TLS1_2 }) do |faraday|
  # ...
end
