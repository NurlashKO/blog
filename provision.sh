#!/bin/bash

DOCKER_DEFAULT_ARGS() {
  local name="$1";
  echo "-d --rm --network internal --hostname ${name} --name ${name}"
}

# Cleanup everything
docker kill $(docker ps -q)
docker system prune -af

# Setup networking
docker network create internal

# Deploy container watcher
docker run $(DOCKER_DEFAULT_ARGS watcher) \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -v /home/"$USER"/.docker/config.json:/config.json \
    containrrr/watchtower --revive-stopped --interval 5

# Deploy microservices
mkdir -p certificates && \
  docker run $(DOCKER_DEFAULT_ARGS loadbalancer) \
      -p 80:80 \
      -p 443:443 \
      --env CERTBOT_EMAIL=zh.nurlan96@gmail.com \
      -v $(pwd)/certificates:/etc/letsencrypt \
      gcr.io/kouzoh-p-nurlashko/nurlashko/nginx


docker run $(DOCKER_DEFAULT_ARGS blog) \
    gcr.io/kouzoh-p-nurlashko/nurlashko/blog
