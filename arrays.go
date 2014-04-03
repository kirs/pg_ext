package pg_ext

import (
	"fmt"
	"strings"
)

func ToArray(elements []string) string {
	return fmt.Sprintf("{%s}", strings.Join(elements, ", "))
}

func ParseArray(raw string) []string {
	// cut first and last symbols
	items := raw[1 : len(raw)-1]
	return strings.Split(items, ",")
}
