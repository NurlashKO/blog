#!/bin/bash
source _utils.sh

docker run $(DOCKER_DEFAULT_ARGS blog) \
    gcr.io/kouzoh-p-nurlashko/nurlashko/blog
