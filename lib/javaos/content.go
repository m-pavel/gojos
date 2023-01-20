package javaos

func readClassData(s *Stream, desc *ClassDesc, bd *blockData) interface{} {
	cr := javaClassReaderFor(desc)
	if bd == nil {
		return cr.Read(s, desc)
	} else {
		return cr.ReadFromBlock(bd)
	}
}
