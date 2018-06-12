package main // import "vimagination.zapto.org/game"

import (
	"fmt"
	"os"
	"sync"

	"vimagination.zapto.org/engine"
	_ "vimagination.zapto.org/engine/graphics/glfw31_gl21"
)

var (
	mainMenu = Menu{
		items: []MenuItem{
			{"Item 1", func() {}},
			{"Item 2", func() {}},
		},
		top:    0.1,
		left:   0.1,
		gap:    0.1,
		width:  0.8,
		height: 0.05,
	}
	allInit sync.Once
)

func run(width, height int, time float64) bool {
	allInit.Do(func() {
		menuFontInit()
	})
	t := triangle{
		Angle: time,
	}
	t.Angle = time
	StartRender(width, height)
	t.Render(time)
	mainMenu.Render(time)
	FinishRender(time)
	if engine.KeyPressed(engine.KeyEscape) {
		return false
	}
	return true
}

func main() {
	var c UserConfig
	c.Width = 640
	c.Height = 480
	c.Title = "Test"

	f, err := os.Open("config.toml")
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "Error opening configuration file, using defaults. Err: %s", err)
		}
	} else {
		if err := c.Load(f); err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing configuration file, using defaults. Err: %s", err)
		}
		f.Close()
	}

	if err := engine.Loop(c.Config, run); err != nil {
		fmt.Fprintf(os.Stderr, "Error occurred: %s", err)
		return
	}
}
