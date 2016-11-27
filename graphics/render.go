package graphics

import (
	"github.com/go-gl/gl/v3.2-compatibility/gl"
)

func Draw() {
		gl.ClearColor(0.0, 1.0, 0.0, 0.0)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}