package node

import (
	"crypto/tls"
	"encoding/binary"
	"fmt"
	"io"
	"net"

	"github.com/GirigiriG/cluster/cmd/node"
)

func Run() net.Listener {
	socket := "localhost:5341"
	cfg := node.TLSConfig()
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
