package main

import (
	"lvl2/develop/dev03/internal/sortfile"
)

func main() {
	file := sortfile.NewTextFile()
	file.ParseFlags()
	file.SetInputPath()
	file.SetOutputPath()
	file.Read()
	sortfile.SortFile(file)
	file.Write()
}
