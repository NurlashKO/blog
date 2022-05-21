docker run $(DOCKER_DEFAULT_ARGS watcher) \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -v $(pwd)/.docker/config.json:/config.json \
    gcr.io/kouzoh-p-nurlashko/nurlashko/watcher
