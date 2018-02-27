package main

import (
	"fmt"

	"github.com/adamjace/cfganalyze"
)

func main() {

	a, err := cfganalyze.Connect("host-alias")
	if err != nil {
		fmt.Println(err)
		return
	}

	missingKeys, err := a.AnalyzeJson("test/config.json", "~/home/ubuntu/config.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, key := range missingKeys {
		fmt.Printf("warning: missing key in config '%s'\n", key)
	}
}
