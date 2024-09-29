FROM filebrowser/filebrowser:v2.27.0

COPY src/config.json /.filebrowser.json

ENTRYPOINT /filebrowser config set --auth.method=proxy --auth.header=X-AUTH-STATIKA
