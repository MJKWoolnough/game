package main

import (
	"fmt"
	"os"

	"github.com/MJKWoolnough/engine"
	_ "github.com/MJKWoolnough/engine/graphics/glfw31_gl21"
)

func run(width, height int, time float64) bool {
	t := triangle{
		Angle: time,
	}
	t.Angle = time
	t.Render(width, height, time)
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
