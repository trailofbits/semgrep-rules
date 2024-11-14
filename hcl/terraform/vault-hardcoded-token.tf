provider "vault" {
  address = "https://vault.yourcompany.com:8200"
  # ruleid: vault-hardcoded-token
  token = "HARDCODEDTOKEN"
}

provider "vault" {
  address = "https://vault.yourcompany.com:8200"
  # ok: vault-hardcoded-token
  token = local.vault_token
}

provider "vault" {
  address = "https://vault.yourcompany.com:8200"
  # ok: vault-hardcoded-token
  token = "${var.token}"
}

provider "vault" {
  address = "https://vault.yourcompany.com:8200"
  # ok: vault-hardcoded-token
  token = var.token
}
