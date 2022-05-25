ui = true

listener "tcp" {
  tls_disable = 1
  address = "0.0.0.0:8200"
}

storage "file" {
  path = "/vault/file"
}
