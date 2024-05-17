package dto

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/go-playground/validator/v10"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var validate = validator.New()

func ValidateParamsToDTO(data map[string]string, dto interface{}) error {
	for key, value := range data {
		caser := cases.Title(language.English)
		fieldName := caser.String(key)

		field := reflect.ValueOf(dto).Elem().FieldByName(fieldName)

		if !field.IsValid() {
			return fmt.Errorf("no such field: %s in dto struct", fieldName)
		}

		if !field.CanSet() {
			return fmt.Errorf("cannot set %s field value", fieldName)
		}

		switch field.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			intValue, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return fmt.Errorf("%s field is not a valid int", field.Type().Name())
			}
			field.SetInt(intValue)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			uintValue, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return fmt.Errorf("%s field is not a valid uint", field.Type().Name())
			}
			field.SetUint(uintValue)
		case reflect.Float32, reflect.Float64:
			floatValue, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return fmt.Errorf("%s field is not a valid float", field.Type().Name())
			}
			field.SetFloat(floatValue)
		case reflect.Bool:
			boolValue, err := strconv.ParseBool(value)
			if err != nil {
				return fmt.Errorf("%s field is not a valid bool", field.Type().Name())
			}
			field.SetBool(boolValue)
		case reflect.String:
			field.SetString(value)
		default:
			return fmt.Errorf("unsupported type: %s", field.Type().Name())
		}

	}

	if err := validate.Struct(dto); err != nil {
		return fmt.Errorf("validation error for key %s", err.Error())
	}

	return nil
}

func ValidateDTO(dto interface{}) error {
	return validate.Struct(dto)
}
