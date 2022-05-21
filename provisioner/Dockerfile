FROM debian:buster-slim

WORKDIR /opts/provision

COPY . .
RUN mkdir tmp

ENTRYPOINT cp /opts/provision/run.sh tmp/ && cp /opts/provision/run.sh tmp/run-backup.sh && sleep infinity
