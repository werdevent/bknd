package utils

import (
	"fmt"
	"reflect"
	"strings"
)

// FilterEmptyStructs will remove the empty filters and leave a map with all the fields that will be replaced
func FilterEmptyStructs(structure interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	value := reflect.ValueOf(structure).Elem()
	typeOfValue := value.Type()

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fieldName := typeOfValue.Field(i).Name

		if field.Type().Kind() != reflect.Struct {
			if field.Interface() != reflect.Zero(field.Type()).Interface() {
				result[strings.ToLower(fieldName)] = field.Interface()
			}
		}

	}
	return result
}

// FilterEmptyStructsWithSuffix will remove the empty filters and leave a map with all the fields that will be replaced including a suffix
func FilterEmptyStructsWithSuffix(structure interface{}, suffix string) map[string]interface{} {
	result := make(map[string]interface{})

	value := reflect.ValueOf(structure).Elem()
	typeOfValue := value.Type()

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fieldName := typeOfValue.Field(i).Name

		if field.Type().Kind() != reflect.Struct {
			if field.Interface() != reflect.Zero(field.Type()).Interface() {
				result[fmt.Sprintf("%v.%v", suffix, strings.ToLower(fieldName))] = field.Interface()
			}
		}

	}
	return result
}
