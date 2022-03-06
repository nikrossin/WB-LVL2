package main

import (
	utelnet "lvl2/develop/dev10/internal"
	"os"
)

func main() {
	config := utelnet.NewConfig()
	config.Init()
	client := utelnet.NewTelnetClient(config, os.Stdin, os.Stdout)
	client.Run()
}
