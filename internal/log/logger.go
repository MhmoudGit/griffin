package log

import (
	"fmt"
	"log"
	"strings"
)

// Define color constants
const (
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorReset  = "\033[0m"
)

// Custom log writer struct
type logWriter struct {
	prefix string
	color  string
}

// Implement io.Writer for logWriter
func (w logWriter) Write(p []byte) (n int, err error) {
	message := strings.TrimSpace(string(p)) // Remove extra spaces
	log.Printf("%s%s%s %s", w.color, w.prefix, message, ColorReset)
	return len(p), nil
}

// Predefined writers for stdout and stderr
var (
	StdWriter = logWriter{prefix: "", color: ColorYellow}
)

// Initialize log settings
func init() {
	log.SetFlags(0) // Remove timestamp
}

// Error logs an error message with a red prefix
func Error(data ...interface{}) {
	log.Printf("%s[ERROR]%s %v", ColorRed, ColorReset, fmt.Sprint(data...))
}

// Success logs a success message with a green prefix
func Success(data ...interface{}) {
	log.Printf("%s[SUCCESS]%s %v", ColorGreen, ColorReset, fmt.Sprint(data...))
}

// Warning logs a warning message with a yellow prefix
func Warning(data ...interface{}) {
	log.Printf("%s[WARNING]%s %v", ColorYellow, ColorReset, fmt.Sprint(data...))
}

// Info logs an informational message with a blue prefix
func Info(data ...interface{}) {
	log.Printf("%s[INFO]%s %v", ColorBlue, ColorReset, fmt.Sprint(data...))
}

// Panic logs a panic message with a red prefix and calls log.Panic
func Panic(data ...interface{}) {
	log.Panicf("%s[PANIC]%s %v", ColorRed, ColorReset, fmt.Sprint(data...))
}

// Print logs a normal message without a colored prefix
func Print(data ...interface{}) {
	log.Print(data...)
}
