package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// Global variable
var CLEAR_SCREEN bool = true

func ClearScreen() {
	if !CLEAR_SCREEN {
		return
	}
	// Check the operating system to determine the appropriate clear command
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("clear") // Linux and macOS
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls") // Windows
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("Clearing the screen is not supported on this platform.")
	}
}
