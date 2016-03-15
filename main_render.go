package main

import "github.com/go-gl/gl/v2.1/gl"

func StartRender(width, height int) {
	ratio := float64(width) / float64(height)
	gl.Viewport(0, 0, int32(width), int32(height))
	gl.Clear(gl.COLOR_BUFFER_BIT)

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(-ratio, ratio, -1, 1, 1, -1)
	gl.MatrixMode(gl.MODELVIEW)
}

func FinishRender(time float64) {

}
