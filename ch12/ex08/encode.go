// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 339.

package sexpr

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"reflect"
)

type Encoder struct {
	w io.Writer
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}
func (enc *Encoder) Encode(v interface{}) error {
	return encode(enc.w, reflect.ValueOf(v))
}

//!+Marshal
// Marshal encodes a Go value in S-expression form.
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

//!-Marshal

// encode writes to buf an S-expression representation of v.
//!+encode
func encode(buf io.Writer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Fprint(buf, "nil")

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return encode(buf, v.Elem())

	case reflect.Array, reflect.Slice: // (value ...)
		fmt.Fprint(buf, "(")
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				fmt.Fprint(buf, " ")
			}
			if err := encode(buf, v.Index(i)); err != nil {
				return err
			}
		}
		fmt.Fprint(buf, ")")

	case reflect.Struct: // ((name value) ...)
		fmt.Fprint(buf, "(")
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				fmt.Fprint(buf, " ")
			}
			fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i)); err != nil {
				return err
			}
			fmt.Fprint(buf, ")")
		}
		fmt.Fprint(buf, ")")

	case reflect.Map: // ((key value) ...)
		fmt.Fprint(buf, "(")
		for i, key := range v.MapKeys() {
			if i > 0 {
				fmt.Fprint(buf, " ")
			}
			fmt.Fprint(buf, "(")
			if err := encode(buf, key); err != nil {
				return err
			}
			fmt.Fprint(buf, " ")
			if err := encode(buf, v.MapIndex(key)); err != nil {
				return err
			}
			fmt.Fprint(buf, ")")
		}
		fmt.Fprint(buf, ")")

	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%g", v.Float())
	case reflect.Complex64, reflect.Complex128:
		fmt.Fprintf(buf, "#C(%g, %g)", real(v.Complex()), imag(v.Complex()))
	case reflect.Bool:
		if v.Bool() {
			fmt.Fprint(buf, "t")
		} else {
			fmt.Fprint(buf, "nil")
		}
	case reflect.Interface:
		fmt.Fprint(buf, "(")
		if v.Type().Name() == "" {
			fmt.Fprint(buf, "")
		} else {
			fmt.Fprintf(buf, "%s.%s", v.Type().PkgPath(), v.Type().Name())
		}
		if err := encode(buf, v.Elem()); err != nil {
			return errors.Wrap(err, "reflect.Interface failed")
		}
		fmt.Fprint(buf, ")")
	default:
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

//!-encode
