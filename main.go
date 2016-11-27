package main

import (
	"fmt"
	"log"
    "bohica/graphics"
    "bohica/scene"
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

    menu := scene.MenuScene{"Header title"}

    fmt.Println(menu.Title)

	window.MakeContextCurrent()
	if err := gl.Init(); err != nil {
		log.Fatalln("failed to initialize gl bindings:", err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)
	
    for !window.ShouldClose() {
        graphics.Draw()
        
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func init() {

}
