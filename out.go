/*
Package out provides a simple interface to terminal outputting. Suitable for terminal
applications.

View sample screenshot at https://github.com/GochoMugo/go-out/blob/master/screenshot.png

All the exported functions, unless noted otherwise, format the message string
using Printf. Also, color codes, as expected by colorstring (https://github.com/mitchellh/colorstring),
found within the message are replaced with the respective colors. For example `"hello, [red] mars"`
will have the "[red]" replaced with color red.
*/
package out

import (
	"io"
	"os"

	"github.com/mitchellh/colorstring"
)

const marker = "[magenta] >>> "

// wrap wraps a message with the designated color codes
func printout(out io.Writer, message string, colorcode string, args ...interface{}) {
	colorstring.Fprintf(out, marker+"["+colorcode+"]"+message+"\n", args...)
}

// Success outputs a success message
func Success(message string, args ...interface{}) {
	printout(os.Stdout, message, "green", args...)
}

// Error outputs an error message
func Error(message string, args ...interface{}) {
	printout(os.Stderr, message, "red", args...)
}

// Warn outputs a warning message
func Warn(message string, args ...interface{}) {
	printout(os.Stderr, message, "yellow", args...)
}

// Info outputs a informatory message
func Info(message string, args ...interface{}) {
	printout(os.Stdout, message, "blue", args...)
}

// Debug outputs a debug message if ${DEBUG} is truthy
func Debug(message string, args ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		printout(os.Stderr, message, "cyan", args...)
	}
}
