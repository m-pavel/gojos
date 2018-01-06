package gojos

import (
	"io"

	"github.com/m-pavel/gojos/lib/javaos"
	"github.com/m-pavel/gojos/lib/unmarshaller"
)

func Deserialize(reader io.Reader) (*javaos.JavaModel, error) {
	return javaos.Deserialize(reader)
}

func Unmarshall(model *javaos.JavaModel, dest interface{}) error {
	return javaum.Unmarshall(model, dest)
}

func UnmarshallStream(reader io.Reader, dest interface{}) error {
	jm, err := Deserialize(reader)
	if err != nil {
		return err
	}
	return javaum.Unmarshall(jm, dest)
}
