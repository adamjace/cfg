package cfganalyze

import (
	"testing"
)

func TestEnvAnalyze(t *testing.T) {
	analyzer := NewAnalyzer()

	missingKeys, err := analyzer.AnalyzeEnv("test/a.env", "test/b.env")
	if err != nil {
		t.Fatal(err)
	}

	if len(missingKeys) == 0 {
		t.Fatal("expected to have missing keys")
	}
}
