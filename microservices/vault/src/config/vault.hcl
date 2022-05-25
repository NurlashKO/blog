ui = true

listener "tcp" {
  tls_disable = 1
}

storage "file" {
  path = "/vault/file"
}
