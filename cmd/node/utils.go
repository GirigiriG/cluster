package node

import (
	"crypto/tls"
	"log"
	"os"
	"path/filepath"
)

func TLSConfig() *tls.Config {
	wd, _ := os.Getwd()
	serverCert := filepath.Join(wd, "certs", "server.crt")
	serverKey := filepath.Join(wd, "certs", "server.key")

	cert, err := tls.LoadX509KeyPair(serverCert, serverKey)
	if err != nil {
		log.Panic(err)
	}

	return &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

}
