package reflection

import "reflect"

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func walk(x any, fn func(string)) {
	val := getValue(x)

	if val.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
		return
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		// si es un string se llama al método
		case reflect.String:
			fn(field.String())
		// si es un struct se llama a walk utilizando los atributos del struct
		case reflect.Struct:
			walk(field.Interface(), fn)
		case reflect.Slice:
			for i := 0; i < val.Len(); i++ {
				walk(val.Index(i).Interface(), fn)
			}
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}
