package cfg

import (
	"testing"
)

func TestEnvMaster(t *testing.T) {
	c := Config{
		WorkingPath: "test/a.env",
		MasterPath:  "test/b.env",
	}

	analyzer, err := newEnvAnalyzer(c)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		key   string
		value string
	}{
		{"FRUIT", "Mango"},
		{"FOOD", "Pizza"},
		{"LANG", "Go"},
		{"ANIMAL", "Koala"},
		{"SPORT", "Football"},
		{"DRINK", "Soda"},
	}

	for i, tt := range tests {
		if analyzer.envMaster[i].Key != tt.key {
			t.Fatalf("expected=%s actual=%s", tt.key, analyzer.envMaster[i].Key)
		}
		if analyzer.envMaster[i].Value != tt.value {
			t.Fatalf("expected=%s actual=%s", tt.value, analyzer.envMaster[i].Value)
		}
	}
}

func TestEnvWorking(t *testing.T) {
	c := Config{
		WorkingPath: "test/a.env",
		MasterPath:  "test/b.env",
	}

	analyzer, err := newEnvAnalyzer(c)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		key   string
		value string
	}{
		{"FRUIT", "Mango"},
		{"ANIMAL", "Koala"},
		{"SPORT", "Football"},
	}

	for i, tt := range tests {
		if analyzer.envWorking[i].Key != tt.key {
			t.Fatalf("expected=%s actual=%s", tt.key, analyzer.envWorking[i].Key)
		}
		if analyzer.envWorking[i].Value != tt.value {
			t.Fatalf("expected=%s actual=%s", tt.value, analyzer.envWorking[i].Value)
		}
	}
}
