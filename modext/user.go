package modext

import (
	"strings"
)

func DisplayableFields() []string {
	return []string{"id", "name", "email", "mobile"}
}

func VisibleFieldsForRawSelect() string {
	return strings.Join(DisplayableFields(), ",")
}
