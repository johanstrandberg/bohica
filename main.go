package main

import (
	"fmt"
	"log"

	"github.com/go-gl/gl/v3.2-compatibility/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func main() {
	//fmt.Printf("kebab")
	defer glfw.Terminate()
	glfw.Init()
	window, err := glfw.CreateWindow(640, 480, "Bohica", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()
	fmt.Println(glfw.VulkanSupported())
	fmt.Println(glfw.GetVersion())
	if err := gl.Init(); err != nil {
		log.Fatalln("failed to initialize gl bindings:", err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)
	for !window.ShouldClose() {
		gl.ClearColor(0.0, 1.0, 0.0, 0.0)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
