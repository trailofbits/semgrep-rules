require 'json'
require 'rails'

module JsonSerializer
  def dump(value)
    JSON.dump(value)
  end

  def load(dumped)
    JSON.load(dumped)
  end
end

Rails.application.configure do
  # ruleid: rails-cache-store-marshal
  config.cache_store = :mem_cache_store

  # ruleid: rails-cache-store-marshal
  config.cache_store = :mem_cache_store, ENV['MEMCACHED_SERVERS']

  # ruleid: rails-cache-store-marshal
  config.cache_store = :mem_cache_store, ENV['MEMCACHED_SERVERS'], {
    failover: true,
    socket_timeout: 1.5,
    socket_failure_delay: 0.2,
    compress: true,
  }

  # ok: rails-cache-store-marshal
  config.cache_store = :mem_cache_store, ENV['MEMCACHED_SERVERS'], {
    failover: true,
    socket_timeout: 1.5,
    socket_failure_delay: 0.2,
    compress: true,
    serializer: JsonSerializer,
  }

  # ruleid: rails-cache-store-marshal
  config.cache_store = :redis_cache_store, { url: ENV["REDIS_URL"] }

  # ok: rails-cache-store-marshal
  config.cache_store = :redis_cache_store, { url: ENV["REDIS_URL"], serializer: JsonSerializer }

  # ok: rails-cache-store-marshal
  config.cache_store = :null_store
end
