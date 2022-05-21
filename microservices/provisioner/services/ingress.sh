mkdir -p ./data/certificates

docker run $(DOCKER_DEFAULT_ARGS ingress) \
    -p 80:80 \
    -p 443:443 \
    -v $(pwd)/data/certificates:/etc/letsencrypt \
    gcr.io/kouzoh-p-nurlashko/nurlashko/ingress
