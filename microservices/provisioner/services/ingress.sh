mkdir -p ./data/certificates

docker run $(DOCKER_DEFAULT_ARGS ingress) \
    -p 80:80 \
    -p 443:443 \
    -p 943:943 \
    -p 1194:1194 \
    -v $(pwd)/data/certificates:/etc/letsencrypt \
    -v $(pwd)/data/statika/srv/images:/www/data/images \
    gcr.io/kouzoh-p-nurlashko/nurlashko/ingress
