package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"utils"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		bufio.NewReader(request.Body).WriteTo(writer)
	})

	config := &tls.Config{
		CurvePreferences: []tls.CurveID{
			tls.CurveP256,
			tls.X25519,
		},
		MinVersion: tls.VersionTLS12,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
	}

	customServer := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		TLSConfig:         config,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      100 * time.Second,
		IdleTimeout:       120 * time.Second,
	}
	cert, key, err := utils.CreateSimpleSSLCert(
		"2033-Jan-02",
		"192.168.50.180",
		[]string{"private org"},
		[]string{"nothing"},
		"test demo")
	if err != nil {
		fmt.Printf("https server tls error : %v \n", err)
		return
	}

	certFileName, _ := filepath.Abs("./cert.certificate")
	certFile, _ := os.OpenFile(certFileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	cert.WriteTo(certFile)

	keyFileName, _ := filepath.Abs("./key.certificate")
	keyFile, _ := os.OpenFile(keyFileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	key.WriteTo(keyFile)

	err = customServer.ListenAndServeTLS(certFileName, keyFileName)
	if err != nil {
		fmt.Printf("http server error : %v \n", err)
	}
}
