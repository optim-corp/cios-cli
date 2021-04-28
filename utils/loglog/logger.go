package log

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/optim-corp/cios-cli/utils/go_advance_type/convert"
)

type (
	LogLevel    int
	FilePath    string
	ProgramLine int
	MethodName  string
)

const (
	LOG_LEVEL_TRACE     LogLevel = 401
	LOG_LEVEL_DEBUG     LogLevel = 400
	LOG_LEVEL_INFO      LogLevel = 300
	LOG_LEVEL_WARN      LogLevel = 200
	LOG_LEVEL_ERROR     LogLevel = 100
	LOG_LEVEL_EMERGENCY LogLevel = -1
	LOG_LEVEL_FATAL     LogLevel = -2
	LOG_LEVEL_DISRUPT   LogLevel = -999
	DEFAULT_SKIP        int      = 2
	DEFAULT_LOG_LEVEL   LogLevel = LOG_LEVEL_INFO
)

var (
	logLevel    = DEFAULT_LOG_LEVEL
	skip        = DEFAULT_SKIP
	keys        = []interface{}{}
	TraceOutput = func(values []interface{}, filePath FilePath, line ProgramLine, name MethodName) []string {
		return []string{"[", string(filePath), " / ", string(name), "]: ", strings.Join(InterfacesToStrings(values), " ")}
	}
	DebugOutput = func(values []interface{}, filePath FilePath, line ProgramLine, name MethodName) []string {
		return []string{"[", string(filePath), " / ", string(name), "]: ", strings.Join(InterfacesToStrings(values), " ")}
	}
	InfoOutput = func(values []interface{}, _ FilePath, line ProgramLine, name MethodName) []string {
		return []string{"[", string(name), "]: ", strings.Join(InterfacesToStrings(values), " ")}
	}
	WarnOutput = func(values []interface{}, _ FilePath, line ProgramLine, name MethodName) []string {
		return []string{"[", string(name), "]: ", strings.Join(InterfacesToStrings(values), " ")}
	}
	ErrorOutput = func(values []interface{}, _ FilePath, line ProgramLine, name MethodName) []string {
		return []string{"[", string(name), "]: ", strings.Join(InterfacesToStrings(values), " ")}
	}
	EmergencyOutput = func(values []interface{}, _ FilePath, line ProgramLine, name MethodName) []string {
		return []string{"[", string(name), "]: ", strings.Join(InterfacesToStrings(values), " ")}
	}
	FatalOutput = func(values []interface{}, _ FilePath, line ProgramLine, name MethodName) []string {
		return []string{"[", string(name), "]: ", strings.Join(InterfacesToStrings(values), " ")}
	}
	DisruptOutput = func(values []interface{}, _ FilePath, line ProgramLine, name MethodName) []string {
		return []string{"[", string(name), "]: ", strings.Join(InterfacesToStrings(values), " ")}
	}
)

func ParseLevel(level string) (LogLevel, error) {
	switch strings.ToLower(level) {
	case "trace", "t", "tra":
		return LOG_LEVEL_TRACE, nil
	case "debug", "deb", "d", "develop", "dev":
		return LOG_LEVEL_DEBUG, nil
	case "info", "inf", "i":
		return LOG_LEVEL_INFO, nil
	case "warn", "waring", "w":
		return LOG_LEVEL_WARN, nil
	case "error", "err", "e":
		return LOG_LEVEL_ERROR, nil
	case "emergency":
		return LOG_LEVEL_EMERGENCY, nil
	case "fatal", "f":
		return LOG_LEVEL_FATAL, nil
	case "disrupt":
		return LOG_LEVEL_DISRUPT, nil
	default:
		return LogLevel(-1000), errors.New("No match pattern")
	}
}
func SetFunctionSkip(_skip int) {
	skip = _skip
}
func SetLevel(level LogLevel) {
	logLevel = level
}
func SetLevelOrDefault(level interface{}, defaultLevel LogLevel) {
	if _level, err := ParseLevel(convert.MustStr(level)); err != nil {
		SetLevel(defaultLevel)
	} else {
		SetLevel(_level)
	}
}
func SetContextKey(_keys ...interface{}) {
	for _, key := range _keys {
		keys = append(keys, key)
	}
}
func CleanContextKey() {
	keys = []interface{}{}
}
func DeleteContextKey(_key interface{}) {
	done := false
	for !done {
		done = true
		for index, key := range keys {
			if key == _key {
				done = false
				keys = append(keys[:index], keys[index+1:]...)
			}
			break
		}
	}

}

func GetLevel() LogLevel {
	return logLevel
}

