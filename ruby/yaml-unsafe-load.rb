require 'yaml'

data = get_data()

# ruleid: yaml-unsafe-load
YAML.unsafe_load(data)

# ruleid: yaml-unsafe-load
Psych.unsafe_load(data)

# ok: yaml-unsafe-load
YAML.unsafe_load("---foo: bar")

class User < ActiveRecord::Base
  # ruleid: yaml-unsafe-load
  serialize :preferences, coder: YAML, yaml: { unsafe_load: true }
end
