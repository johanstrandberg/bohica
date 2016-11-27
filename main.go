package main

import (
    "fmt"
    "github.com/go-gl/glfw/v3.2/glfw"
)

func main() {
    //fmt.Printf("kebab")
    glfw.Init()
    defer glfw.Terminate()
    window, err := glfw.CreateWindow(640,480, "Bohica", nil, nil)
    if err != nil {
        panic(err)
    }
    window.MakeContextCurrent()
    fmt.Println(glfw.VulkanSupported())
        fmt.Println(glfw.GetVersion())

    for !window.ShouldClose() {
        window.SwapBuffers()
        glfw.PollEvents()
    }
}