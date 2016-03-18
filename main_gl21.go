package main

import "github.com/go-gl/gl/v2.1/gl"

func StartRender(width, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
	gl.Clear(gl.COLOR_BUFFER_BIT)

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0, 1, 1, 0, 0, 1)
	gl.MatrixMode(gl.MODELVIEW)
}

func FinishRender(time float64) {

}
