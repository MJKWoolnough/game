package main

import "github.com/go-gl/gl/v2.1/gl"

func (t triangle) Render(time float64) {

	gl.LoadIdentity()
	gl.PushMatrix()
	gl.Translatef(0.5, 0.5, 0)
	gl.Rotated(time*50, 0, 0, 1)

	gl.Begin(gl.TRIANGLES)
	gl.Color3f(1, 0, 0)
	gl.Vertex3f(-0.4, -0.231, 0.)
	gl.Color3f(0, 1, 0)
	gl.Vertex3f(0.4, -0.231, 0)
	gl.Color3f(0, 0, 1)
	gl.Vertex3f(0, 0.46182, 0)
	gl.End()
	gl.PopMatrix()
}
