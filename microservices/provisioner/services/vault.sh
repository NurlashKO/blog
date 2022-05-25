docker run $(DOCKER_DEFAULT_ARGS vault) \
  --cap-add=IPC_LOCK \
  -v "$(pwd)/credentials/vault:/vault/file"
  gcr.io/kouzoh-p-nurlashko/nurlashko/vault server
