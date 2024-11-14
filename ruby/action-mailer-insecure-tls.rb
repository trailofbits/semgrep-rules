# ruleid: action-mailer-insecure-tls
ActionMailer::Base.smtp_settings = {
  address:              'smtp.test.net',
  port:                 587,
  user_name:            'user',
  password:             'pass',
  domain:               'mailer.test.com',
  authentication:       :plain,
  enable_starttls_auto: true
}

# ok: action-mailer-insecure-tls
ActionMailer::Base.smtp_settings = {
  address:              'smtp.test.net',
  port:                 587,
  user_name:            'user',
  password:             'pass',
  domain:               'mailer.test.com',
  authentication:       :plain,
  enable_starttls:      true
}

# ruleid: action-mailer-insecure-tls
ActionMailer::Base.smtp_settings = {
  address:              'smtp.test.net',
  port:                 587,
  user_name:            'user',
  password:             'pass',
  domain:               'mailer.test.com',
  authentication:       :plain,
  openssl_verify_mode:  OpenSSL::SSL::VERIFY_NONE
}

# ok: action-mailer-insecure-tls
ActionMailer::Base.smtp_settings = {
  address:              'smtp.test.net',
  port:                 587,
  user_name:            'user',
  password:             'pass',
  domain:               'mailer.test.com',
  authentication:       :plain,
  openssl_verify_mode:  OpenSSL::SSL::VERIFY_PEER
}
