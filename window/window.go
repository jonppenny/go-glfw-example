package window

import "github.com/go-gl/glfw/v3.3/glfw"

const (
	width  int = 640
	height int = 480
)

func InitGlfw() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	window, err := glfw.CreateWindow(width, height, "Test", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	return window
}
