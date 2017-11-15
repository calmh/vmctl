package pretty

import (
	"fmt"
	"reflect"

	"github.com/fatih/structs"
)

func Print(v interface{}) {
	fields := structs.Fields(v)
nextField:
	for _, f := range fields {
		name := f.Name()
		if tag := f.Tag("pretty"); tag == "-" {
			continue nextField
		} else if tag != "" {
			name = tag
		}

		if f.Kind() == reflect.Struct {
			intf := f.Value()
			if s, ok := intf.(fmt.Stringer); ok {
				fmt.Printf("%s: %s\n", name, s.String())
			} else {
				Print(intf)
			}
			return
		}

		format := "%v"
		if tag := f.Tag("format"); tag != "" {
			format = tag
		}

		value := fmt.Sprintf(format, f.Value())
		fmt.Printf("%s: %s\n", name, value)
	}
}
