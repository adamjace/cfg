package cfg

import "testing"

func TestNewBash(t *testing.T) {
	testHost := "test-host"
	bash := newBash(testHost)

	if bash.hostAlias != testHost {
		t.Fatalf("expected=%s actual=%s", testHost, bash.hostAlias)
	}
}
