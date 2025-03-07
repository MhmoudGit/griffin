package log

import (
	"fmt"
	"log"
)

// Define color constants
const (
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorReset  = "\033[0m"
)

// Initialize the log package with no flags to have cleaner output
func init() {
	log.SetFlags(0) // Remove default timestamp
}

// Error logs an error message with a red prefix
func Error(data ...interface{}) {
	log.Println(fmt.Sprintf("%s[ERROR]%s %v", ColorRed, ColorReset, fmt.Sprint(data...)))
}

// Success logs a success message with a green prefix
func Success(data ...interface{}) {
	log.Println(fmt.Sprintf("%s[SUCCESS]%s %v", ColorGreen, ColorReset, fmt.Sprint(data...)))
}

// Warning logs a warning message with a yellow prefix
func Warning(data ...interface{}) {
	log.Println(fmt.Sprintf("%s[WARNING]%s %v", ColorYellow, ColorReset, fmt.Sprint(data...)))
}

// Info logs an informational message with a blue prefix
func Info(data ...interface{}) {
	log.Println(fmt.Sprintf("%s[INFO]%s %v", ColorBlue, ColorReset, fmt.Sprint(data...)))
}

// Panic logs a panic message with a red prefix and calls log.Panic
func Panic(data ...interface{}) {
	log.Panic(fmt.Sprintf("%s[PANIC]%s %v", ColorRed, ColorReset, fmt.Sprint(data...)))
}

// Print logs a normal message without a colored prefix
func Print(data ...interface{}) {
	log.Println(fmt.Sprint(data...))
}
