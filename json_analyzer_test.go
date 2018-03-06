package cfg

import (
	"testing"
)

func TestJsonAnalyze(t *testing.T) {

	c := Config{
		WorkingPath: "test/a.json",
		MasterPath:  "test/b.json",
	}

	keys, err := ScanJson(c)
	if err != nil {
		t.Fatal(err)
	}

	if len(keys) == 0 {
		t.Fatal("expected missing keys to be 0")
	}
}
