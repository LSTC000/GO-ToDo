package main

import "todo/internal/server"

func main() {
	s := server.GetServer()
	s.DotEnvLoad()
	s.Run()
}
