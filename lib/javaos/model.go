package javaos

type ClassDesc struct {
	Name             string
	SerialVersionUID uint64
	Fields           []FieldDesc
}

func (cd *ClassDesc) Field(name string) *FieldDesc {
	for idx, f := range cd.Fields {
		if f.Name == name {
			return &cd.Fields[idx]
		}
	}
	return nil
}

type FieldDesc struct {
	Typ      byte
	Name     string
	Class    string
	ClassDef JavaModel
	Val      RR
}

type JavaModel struct {
	Classes []ClassDesc
}
