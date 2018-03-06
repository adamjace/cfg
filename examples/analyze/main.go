package main

import (
	"log"

	"github.com/adamjace/cfg"
)

func main() {

	c := cfg.Config{
		WorkingPath: "test/a.json",
		MasterPath:  "test/b.json",
	}

	if err := cfg.AnalyzeJson(c); err != nil {
		log.Print(err)
	}
}
