FROM grafana/loki:2.5.0

COPY promtail-config.yaml /cfg/loki-config.yaml

CMD ["-config.file", "/cfg/loki-config.yaml"]
