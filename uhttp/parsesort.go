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
