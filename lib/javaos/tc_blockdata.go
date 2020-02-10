package javaos

import (
	"encoding/binary"
	"math"
)

type blockDataReader struct{}
type blockLongDataReader struct{}

type blockData struct {
	val []byte
	pos int
}

func (*blockDataReader) Type() byte {
	return TC_BLOCKDATA
}
func (*blockDataReader) Process(s *Stream) RR {
	size, _ := s.ReadOne()
	bd := blockData{pos: 0}
	bd.val, _ = s.read(int(size))
	return RR{Type: TC_BLOCKDATA, Value: &bd}
}

func (*blockLongDataReader) Type() byte {
	return TC_BLOCKDATALONG
}
func (*blockLongDataReader) Process(s *Stream) RR {
	size, _ := s.ReadUint32()
	buffer, _ := s.read(int(size))
	return RR{Type: TC_BLOCKDATALONG, Value: buffer}
}

func (bd *blockData) ReadUint32() uint32 {
	bd.pos += 4
	return binary.BigEndian.Uint32(bd.val[bd.pos-4 : bd.pos])
}

func (bd *blockData) ReadByte() byte {
	bd.pos += 1
	return (bd.val[bd.pos-1 : bd.pos])[0]
}

func (bd *blockData) ReadUint16() uint16 {
	bd.pos += 2
	return binary.BigEndian.Uint16(bd.val[bd.pos-2 : bd.pos])
}

func (bd *blockData) ReadUint64() uint64 {
	bd.pos += 8
	return binary.BigEndian.Uint64(bd.val[bd.pos-8 : bd.pos])
}

func (bd *blockData) ReadFloat32() float32 {
	return math.Float32frombits(bd.ReadUint32())
}

func (r *RR) blockReadUint32() uint32 {
	return r.Value.(*blockData).ReadUint32()
}

func (r *RR) blockReadUint64() uint64 {
	return r.Value.(*blockData).ReadUint64()
}

//func (r *RR) blockReadFloat32() float32 {
//	return r.Value.(*blockData).ReadFloat32()
//}
