require 'rest-client'

# ruleid: rest-client-disable-verification
RestClient::Resource.new(
  'https://example.com',
  :ssl_client_cert  =>  OpenSSL::X509::Certificate.new(File.read("cert.pem")),
  :ssl_client_key   =>  OpenSSL::PKey::RSA.new(File.read("key.pem"), "passphrase, if any"),
  :ssl_ca_file      =>  "ca_certificate.pem",
  :verify_ssl       =>  OpenSSL::SSL::VERIFY_NONE
).get

# ruleid: rest-client-disable-verification
RestClient::Resource.new(
  'https://example.com',
  :ssl_client_cert  =>  OpenSSL::X509::Certificate.new(File.read("cert.pem")),
  :ssl_client_key   =>  OpenSSL::PKey::RSA.new(File.read("key.pem"), "passphrase, if any"),
  :ssl_ca_file      =>  "ca_certificate.pem",
  :verify_ssl       =>  false
).get

# ruleid: rest-client-disable-verification
RestClient::Resource.new(
  'https://example.com',
  ssl_client_cert:  OpenSSL::X509::Certificate.new(File.read("cert.pem")),
  ssl_client_key:  OpenSSL::PKey::RSA.new(File.read("key.pem"), "passphrase, if any"),
  ssl_ca_file:  "ca_certificate.pem",
  verify_ssl:  false
).get

# ok: rest-client-disable-verification
RestClient::Resource.new(
  'https://example.com',
  :ssl_client_cert  =>  OpenSSL::X509::Certificate.new(File.read("cert.pem")),
  :ssl_client_key   =>  OpenSSL::PKey::RSA.new(File.read("key.pem"), "passphrase, if any"),
  :ssl_ca_file      =>  "ca_certificate.pem",
  :verify_ssl       =>  OpenSSL::SSL::VERIFY_PEER
).get
