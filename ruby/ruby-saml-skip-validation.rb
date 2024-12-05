require 'onelogin/ruby-saml'

# ruleid: ruby-saml-skip-validation
response = OneLogin::RubySaml::Response.new(params[:SAMLResponse], {skip_audience: true})

# ok: ruby-saml-skip-validation
response = OneLogin::RubySaml::Response.new(params[:SAMLResponse])

settings = OneLogin::RubySaml::Settings.new

# ruleid: ruby-saml-skip-validation
settings.skip_audience = true

# ok: ruby-saml-skip-validation
settings.name_identifier_format = "urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress"
