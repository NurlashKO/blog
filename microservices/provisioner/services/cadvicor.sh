docker run $(DOCKER_DEFAULT_ARGS cadvisor) \
    --volume=/:/rootfs:ro \
    --volume=/var/run:/var/run:ro \
    --volume=/sys:/sys:ro \
    --volume=/var/lib/docker/:/var/lib/docker:ro \
    --volume=/dev/disk/:/dev/disk:ro \
    --privileged \
    --device=/dev/kmsg \
  gcr.io/kouzoh-p-nurlashko/nurlashko/cadvisor
