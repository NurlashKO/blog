docker run $(DOCKER_DEFAULT_ARGS grafana) \
    -v grafana-storage:/var/lib/grafana \
    gcr.io/kouzoh-p-nurlashko/nurlashko/grafana
