package main

import (
	"github.com/adamjace/cfg"
)

func main() {

	c := cfg.Config{
		WorkingPath: "test/a.json",
		MasterPath:  "test/b.json",
	}

	cfg.PrintJson(c)
}
