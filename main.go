package main

import (
	"fmt"

	node "github.com/GirigiriG/cluster/cmd/node/server"
)

func main() {
	ln := node.Run()
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("connect failure :: host %v\n", conn.RemoteAddr())
			continue
		}

		go node.HandleConnection(conn)
	}
}
