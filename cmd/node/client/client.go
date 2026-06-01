package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/binary"
	"fmt"
	"log"
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
		RootCAs:    certPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
		
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

	SendMessage(conn, payload)
}

func SendMessage(conn net.Conn, message []byte) error {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, uint32(len(message)))
	Must(err)

	_, err = buf.Write(message)
	Must(err)

	_, err = conn.Write(buf.Bytes())
	Must(err)

	return nil
}

func Must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
