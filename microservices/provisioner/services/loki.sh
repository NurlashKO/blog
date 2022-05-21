#!/bin/bash
source _utils.sh

docker run $(DOCKER_DEFAULT_ARGS loki) \
    --mount source=loki-data,target=/loki \
    gcr.io/kouzoh-p-nurlashko/nurlashko/loki
