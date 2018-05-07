package ajson

import (
	"strings"
	"unicode"
)

// Object simulates a Javascript Object
type Object map[string]interface{}

func (o Object) StringAt(path string) (string, bool) {
	val, ok := o.ValueAt(path)
	if !ok {
		return "", false
	}

	str, ok := val.(string)
	return str, ok
}

func (o Object) BytesAt(path string) ([]byte, bool) {
	val, ok := o.ValueAt(path)
	if !ok {
		return nil, false
	}

	b, ok := val.([]byte)
	return b, ok
}

func (o Object) BoolAt(path string) (bool, bool) {
	val, ok := o.ValueAt(path)
	if !ok {
		return false, false
	}

	b, ok := val.(bool)
	return b, ok
}

func (o Object) Float32At(path string) (float32, bool) {
	val, ok := o.ValueAt(path)
	if !ok {
		return 0, false
	}

	f, ok := val.(float32)
	return f, ok
}

func (o Object) Float64At(path string) (float64, bool) {
	val, ok := o.ValueAt(path)
	if !ok {
		return 0, false
	}

	f, ok := val.(float64)
	return f, ok
}

func (o Object) Int8At(path string) (int8, bool) {
	val, ok := o.ValueAt(path)
	if !ok {
		return 0, false
	}

	i, ok := val.(int8)
	return i, ok
}

func (o Object) Int16At(path string) (int16, bool) {
	val, ok := o.ValueAt(path)
	if !ok {
		return 0, false
	}

	i, ok := val.(int16)
	return i, ok
}

func (o Object) Int32At(path string) (int32, bool) {
	val, ok := o.ValueAt(path)
	if !ok {
		return 0, false
	}

	i, ok := val.(int32)
	return i, ok
}

func (o Object) Int64At(path string) (int64, bool) {
	val, ok := o.ValueAt(path)
	if !ok {
		return 0, false
	}

	i, ok := val.(int64)
	return i, ok
}

func (o Object) IntAt(path string) (int, bool) {
	val, ok := o.ValueAt(path)
	if !ok {
		return 0, false
	}

	i, ok := val.(int)
	return i, ok
}

func (o Object) Uint8At(path string) (uint8, bool) {
	val, ok := o.ValueAt(path)
	if !ok {
		return 0, false
	}

	u, ok := val.(uint8)
	return u, ok
}

func (o Object) Uint16At(path string) (uint16, bool) {
	val, ok := o.ValueAt(path)
	if !ok {
		return 0, false
	}

	u, ok := val.(uint16)
	return u, ok
}

func (o Object) Uint32At(path string) (uint32, bool) {
	val, ok := o.ValueAt(path)
	if !ok {
		return 0, false
	}

	u, ok := val.(uint32)
	return u, ok
}

func (o Object) Uint64At(path string) (uint64, bool) {
	val, ok := o.ValueAt(path)
	if !ok {
		return 0, false
	}

	u, ok := val.(uint64)
	return u, ok
}

func (o Object) UintAt(path string) (uint, bool) {
	val, ok := o.ValueAt(path)
	if !ok {
		return 0, false
	}

	u, ok := val.(uint)
	return u, ok
}

func (o Object) ObjectAt(path string) (Object, bool) {
	val, ok := o.ValueAt(path)
	if !ok {
		return nil, false
	}

	obj, ok := val.(Object)
	return obj, ok
}

func (o Object) ObjectArrayAt(path string) (ObjectArray, bool) {
	val, ok := o.ValueAt(path)
	if !ok {
		return nil, false
	}

	arr, ok := val.(ObjectArray)
	return arr, ok
}

func (o Object) ArrayAt(path string) (Array, bool) {
	val, ok := o.ValueAt(path)
	if !ok {
		return nil, false
	}

	arr, ok := val.(Array)
	return arr, ok
}

func (o Object) ValueAt(path string) (interface{}, bool) {
	if o == nil {
		return nil, false
	}

	path = strings.TrimFunc(path, trimPath)
	if path == "" {
		return o, true
	}

	step := o
	parts := strings.Split(path, ".")
	lp := len(parts)
	for i := 0; i < lp; i++ {
		istep, ok := step[parts[i]]
		if !ok {
			return nil, false
		}

		if i != lp-1 {
			step, ok = istep.(Object)
			if !ok {
				return nil, false
			}

			continue
		}

		return istep, true
	}

	return nil, false
}

func trimPath(c rune) bool {
	return c == '.' || unicode.IsSpace(c)
}

type ObjectArray []Object

type Array []interface{}
