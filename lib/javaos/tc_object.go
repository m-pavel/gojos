package javaos

import (
	"fmt"
	"io"
)

type objectReader struct{}

func (*objectReader) Type() byte {
	return TC_OBJECT
}
func (*objectReader) Process(s *Stream) RR {
	classes := make([]ClassDesc, 0)
	done := true
	var gov interface{}
	for done {
		typ, err := s.ReadOne()
		if io.EOF == err {
			done = false
			break
		}
		switch typ {
		case TC_CLASSDESC:
			rr := stateFor(typ).Process(s)
			cd := rr.Value.(ClassDesc)
			s.h.assgn(cd)
			classes = append(classes, cd)
			break
		case TC_ENDBLOCKDATA:
			continue
		case TC_NULL:
			done = false
			break
		case TC_REFERENCE:
			rr := stateFor(typ).Process(s)
			cd := s.h.get(rr.Value.(uint32)).(*ClassDesc)
			gov = readClassData(s, cd)
			break
		default:
			panic(fmt.Sprintf("objectReader %x", typ))
		}
	}

	for i := len(classes) - 1; i >= 0; i-- {
		gov = readClassData(s, &classes[i])
	}
	return RR{Type: TC_OBJECT, Value: classes, GoValue: gov}
}
