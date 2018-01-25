package javaos

import (
	"time"
)

type java_util_Date struct{}

func (java_util_Date) Name() string {
	return "java.util.Date"
}
func (d *java_util_Date) Read(s *Stream, cd *ClassDesc) interface{} {
	blk := readFor(s)
	if blk.Type != TC_BLOCKDATA {
		panic("Unexpectable!!!")
	}
	val := blk.blockReadUint64()
	return time.Unix(0, int64(val)*int64(time.Millisecond))
}

func (java_util_Date) newInstance() JavaClassReader {
	return &java_util_Date{}
}
