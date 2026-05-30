package node

import (
	"crypto/tls"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
)

func Run() net.Listener {
	socket := "localhost:5341"
	cfg := tlsConfig()
	listener, err := tls.Listen("tcp", socket, cfg)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Println("started on " + socket)
	return listener
}

func HandleConnection(conn net.Conn) {
	for {
		var length uint32
		err := binary.Read(conn, binary.BigEndian, &length)

		if err != nil {
			if err == io.EOF {
				// fmt.Println("client disconnected")
			} else {
				fmt.Println("error reading message length :: ", err)

			}
			return
		}

		buffer := make([]byte, length)
		_, err = io.ReadFull(conn, buffer)

		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(string(buffer))
	}
}

func tlsConfig() *tls.Config {
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
