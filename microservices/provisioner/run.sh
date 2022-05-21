#!/bin/sh

cd /home/nurlashko

PROVISIONER_MARK="__provisioner-managed__"

DOCKER_DEFAULT_ARGS() {
  name="$1";
  echo "-detach --log-driver json-file --log-opt tag=\"{{.ImageName}}|{{.Name}}|{{.ImageFullID}}|{{.FullID}}\" --network internal --hostname ${name} --name ${name} --label ${PROVISIONER_MARK}"
}

# Cleanup everything
docker kill $(docker ps -f "label=$PROVISIONER_MARK" -q)
docker system prune -f

# Setup networking
docker network create internal

# Setup filesystem shared with microservices
mkdir -p data

# Deploy microservices
docker run $(DOCKER_DEFAULT_ARGS watcher) \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -v $(pwd)/.docker/config.json:/config.json \
    gcr.io/kouzoh-p-nurlashko/nurlashko/watcher

docker run $(DOCKER_DEFAULT_ARGS loki) \
    --mount source=loki-data,target=/loki \
    gcr.io/kouzoh-p-nurlashko/nurlashko/loki

docker run $(DOCKER_DEFAULT_ARGS promtail) \
    -v /var/lib/docker/:/var/lib/docker:ro \
    gcr.io/kouzoh-p-nurlashko/nurlashko/promtail

mkdir -p ./data/certificates && \
  docker run $(DOCKER_DEFAULT_ARGS ingress) \
      -p 80:80 \
      -p 443:443 \
      --env CERTBOT_EMAIL=zh.nurlan96@gmail.com \
      -v $(pwd)/data/certificates:/etc/letsencrypt \
      gcr.io/kouzoh-p-nurlashko/nurlashko/ingress


docker run $(DOCKER_DEFAULT_ARGS blog) \
    gcr.io/kouzoh-p-nurlashko/nurlashko/blog
