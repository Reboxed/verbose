package main

import "github.com/Reboxed/verbose"

func main() {
	verbose.Success("TEST", "This is a test success log")
	verbose.Info("TEST", "This is a test info log")
	verbose.Print("TEST", "This is a test print log")
	verbose.Debug("TEST", "This is a test debug log")
	verbose.Warn("TEST", "This is a test warn log")
	verbose.Error("TEST", "This is a test error log")
}
