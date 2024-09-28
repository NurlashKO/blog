| Microservice | Description |
|--------------|-------------|
[![auth CI](https://github.com/NurlashKO/blog/actions/workflows/auth-ci.yml/badge.svg)](https://github.com/NurlashKO/blog/actions/workflows/auth-ci.yml)|authentication service
[![statika CI](https://github.com/NurlashKO/blog/actions/workflows/statika-ci.yml/badge.svg)](https://github.com/NurlashKO/blog/actions/workflows/statika-ci.yml)|Web interface to manage server static files
[![database CI](https://github.com/NurlashKO/blog/actions/workflows/database-ci.yml/badge.svg)](https://github.com/NurlashKO/blog/actions/workflows/database-ci.yml)|Persistant storage for data like blog posts
[![vault CI](https://github.com/NurlashKO/blog/actions/workflows/vault-ci.yml/badge.svg)](https://github.com/NurlashKO/blog/actions/workflows/vault-ci.yml)|Tool for securely storing and accessing secrets
[![Provisioner CI](https://github.com/NurlashKO/blog/actions/workflows/provisioner-ci.yml/badge.svg)](https://github.com/NurlashKO/blog/actions/workflows/provisioner-ci.yml) | Combines all provisioning scripts. Used as versioned package delivery.
[![Watcher CI](https://github.com/NurlashKO/blog/actions/workflows/watcher-ci.yml/badge.svg)](https://github.com/NurlashKO/blog/actions/workflows/watcher-ci.yml) | Pulls then restarts microservice when new version is available
[![Ingress CI](https://github.com/NurlashKO/blog/actions/workflows/ingress-ci.yml/badge.svg)](https://github.com/NurlashKO/blog/actions/workflows/ingress-ci.yml) | All requests from `public` to `internal` network pass through here
[![Loki CI](https://github.com/NurlashKO/blog/actions/workflows/loki-ci.yml/badge.svg)](https://github.com/NurlashKO/blog/actions/workflows/loki-ci.yml) | Logs aggregator. Collects logs from both containers and machine host 
[![Blog CI](https://github.com/NurlashKO/blog/actions/workflows/blog-ci.yml/badge.svg)](https://github.com/NurlashKO/blog/actions/workflows/blog-ci.yml) | blog microservice
[![cAdvisor CI](https://github.com/NurlashKO/blog/actions/workflows/cadvisor.yml/badge.svg)](https://github.com/NurlashKO/blog/actions/workflows/cadvisor.yml) | Provides container metrics like resource usage and performance
[![Promtail CI](https://github.com/NurlashKO/blog/actions/workflows/promtail-ci.yml/badge.svg)](https://github.com/NurlashKO/blog/actions/workflows/promtail-ci.yml) | Provide system logs to Loki
[![Prometheus CI](https://github.com/NurlashKO/blog/actions/workflows/prometheus-ci.yml/badge.svg)](https://github.com/NurlashKO/blog/actions/workflows/prometheus-ci.yml) | Scrape metrics, stores as time series data
[![Node Exporter CI](https://github.com/NurlashKO/blog/actions/workflows/node-exporter-ci.yml/badge.svg)](https://github.com/NurlashKO/blog/actions/workflows/node-exporter-ci.yml) | Provide system metrics
[![Grafana CI](https://github.com/NurlashKO/blog/actions/workflows/grafana-ci.yml/badge.svg)](https://github.com/NurlashKO/blog/actions/workflows/grafana-ci.yml) | Analyrics and visialisaition tool for collected logs and metrics
[![Docs CI](https://github.com/NurlashKO/blog/actions/workflows/docs-ci.yml/badge.svg)](https://github.com/NurlashKO/blog/actions/workflows/docs-ci.yml) | Providing documentation for My Personal Blog

# My Personal Blog

![Untitled drawing (5)](https://user-images.githubusercontent.com/10639020/169705462-f48bc1b2-8883-4b5c-a116-37294ec3c40e.png)

Full documentation can be found here: https://docs.nurlashko.dev
