package main

import (
	"io"

	"github.com/BurntSushi/toml"
	"github.com/MJKWoolnough/engine"
)

type UserConfig struct {
	engine.Config
}

func (c *UserConfig) Load(r io.Reader) error {
	_, err := toml.DecodeReader(r, c)
	return err
}

func (c *UserConfig) Save(w io.Writer) error {
	return toml.NewEncoder(w).Encode(c)
}
