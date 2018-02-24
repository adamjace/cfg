package cfganalyze

import (
	"fmt"
	"io/ioutil"
)

const (
	ConfigTypeJSON = "json"
	ConfigTypeENV  = "env"
)

type CfgAnalyze struct {
	a          string
	b          string
	configType string
}

func NewCfgAnalyze(a string, b string, configType string) CfgAnalyze {
	return CfgAnalyze{
		a:          a,
		b:          b,
		configType: configType,
	}
}

func (c CfgAnalyze) Analyze() ([]string, error) {
	a, err := ioutil.ReadFile(c.a)
	if err != nil {
		return nil, fmt.Errorf("could not open %s", c.a)
	}

	b, err := ioutil.ReadFile(c.b)
	if err != nil {
		return nil, fmt.Errorf("could not open %s", c.b)
	}

	if c.configType == ConfigTypeJSON {
		cfgAnalyzeJson, err := newCfgAnalyzeJSON(a, b)
		if err != nil {
			return nil, err
		}

		cfgAnalyzeJson.analyze()

		return cfgAnalyzeJson.missingKeys, nil
	}

	return nil, nil
}
