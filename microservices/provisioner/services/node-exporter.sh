
docker run $(DOCKER_DEFAULT_ARGS node-exporter) \
    -v /proc:/host/proc:ro \
    -v /sys:/host/sys:ro \
    -v /:/rootfs:ro \
    gcr.io/kouzoh-p-nurlashko/nurlashko/node-exporter --collector.systemd --collector.processes
