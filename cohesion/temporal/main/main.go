package main

import (
	"github.com/mrp130/fazztech-orientation/cohesion/temporal/init"
	"github.com/mrp130/fazztech-orientation/cohesion/temporal/shutdown"
)

func main() {
	init.InitFoo()
	init.InitBar()
	init.InitBaz()

	//run service..

	shutdown.ShutdownFoo()
	shutdown.ShutdownBar()
	shutdown.ShutdownBaz()
}
