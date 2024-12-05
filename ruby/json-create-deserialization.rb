require 'json'

class Foo
  def to_json(*args)
    {
      JSON.create_id  => self.class.name,
      'a'             => [ bar, baz ]
    }.to_json(*args)
  end

  # ruleid: json-create-deserialization
  def self.json_create(object)
    new(*object['a'])
  end
end

# ok: json-create-deserialization
class Bar
  def to_json(*args)
    {
      JSON.create_id  => self.class.name,
      'a'             => [ bar, baz ]
    }.to_json(*args)
  end
end
