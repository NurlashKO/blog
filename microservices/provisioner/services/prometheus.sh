docker run $(DOCKER_DEFAULT_ARGS prometheus) \
    gcr.io/kouzoh-p-nurlashko/nurlashko/prometheus \
      --storage.tsdb.retention.size=256MB
