package cfganalyze

import (
	"encoding/json"
)

type configJSON map[string]interface{}

type cfgAnalyzeJSON struct {
	a           configJSON
	b           configJSON
	missingKeys []string
}

func newCfgAnalyzeJSON(a, b []byte) (*cfgAnalyzeJSON, error) {
	jsonA := configJSON{}
	if err := json.Unmarshal(a, &jsonA); err != nil {
		return nil, err
	}

	jsonB := configJSON{}
	if err := json.Unmarshal(b, &jsonB); err != nil {
		return nil, err
	}

	cfgAnalyzeJSON := &cfgAnalyzeJSON{
		a: jsonA,
		b: jsonB,
	}

	return cfgAnalyzeJSON, nil
}

func (c *cfgAnalyzeJSON) analyze() {
	c.analyzeKeys(c.a, c.b)
}

func (c *cfgAnalyzeJSON) analyzeKeys(a configJSON, b configJSON) {

	keysA := []string{}
	keysB := []string{}

	c.diff(keysA, a, b)
	c.diff(keysB, b, a)

	for _, str := range keysA {
		if !c.contains(keysB, str) && !c.contains(c.missingKeys, str) {
			c.missingKeys = append(c.missingKeys, str)
		}
	}
}

func (c *cfgAnalyzeJSON) diff(keys []string, a configJSON, b configJSON) {

	for k, _ := range a {
		keys = append(keys, k)

		if c.isMap(a[k]) {
			if !c.isMap(b[k]) {
				c.missingKeys = append(c.missingKeys, k)
				continue
			}

			c.analyzeKeys(a[k].(map[string]interface{}),
				b[k].(map[string]interface{}))
		}
	}
}

func (c cfgAnalyzeJSON) equalKeys() (bool, error) {

	c.clearValues(c.a)
	c.clearValues(c.b)

	bytesA, err := json.Marshal(c.a)
	if err != nil {
		return false, err
	}

	bytesB, err := json.Marshal(c.b)
	if err != nil {
		return false, err
	}

	return string(bytesA) == string(bytesB), nil
}

func (c *cfgAnalyzeJSON) clearValues(m map[string]interface{}) {
	for k, _ := range m {
		if !c.isMap(m[k]) {
			m[k] = ""
		} else {
			c.clearValues(m[k].(map[string]interface{}))
		}
	}
}

func (c cfgAnalyzeJSON) isMap(m interface{}) bool {
	_, ok := m.(map[string]interface{})
	return ok
}

func (c cfgAnalyzeJSON) contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
