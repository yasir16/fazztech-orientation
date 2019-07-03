package main

import (
	"github.com/mrp130/fazztech-orientation/cohesion/procedural/init"
	"github.com/mrp130/fazztech-orientation/cohesion/procedural/shutdown"
)

func main() {
	init.Init()

	//run service..

	shutdown.Shutdown()
}
