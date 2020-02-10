package javaos

import (
	"fmt"

	"github.com/cenkalti/log"
)

type defaultJavaClassReader struct {
}

func (defaultJavaClassReader) newInstance() JavaClassReader {
	return &defaultJavaClassReader{}
}
func (defaultJavaClassReader) Name() string {
	return "__ANY__"
}

func (dr *defaultJavaClassReader) Read(s *Stream, cd *ClassDesc) interface{} {
	for idx, fld := range cd.Fields {
		log.Debugf("Reading field %s\n", fld.Name)
		//cd.Fields[idx].Val = RR{}
		switch fld.Typ {
		case 0x4c: // L - refference
			cd.Fields[idx].Val = procObject(s)
		case 'Z':
			b, _ := s.ReadOne()
			if b == 1 {
				cd.Fields[idx].Val.Value = true
			} else {
				cd.Fields[idx].Val.Value = false
			}
		case 'B':
			cd.Fields[idx].Val.Value, _ = s.ReadOne()
		case 'C':
			cd.Fields[idx].Val.Value, _ = s.ReadUint16()
		case 'S':
			cd.Fields[idx].Val.Value, _ = s.ReadUint16()
		case 'I':
			cd.Fields[idx].Val.Value, _ = s.ReadUint32()
		case 'F':
			cd.Fields[idx].Val.Value, _ = s.ReadFloat32()
		case 'J':
			cd.Fields[idx].Val.Value, _ = s.ReadUint64()
		case 'D': // double
			cd.Fields[idx].Val.Value, _ = s.ReadFloat64()
		default:
			panic(fmt.Sprintf("Unknown field %x %s\n", fld.Typ, fld.Name))
		}
	}
	return cd
}

func procObject(s *Stream) RR {
	rr := readFor(s)
	switch rr.Type {
	case TC_OBJECT:
		return rr
	case TC_STRING:
		return rr
	case TC_NULL:
		return RR{Value: nil, Type: TC_NULL} // TODO special value ?
	case TC_ENDBLOCKDATA:
		return procObject(s)
	case TC_REFERENCE:
		rv := s.h.get(rr.Value.(uint32))
		log.Panic(rv)
		return rr
	default:
		panic(fmt.Sprintf("Unknown type 0x%x", rr.Type))
	}
}
