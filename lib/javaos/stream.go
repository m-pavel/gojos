package javaos

import (
	"encoding/binary"
	"io"
	"math"
)

type Stream struct {
	r   io.Reader
	h   Handles
	blk *blockData
	c   map[string]*ClassDesc
}

func NewStream(r io.Reader) Stream {
	return Stream{r: r, h: Handles{}, c: map[string]*ClassDesc{}}
}

func (s *Stream) read(count int) ([]byte, error) {
	res := make([]byte, count)
	_, err := s.r.Read(res)
	return res, err
}

func (s *Stream) ReadOne() (byte, error) {
	buffer, err := s.read(1)
	return buffer[0], err
}

func (s *Stream) ReadTwo() (byte, byte, error) {
	buffer, err := s.read(2)
	return buffer[0], buffer[1], err
}

func (s *Stream) ReadUint16() (uint16, error) {
	buffer, err := s.read(2)
	if err == nil {
		return binary.BigEndian.Uint16(buffer), nil
	} else {
		return 0, err
	}

}

func (s *Stream) ReadUint32() (uint32, error) {
	buffer, err := s.read(4)
	if err == nil {
		return binary.BigEndian.Uint32(buffer), nil
	} else {
		return 0, err
	}
}

func (s *Stream) ReadFloat32() (float32, error) {
	x, err := s.ReadUint32()
	if err == nil {
		return math.Float32frombits(x), nil
	} else {
		return 0, nil
	}

}

func (s *Stream) ReadFloat64() (float64, error) {
	x, err := s.ReadUint64()
	if err == nil {
		return math.Float64frombits(x), nil
	} else {
		return 0, nil
	}

}

func (s *Stream) ReadUint64() (uint64, error) {
	buffer, err := s.read(8)
	if err == nil {
		return binary.BigEndian.Uint64(buffer), nil
	} else {
		return 0, err
	}
}

func (s *Stream) ReadString() (string, error) {
	str, err := s.ReadStringNoHandle()
	if err != nil {
		return "", err
	} else {
		s.h.assgn(str)
		return str, nil
	}
}

func (s *Stream) ReadStringNoHandle() (string, error) {
	length, err := s.ReadUint16()
	if err != nil {
		return "", err
	}
	buffer, err := s.read(int(length))
	if err != nil {
		return "", err
	}
	str := string(buffer)
	return str, nil
}

func (s *Stream) HasBlock() bool {
	if s.blk == nil {
		return false
	}
	if s.blk.pos >= len(s.blk.val) {
		return false
	}
	return true
}
