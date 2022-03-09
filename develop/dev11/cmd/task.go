package main

import "lvl2/develop/dev11/internal/server"

func main() {
	cfg := server.NewConfig()
	cfg.DefaultSet()
	serv := server.NewServer(cfg)
	serv.Run()
}
