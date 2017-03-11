package utils

import "strings"

func Indent(n int, s string) string {
	spaces := strings.Repeat(" ", n)
	return strings.Join(strings.Split(s, "\n"), "\n"+spaces)
}
