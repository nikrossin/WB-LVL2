package main

import (
	"fmt"
	"os"
)

func main() {
	if a, err := os.Stat("pattern"); err != nil {
		fmt.Println("kek")
	} else {
		fmt.Println(a)
	}
}
