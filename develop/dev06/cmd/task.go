package main

import (
	"log"
	"lvl2/develop/dev06/internal"
	"os"
)

func main() {
	c := cut.NewCut(os.Stdin, os.Stdout)
	c.InitConfig()
	if err := c.Run(); err != nil {
		log.Fatalln(err)
	}
}
