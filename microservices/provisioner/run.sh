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

# Provision microservices
source <(cat provision/services/*)
