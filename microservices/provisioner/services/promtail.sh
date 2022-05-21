docker run $(DOCKER_DEFAULT_ARGS promtail) \
    -v /var/lib/docker/:/var/lib/docker:ro \
    -v /var/log/journal/:/var/log/journal/:ro \
    -v /run/log/journal/:/run/log/journal/:ro \
    -v /etc/machine-id:/etc/machine-id:ro \
    gcr.io/kouzoh-p-nurlashko/nurlashko/promtail
