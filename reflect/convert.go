/*
 * Licensed to the AcmeStack under one or more contributor license
 * agreements. See the NOTICE file distributed with this work for
 * additional information regarding copyright ownership.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package reflect

import (
	"reflect"
	
	"github.com/acmestack/devstack/logging"
)

const (
	AlignToTag = "alignTo"
)

func Convert(target any, source any) {
	sourceValueOf := reflect.ValueOf(source)
	targetValueOf := reflect.ValueOf(target).Elem()
	
	sourceTypeOf := reflect.TypeOf(source)
	sourceFields := reflect.VisibleFields(sourceTypeOf)
	for i, field := range sourceFields {
		logging.Logger.Infof("file index %d, field %v", i, field)
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
