package cfganalyze

import (
	"fmt"
	"io/ioutil"
)

const (
	ConfigTypeJSON = "json"
	ConfigTypeENV  = "env"
)

// Analyzer contains data for analyzing config files
type Analyzer struct {
	configPathA string
	configPathB string
	configType  string
}

// NewAnalyzer returns a new Analyzer
func NewAnalyzer(a string, b string, configType string) Analyzer {
	return Analyzer{
		configPathA: a,
		configPathB: b,
		configType:  configType,
	}
}

// Analyze will compare two configurations highlighting keys that are missing
func (c Analyzer) Analyze() ([]string, error) {
	a, err := ioutil.ReadFile(c.configPathA)
	if err != nil {
		return nil, fmt.Errorf("could not open %s", c.configPathA)
	}

	b, err := ioutil.ReadFile(c.configPathB)
	if err != nil {
		return nil, fmt.Errorf("could not open %s", c.configPathB)
	}

	if c.configType == ConfigTypeJSON {
		analyzer, err := newJsonAnalyzer(a, b)
		if err != nil {
			return nil, err
		}

		analyzer.analyze()

		return analyzer.missingKeys, nil
	}

	analyzer, err := newEnvAnalyzer(a, b)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%+v\n", analyzer.envConfigB)

	analyzer.analyze()

	return analyzer.missingKeys, nil
}

// EqualKeys will compare two configurations identifying whether they are the same
func (c Analyzer) EqualKeys() (bool, error) {
	a, err := ioutil.ReadFile(c.configPathA)
	if err != nil {
		return false, fmt.Errorf("could not open %s", c.configPathB)
	}

	b, err := ioutil.ReadFile(c.configPathB)
	if err != nil {
		return false, fmt.Errorf("could not open %s", c.configPathB)
	}

	if c.configType == ConfigTypeJSON {
		cfgAnalyzeJson, err := newJsonAnalyzer(a, b)
		if err != nil {
			return false, err
		}

		return cfgAnalyzeJson.equalKeys()
	}

	return false, nil
}
