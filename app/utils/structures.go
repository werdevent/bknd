package utils

import "reflect"

// FilterEmptyStructs will remove the empty filters and leave a map with all the fields that will be replaced
func FilterEmptyStructs(structure interface{}) map[string]interface{} {

	result := make(map[string]interface{})

	value := reflect.ValueOf(structure).Elem()
	typeOfValue := value.Type()

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fieldName := typeOfValue.Field(i).Name

		if field.Interface() != reflect.Zero(field.Type()).Interface() {
			result[fieldName] = field.Interface()
		}

	}
	return result
}
