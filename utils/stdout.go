package utils

import (
	"bufio"
	"fmt"
)

type out struct{ writer *bufio.Writer }

func (o *out) Fprintln(v ...interface{}) {
	fmt.Fprintln(o.writer, v...)
}
func (o *out) Fprint(v ...interface{}) {
	fmt.Fprint(o.writer, v...)
}
func (o *out) Fprintf(format string, v ...interface{}) {
	fmt.Fprintf(o.writer, format, v...)
}
func (o *out) Printf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

func (o *out) Flush() error {
	return o.writer.Flush()
}

func (o *out) F() {
	o.Flush()
}
