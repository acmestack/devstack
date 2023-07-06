package structcp

import (
	"reflect"
	"testing"
	
	"github.com/acmestack/devstack/logging"
)

func Test(t *testing.T) {
	type Target struct {
		Name   string
		Addr   int
		Other1 string
		Other2 string
	}
	tests := []struct {
		name   string
		source any
		target Target
		align  bool
	}{
		{
			name: "convert case 0",
			source: struct {
				Name   string
				Addr   string
				Other0 string
				Other2 string
			}{Name: "name value", Addr: "addr value", Other2: "other2 value"},
			target: Target{Name: "bb", Other1: "other1"},
			align:  true,
		},
		{
			name: "convert case 1",
			source: struct {
				Name1   string
				Addr1   string
				Other01 string
				Other21 string
			}{Name1: "name value", Addr1: "addr value", Other21: "other2 value"},
			target: Target{Name: "bb", Other1: "other1"},
			align:  false,
		},
		{
			name: "convert case 2",
			source: struct {
				Name1   string `alignTo:"Name"`
				Addr1   string
				Other01 string `alignTo:"Other1"`
				Other21 string
			}{Name1: "name value", Addr1: "addr value", Other01: "other01 value"},
			target: Target{Name: "bb", Other1: "other1"},
			align:  true,
		},
	}
	
	logger := logging.InitLogger("info")
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			logger.Sugar.Infof("case name %s", test.name)
			Align(&test.target, test.source)
			if test.align {
				sourceValueOf := reflect.ValueOf(test.source)
				targetValueOf := reflect.ValueOf(&test.target).Elem()
				
				sourceTypeOf := reflect.TypeOf(test.source)
				sourceFields := reflect.VisibleFields(sourceTypeOf)
				for i, field := range sourceFields {
					logger.Sugar.Infof("file index %d, field %v", i, field)
					value := targetValueOf.FieldByName(field.Name)
					
					sourceValue := sourceValueOf.FieldByName(field.Name).String()
					if value.Kind() != reflect.Invalid && value.Type() == field.Type && !reflect.DeepEqual(value.String(), sourceValue) {
						t.Error("align error")
					}
					logger.Sugar.Infof("value %s, sourcevalue %s", value.String(), sourceValue)
				}
			}
			logger.Sugar.Infof("bb addr: %d, name: %s, others1 %s", test.target.Addr, test.target.Name, test.target.Other1)
		})
	}
}
