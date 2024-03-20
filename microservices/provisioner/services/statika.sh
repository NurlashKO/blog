mkdir -p "$(pwd)/data/statika/srv"

# Create db file if doesn't exist
db_file="$(pwd)/data/statika/filebrowser.db"
test -f $db_file || touch $db_file

docker run $(DOCKER_DEFAULT_ARGS statika) \
    -v "$(pwd)/data/statika/filebrowser.db:/database.db" \
    -v "$(pwd)/data/statika/srv:/srv" \
    gcr.io/kouzoh-p-nurlashko/nurlashko/statika
