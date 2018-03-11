package cfg

import (
	"testing"
)

func TestJsonMaster(t *testing.T) {
	c := Config{
		WorkingPath: "test/a.json",
		MasterPath:  "test/b.json",
	}

	analyzer, err := newJsonAnalyzer(c)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		key   string
		value interface{}
	}{
		{"1", true},
		{"2", false},
		{"6", true},
	}

	for _, tt := range tests {
		if analyzer.jsonMaster[tt.key] != tt.value {
			t.Fatalf("expected=%s actual=%s", tt.value, analyzer.jsonMaster[tt.key])
		}
	}
}

func TestJsonEqual(t *testing.T) {
	c := Config{
		WorkingPath: "test/c.json",
		MasterPath:  "test/d.json",
	}

	analyzer, err := newJsonAnalyzer(c)
	if err != nil {
		t.Fatal(err)
	}

	analyzer.scan()

	if len(analyzer.missing) > 0 {
		t.Fatal("keys should be equal")
	}

	equal, err := analyzer.equality()
	if err != nil {
		t.Fatal(err)
	}

	if equal {
		t.Fatal("values should not be equal")
	}
}
