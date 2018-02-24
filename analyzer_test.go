package cfganalyze

import (
	"testing"
)

func TestAnalyze(t *testing.T) {
	cfganalyze := NewAnalyzer("fixtures/a.json", "fixtures/b.json", "json")

	missingKeys, err := cfganalyze.Analyze()
	if err != nil {
		t.Fatal(err)
	}

	if len(missingKeys) > 0 {
		t.Fatal("expected missing keys to be 0")
	}
}

func TestEqualKeys(t *testing.T) {
	cfganalyze := NewAnalyzer("fixtures/a.json", "fixtures/b.json", "json")

	equal, err := cfganalyze.EqualKeys()
	if err != nil {
		t.Fatal(err)
	}

	if equal {
		t.Fatal("expected to be equal")
	}
}
