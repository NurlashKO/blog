ui = true

listener "tcp" {
  tls_disable = 1
  address = "127.0.0.1:8200"
}

storage "file" {
  path = "/vault/file"
}
