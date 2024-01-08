package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	CACertFilePath     = "/tmp/ca-trust/ca-bundle.crt"
	ServerCertFilePath = "/tmp/tls/tls.crt"
	ServerKeyFilePath  = "/tmp/tls/tls.key"
	ServerName         = "hello-world"
	ServerPort         = "8082"
)

var (
	cacert     string
	cert       string
	key        string
	port       string
	servername string
)

func httpRequestHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello,World!\n"))
}

func main() {

	flag.StringVar(&cacert, "cacert", CACertFilePath, "ca certificate filepath")
	flag.StringVar(&cert, "cert", ServerCertFilePath, "certificate filepath")
	flag.StringVar(&key, "key", ServerKeyFilePath, "private key filepath")
	flag.StringVar(&port, "port", ServerPort, "server port")
	flag.StringVar(&servername, "servername", ServerName, "server name for SNI validation")
	flag.Parse()

	log.Printf("Parsed flags: cacert: %s, cert: %s, key: %s, port: %s, servername: %s",
		cacert, cert, key, port, servername)

	// load tls certificates
	serverTLSCert, err := tls.LoadX509KeyPair(cert, key)
	if err != nil {
		log.Fatalf("failed to load key pair: %v", err)
	}

	// Configure the server to trust TLS client cert issued by your CA.
	certPool := x509.NewCertPool()
	if caCertPEM, err := os.ReadFile(cacert); err != nil {
		log.Fatalf("failed to read CA certificate file: %v", err)
	} else if ok := certPool.AppendCertsFromPEM(caCertPEM); !ok {
		log.Fatalf("failed to load CA certificates")
	}

	tlsConfig := &tls.Config{
		ClientAuth:   tls.VerifyClientCertIfGiven,
		ClientCAs:    certPool,
		Certificates: []tls.Certificate{serverTLSCert},
		MinVersion:   tls.VersionTLS12,
		MaxVersion:   tls.VersionTLS13,
		RootCAs:      certPool,
		ServerName:   servername,
	}
	log.Printf("Server TLS config: %+v\n", tlsConfig)
	server := http.Server{
		Addr:      fmt.Sprintf(":%s", port),
		Handler:   http.HandlerFunc(httpRequestHandler),
		TLSConfig: tlsConfig,
	}
	defer server.Close()

	log.Println("Server accessible at address", server.Addr)
	log.Fatal(server.ListenAndServeTLS("", ""))
}
