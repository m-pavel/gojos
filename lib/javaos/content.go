package javaos

func readClassData(s *Stream, desc *ClassDesc) interface{} {
	cr := javaClassReaderFor(desc)
	return cr.Read(s, desc)
}
