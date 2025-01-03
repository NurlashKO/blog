package proxy

import (
	"log/slog"
	"net/http"
	"net/http/httputil"
)

type ProxyServer interface {
	StartProxy()
}

type ProxyTarget struct {
	Host  string
	proxy *httputil.ReverseProxy
}

func StartProxy(proxyMap map[string]ProxyTarget) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if proxy, ok := proxyMap[r.Host]; ok {
			slog.Info("proxying to %s", r.Host)
			proxy.proxy.ServeHTTP(w, r)
		} else {
			slog.Error("unknown forward target %s", r.Host)
			w.WriteHeader(http.StatusNotFound)
		}
	})

	slog.Info("Listening on 0.0.0.0:9000")
	if err := http.ListenAndServe("0.0.0.0:9000", mux); err != nil {
		slog.Error("error listening: %v", err)
	}
}
