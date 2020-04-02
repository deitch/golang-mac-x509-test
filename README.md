# Golang x509 macOS Timing Issue

To test:

1. Clone this repo
1. Ensure you have golang and make installed
1. Run `make`
1. Run `./bin/certs-cgo`
1. Run `./bin/certs-nocgo`

Note that it enables `GODEBUG=x509roots=1` for you.
