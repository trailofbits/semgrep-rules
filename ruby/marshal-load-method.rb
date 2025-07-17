require 'marshal'

class Foo
  # ruleid: marshal-load-method
  def marshal_load(array)
    initialize array[0]
  end
end

# ok: marshal-load-method
class Bar
end
