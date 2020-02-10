package javaos

import (
	"fmt"
	"io"

	"github.com/pkg/errors"
)

func Deserialize(reader io.Reader) (*JavaModel, error) {
	stream := NewStream(reader)
	buffer, _ := stream.read(2)
	magic := fmt.Sprintf("%x%x", buffer[0], buffer[1])
	if STREAM_MAGIC != magic {
		return nil, errors.New(fmt.Sprintf("Wrong magic %s", magic))
	}
	// version
	if _, _, err := stream.ReadTwo(); err != nil {
		return nil, err
	}

	rr := readFor(&stream)
	return &JavaModel{Classes: rr.Value.([]ClassDesc)}, nil
}
