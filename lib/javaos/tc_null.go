package javaos

type nullReader struct{}

func (*nullReader) Type() byte {
	return TC_NULL
}
func (*nullReader) Process(s *Stream) RR {
	return RR{Type: TC_NULL, Value: nil}
}
