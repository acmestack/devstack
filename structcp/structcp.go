package structcp

import (
	"reflect"
)

const (
	AlignToTag = "alignTo"
)

func Align(target any, source any) {
	sourceValueOf := reflect.ValueOf(source)
	targetValueOf := reflect.ValueOf(target).Elem()

	sourceTypeOf := reflect.TypeOf(source)
	sourceFields := reflect.VisibleFields(sourceTypeOf)
	for _, field := range sourceFields {
		targetFieldName := field.Name
		tag := field.Tag.Get(AlignToTag)
		if tag != "" {
			targetFieldName = tag
		}
		value := targetValueOf.FieldByName(targetFieldName)
		if value.Kind() != reflect.Invalid && value.Type() == field.Type {
			value.Set(sourceValueOf.FieldByName(field.Name))
		}
	}
}
