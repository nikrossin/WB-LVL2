package main

import wget "lvl2/develop/dev09/internal"

func main() {

	util := wget.NewWget()
	util.InitConfig()
	util.Run()

}
