package cfg

import (
	"testing"
)

func TestEnvAnalyze(t *testing.T) {

	c := Config{
		WorkingPath: "test/a.env",
		MasterPath:  "test/b.env",
	}

	keys, err := ScanEnv(c)
	if err != nil {
		t.Fatal(err)
	}

	if len(keys) == 0 {
		t.Fatal("expected to have missing keys")
	}
}
