package proxy

import (
	"log"
	"log/slog"
	"net/http/httputil"
	"net/url"

	"nurlashko.dev/auth/internal/jwt"
)

func NewStatikaProxyTarget(u string, jwt *jwt.Client) ProxyTarget {
	target, err := url.Parse(u)
	if err != nil {
		log.Fatalf("failed to parse url for statika: %v", err)
	}

	proxy := &httputil.ReverseProxy{
		Rewrite: func(r *httputil.ProxyRequest) {
			forward, err := url.Parse(r.In.Header.Get("X-AUTH-PROXY-FORWARD"))
			if err != nil {
				slog.Error("[statika proxy] failed to parse next proxy jump url: %v", err)
				return
			}
			r.SetURL(forward)
			r.Out.Header.Set("X-AUTH-STATIKA", "anonymous")

			if c, err := r.Out.Cookie("X-AUTH-TOKEN"); err == nil {
				if jwt.VerifyToken(c.Value) {
					r.Out.Header.Set("X-AUTH-STATIKA", "admin")
				}
			}
		},
	}
	proxy.Director = nil

	return ProxyTarget{
		proxy: proxy,
		Host:  target.Host,
	}
}
