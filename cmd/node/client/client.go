package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"os"

	"github.com/GirigiriG/cluster/internal/protocol"
)

func main() {
	cert, err := os.ReadFile("../../../certs/server.crt")
	if err != nil {
		panic(err)
	}

	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(cert)

	cfg := &tls.Config{
		RootCAs: certPool,
	}

	conn, err := tls.Dial("tcp", "localhost:5341", cfg)
	if err != nil {
		fmt.Println("error connection to remote :: ", err)
		return
	}
	defer conn.Close()

	payload, err := protocol.EncodeNode(protocol.Node{
		ID:   42,
		Name: "Gideon",
	})

	fmt.Println(protocol.DecodeNode(payload))
	SendMessage(conn, payload)
}

func SendMessage(conn net.Conn, message []byte) error {

	_, err := conn.Write(message)
	if err != nil {
		return err
	}

	return nil
}
