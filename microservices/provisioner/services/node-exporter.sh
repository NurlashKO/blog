
docker run $(DOCKER_DEFAULT_ARGS node-exporter) \
    -v /proc:/host/proc:ro \
    -v /sys:/host/sys:ro \
    -v /:/rootfs:ro \
    -v /run/dbus/system_bus_socket:/var/run/dbus/system_bus_socket:ro \
    gcr.io/kouzoh-p-nurlashko/nurlashko/node-exporter \
     --path.procfs=/host/proc \
     --path.sysfs=/host/sys \
     --path.rootfs=/rootfs \
     --collector.filesystem.ignored-mount-points='^/(sys|proc|dev|host|etc)($$|/)' \
     --collector.systemd \
     --collector.processes
