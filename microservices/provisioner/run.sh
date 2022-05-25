#!/bin/bash

cd /home/nurlashko

PROVISIONER_MARK="__provisioner-managed__"

DOCKER_DEFAULT_ARGS() {
  name="$1";
  echo "-detach --log-driver json-file --log-opt tag=\"{{.ImageName}}|{{.Name}}|{{.ImageFullID}}|{{.FullID}}\" --network internal --hostname ${name} --name ${name} --label ${PROVISIONER_MARK}"
}

# Choose microservices to re-provision
RESTART=(
blog
# Everything
# `docker ps -f "label=$PROVISIONER_MARK" -q`
)
docker kill $"${RESTART[@]}"

# Cleanup
docker system prune -f

# Setup networking
docker network create internal

# Setup filesystem shared with microservices
mkdir -p data
mkdir -p credentials

# Provision microservices
source <(cat provision/services/*)
