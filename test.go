package main

import "net/http"

type kek func(int) int

func (k kek) test(g int) {
	k(5)
}

func main() {
	var a kek
	a = func(b int) int {
		return b + 1
	}
	http.HandlerFunc()
	a.test(5)
	http.ListenAndServe(":80", nil)
}
