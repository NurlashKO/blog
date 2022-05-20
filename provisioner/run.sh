#!/bin/sh

su - nurlashko
cd /home/nurlashko

PROVISIONER_MARK="__provisioner-managed__"

DOCKER_DEFAULT_ARGS() {
  name="$1";
  echo "-detach --network internal --hostname ${name} --name ${name} --label ${PROVISIONER_MARK} --label test --label test2"
}

# Cleanup everything
docker kill $(docker ps -f "label=$PROVISIONER_MARK" -q)
docker system prune -f

# Setup networking
docker network create internal

# Setup filesystem shared with microservices
mkdir -p data

# Deploy container watcher
docker run $(DOCKER_DEFAULT_ARGS watcher) \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -v $(pwd).docker/config.json:/config.json \
    containrrr/watchtower --include-stopped --revive-stopped --cleanup --interval 5

# Deploy microservices
mkdir -p ./data/certificates && \
  docker run $(DOCKER_DEFAULT_ARGS loadbalancer) \
      -p 80:80 \
      -p 443:443 \
      --env CERTBOT_EMAIL=zh.nurlan96@gmail.com \
      -v $(pwd)/data/certificates:/etc/letsencrypt \
      gcr.io/kouzoh-p-nurlashko/nurlashko/nginx


docker run $(DOCKER_DEFAULT_ARGS blog) \
    gcr.io/kouzoh-p-nurlashko/nurlashko/blog
