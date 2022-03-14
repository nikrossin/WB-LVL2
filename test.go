package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, _ := os.Open("a.txt")

	data := make([]byte, 2)

	for {
		n, err := file.Read(data)
		fmt.Println(1)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		fmt.Print(string(data[:n]))
	}

}
