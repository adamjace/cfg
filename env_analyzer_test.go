package cfganalyze

import (
	"fmt"
	"testing"
)

func TestEnvAnalyze(t *testing.T) {
	cfganalyze := NewAnalyzer("fixtures/a.env", "fixtures/b.env", ConfigTypeENV)

	missingKeys, err := cfganalyze.Analyze()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(missingKeys)

	if len(missingKeys) == 0 {
		t.Fatal("expected to have missing keys")
	}
}
