Rails.application.configure do
  # ruleid: action-dispatch-insecure-ssl
  config.force_ssl = false

  # ruleid: action-dispatch-insecure-ssl
  config.ssl_options = { hsts: false }

  # ruleid: action-dispatch-insecure-ssl
  config.ssl_options = { hsts: { subdomains: false } }

  # ok: action-dispatch-insecure-ssl
  config.ssl_options = { redirect: { exclude: -> request { /healthcheck/.match?(request.path) } } }
end
