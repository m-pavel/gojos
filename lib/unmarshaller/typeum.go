package javaum

import (
	"log"
	"reflect"

	"strings"

	"github.com/m-pavel/gojos/lib/javaos"
)

type TypeUnmarshaller interface {
	MyType(t reflect.Type) bool
	Unmarshal(value reflect.Value, cls interface{})
}

type defaultStructUm struct{}

func (defaultStructUm) MyType(t reflect.Type) bool {
	return t.Kind() == reflect.Struct
}
func (defaultStructUm) Unmarshal(value reflect.Value, cls interface{}) {
	t := value.Type()
	for i := 0; i < t.NumField(); i++ {
		ftype := t.Field(i).Type
		tgt := value.Field(i)
		ptr := false
		if ftype.Kind() == reflect.Ptr {
			ftype = t.Field(i).Type.Elem()
			if value.Field(i).IsNil() {
				tgt = reflect.New(ftype)
				value.Field(i).Set(tgt)
				tgt = tgt.Elem()
			}
			ptr = true
		}
		fu := unmarshallerFor(ftype)
		if fu == nil {
			log.Printf("No unmarshaller for filed %s of type %s(%T)", t.Field(i).Name, t.Field(i).Type.Name(), t.Field(i).Type)
		} else {
			jf := findField(t.Field(i), cls)
			if jf == nil {
				log.Printf("Unable to find java mapping for %s. Is it lowercase field ?", t.Field(i).Name)
			} else {
				if !ptr && !tgt.CanSet() {
					log.Printf("Unable to set field %s", t.Field(i).Name)
				} else {
					fu.Unmarshal(tgt, *jf)
				}
			}
		}
	}
}

func goFieldName(gf reflect.StructField) *string {
	tag := gf.Tag.Get("java")
	if tag != "" {
		return &tag
	}
	// Ignore GO 'private' fields
	if strings.ToLower(gf.Name[0:1]) == gf.Name[0:1] {
		return nil
	}
	lname := strings.ToLower(gf.Name[0:1]) + gf.Name[1:]
	return &lname
}
func findField(gf reflect.StructField, jf interface{}) *javaos.FieldDesc {
	gname := goFieldName(gf)
	if gname == nil {
		return nil
	}
	if jm, ok := jf.(*javaos.JavaModel); ok {
		for _, cls := range jm.Classes {
			for _, fld := range cls.Fields {
				if *gname == fld.Name {
					return &fld
				}
			}
		}
	}
	// sub object
	if fd, ok := jf.(javaos.FieldDesc); ok {
		if cls, ok := fd.Val.Value.([]javaos.ClassDesc); ok {
			for _, cls := range cls {
				for _, fld := range cls.Fields {
					if *gname == fld.Name {
						return &fld
					}
				}
			}
		} else {
			return findField(gf, &fd.ClassDef)
		}
	}
	log.Println(gf)
	return nil
}

type dateUm struct{}

func (dateUm) MyType(t reflect.Type) bool {
	return t.Name() == "Time" && t.PkgPath() == "time"
}
func (dateUm) Unmarshal(value reflect.Value, cls interface{}) {
	if fd, ok := cls.(javaos.FieldDesc); ok {
		cd := fd.Val.GoValue
		cdt := reflect.TypeOf(cd)
		if cdt != nil {
			if cdt.AssignableTo(value.Type()) {
				value.Set(reflect.ValueOf(cd))
			} else {
				log.Printf("value of type %s is not assignable to type %s", cdt.Name(), value.Type().Name())
			}
		}
	}
}

type stringUm struct{}

func (stringUm) MyType(t reflect.Type) bool {
	return t.Name() == "string"
}
func (stringUm) Unmarshal(value reflect.Value, cls interface{}) {
	if reflect.TypeOf(cls).Name() == "FieldDesc" {
		cd := cls.(javaos.FieldDesc).Val.Value
		if reflect.TypeOf(cd) != nil {
			value.SetString(cls.(javaos.FieldDesc).Val.Value.(string))
		}
	}
}

type mapUm struct{}

func (mapUm) MyType(t reflect.Type) bool {
	return t.Kind() == reflect.Map
}
func (mapUm) Unmarshal(value reflect.Value, cls interface{}) {
	cd := cls.(javaos.FieldDesc).Val.GoValue
	if reflect.TypeOf(cd) != nil && reflect.TypeOf(cd).Kind() == reflect.Struct {
		if cd.(reflect.Value).Type().Kind() == reflect.Map {
			value.Set(reflect.MakeMap(cd.(reflect.Value).Type()))
			keys := cd.(reflect.Value).MapKeys()
			for _, key := range keys {
				value.SetMapIndex(key, cd.(reflect.Value).MapIndex(key))
			}
		}
	}

}

type intUm struct{}
type uintUm struct{}

