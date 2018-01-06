package javaos

type endofblockReader struct{}

func (*endofblockReader) Type() byte {
	return TC_ENDBLOCKDATA
}
func (*endofblockReader) Process(s *Stream) RR {
	return RR{Type: TC_ENDBLOCKDATA, Value: nil}
}
