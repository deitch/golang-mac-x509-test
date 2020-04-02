package main

import (
	_ "github.com/deitch/golang-mac-x509-test/init"
	"crypto/x509"
	"log"
	"time"
)

func main() {
	start := time.Now()
	if _, err := x509.SystemCertPool(); err != nil {
		log.Fatalf("failed: %v", err)
	}
	elapsed := time.Since(start)
	log.Printf("complete: took %s", elapsed)
}
