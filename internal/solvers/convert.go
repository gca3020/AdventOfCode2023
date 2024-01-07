package solvers

import "strings"

func toLines(bytes []byte) []string {
	return strings.Split(strings.TrimSpace(string(bytes)), "\n")
}
