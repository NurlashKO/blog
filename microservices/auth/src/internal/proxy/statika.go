package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func NewStatikaReverseProxy() ProxyServer {
	mux := http.NewServeMux()
	target, err := url.Parse("https://static.nurlashko.dev")
	if err != nil {
		log.Fatal("failed to parse url for statika: %v", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.Rewrite = func(r *httputil.ProxyRequest) {
		r.SetURL(target)
		r.Out.Header.Set("X-AUTH-STATIKA", "admin")
		r.SetXForwarded()
	}
	proxy.Director = nil

	return &ReverseProxy{
		mux:   mux,
		proxy: proxy,
		addr:  "0.0.0.0:8001",
	}
}
