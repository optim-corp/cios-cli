package console

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/AlecAivazis/survey/v2"
)

var (
	writer = bufio.NewWriter(os.Stdout)
)

func Fprintln(v ...interface{}) {
	fmt.Fprintln(writer, v...)
}
func Fprint(v ...interface{}) {
	fmt.Fprint(writer, v...)
}
func Fprintf(format string, v ...interface{}) {
	fmt.Fprintf(writer, format, v...)
}
func Printf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

func Flush() error {
	return writer.Flush()
}

func F() {
	Flush()
}
func SetWriter(w io.Writer) {
	writer = bufio.NewWriter(w)
}

func ReadMultiLine(message string) string {
	ans := struct{ Body string }{}
	Question([]*survey.Question{
		{Name: "body", Prompt: &survey.Multiline{Message: message}},
	}, &ans)
	return ans.Body
}
