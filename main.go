package main

import (
	"github.com/LucasJosivan/desafio-simplificado/domains"
	"github.com/nuveo/log"
)

func main() {
	log.Debugln("Starting application...")
	domains.StartGame()
	log.Debugln("Program finished")
}
