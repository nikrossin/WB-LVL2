package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	name := "echo"
	arg := "aaaaaaaaA"
	cmd := exec.Command(name, arg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	fmt.Println(err)

}
