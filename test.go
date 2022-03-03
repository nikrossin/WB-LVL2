package main

import (
	"fmt"
)

func main() {
	c := func() int {
		a := 5
		b := 6
		if a == 5 {
			return
		}
		return b
	}
	fmt.Println(c)
}
