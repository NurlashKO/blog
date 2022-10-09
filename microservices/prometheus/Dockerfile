FROM prom/prometheus:v2.35.0

COPY prometheus.yml /etc/prometheus/prometheus.yml

 CMD [
    "--config.file=/etc/prometheus/prometheus.yml",
    "--web.console.libraries=/usr/share/prometheus/console_libraries",
    "--web.console.templates=/usr/share/prometheus/consoles",
    "--storage.tsdb.path=/prometheus",
    "--storage.tsdb.retention.size=256MB"
 ]