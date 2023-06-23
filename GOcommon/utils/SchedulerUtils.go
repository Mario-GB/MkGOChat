package gocomm_utils

import (
	"fmt"
	"log"
	"os/exec"
)

// TODO
func GetTerminalIP() string {
	cmd := exec.Command("ls", "-l")
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	// Imprimir la salida del comando
	fmt.Println(string(output))

	return ""
}
