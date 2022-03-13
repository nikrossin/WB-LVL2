package main

import (
	shell "lvl2/develop/dev08/internal"
	"os"
)

func main() {
	sh := shell.NewShell(os.Stdout, os.Stdin)
	sh.SetColors("\033[33m", "\033[34m", "\033[0m")
	sh.SetSystem("admin", "admin")
	sh.Run()
}
