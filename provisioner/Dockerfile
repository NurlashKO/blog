FROM debian:buster-slim

WORKDIR /opts/provision

COPY . .
RUN mkdir tmp

ENTRYPOINT cp /opts/provision/provision.sh tmp/ && sleep infinity

#docker run \
#	-v /home/"$USER"/.docker/config.json:/root/.docker/config.json \
#	-v /var/run/docker.sock:/var/run/docker.sock \
#	-v /home/"$USER"/data:/opts/provision/data \
#	 gcr.io/kouzoh-p-nurlashko/nurlashko/provisioner