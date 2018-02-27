package cfganalyze

import (
	"testing"
)

func TestJsonAnalyze(t *testing.T) {
	analyzer := NewAnalyzer()

	missingKeys, err := analyzer.MissingJsonKeys("test/a.json", "test/b.json")
	if err != nil {
		t.Fatal(err)
	}

	analyzer.AnalyzeJson("test/a.json", "test/b.json")

	if len(missingKeys) > 0 {
		t.Fatal("expected missing keys to be 0")
	}
}

func TestJsonEqualKeys(t *testing.T) {
	analyzer := NewAnalyzer()

	equal, err := analyzer.EqualKeys("test/a.json", "test/b.json")
	if err != nil {
		t.Fatal(err)
	}

	if equal {
		t.Fatal("expected to be equal")
	}
}
