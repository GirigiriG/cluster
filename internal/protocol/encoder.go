package protocol

import (
	"bytes"
	"encoding/binary"
)

func EncodeNode(n Node) ([]byte, error) {
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.BigEndian, n.ID)
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.BigEndian, uint16(len(n.Name)))
	if err != nil {
		return nil, err
	}

	buf.WriteString(n.Name)

	return buf.Bytes(), nil
}
