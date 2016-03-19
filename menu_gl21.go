package main

import (
	"os"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/gltext"
)

var menuFont *gltext.Font

func menuFontInit() {
	f, err := os.Open("/usr/share/fonts/corefonts/times.ttf")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	menuFont, err = gltext.LoadTruetype(f, 20, 32, 127, gltext.LeftToRight)
	if err != nil {
		panic(err)
	}
}

func (m *Menu) Render(time float64) {
	for num, item := range m.items {
		//menuFont.Metrics(item.Name)
		top := m.top + float64(num)*m.gap
		left := m.left
		gl.Begin(gl.TRIANGLE_STRIP)
		gl.Color3f(1, 0, 0)
		gl.Vertex3d(left+m.width, top, 0)
		gl.Vertex3d(left, top, 0)
		gl.Vertex3d(left+m.width, top+m.height, 0)
		gl.Vertex3d(left, top+m.height, 0)
		gl.End()
		gl.PushMatrix()
		var vp [4]int32
		gl.GetIntegerv(gl.VIEWPORT, &vp[0])
		gl.Translated(left*float64(vp[2]), -top*float64(vp[3]), 0)
		gl.Color3f(0, 1, 0)
		menuFont.Printf(0, 0, item.Name)
		gl.PopMatrix()
	}
}
