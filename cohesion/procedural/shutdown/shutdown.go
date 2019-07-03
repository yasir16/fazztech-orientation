package shutdown

func shutdownFoo() {}
func shutdownBar() {}
func shutdownBaz() {}

func Shutdown() {
	shutdownFoo()
	shutdownBar()
	shutdownBaz()
}