func Trace(values ...interface{}) {
	if logLevel >= LOG_LEVEL_TRACE {
		filePath, line, name := getInfo()
		log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
		log.Println("[Trace] ", TraceOutput(values, filePath, line, name))
	}
}
func Debug(values ...interface{}) {
	if logLevel >= LOG_LEVEL_DEBUG {
		filePath, line, name := getInfo()
		log.SetFlags(log.Ldate | log.Ltime)
		log.Println("[DEBUG] ", DebugOutput(values, filePath, line, name))
	}
}
func Info(values ...interface{}) {
	if logLevel >= LOG_LEVEL_INFO {
		filePath, line, name := getInfo()
		log.SetFlags(log.Ldate | log.Ltime)
		log.Println("[INFO] ", InfoOutput(values, filePath, line, name))
	}
}
func Warn(values ...interface{}) {
	if logLevel >= LOG_LEVEL_WARN {
		filePath, line, name := getInfo()
		log.SetFlags(log.Ldate | log.Ltime)
		log.Println("[\u001B[33mWARN\u001B[0m] ", WarnOutput(values, filePath, line, name))
	}
}
func Error(values ...interface{}) {
	if logLevel >= LOG_LEVEL_ERROR {
		filePath, line, name := getInfo()
		log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
		log.Println("[\u001B[31mERROR\u001B[0m] ", ErrorOutput(values, filePath, line, name))
	}
}
func Emergency(values ...interface{}) {
	if logLevel >= LOG_LEVEL_EMERGENCY {
		filePath, line, name := getInfo()
		log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
		log.Println("\u001B[31m[EMERGENCY] ", EmergencyOutput(values, filePath, line, name), "\x1b[0m")
	}
}
func Disrupt(values ...interface{}) {
	if logLevel >= LOG_LEVEL_DISRUPT {
		filePath, line, name := getInfo()
		log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
		log.Println("\u001B[31m[DISRUPT] ", DisruptOutput(values, filePath, line, name), "\u001B[0m")
	}
}
func Fatal(values ...interface{}) {
	if logLevel >= LOG_LEVEL_FATAL {
		filePath, line, name := getInfo()
		log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
		log.Fatalln("\u001B[31m[FATAL] ", FatalOutput(values, filePath, line, name), "\x1b[0m")
	}
}

func TraceCtx(ctx context.Context, values ...interface{}) {
	if logLevel >= LOG_LEVEL_TRACE {
		filePath, line, name := getInfo()
		prefix := getContextValues(ctx, keys)
		log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
		log.Println("[Trace] ", fmt.Sprintf("(%s)", prefix), TraceOutput(values, filePath, line, name))
	}
}
func DebugCtx(ctx context.Context, values ...interface{}) {
	if logLevel >= LOG_LEVEL_DEBUG {
		filePath, line, name := getInfo()
		prefix := getContextValues(ctx, keys)
		log.SetFlags(log.Ldate | log.Ltime)
		log.Println("[DEBUG]", fmt.Sprintf("(%s)", prefix), DebugOutput(values, filePath, line, name))
	}
}
func InfoCtx(ctx context.Context, values ...interface{}) {
	if logLevel >= LOG_LEVEL_INFO {
		filePath, line, name := getInfo()
		prefix := getContextValues(ctx, keys)
		log.SetFlags(log.Ldate | log.Ltime)
		log.Println("[INFO]", fmt.Sprintf("(%s)", prefix), InfoOutput(values, filePath, line, name))
	}
}
func WarnCtx(ctx context.Context, values ...interface{}) {
	if logLevel >= LOG_LEVEL_WARN {
		filePath, line, name := getInfo()
		prefix := getContextValues(ctx, keys)
		log.SetFlags(log.Ldate | log.Ltime)
		log.Println("[\u001B[33mWARN\u001B[0m] ", fmt.Sprintf("(%s)", prefix), WarnOutput(values, filePath, line, name))
	}
}
func ErrorCtx(ctx context.Context, values ...interface{}) {
	if logLevel >= LOG_LEVEL_ERROR {
		filePath, line, name := getInfo()
		prefix := getContextValues(ctx, keys)
		log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
		log.Println("[\u001B[31mERROR\u001B[0m] ", fmt.Sprintf("(%s)", prefix), ErrorOutput(values, filePath, line, name))
	}
}
func EmergencyCtx(ctx context.Context, values ...interface{}) {
	if logLevel >= LOG_LEVEL_EMERGENCY {
		filePath, line, name := getInfo()
		prefix := getContextValues(ctx, keys)
		log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
		log.Println("\u001B[31m[EMERGENCY] ", fmt.Sprintf("(%s)", prefix), EmergencyOutput(values, filePath, line, name), "\x1b[0m")
	}
}
func FatalCtx(ctx context.Context, values ...interface{}) {
	if logLevel >= LOG_LEVEL_FATAL {
		filePath, line, name := getInfo()
		prefix := getContextValues(ctx, keys)
		log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
		log.Println("\u001B[31m[FATAL] ", fmt.Sprintf("(%s)", prefix), FatalOutput(values, filePath, line, name), "\x1b[0m")
	}
}
func DisruptCtx(ctx context.Context, values ...interface{}) {
	if logLevel >= LOG_LEVEL_DISRUPT {
		filePath, line, name := getInfo()
		prefix := getContextValues(ctx, keys)
		log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
		log.Println("\u001B[31m[DISRUPT] ", fmt.Sprintf("(%s)", prefix), DisruptOutput(values, filePath, line, name), "\u001B[0m")
	}
}
