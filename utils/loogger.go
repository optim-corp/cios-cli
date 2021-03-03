package utils

import (
	"fmt"
	"log"
	"strings"
)

func NewLogging(level *int) (self Logging) {
	instance := Logging{}
	if level == nil {
		instance.LogLevel = 300
	} else {
		instance.LogLevel = *level
	}
	return instance
}

// Log.Info
func (self Logging) Info(values ...string) {
	if self.LogLevel >= LEVEL["INFO"] {
		log.SetPrefix("[INFO]: ")
		log.Println(strings.Join(values, " "))
	}

}

// Log.Warn
func (self Logging) Warn(values ...string) {
	if self.LogLevel >= LEVEL["WARN"] {
		log.SetPrefix("[\x1b[33mWARN\x1b[0m]: ")
		log.Println(strings.Join(values, " "))

	}
}

// Log.Error
func (self Logging) Error(values ...string) {
	if self.LogLevel >= LEVEL["ERROR"] {
		// log.SetFlags(log.Ldate | log.Ltime)
		log.SetPrefix("[\x1b[31mERROR\x1b[0m]: ")
		log.Println(strings.Join(values, " "))
		log.SetFlags(log.Ldate | log.Ltime)
	}
}

// Log.Debug
func (self Logging) Debug(values ...interface{}) {
	if self.LogLevel >= LEVEL["DEBUG"] {
		log.SetPrefix("[DEBUG]: ")
		log.Println(values...)
	}
}

// Log.Emergency
func (self Logging) Emergency(values ...string) {
	if self.LogLevel >= LEVEL["EMERGENCY"] {
		log.Print("\x1b[31m")
		log.SetPrefix("[EMERGENCY]: ")
		log.Print(strings.Join(values, " "))
		log.Println("\x1b[0m")
	}
}

// Log.Disrupt
func (self Logging) Disrupt(values ...string) {
	if self.LogLevel >= LEVEL["DISRUPT"] {
		log.Print("\x1b[31m")
		log.SetPrefix("[DISRUPT]: ")
		log.Fatal(strings.Join(values, " "))
		log.Println("\x1b[0m")
	}
}

func (self Logging) Println(values ...string) {
	log.SetPrefix("")
	log.Println(strings.Join(values, " "))
}

func (self Logging) Print(values ...string) {
	log.SetPrefix("")
	log.Print(strings.Join(values, " "))
}

func Fprintln(v ...interface{}) {
	fmt.Fprintln(Out, v...)
}
func Fprint(v ...interface{}) {
	fmt.Fprint(Out, v...)
}
func Fprintf(format string, v ...interface{}) {
	fmt.Fprintf(Out, format, v...)
}
func Println(v ...interface{}) {
	fmt.Println(v...)
}
func Print(v ...interface{}) {
	fmt.Print(v...)
}
func Printf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

var (
	Log   = NewLogging(nil)
	LEVEL = map[string]int{
		"DISRUPT":   DISRUPT_LOG_LEVEL,
		"EMERGENCY": EMERGENCY_LOG_LEVEL,
		"ERROR":     ERROR_LOG_LEVEL,
		"WARN":      WARN_LOG_LEVEL,
		"INFO":      INFO_LOG_LEVEL,
		"DEBUG":     DEBUG_LOG_LEVEL,
	}
	DISRUPT_LOG_LEVEL   = -999
	EMERGENCY_LOG_LEVEL = -1
	ERROR_LOG_LEVEL     = 100
	WARN_LOG_LEVEL      = 200
	INFO_LOG_LEVEL      = 300
	DEBUG_LOG_LEVEL     = 400
)
