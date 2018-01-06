package javaos

import (
	"fmt"

	"github.com/cenkalti/log"
)

type defaultJavaClassReader struct {
}

func (*defaultJavaClassReader) Name() string {
	return "__ANY__"
}

func (dr *defaultJavaClassReader) Read(s *Stream, cd *ClassDesc) interface{} {
	for idx, fld := range cd.Fields {
		log.Debugf("Reading field %s\n", fld.Name)
		//cd.Fields[idx].Val = RR{}
		switch fld.Typ {
		case 0x4c: // L - refference
			cd.Fields[idx].Val = procObject(s)
			break
		case 'Z':
			b, _ := s.ReadOne()
			if b == 1 {
				cd.Fields[idx].Val.Value = true
			} else {
				cd.Fields[idx].Val.Value = false
			}
			break
		case 'B':
			cd.Fields[idx].Val.Value, _ = s.ReadOne()
			break
		case 'C':
			cd.Fields[idx].Val.Value, _ = s.ReadUint16()
			break
		case 'S':
			cd.Fields[idx].Val.Value, _ = s.ReadUint16()
			break
		case 'I':
			cd.Fields[idx].Val.Value, _ = s.ReadUint32()
			break
		case 'F':
			cd.Fields[idx].Val.Value, _ = s.ReadFloat32()
			break
		case 'J':
			cd.Fields[idx].Val.Value, _ = s.ReadUint64()
			break
		case 'D': // double
			cd.Fields[idx].Val.Value, _ = s.ReadFloat64()
			break
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
		//cd.Fields[idx].ClassDef = JavaModel{rr.Value.([]ClassDesc)}
		////creader := JavaClassReaderFor(fld.ClassDef.ClassName())
		//if len((cd.Fields[idx].ClassDef.Classes)) != 1 {
		//	panic("Unexpected number of classes.")
		//}
		//creader := classReaderFor(&cd.Fields[idx].ClassDef.Classes[0])
		//if creader != nil {
		//	creader.Read(s)
		//} else {
		//	fmt.Printf("No reader for class %s\n", fld.ClassDef.ClassName())
		//	panic("sad")
		//}
	case TC_STRING:
		return rr
	case TC_NULL:
		return RR{Value: nil, Type: TC_NULL} // TODO special value ?
	case TC_ENDBLOCKDATA:
		return procObject(s)
	default:
		panic(fmt.Sprintf("??? %x", rr.Type))
	}
}
