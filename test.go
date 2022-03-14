package main

import "fmt"

func main() {
	a := make(map[int]int)
	a[3] = 5
	_, ok := a[4]
	fmt.Println(ok)
	fmt.Println(a[3], a[4])
	_, ok = a[4]
	fmt.Println(ok)

}
