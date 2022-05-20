FROM docker:20.10

WORKDIR /opt/microservices

COPY . .

RUN find */. -type f ! -name '.pro' -delete

ENTRYPOINT ["sh", "provisioner"]