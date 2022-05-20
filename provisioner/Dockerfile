FROM debian:buster-slim

WORKDIR /opts/provision

COPY . .
RUN mkdir tmp

ENTRYPOINT cp /opts/provision/run.sh tmp/ && sleep infinity
