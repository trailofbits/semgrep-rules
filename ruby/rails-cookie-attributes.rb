require 'rails'

# ruleid: rails-cookie-attributes
cookies[:key] = { value: 'value', same_site: :none }

# ruleid: rails-cookie-attributes
cookies[:key] = { value: 'value', secure: false }

# ruleid: rails-cookie-attributes
cookies[:key] = { value: 'value', httponly: false }

# ruleid: rails-cookie-attributes
cookies.permanent.encrypted[:key] = { value: 'value', secure: false }

# ruleid: rails-cookie-attributes
cookies.signed[:key] = { value: 'value', secure: false }

# ok: rails-cookie-attributes
cookies[:key] = { value: 'value', secure: true, httponly: true, same_site: :strict }
