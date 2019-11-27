// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 349.

// Package params provides a reflection-based parser for URL parameters.
package params

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

//!+Unpack

// Unpack populates the fields of the struct pointed to by ptr
// from the HTTP request parameters in req.
func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	type query struct {
		value     reflect.Value
		validator string
	}

	// Build map of fields keyed by effective name.
	fields := make(map[string]query)
	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}

		validator := tag.Get("validate")
		fields[name] = query { v.Field(i), validator }
	}

	// Update struct field for each parameter in the request.
	for name, values := range req.Form {
		f := fields[name]
		if !f.value.IsValid() {
			continue // ignore unrecognized HTTP parameters
		}
		for _, value := range values {
			if f.validator != "" {
				rules := strings.Split(f.validator, ",")
				for _, r := range rules {
					s := strings.Split(r, "=")

					// max=10のような形式でない場合はpanic
					if len(s) != 2 {
						panic(fmt.Sprintf("unexpected tag value: %s", r))
					}

					switch s[0] {
					case "max":
						limit, err := strconv.ParseInt(s[1], 10, 64)
						if err != nil {
							panic(fmt.Sprintf("invalid number: %s", r))
						}
						if f.value.Kind() != reflect.Int {
							panic(fmt.Sprintf("unexpected type found. expectd: int actual: %s", f.value.Kind().String()))
						}
						fv, err := strconv.ParseInt(value, 10, 64)
						if err != nil {
							panic(fmt.Sprintf("invalid value: %s", value))
						}
						if fv > limit {
							panic(fmt.Sprintf("%s is exceeded limit. limit: %d value: %d", name, limit, f.value.Int()))
						}
					case "min":
						limit, err := strconv.ParseInt(s[1], 10, 64)
						if err != nil {
							panic(fmt.Sprintf("invalid number: %s", r))
						}
						if f.value.Kind() != reflect.Int {
							panic(fmt.Sprintf("unexpected type found. expectd: int actual: %s", f.value.Kind().String()))
						}
						fv, err := strconv.ParseInt(value, 10, 64)
						if err != nil {
							panic(fmt.Sprintf("invalid value: %s", value))
						}
						if fv < limit {
							panic(fmt.Sprintf("%s is falled limit. limit: %d value: %d", name, limit, f.value.Int()))
						}
					default:
						panic(fmt.Sprintf("unexpected token: %s", s[0]))
					}
				}
			}

			if f.value.Kind() == reflect.Slice {
				elem := reflect.New(f.value.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.value.Set(reflect.Append(f.value, elem))
			} else {
				if err := populate(f.value, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

//!-Unpack

//!+populate
func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}

//!-populate

//!+Pack

func Pack(ptr interface{}) (string, error) {
	v := reflect.ValueOf(ptr).Elem()
	q := strings.Builder{}
	for i := 0; i < v.NumField(); i++ {
		if i > 0 {
			q.WriteString("&")
		}
		tag := v.Type().Field(i).Tag.Get("http")
		var name string
		if tag == "" {
			name = v.Type().Field(i).Name
		} else {
			name = tag
		}

		if v.Field(i).Kind() == reflect.Slice {
			for j := 0; j < v.Field(i).Len(); j++ {
				if j > 0 {
					q.WriteString("&")
				}
				q.WriteString(fmt.Sprintf("%s=%s", name, valueToString(v.Field(i).Index(j))))
			}
		} else {
			q.WriteString(fmt.Sprintf("%s=%s", name, valueToString(v.Field(i))))
		}
	}
	return q.String(), nil
}

//!-Pack

func valueToString(v reflect.Value) string {
	switch v.Kind() {
	case reflect.String:
		return v.String()
	case reflect.Int:
		return fmt.Sprintf("%d", v.Int())
	case reflect.Bool:
		return fmt.Sprintf("%t", v.Bool())
	default:
		panic(fmt.Sprintf("unsupported type: %v", v.Kind()))
	}
}
