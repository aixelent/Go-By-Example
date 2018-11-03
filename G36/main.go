package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	p := os.Getpid()
	fmt.Printf("Process ID: %d", p)

	proc := exec.Command("ps", "-p", strconv.Itoa(p), "-v")
	outpur, err := proc.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(outpur))
}
