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

import (
	"bufio"
	"strings"
)

// TemplateIndent add indent of `width` to each non empty line
func TemplateIndent(in string, width int) string {
	out := ""
	indent := ""
	for i := 0; i < width; i++ {
		indent += " "
	}
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		text := scanner.Text()
		if text != "" {
			out += indent
		}
		out += text
		out += "\n"
	}
	return out
}
