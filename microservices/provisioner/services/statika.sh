mkdir -p "$(pwd)/data/statika/srv"
mkdir -p "$(pwd)/data/statika/database"

docker run $(DOCKER_DEFAULT_ARGS statika) \
    -v "$(pwd)/data/statika/database:/database" \
    -v "$(pwd)/data/statika/srv:/srv" \
    gcr.io/kouzoh-p-nurlashko/nurlashko/statika
