package main

import (
	"fmt"

	"github.com/adamjace/cfganalyze"
)

func main() {

	analyzer, err := cfganalyze.NewAnalyzer("fixtures/a.json", "fixtures/a.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	missing, _ := analyzer.AnalyzeJson()

	for _, key := range missing {
		fmt.Printf("missing key in config %s", key)
	}
}
