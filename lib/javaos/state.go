package javaos

import (
	"fmt"
)

type RR struct {
	Value   interface{}
	GoValue interface{}
	Type    byte
}
type State interface {
	Type() byte
	Process(s *Stream) RR
}

func stateFor(b byte) State {
	switch b {
	case TC_CLASSDESC:
		return &classDescReader{}
	case TC_STRING:
		return &stringReader{}
	case TC_REFERENCE:
		return &refferenceReader{}
	case TC_OBJECT:
		return &objectReader{}
	case TC_ENDBLOCKDATA:
		return &endofblockReader{}
	case TC_BLOCKDATA:
		return &blockDataReader{}
	case TC_BLOCKDATALONG:
		return &blockLongDataReader{}
	case TC_NULL:
		return &nullReader{}
	default:
		return nil
	}
}

func readFor(s *Stream) RR {
	typ, _ := s.ReadOne()
	state := stateFor(typ)
	if state == nil {
		panic(fmt.Sprintf("Unknown type %x", typ))
	}
	return state.Process(s)
}
