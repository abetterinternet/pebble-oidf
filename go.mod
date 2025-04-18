module github.com/letsencrypt/pebble/v2

go 1.23.5

require (
	github.com/go-jose/go-jose/v4 v4.0.4
	github.com/letsencrypt/challtestsrv v1.3.2
	github.com/miekg/dns v1.1.62
	github.com/tgeoghegan/oidf-box v0.0.1
)

require (
	github.com/go-errors/errors v1.5.1 // indirect
	golang.org/x/crypto v0.32.0 // indirect
	golang.org/x/mod v0.22.0 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/tools v0.29.0 // indirect
)

replace github.com/tgeoghegan/oidf-box => ../oidf-box
