package javaos

type Handles struct {
	data []interface{}
}

func NewHandles() *Handles {
	h := Handles{}
	h.data = make([]interface{}, 0)
	return &h
}

func (h *Handles) assgn(entry interface{}) int {
	h.data = append(h.data, entry)
	return len(h.data)
}

func (h *Handles) get(idx uint32) interface{} {
	return h.data[idx]
}
