package proxy

import "C"
import (
	"log"
	"net/http/httputil"
	"net/url"

	"nurlashko.dev/auth/internal/jwt"
)

func NewStatikaProxyTarget(u string, jwt *jwt.Client) ProxyTarget {
	target, err := url.Parse(u)
	if err != nil {
		log.Fatalf("failed to parse url for statika: %v", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.Rewrite = func(r *httputil.ProxyRequest) {
		r.SetURL(target)
		r.Out.Header.Set("X-AUTH-STATIKA", "anonymous")

		if c, err := r.Out.Cookie("X-AUTH-TOKEN"); err == nil {
			if jwt.VerifyToken(c.Value) {
				r.Out.Header.Set("X-AUTH-STATIKA", "admin")
			}
		}
		r.SetXForwarded()
	}
	proxy.Director = nil

	return ProxyTarget{
		proxy: proxy,
		Host:  target.Host,
	}
}
