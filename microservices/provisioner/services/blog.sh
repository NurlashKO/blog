docker run $(DOCKER_DEFAULT_ARGS blog) \
      -v $(pwd)/data/statika/srv/images:/www/data/images \
    gcr.io/kouzoh-p-nurlashko/nurlashko/blog

