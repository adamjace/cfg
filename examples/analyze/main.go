package main

import (
	"log"

	"github.com/adamjace/cfganalyze"
)

func main() {

	a := cfganalyze.NewAnalyzer()

	missingKeys, err := a.AnalyzeJson("sampleA.json", "sampleB.json")
	if err != nil {
		log.Println(err)
		return
	}

	for _, key := range missingKeys {
		log.Printf("Found missing key: %s", key)
	}
}
