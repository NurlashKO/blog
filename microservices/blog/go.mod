module nurlashko.dev/blog

go 1.23

toolchain go1.23.4

require (
	github.com/a-h/templ v0.3.819
	github.com/gomarkdown/markdown v0.0.0-20231222211730-1d6d20845b47
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/lib/pq v1.10.9
	github.com/xeonx/timeago v1.0.0-rc5
	nurlashko.dev/auth v0.0.0-00010101000000-000000000000
)

require (
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.1 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/go-secure-stdlib/strutil v0.1.2 // indirect
	github.com/hashicorp/vault-client-go v0.4.3 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/ryanuber/go-glob v1.0.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/time v0.0.0-20220922220347-f3bd1da661af // indirect
)

replace nurlashko.dev/auth => github.com/nurlashko/blog/microservices/auth/src v0.0.0-20250103094058-60cd8fe15dbe
