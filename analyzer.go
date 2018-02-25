package cfganalyze

import (
	"fmt"
	"io/ioutil"
)

// Analyzer contains data for analyzing config files
type Analyzer struct {
	configA []byte
	configB []byte
	host    string
}

// NewAnalyzer returns a new Analyzer
func NewAnalyzer(pathA string, pathB string) (*Analyzer, error) {
	a, err := ioutil.ReadFile(pathA)
	if err != nil {
		return nil, fmt.Errorf("could not open %s", pathA)
	}

	b, err := ioutil.ReadFile(pathB)
	if err != nil {
		return nil, fmt.Errorf("could not open %s", pathB)
	}

	analyzer := Analyzer{
		configA: a,
		configB: b,
	}

	return &analyzer, nil
}

// Connect will return a new connected Analyzer to an external host via SSH
func Connect(pathA string, pathB string, sshHost string) (*Analyzer, error) {
	analyzer, err := NewAnalyzer(pathA, pathB)
	if err != nil {
		return nil, err
	}

	analyzer.host = sshHost

	return analyzer, nil
}

// AnalyzeJson will compare two .json configuration files
// highlighting keys that are missing
func (c Analyzer) AnalyzeJson() ([]string, error) {
	analyzer, err := newJsonAnalyzer(c.configA, c.configB)
	if err != nil {
		return nil, err
	}

	analyzer.analyze()

	return analyzer.missingKeys, nil
}

// AnalyzeEnv will compare two .env configuration files
// highlighting keys that are missing
func (c Analyzer) AnalyzeEnv() ([]string, error) {
	analyzer, err := newEnvAnalyzer(c.configA, c.configB)
	if err != nil {
		return nil, err
	}

	analyzer.analyze()

	return analyzer.missingKeys, nil
}

// EqualKeys will compare two configurations identifying whether they are the same
func (c Analyzer) EqualKeys() (bool, error) {
	analyzer, err := newJsonAnalyzer(c.configA, c.configB)
	if err != nil {
		return false, err
	}

	return analyzer.equalKeys()
}
