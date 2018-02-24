package cfganalyze

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type CfgAnalyze struct {
	a          string
	b          string
	configType string
}

type jsonConfig map[string]interface{}

func NewCfgAnalyze(a string, b string, configType string) CfgAnalyze {
	return CfgAnalyze{
		a:          a,
		b:          b,
		configType: configType,
	}
}

func (c CfgAnalyze) Analyze() (bool, error) {
	a, err := ioutil.ReadFile(c.a)
	if err != nil {
		return false, fmt.Errorf("could not open %s", c.a)
	}

	b, err := ioutil.ReadFile(c.b)
	if err != nil {
		return false, fmt.Errorf("could not open %s", c.b)
	}

	if c.configType == "json" {
		return c.analyzeJson(a, b)
	}

	return false, nil
}

func (c CfgAnalyze) analyzeJson(a, b []byte) (bool, error) {
	cA := jsonConfig{}
	if err := json.Unmarshal(a, &cA); err != nil {
		return false, err
	}

	cB := jsonConfig{}
	if err := json.Unmarshal(b, &cB); err != nil {
		return false, err
	}

	c.clearValues(cA)
	c.clearValues(cB)

	bytesA, err := json.Marshal(cA)
	if err != nil {
		return false, err
	}

	bytesB, err := json.Marshal(cB)
	if err != nil {
		return false, err
	}

	strA := string(bytesA)
	strB := string(bytesB)

	fmt.Printf("%s\n", strA)
	fmt.Printf("%s\n", strB)

	return strA == strB, nil
}

func (c *CfgAnalyze) clearValues(m map[string]interface{}) {
	for k, _ := range m {
		if !c.isMap(m[k]) {
			m[k] = ""
		} else {
			c.clearValues(m[k].(map[string]interface{}))
		}
	}
}

func (c CfgAnalyze) isMap(m interface{}) bool {
	_, ok := m.(map[string]interface{})
	return ok
}
