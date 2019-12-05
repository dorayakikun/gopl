package cyclic

import (
	"fmt"
	"reflect"
	"unsafe"
)

type comparison struct {
	ptr unsafe.Pointer
	t   reflect.Type
}

func IsCyclic(x interface{}) bool {
	v := reflect.ValueOf(x)
	seen := make(map[comparison]bool)
	return isCyclic(v, seen)
}

func isCyclic(x reflect.Value, seen map[comparison]bool) bool {
	if !x.IsValid() {
		panic(fmt.Sprintf("x is invalid: %v", x))
	}

	//!-
	//!+cyclecheck
	// cycle check
	if x.CanAddr() {
		xptr := unsafe.Pointer(x.UnsafeAddr())
		c := comparison{xptr, x.Type()}
		if seen[c] {
			return true // already seen
		}
		seen[c] = true
	}

	switch x.Kind() {
	case reflect.Ptr:
		return isCyclic(x.Elem(), seen)
	case reflect.Array, reflect.Slice:
		for i := 0; i < x.Len(); i++ {
			ret := isCyclic(x.Index(i), seen)
			if ret {
				return ret
			}
		}
	case reflect.Struct:
		for i, n := 0, x.NumField(); i < n; i++ {
			ret := isCyclic(x.Field(i), seen)
			if ret {
				return ret
			}
		}
	case reflect.Map:
		for _, k := range x.MapKeys() {
			ret := isCyclic(x.MapIndex(k), seen)
			if ret {
				return ret
			}
		}
	}
	return false
}
