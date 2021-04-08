package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
)

func encode(buf *bytes.Buffer, v reflect.Value, prefixSpacesWhenNewLine int) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")

	case reflect.Bool:
		boolGetter := func() string {
			if v.Bool() {
				return "t"
			} else {
				return "nil"
			}
		}

		fmt.Fprintf(buf, "%s", boolGetter())

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%f", v.Float())

	case reflect.Complex64, reflect.Complex128:
		vc := v.Complex()
		fmt.Fprintf(buf, "#C(%f %f)", real(vc), imag(vc))

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return encode(buf, v.Elem(), prefixSpacesWhenNewLine)

	case reflect.Array, reflect.Slice: // (value ...)
		buf.WriteByte('(')
		prefixSpacesWhenNewLine2 := prefixSpacesWhenNewLine + 1
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				writePrefixSpaces(buf, prefixSpacesWhenNewLine2)
			}
			if err := encode(buf, v.Index(i), prefixSpacesWhenNewLine2); err != nil {
				return err
			}
		}
		buf.WriteByte(')')

	case reflect.Struct: // ((name value) ...)
		var fieldName string
		var prefixSpacesWhenNewLine2 int

		buf.WriteByte('(')
		for i := 0; i < v.NumField(); i++ {
			fieldName = fmt.Sprintf("(%s ", v.Type().Field(i).Name)

			if i > 0 {
				writePrefixSpaces(buf, prefixSpacesWhenNewLine+1)
			} else {
				prefixSpacesWhenNewLine2 = prefixSpacesWhenNewLine + len(fieldName)
			}

			buf.WriteString(fieldName)

			if err := encode(buf, v.Field(i), prefixSpacesWhenNewLine2); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')

	case reflect.Map: // ((key value) ...)
		buf.WriteByte('(')
		prefixSpacesWhenNewLine2 := prefixSpacesWhenNewLine + 1
		for i, key := range v.MapKeys() {
			if i > 0 {
				writePrefixSpaces(buf, prefixSpacesWhenNewLine2)
			}
			buf.WriteByte('(')
			if err := encode(buf, key, prefixSpacesWhenNewLine2); err != nil {
				return err
			}
			buf.WriteByte(' ')
			if err := encode(buf, v.MapIndex(key), prefixSpacesWhenNewLine2); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')

	case reflect.Interface:
		buf.WriteByte('(')

		elem := v.Elem()
		typeString := fmt.Sprintf("%q ", elem.Type().String())

		fmt.Fprint(buf, typeString)
		encode(buf, elem, prefixSpacesWhenNewLine+len(typeString))

		buf.WriteByte(')')

	default: // float, complex, bool, chan, func, interface
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

func writePrefixSpaces(buf *bytes.Buffer, prefixSpacesWhenNewLine int) {
	buf.WriteByte('\n')
	for i := 0; i < prefixSpacesWhenNewLine; i++ {
		buf.WriteByte(' ')
	}
}

// Marshal encodes a Go value in S-expression form.
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v), 0); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
