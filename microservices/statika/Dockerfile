FROM filebrowser/filebrowser:v2.27.0

COPY src/config.json /.filebrowser.json

ADD src/run.sh run.sh
RUN chmod +x run.sh

ENTRYPOINT ./run.sh
