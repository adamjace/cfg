package cfganalyze

import (
	"testing"
)

func TestAnalyze(t *testing.T) {
	cfganalyze := NewCfgAnalyze("fixtures/a.json", "fixtures/b.json", "json")

	equal, err := cfganalyze.Analyze()
	if err != nil {
		t.Fatal(err)
	}

	if equal {
		t.Fatal("expected to be equal")
	}

}
