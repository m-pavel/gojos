package javaos

import (
	"log"
	"reflect"
)

type Java_util_HashMap_Entry struct {
	Key   interface{}
	Value interface{}
}

type java_util_HashMap struct {
	buckets uint32
	size    int
	entries []Java_util_HashMap_Entry
}

func (java_util_HashMap) Name() string {
	return "java.util.HashMap"
}
func (java_util_HashMap) newInstance() JavaClassReader {
	return &java_util_HashMap{}
}

func (hm *java_util_HashMap) Read(s *Stream, cd *ClassDesc) interface{} {
	//
	cd.Field("loadFactor").Val = RR{}
	cd.Field("loadFactor").Val.Value, _ = s.ReadFloat32()
	cd.Field("threshold").Val = RR{}
	cd.Field("threshold").Val.Value, _ = s.ReadUint32()
	bd := readFor(s)
	block := bd.Value.(*blockData)

	hm.buckets = block.ReadUint32()
	hm.size = int(block.ReadUint32())
	hm.entries = make([]Java_util_HashMap_Entry, hm.size)
	if hm.size > 0 {
		for i := 0; i < hm.size; i++ {
			key := readFor(s)
			var keyv interface{}
			if key.Type == TC_REFERENCE {
				keyv = s.h.get(key.Value.(uint32))
			} else {
				keyv = key.Value
			}
			value := readFor(s)
			var vv interface{}
			if value.Type == TC_REFERENCE {
				vv = s.h.get(value.Value.(uint32))
			} else {
				vv = value.Value
			}
			hm.entries[i] = Java_util_HashMap_Entry{Key: keyv, Value: vv}
		}
	}
	gomap := tryToMakeMap(*hm)
	return gomap
}

func (hm *java_util_HashMap) ReadFromBlock(bd *blockData) interface{} {
	panic("TODO Implement")
}

func tryToMakeMap(hm java_util_HashMap) interface{} {
	var mp reflect.Value
	var ktyp reflect.Type
	if len(hm.entries) == 0 {
		log.Println("Empty map")
	} else {
		ktyp = reflect.TypeOf(hm.entries[0].Key)
		if ktyp == nil {
			if len((hm.entries)) == 1 {
				panic("Unable to determine key type")
			}
			ktyp = reflect.TypeOf(hm.entries[1].Key)
		}
		// TODO validate for valid key type
		var vtype reflect.Type
		for i := range hm.entries {
			if hm.entries[i].Value != nil {
				vtype = reflect.MapOf(ktyp, reflect.TypeOf(hm.entries[i].Value))
			}
		}
		if vtype == nil {
			vtype = reflect.MapOf(ktyp, reflect.TypeOf(reflect.Interface))
		}
		mp = reflect.MakeMap(vtype)
	}
	//mp := make(map[interface{}]interface{}, 0)
	for _, e := range hm.entries {
		if e.Key == nil {
			if ktyp.Kind() == reflect.String {
				mp.SetMapIndex(reflect.ValueOf("__NULL__"), reflect.ValueOf(e.Value))
			} else {
				panic("TODO add type")
			}
		} else {
			mp.SetMapIndex(reflect.ValueOf(e.Key), reflect.ValueOf(e.Value))
		}
	}
	return mp
}
