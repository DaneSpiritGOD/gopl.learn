package display

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
)

const maxDepth = 10

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x), os.Stdout, 0)
}

func display(path string, v reflect.Value, writer io.Writer, depth int) {
	if depth >= maxDepth {
		fmt.Fprintf(writer, "%s...\n", path)
		return
	}

	switch v.Kind() {
	case reflect.Invalid:
		fmt.Fprintf(writer, "%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i), writer, depth+1)
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i), writer, depth+1)
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			var fieldPath string
			switch key.Kind() {
			case reflect.Array, reflect.Struct:
				buf := new(bytes.Buffer)

				fmt.Fprintf(buf, "%s[\n", path)
				display("$", key, buf, depth+1)
				fmt.Fprint(buf, "]")

				fieldPath = buf.String()
			default:
				fieldPath = fmt.Sprintf("%s[%s]", path, formatAtom(key))
			}

			display(fieldPath, v.MapIndex(key), writer, depth+1)
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Fprintf(writer, "%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem(), writer, depth) // do not increase pointer's depth
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Fprintf(writer, "%s = nil\n", path)
		} else {
			fmt.Fprintf(writer, "%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem(), writer, depth+1)
		}
	default: // basic types, channels, funcs
		fmt.Fprintf(writer, "%s = %s\n", path, formatAtom(v))
	}
}

// formatAtom formats a value without inspecting its internal structure.
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}
