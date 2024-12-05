# ok: active-record-hardcoded-encryption-key
config.active_record.encryption.primary_key = ENV['ACTIVE_RECORD_ENCRYPTION_PRIMARY_KEY']

# ok: active-record-hardcoded-encryption-key
config.active_record.encryption.deterministic_key = ENV['ACTIVE_RECORD_ENCRYPTION_DETERMINISTIC_KEY']

# ruleid: active-record-hardcoded-encryption-key
config.active_record.encryption.primary_key = 'AAAAAAAAAAAAAAAA'

# ruleid: active-record-hardcoded-encryption-key
config.active_record.encryption.primary_key = [
  'AAAAAAAAAAAAAAAA',
  ENV['ACTIVE_RECORD_ENCRYPTION_PRIMARY_KEY']
]

# ruleid: active-record-hardcoded-encryption-key
config.active_record.encryption.deterministic_key = 'AAAAAAAAAAAAAAAA'
