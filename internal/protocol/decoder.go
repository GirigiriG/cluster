package protocol

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
)

func DecodeNode(data []byte) (*Node, error) {
	var n Node
	r := bytes.NewReader(data)

	if err := binary.Read(r, binary.BigEndian, &n.ID); err != nil {
		log.Fatal(err)
	}

	var nameLen uint16
	if err := binary.Read(r, binary.BigEndian, &nameLen); err != nil {
		log.Fatal(err)
	}

	name := make([]byte, nameLen)
	if _, err := io.ReadFull(r, name); err != nil {
		log.Fatal(err)
	}

	n.Name = string(name)
	return &n, nil
}
