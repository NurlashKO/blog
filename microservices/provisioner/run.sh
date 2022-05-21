#!/bin/sh

cd /home/nurlashko

# Cleanup everything
docker kill $(docker ps -f "label=$PROVISIONER_MARK" -q)
docker system prune -f

# Setup networking
docker network create internal

# Setup filesystem shared with microservices
mkdir -p data

# Provision microservices
for f in services/*.sh; do
  ./"$f"
done