func (intUm) MyType(t reflect.Type) bool {
	return t.Kind() == reflect.Int ||
		t.Kind() == reflect.Int8 ||
		t.Kind() == reflect.Int16 ||
		t.Kind() == reflect.Int32 ||
		t.Kind() == reflect.Int64
}

func (uintUm) MyType(t reflect.Type) bool {
	return t.Kind() == reflect.Uint ||
		t.Kind() == reflect.Uint8 ||
		t.Kind() == reflect.Uint16 ||
		t.Kind() == reflect.Uint32 ||
		t.Kind() == reflect.Uint64
}

func (intUm) Unmarshal(value reflect.Value, cls interface{}) {
	if reflect.TypeOf(cls).Name() == "FieldDesc" {
		cd := cls.(javaos.FieldDesc).Val.Value
		if reflect.TypeOf(cd) != nil {
			var intVal int64
			switch reflect.TypeOf(cd).Kind() {
			case reflect.Int:
				intVal = int64(cls.(javaos.FieldDesc).Val.Value.(int))
			case reflect.Int8:
				intVal = int64(cls.(javaos.FieldDesc).Val.Value.(int8))
			case reflect.Int16:
				intVal = int64(cls.(javaos.FieldDesc).Val.Value.(int16))
			case reflect.Int32:
				intVal = int64(cls.(javaos.FieldDesc).Val.Value.(int32))
			case reflect.Int64:
				intVal = int64(cls.(javaos.FieldDesc).Val.Value.(int64))
			case reflect.Uint:
				intVal = int64(cls.(javaos.FieldDesc).Val.Value.(uint))
			case reflect.Uint8:
				intVal = int64(cls.(javaos.FieldDesc).Val.Value.(uint8))
			case reflect.Uint16:
				intVal = int64(cls.(javaos.FieldDesc).Val.Value.(uint16))
			case reflect.Uint32:
				intVal = int64(cls.(javaos.FieldDesc).Val.Value.(uint32))
			case reflect.Uint64:
				intVal = int64(cls.(javaos.FieldDesc).Val.Value.(uint64))
			}
			value.SetInt(intVal)
		}
	}
}

func (uintUm) Unmarshal(value reflect.Value, cls interface{}) {
	if reflect.TypeOf(cls).Name() == "FieldDesc" {
		cd := cls.(javaos.FieldDesc).Val.Value
		if reflect.TypeOf(cd) != nil {
			var intVal uint64
			switch reflect.TypeOf(cd).Kind() {
			case reflect.Uint:
				intVal = uint64(cls.(javaos.FieldDesc).Val.Value.(uint))
			case reflect.Uint8:
				intVal = uint64(cls.(javaos.FieldDesc).Val.Value.(uint8))
			case reflect.Uint16:
				intVal = uint64(cls.(javaos.FieldDesc).Val.Value.(uint16))
			case reflect.Uint32:
				intVal = uint64(cls.(javaos.FieldDesc).Val.Value.(uint32))
			case reflect.Uint64:
				intVal = uint64(cls.(javaos.FieldDesc).Val.Value.(uint64))
			}
			value.SetUint(intVal)
		}
	}
}

type boolUm struct{}

func (boolUm) MyType(t reflect.Type) bool {
	return t.Kind() == reflect.Bool
}

func (boolUm) Unmarshal(value reflect.Value, cls interface{}) {
	if reflect.TypeOf(cls).Name() == "FieldDesc" {
		cd := cls.(javaos.FieldDesc).Val.Value
		if reflect.TypeOf(cd) != nil {
			switch reflect.TypeOf(cd).Kind() {
			case reflect.Bool:
				value.SetBool(cls.(javaos.FieldDesc).Val.Value.(bool))
			}

		}
	}
}

type floatUm struct{}

func (floatUm) MyType(t reflect.Type) bool {
	return t.Kind() == reflect.Float32 ||
		t.Kind() == reflect.Float64
}

func (floatUm) Unmarshal(value reflect.Value, cls interface{}) {
	if reflect.TypeOf(cls).Name() == "FieldDesc" {
		cd := cls.(javaos.FieldDesc).Val.Value
		if reflect.TypeOf(cd) != nil {
			var fVal float64
			switch reflect.TypeOf(cd).Kind() {
			case reflect.Float32:
				fVal = float64(cls.(javaos.FieldDesc).Val.Value.(float32))
			case reflect.Float64:
				fVal = float64(cls.(javaos.FieldDesc).Val.Value.(float64))
			}
			value.SetFloat(fVal)
		}
	}
}

var typeUm = []TypeUnmarshaller{
	&dateUm{},
	&stringUm{},
	&mapUm{},
	&intUm{},
	&uintUm{},
	&boolUm{},
	&floatUm{},
	&defaultStructUm{}, //last option
}

func unmarshallerFor(p reflect.Type) TypeUnmarshaller {
	for _, tu := range typeUm {
		if tu.MyType(p) {
			return tu
		}
	}
	return nil
}
