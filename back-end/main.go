package main

import (
	"TechVeritasMultiplex/model"
	"TechVeritasMultiplex/server"
)

func main() {

	model.DBInit()
	s := server.NewServer()
	s.Init("8000")
	s.Start()

}
