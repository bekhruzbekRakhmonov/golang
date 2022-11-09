package main

import (
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

func main() {
	m := &autocert.Manager{
		Cache:      autocert.DirCache("secret-dir"),
		Prompt:     autocert.AcceptTOS,
		Email:      "example@example.org",
		HostPolicy: autocert.HostWhitelist("example.org", "www.example.org"),
	}
	s := &http.Server{
		Addr:      ":https",
		TLSConfig: m.TLSConfig(),
	}
	s.ListenAndServeTLS(":8001", "")
}