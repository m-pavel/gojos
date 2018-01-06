package javaos

type stringReader struct{}

func (*stringReader) Type() byte {
	return TC_STRING
}
func (*stringReader) Process(s *Stream) RR {
	str, _ := s.ReadString()
	return RR{Type: TC_STRING, Value: str}
}
