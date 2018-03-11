package main

import (
	"github.com/adamjace/cfg"
)

func main() {

	c := cfg.Config{
		WorkingPath: "test/c.env",
		MasterPath:  "test/d.env",
	}

	cfg.PrintEnv(c)
}
