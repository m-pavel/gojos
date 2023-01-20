package javaos

type JavaClassReader interface {
	Name() string
	Read(s *Stream, cd *ClassDesc) interface{}
	ReadFromBlock(bd *blockData) interface{}
	newInstance() JavaClassReader
}

var javaClasses = []JavaClassReader{
	&java_util_Date{},
	&java_util_HashMap{},
}

func javaClassReaderFor(desc *ClassDesc) JavaClassReader {
	for _, jc := range javaClasses {
		if jc.Name() == desc.Name {
			return jc.newInstance()
		}
	}
	return &defaultJavaClassReader{}
}

func hasCustomClassReader(class string) bool {
	for _, jc := range javaClasses {
		if jc.Name() == class {
			return true
		}
	}
	return false
}
