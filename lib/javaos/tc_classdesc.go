package javaos

import (
	"fmt"
)

type classDescReader struct{}

func (*classDescReader) Type() byte {
	return TC_CLASSDESC
}
func (*classDescReader) Process(s *Stream) RR {
	var res ClassDesc
	s.h.assgn(&res)
	res.Name, _ = s.ReadStringNoHandle() // TODO could be a ref ???
	res.SerialVersionUID, _ = s.ReadUint64()
	s.ReadOne() // flag
	nof, _ := s.ReadUint16()
	res.Fields = make([]FieldDesc, int(nof))
	var fi uint16
	for fi = 0; fi < nof; fi++ {
		res.Fields[fi] = readField(s)
	}
	return RR{Type: TC_CLASSDESC, Value: res}
}

func readField(s *Stream) FieldDesc {
	typ, _ := s.ReadOne()
	fieldName, _ := s.ReadStringNoHandle()
	switch typ {
	case 0x4c: // L
		var className string
		val := readFor(s)
		switch val.Type {
		case TC_STRING:
			className = val.Value.(string)
			break
		case TC_REFERENCE:
			className = s.h.get(val.Value.(uint32)).(string)
			break
		}
		return FieldDesc{Class: className, Name: fieldName, Typ: typ}
	case 0x42: // byte
	case 0x43: // char
	case 0x44: // double
	case 0x46: // float
	case 0x49: // int
	case 0x4a: // long
	case 0x5a: // bool
	case 0x53: // short
	default:
		fmt.Printf("Unknown %x %c %s\n", typ, typ, fieldName)
		break
	}
	return FieldDesc{Name: fieldName, Typ: typ}
}
