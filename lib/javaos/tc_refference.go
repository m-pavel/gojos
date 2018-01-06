package javaos

type refferenceReader struct{}

func (*refferenceReader) Type() byte {
	return TC_REFERENCE
}
func (*refferenceReader) Process(s *Stream) RR {
	handle, _ := s.ReadUint32()
	return RR{Type: TC_REFERENCE, Value: handle - baseWireHandle}
}
