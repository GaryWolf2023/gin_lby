package utils

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

func ProcessValidataErrors(errors validator.ValidationErrors, target any) error {
	var errorResult error
	// 通过反射获取只针对某个字段的错误信息
	fields := reflect.TypeOf(target).Elem()
	for _, fieldErr := range errors {
		field, _ := fields.FieldByName(fieldErr.Field())
		errorMessageTag := fmt.Sprintf("%s_err", fieldErr.Tag())
		errorMessage := field.Tag.Get(errorMessageTag)
		if errorMessage == "" {
			errorMessage = field.Tag.Get("message")
		}
		if errorMessage == "" {
			errorMessage = fmt.Sprintf("%s,%s Error", fieldErr.Field(), fieldErr.Tag())
		}
		errorResult = AppendError(errorResult, fmt.Errorf(errorMessage))
	}
	return errorResult
}
