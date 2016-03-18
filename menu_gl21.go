package main

import "github.com/go-gl/gl/v2.1/gl"

func (m *Menu) Render(time float64) {
	for num := range m.items {
		top := m.top + float64(num)*m.gap
		left := m.left
		gl.Begin(gl.TRIANGLE_STRIP)
		gl.Color3f(1, 0, 0)
		gl.Vertex3d(left+m.width, top, 0)
		gl.Vertex3d(left, top, 0)
		gl.Vertex3d(left+m.width, top+m.height, 0)
		gl.Vertex3d(left, top+m.height, 0)
		gl.End()
	}
}
