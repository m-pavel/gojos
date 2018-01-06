package javaos

func classReaderFor(desc *ClassDesc) JavaClassReader {
	creader := javaClassReaderFor(desc)
	if creader == nil {
		return &defaultJavaClassReader{}
	} else {
		return creader
	}
}

func readClassData(s *Stream, desc *ClassDesc) interface{} {
	cr := classReaderFor(desc)
	return cr.Read(s, desc)
}
