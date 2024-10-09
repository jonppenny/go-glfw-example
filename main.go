package main

import (
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"

	"jonppenny.co.uk/go-glfw-example/window"
)

func init() {
	// Ensure that main() runs on main thread.
	// See documentation for functions that can be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	defer glfw.Terminate()

	w := window.InitGlfw()

	w.MakeContextCurrent()

	for !w.ShouldClose() {
		w.SwapBuffers()
		glfw.PollEvents()
	}
}
