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
docker run $(DOCKER_ARGS watcher) \
    -v /var/run/docker.sock:/var/run/docker.sock \
    containrrr/watchtower --revive-stopped --interval 5

# Deploy microservices
docker run $(DOCKER_ARGS loadbalancer) \
    -p 80:80 \
    -p 443:443 \
    --entrypoint /bin/sh -c 'while :; do sleep 6h & wait $${!}; nginx -s reload; done & nginx -g \"daemon off;\"' \
    gcr.io/kouzoh-p-nurlashko/nurlashko/nginx

docker run $(DOCKER_ARGS blog) \
    crccheck/hello-world
