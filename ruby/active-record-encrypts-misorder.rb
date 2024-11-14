class Article < ApplicationRecord
  # ok: active-record-encrypts-misorder
  serialize :title, type: Title
  encrypts :title
end

class Article < ApplicationRecord
  # ruleid: active-record-encrypts-misorder
  encrypts :title
  serialize :title, type: Title
end
