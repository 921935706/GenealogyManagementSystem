package validator

import (
	"reflect"
	"strings"
)

// Validate 校验结构体
func Validate(v interface{}) error {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return nil
	}

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		// 检查必填字段
		tag := fieldType.Tag.Get("validate")
		if tag == "required" && isEmpty(field) {
			return ErrRequiredField{Field: fieldType.Name}
		}
	}

	return nil
}

// isEmpty 检查字段是否为空
func isEmpty(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	case reflect.Array, reflect.Slice:
		return v.Len() == 0
	case reflect.Map:
		return v.Len() == 0
	default:
		return false
	}
}

// ErrRequiredField 必填字段错误
type ErrRequiredField struct {
	Field string
}

// Error 实现error接口
func (e ErrRequiredField) Error() string {
	return strings.Title(e.Field) + " is required"
}
