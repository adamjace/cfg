package cfg

import (
	"testing"
)

func TestNewAnalzer(t *testing.T) {
	c := Config{
		WorkingPath: "test/a.env",
		MasterPath:  "test/b.env",
	}

	analyzer, _ := newAnalyzer(c)

	if analyzer.bash != nil {
		t.Fatal("expected bash to be nil")
	}
}

func TestScanEnv(t *testing.T) {
	c := Config{
		WorkingPath: "test/a.env",
		MasterPath:  "test/b.env",
	}

	keys, err := ScanEnv(c)
	if err != nil {
		t.Fatal(err)
	}

	expected := 3
	actual := len(keys)

	if actual != expected {
		t.Fatalf("expected=%d actual", expected, actual)
	}
}

func TestPrintEnv(t *testing.T) {
	c := Config{
		WorkingPath: "test/a.env",
		MasterPath:  "test/b.env",
	}

	if err := PrintEnv(c); err != nil {
		t.Fatal(err)
	}
}

func TestScanJson(t *testing.T) {
	c := Config{
		WorkingPath: "test/a.json",
		MasterPath:  "test/b.json",
	}

	keys, err := ScanJson(c)
	if err != nil {
		t.Fatal(err)
	}

	expected := 2
	actual := len(keys)

	if actual != expected {
		t.Fatalf("expected=%d actual", expected, actual)
	}
}

func TestPrintJson(t *testing.T) {
	c := Config{
		WorkingPath: "test/a.json",
		MasterPath:  "test/b.json",
	}

	if err := PrintJson(c); err != nil {
		t.Fatal(err)
	}
}
