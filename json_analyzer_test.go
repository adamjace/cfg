package cfganalyze

import (
	"testing"
)

func TestJsonAnalyze(t *testing.T) {
	analyzer, err := NewAnalyzer("fixtures/a.json", "fixtures/b.json")
	if err != nil {
		t.Fatal(err)
	}

	missingKeys, err := analyzer.AnalyzeJson()
	if err != nil {
		t.Fatal(err)
	}

	if len(missingKeys) > 0 {
		t.Fatal("expected missing keys to be 0")
	}
}

func TestJsonEqualKeys(t *testing.T) {
	analyzer, err := NewAnalyzer("fixtures/a.json", "fixtures/b.json")
	if err != nil {
		t.Fatal(err)
	}

	equal, err := analyzer.EqualKeys()
	if err != nil {
		t.Fatal(err)
	}

	if equal {
		t.Fatal("expected to be equal")
	}
}
