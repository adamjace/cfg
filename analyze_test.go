package cfganalyze

import (
	"testing"
)

func TestAnalyze(t *testing.T) {
	cfganalyze := NewCfgAnalyze("fixtures/a.json", "fixtures/b.json", "json")

	missingKeys, err := cfganalyze.Analyze()
	if err != nil {
		t.Fatal(err)
	}

	if len(missingKeys) > 0 {
		t.Fatal("expected to be equal")
	}

}
