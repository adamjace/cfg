package main

import (
	"fmt"

	"github.com/adamjace/cfganalyze"
)

func main() {

	analyzer := cfganalyze.NewAnalyzer("fixtures/a.json", "fixtures/a.json", cfganalyze.ConfigTypeJSON)
	missing, _ := analyzer.Analyze()

	for _, key := range missing {
		fmt.Printf("missing key in config %s", key)
	}
}
