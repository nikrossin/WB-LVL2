package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	_, err := io.Copy(os.Stdin, os.Stdout)
	fmt.Println("ok1")
	fmt.Println(err)

}
