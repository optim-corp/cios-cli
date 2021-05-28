package xstring

import wrp "github.com/fcfcqloow/go-advance/wrapper"

func ToOneLine(arg string) string {
	return wrp.MakeString(arg).
		DeletesAll("\n", "\r", "\t", "  ").
		Str()
}
