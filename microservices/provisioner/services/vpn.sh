docker run $(DOCKER_DEFAULT_ARGS vpn) \
  --device /dev/net/tun  \
  --cap-add=MKNOD \
  --cap-add=NET_ADMIN \
  -v "$(pwd)/credentials/vpn:/openvpn" \
    gcr.io/kouzoh-p-nurlashko/nurlashko/vpn
