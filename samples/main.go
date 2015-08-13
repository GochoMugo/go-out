package main

import (
	"os"

	"github.com/GochoMugo/go-out"
)

func main() {
	var funcName = os.Args[1]

	if funcName == "All" {
		os.Setenv("DEBUG", "truthy")
		out.Success("success message")
		out.Info("informatory message")
		out.Warn("warning message")
		out.Error("error message")
		out.Debug("debug message")
		return
	}

	// invoked by tests
	var message = os.Args[2]

	switch funcName {
	case "Success":
		out.Success(message)
	case "Error":
		out.Error(message)
	case "Warn":
		out.Warn(message)
	case "Info":
		out.Info(message)
	case "Debug":
		out.Debug(message)
	}
}
