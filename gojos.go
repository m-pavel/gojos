// Package gojos provides deserialization of Java objects stored with java ObjectOutpuStream
// Below located high-level functions for deserialization
package gojos

import (
	"io"

	"github.com/m-pavel/gojos/lib/javaos"
	"github.com/m-pavel/gojos/lib/unmarshaller"
)

// Deserialize Java object bytes given to function into JavaModel instance
func Deserialize(reader io.Reader) (*javaos.JavaModel, error) {
	return javaos.Deserialize(reader)
}

// Unmarshal given JavaModel into Go structure provided by dest
// Mapping between java fileds names and Go filds names is 'by name'
// and also can be specified via `java` annotation
func Unmarshal(model *javaos.JavaModel, dest interface{}) error {
	return javaum.Unmarshal(model, dest)
}

// Combination of methods above. Unmarshal stream directly to Go structure
func UnmarshallStream(reader io.Reader, dest interface{}) error {
	jm, err := Deserialize(reader)
	if err != nil {
		return err
	}
	return javaum.Unmarshal(jm, dest)
}
