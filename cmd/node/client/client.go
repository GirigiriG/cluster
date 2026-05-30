package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	args := os.Args

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

	message := strings.Join(args[1:], " ")
	SendMessage(conn, message)
}

func SendMessage(conn net.Conn, message string) error {
	data := []byte(message)

	// Sending message lenghth
	err := binary.Write(conn, binary.BigEndian, uint32(len(data)))
	if err != nil {
		return err
	}
	// send the message
	_, err = conn.Write(data)
	if err != nil {
		return err
	}

	return nil
}
