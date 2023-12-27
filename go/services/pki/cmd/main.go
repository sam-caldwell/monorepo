package main

import (
	"flag"
	"net/http"
)

/*
 * service/pki/cmd/main.go
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * This file defines the pki server used to create and manage a public key cryptography infrastructure, including
 * download the certificate authority (CA) or intermediate certificate authority (ICA) and to submit a certificate-
 * signing request (CSR) and downloading the signed certificate.
 */
func main() {
	listenAddr := flag.String("listenAddr", "0.0.0.0:80", "http listen address:port")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("OK"))
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("OK"))
	})

	http.HandleFunc("/ca", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("OK"))
	})

	http.HandleFunc("/csr", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("OK"))
	})

	http.HandleFunc("/cert", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("OK"))
	})

	err := http.ListenAndServe(*listenAddr, nil)
	if err != nil {
		panic(err)
	}
}
