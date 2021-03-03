package utils

import (
	"fmt"
	"unicode/utf8"
)

func ListUtility(print func()) {
	fmt.Fprintln(Out, "\n********************************************************"+
		"********************************************************\n")
	print()
	fmt.Fprintln(Out, "\n********************************************************"+
		"********************************************************\n")
	Out.Flush()
}

func SpaceRight(val string, len int) string {
	valLen := utf8.RuneCountInString(val)
	for i := 1; 0 < (len - valLen); i++ {
		val += " "
		valLen = utf8.RuneCountInString(val)

	}
	return val
}
