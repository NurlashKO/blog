docker run $(DOCKER_DEFAULT_ARGS database) \
    -v "$(pwd)/data/pgdata:/var/lib/postgresql/data" \
    -e POSTGRES_PASSWORD=tmp \
    -e POSTGRES_USER=nurlashko \
    gcr.io/kouzoh-p-nurlashko/nurlashko/database
