package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	// df -h - shows disc usage
	proc := exec.Command("df", "-h")
	out := bytes.NewBuffer([]byte{})
	proc.Stdout = out

	err := proc.Run()
	if err != nil {
		panic(err)
	}

	if proc.ProcessState.Success() {
		fmt.Println("Process run successfully:", out.String())
	}
}
