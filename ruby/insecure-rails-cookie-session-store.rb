# ruleid: insecure-rails-cookie-session-store
Rails.application.config.session_store :cookie_store, key: 'session_key'

# ok: insecure-rails-cookie-session-store
Rails.application.config.session_store :cookie_store, key: 'session_key', same_site: :secure
