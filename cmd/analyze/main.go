package main

import (
	"fmt"

	"github.com/adamjace/cfganalyze"
)

func main() {

	analyzer := cfganalyze.NewAnalyzer()

	missing, err := analyzer.AnalyzeJson("fixtures/a.json", "fixtures/a.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, key := range missing {
		fmt.Printf("missing key in config %s", key)
	}
}
