package javaum

import (
	"os"
	"testing"

	"log"

	"github.com/m-pavel/gojos/lib/javaos"
	"github.com/stretchr/testify/assert"
)

func readJavaModel(filename string) (*javaos.JavaModel, error) {
	file, err := os.OpenFile("../../target/tdata/"+filename, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	return javaos.Deserialize(file)
}

type Primitives struct {
	ByteVal  byte
	BoolVal  bool
	CharVal  byte
	ShortVal uint8
	IntVal   int
	FloatVal float32
	LongVal  uint64
	DblVal   float64
}

func TestPrimitives(t *testing.T) {
	jm, err := readJavaModel("primitivestest.bin")
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("Java model %s", jm)
	t1 := Primitives{}
	err = Unmarshall(jm, &t1)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("Go model %s", t1)

	assert.Equal(t, true, t1.BoolVal)
	assert.Equal(t, uint8(55), t1.ByteVal)
	assert.Equal(t, byte('A'), t1.CharVal)
	assert.Equal(t, float64(325.55), t1.DblVal)
	assert.Equal(t, float32(11.7), t1.FloatVal)
	assert.Equal(t, 42, t1.IntVal)
	assert.Equal(t, uint64(55555555555555), t1.LongVal)
	assert.Equal(t, uint8(200), t1.ShortVal)
}

type Hierarchy struct {
	StrVal string `java:"c2str"`
	IntVal int    `java:"p1int"`
}

func TestHierarchy(t *testing.T) {
	jm, err := readJavaModel("testhierarchy.bin")
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("Java model %s", jm)
	t1 := Hierarchy{}
	err = Unmarshall(jm, &t1)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("Go model %s", t1)

	assert.Equal(t, "C1", t1.StrVal)
	assert.Equal(t, (55), t1.IntVal)
}
