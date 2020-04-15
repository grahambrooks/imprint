package info

import (
	"fmt"
	"reflect"
)

func Imprint(s interface{}) error {
	if s == nil {
		return nil
	}

	val := reflect.ValueOf(s)

	// If it's an interface or a pointer, unwrap it.
	if val.Kind() == reflect.Ptr && val.Elem().Kind() == reflect.Struct {
		val = val.Elem()
	} else {
		return fmt.Errorf("%s is not a struct", val.Kind())
	}

	valNumFields := val.NumField()

	for i := 0; i < valNumFields; i++ {
		field := val.Field(i)
		fieldKind := field.Kind()

		// Check if it's a pointer to a struct.
		if fieldKind == reflect.Ptr && field.Elem().Kind() == reflect.Struct {
			if field.CanInterface() {
				// Recurse using an interface of the field.
				err := Imprint(field.Interface())
				if err != nil {
					return err
				}
			}

			// Move onto the next field.
			continue
		}

		// Check if it's a struct value.
		if fieldKind == reflect.Struct {
			if field.CanAddr() && field.Addr().CanInterface() {
				// Recurse using an interface of the pointer value of the field.
				err := Imprint(field.Addr().Interface())
				if err != nil {
					return err
				}
			}

			// Move onto the next field.
			continue
		}

		// Check if it's a string or a pointer to a string.
		if fieldKind == reflect.String || (fieldKind == reflect.Ptr && field.Elem().Kind() == reflect.String) {
			typeField := val.Type().Field(i)

			// Get the field tag value.
			tag := typeField.Tag.Get("imprint")

			// Skip if tag is not defined or ignored.
			if tag == "" || tag == "-" {
				continue
			}

			// Set the string value to the sanitized string if it's allowed.
			// It should always be allowed at this point.
			if field.CanSet() {
				replacement, ok := DefinitionMap[tag]
				if ok {
					field.SetString(*replacement.Value)
				} else {
					return fmt.Errorf("missing replacement value for %s", tag)
				}
			}

			continue
		}
	}
	return nil
}
