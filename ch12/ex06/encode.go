// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 339.

package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

//!+Marshal
// Marshal encodes a Go value in S-expression form.
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v), 0); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

//!-Marshal

// encode writes to buf an S-expression representation of v.
//!+encode
func encode(buf *bytes.Buffer, v reflect.Value, depth int) error {
	if v.IsZero() {
		return nil
	}
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("null")

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return encode(buf, v.Elem(), depth)

	case reflect.Array, reflect.Slice: // (value ...)
		buf.WriteByte('[')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte(',')
			}
			buf.WriteByte('\n')
			buf.WriteString(strings.Repeat("\t", depth+1))
			if err := encode(buf, v.Index(i), depth+1); err != nil {
				return err
			}
		}
		buf.WriteByte(']')

	case reflect.Struct: // ((name value) ...)
		buf.WriteString(strings.Repeat("\t", depth))
		buf.WriteByte('{')
		count := 0
		for i := 0; i < v.NumField(); i++ {
			if v.Field(i).IsZero() {
				continue
			}
			if count > 0 {
				buf.WriteByte(',')
			}
			buf.WriteByte('\n')
			buf.WriteString(strings.Repeat("\t", depth+1))
			fmt.Fprintf(buf, "%q: ", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i), depth+1); err != nil {
				return err
			}
			count++
		}
		buf.WriteString("\n")
		buf.WriteString(strings.Repeat("\t", depth))
		buf.WriteByte('}')

	case reflect.Map: // ((key value) ...)
		buf.WriteByte('{')
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteByte(',')
			}

			buf.WriteString("\n")
			buf.WriteString(strings.Repeat("\t", depth+1))

			switch key.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16,
				reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16,
				reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64:
				buf.WriteString("\"")
				if err := encode(buf, key, depth+1); err != nil {
					return err
				}
				buf.WriteString("\"")
			default:
				if err := encode(buf, key, depth+1); err != nil {
					return err
				}
			}

			buf.WriteByte(':')
			if err := encode(buf, v.MapIndex(key), depth+1); err != nil {
				return err
			}
		}
		buf.WriteByte('}')

	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%g", v.Float())
	case reflect.Complex64, reflect.Complex128:
		fmt.Fprintf(buf, "#C(%g, %g)", real(v.Complex()), imag(v.Complex()))
	case reflect.Bool:
		if v.Bool() {
			fmt.Fprint(buf, "true")
		} else {
			fmt.Fprint(buf, "false")
		}
	case reflect.Chan:
		fmt.Fprintf(buf, "%s", v.Type())
	case reflect.Func:
		fmt.Fprintf(buf, "%s", v.Type())
	case reflect.Interface:
		fmt.Fprintf(buf, "%s", v.Type())
	default:
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

//!-encode
