provider "vault" {
  address = "https://vault.yourcompany.com:8200"
  token = local.vault_token
  # ruleid: vault-skip-tls-verify
  skip_tls_verify = true
}

provider "vault" {
  address = "https://vault.yourcompany.com:8200"
  token = local.vault_token
  # ok: vault-skip-tls-verify
  skip_tls_verify = false
}

# ok: vault-skip-tls-verify
provider "vault" {
  address = "https://vault.yourcompany.com:8200"
  token = local.vault_token
}
