package cfganalyze

import (
	"testing"
)

func TestJsonAnalyze(t *testing.T) {
	cfganalyze := NewAnalyzer("fixtures/a.json", "fixtures/b.json", ConfigTypeJSON)

	missingKeys, err := cfganalyze.Analyze()
	if err != nil {
		t.Fatal(err)
	}

	if len(missingKeys) > 0 {
		t.Fatal("expected missing keys to be 0")
	}
}

func TestJsonEqualKeys(t *testing.T) {
	cfganalyze := NewAnalyzer("fixtures/a.json", "fixtures/b.json", ConfigTypeJSON)

	equal, err := cfganalyze.EqualKeys()
	if err != nil {
		t.Fatal(err)
	}

	if equal {
		t.Fatal("expected to be equal")
	}
}
