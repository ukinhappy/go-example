package reflectx

import (
	"errors"
	"go-example/utils/timex"
	"reflect"
	"strconv"
	"time"
)

// ObjectToMap 对象反射成map 结构体成员只支持常用的集中类型
func ObjectToMap(object interface{}, tag string) map[string]string {
	data := make(map[string]string, 0)
	typ := reflect.TypeOf(object).Elem()
	val := reflect.ValueOf(object).Elem()

	if val.Type().Kind() != reflect.Struct {
		return data
	}

	for i := 0; i < typ.NumField(); i++ {
		typeField := typ.Field(i)
		structField := val.Field(i)
		key, ok := typeField.Tag.Lookup(tag)
		if !ok {
			key = typeField.Name
		}
		vType := typeField.Type.Kind()
		switch vType {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			data[key] = strconv.Itoa(int(structField.Int()))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			data[key] = strconv.Itoa(int(structField.Uint()))
		case reflect.String:
			data[key] = structField.String()
		case reflect.Float32, reflect.Float64:
			data[key] = strconv.FormatFloat(structField.Float(), 'f', -1, 64)
		case reflect.Struct:
			if t, ok := structField.Interface().(time.Time); ok {
				if !t.IsZero() {
					data[key] = timex.Timer(t).String()
				} else {
					data[key] = "0000-00-00 00:00:00"
				}

			}
		}
	}

	return data
}

//MapToObject map反射成对象
func MapToObject(ptr interface{}, value map[string]string, tag string) error {
	typ := reflect.TypeOf(ptr).Elem()
	val := reflect.ValueOf(ptr).Elem()
	for i := 0; i < typ.NumField(); i++ {
		typeField := typ.Field(i)
		structField := val.Field(i)
		if !structField.CanSet() {
			continue
		}

		inputFieldName := typeField.Tag.Get(tag)
		if inputFieldName == "" {
			inputFieldName = typeField.Name
		}
		inputValue, exists := value[inputFieldName]
		if !exists {
			continue
		}

		if err := setWithProperType(typeField.Type.Kind(), inputValue, structField); err != nil {
			return err
		}
	}
	return nil
}

func setWithProperType(valueKind reflect.Kind, val string, structField reflect.Value) error {
	switch valueKind {
	case reflect.Int:
		return setIntField(val, 0, structField)
	case reflect.Int8:
		return setIntField(val, 8, structField)
	case reflect.Int16:
		return setIntField(val, 16, structField)
	case reflect.Int32:
		return setIntField(val, 32, structField)
	case reflect.Int64:
		return setIntField(val, 64, structField)
	case reflect.Uint:
		return setUintField(val, 0, structField)
	case reflect.Uint8:
		return setUintField(val, 8, structField)
	case reflect.Uint16:
		return setUintField(val, 16, structField)
	case reflect.Uint32:
		return setUintField(val, 32, structField)
	case reflect.Uint64:
		return setUintField(val, 64, structField)
	case reflect.Bool:
		return setBoolField(val, structField)
	case reflect.Float32:
		return setFloatField(val, 32, structField)
	case reflect.Float64:
		return setFloatField(val, 64, structField)
	case reflect.String:
		structField.SetString(val)
	case reflect.Struct:
		if _, isTime := structField.Interface().(time.Time); isTime {
			structField.Set(reflect.ValueOf(timex.String(val).T))
		}
	default:
		return errors.New("Unknown type")
	}
	return nil
}

func setIntField(val string, bitSize int, field reflect.Value) error {
	if val == "" {
		val = "0"
	}
	intVal, err := strconv.ParseInt(val, 10, bitSize)
	if err == nil {
		field.SetInt(intVal)
	}
	return err
}

func setUintField(val string, bitSize int, field reflect.Value) error {
	if val == "" {
		val = "0"
	}
	uintVal, err := strconv.ParseUint(val, 10, bitSize)
	if err == nil {
		field.SetUint(uintVal)
	}
	return err
}

func setBoolField(val string, field reflect.Value) error {
	if val == "" {
		val = "false"
	}
	boolVal, err := strconv.ParseBool(val)
	if err == nil {
		field.SetBool(boolVal)
	}
	return nil
}

func setFloatField(val string, bitSize int, field reflect.Value) error {
	if val == "" {
		val = "0.0"
	}
	floatVal, err := strconv.ParseFloat(val, bitSize)
	if err == nil {
		field.SetFloat(floatVal)
	}
	return err
}
