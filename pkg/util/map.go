/*
Copyright (c) 2018 Alexander Trost <galexrt@googlemail.com>. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

// ConvertMapInterfaceToMapStringInterface convert map[interface{}]interface{} to map[string]interface{}
func ConvertMapInterfaceToMapStringInterface(in map[interface{}]interface{}) map[string]interface{} {
	out := make(map[string]interface{})
	for key, value := range in {
		switch key := key.(type) {
		case string:
			out[key] = value
		}
	}
	return out
}

// ConvertMapInterfaceToStringSlice convert
func ConvertInterfaceSliceToStringSlice(in []interface{}) []string {
	out := make([]string, len(in))
	for _, value := range in {
		switch value := value.(type) {
		case string:
			out = append(out, value)
		}
	}
	return out
}
