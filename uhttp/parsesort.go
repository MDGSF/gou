// MIT License
//
// Copyright (c) 2019 Huang Jian
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package uhttp

import "strings"

/*
ParseSortParameter parse api sort parameters
example:
sortval = "created_at,-name"
sortNames = ["created_at", "name"]
sortOrder = ["asc", "desc"]
*/
func ParseSortParameter(sortval string) ([]string, []string) {

	sortNames := make([]string, 0)
	sortOrder := make([]string, 0)

	if len(sortval) == 0 {
		return sortNames, sortOrder
	}

	values := strings.Split(sortval, ",")

	for _, v := range values {

		if len(v) == 0 {
			continue
		}

		var name string
		var order string

		if v[0] == '-' {
			name = v[1:]
			order = "desc"
		} else {
			name = v
			order = "asc"
		}

		sortNames = append(sortNames, name)
		sortOrder = append(sortOrder, order)
	}

	return sortNames, sortOrder
}
