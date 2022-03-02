package main

import (
	"log"
	"lvl2/develop/dev05/internal"
)

func main() {
	g := grep.NewGrep()
	if err := g.Init(); err != nil {
		log.Fatalln(err)
	}
	g.Run()

}
