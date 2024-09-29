package proxy

import (
	"log/slog"
	"net/http"
	"net/http/httputil"
)

type ProxyServer interface {
	StartProxy()
}

type ReverseProxy struct {
	mux   *http.ServeMux
	proxy *httputil.ReverseProxy
	addr  string
}

func (r *ReverseProxy) StartProxy() {
	r.mux.HandleFunc("/", r.proxy.ServeHTTP)

	slog.Info("Listening on " + r.addr)
	if err := http.ListenAndServe(r.addr, r.mux); err != nil {
		slog.Error("error listening: %v", err)
	}
}
