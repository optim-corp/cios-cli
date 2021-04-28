package log

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/optim-corp/cios-cli/utils/go_advance_type/check"
	"github.com/optim-corp/cios-cli/utils/go_advance_type/convert"
)

func InterfacesToStrings(arg []interface{}) []string {
	stringArray := make([]string, len(arg))
	for i, v := range arg {
		stringArray[i] = convert.MustStr(v)
	}
	return stringArray
}

func getInfo() (FilePath, ProgramLine, MethodName) {
	pt, file, line, _ := runtime.Caller(skip)
	return FilePath(file + ":" + strconv.Itoa(line)), ProgramLine(line), MethodName(filepath.Base(runtime.FuncForPC(pt).Name()))
}

func getContextValue(ctx context.Context, key interface{}) string {
	value := ctx.Value(key)
	fmt.Println("test", key, value)
	if check.IsNil(value) {
		return fmt.Sprintf("%s: ", convert.MustStr(key))
	}
	return fmt.Sprintf("%s: %s", convert.MustStr(key), convert.MustStr(value))

}
func getContextValues(ctx context.Context, keys []interface{}) (result string) {
	for _, key := range keys {
		result += getContextValue(ctx, key) + ", "
	}
	return
}
