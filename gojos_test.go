package gojos

import (
	"os"
	"testing"
	"time"

	"fmt"

	"path/filepath"

	"log"

	"github.com/m-pavel/gojos/lib/javaos"
	javaum "github.com/m-pavel/gojos/lib/unmarshaller"
	"github.com/stretchr/testify/assert"
)

const (
	fldr = "./target/tdata/"
)

func doTest(fname string, t *testing.T) (res *javaos.JavaModel) {
	log.Printf("Testing %v\n", fname)
	file, err := os.OpenFile(fname, os.O_RDONLY, 0644)
	assert.Nil(t, err)
	res, err = Deserialize(file)
	assert.Nil(t, err)
	fmt.Println(res)
	return res
}

func TestAllParse(t *testing.T) {
	err := filepath.Walk(fldr, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			doTest(path, t)
		}
		return nil
	})
	assert.Nil(t, err)
}

type Tkn struct {
	Gen        time.Time `java:"generatedDateTime"`
	ValidUntil time.Time
	User       string `java:"userId"`
	Metadata   map[string]string
}

func Test55(t *testing.T) {
	file, err := os.OpenFile("./target/tdata/test55.bin", os.O_RDONLY, 0644)
	assert.Nil(t, err)
	t1 := Tkn{}
	err = UnmarshallStream(file, &t1)
	assert.Nil(t, err)
	log.Printf("Go model %v", t1)
}

type MapStruct struct {
	Hm map[string]string
}

func TestHashMap(t *testing.T) {
	file, err := os.OpenFile("./target/tdata/test44.bin", os.O_RDONLY, 0644)
	assert.Nil(t, err)
	ms := MapStruct{}
	jm, err := Deserialize(file)
	assert.Nil(t, err)
	log.Printf("Java model %v", jm)

	err = javaum.Unmarshal(jm, &ms)
	assert.Nil(t, err)

	log.Printf("Go model %v", ms)
}

type A struct {
	Value string
}
type B struct {
	A *A
}
type C struct {
	A *A
	B *B
}

func Test8_Referrence(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	file, err := os.OpenFile("./target/tdata/test8.bin", os.O_RDONLY, 0644)
	assert.Nil(t, err)
	c := C{}
	err = UnmarshallStream(file, &c)
	assert.Nil(t, err)
	log.Printf("Go model C %v", c)
	log.Printf("Go model C.A %p", c.A)
	log.Printf("Go model C.B %p", c.B)
	log.Printf("Go model C.B.A %p", c.B.A)
	log.Println(c.A.Value)
	log.Println(c.B.A.Value)
}

type C9 struct {
	D1 time.Time
	D2 time.Time
	S  string
}

func Test9_Referrence(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	file, err := os.OpenFile("./target/tdata/test9.bin", os.O_RDONLY, 0644)
	assert.Nil(t, err)
	c := C9{}
	err = UnmarshallStream(file, &c)
	assert.Nil(t, err)
	log.Println(c.D1)
	log.Println(c.D2)
	log.Println(c.S)
}
