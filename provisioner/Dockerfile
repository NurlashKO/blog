FROM docker:20.10

WORKDIR /opts/provision

COPY . .

ENTRYPOINT ["./provision.sh"]

#RUN docker run \
#	-v /home/"$USER"/.docker/config.json:/root/.docker/config.json \
#	-v /var/run/docker.sock:/var/run/docker.sock \
#	-v /home/"$USER"/data:/opts/provision/data \
#	 gcr.io/kouzoh-p-nurlashko/nurlashko/provisioner