package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	for i := 0; i < 99; i++ {
		fmt.Print("|")

	}

	clearLinuxShell()

	for i := 0; i < 99; i++ {
		fmt.Print("|")
	}

}

// Clear Windows Shell
func clearLinuxShell() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
