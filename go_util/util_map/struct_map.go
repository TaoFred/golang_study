package util_map

import (
	"reflect"
)

// StructToMap converts a struct to a map
func StructToMap(obj interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	objValue := reflect.ValueOf(obj)
	objType := reflect.TypeOf(obj)

	for i := 0; i < objValue.NumField(); i++ {
		field := objType.Field(i)
		value := objValue.Field(i)
		if value.Kind() == reflect.Struct {
			result[field.Name] = StructToMap(value.Interface())
		} else {
			result[field.Name] = value.Interface()
		}
	}
	return result
}
