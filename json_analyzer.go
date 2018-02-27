package cfganalyze

import (
	"encoding/json"
)

// configJSON is a basic map struct
type configJSON map[string]interface{}

// jsonAnalyzer holds data for both JSON maps
type jsonAnalyzer struct {
	jsonConfigA configJSON
	jsonConfigB configJSON
	missingKeys []string
}

// newJsonAnalyzer returns a new newCfgAnalyzeJSON loaded with JSON maps
func newJsonAnalyzer(a, b []byte) (*jsonAnalyzer, error) {
	jsonA := configJSON{}
	if err := json.Unmarshal(a, &jsonA); err != nil {
		return nil, err
	}

	jsonB := configJSON{}
	if err := json.Unmarshal(b, &jsonB); err != nil {
		return nil, err
	}

	jsonAnalyzer := &jsonAnalyzer{
		jsonConfigA: jsonA,
		jsonConfigB: jsonB,
	}

	return jsonAnalyzer, nil
}

// analyze will analyze two maps identifying keys that exist in map A
// that are missing in map B
func (j *jsonAnalyzer) analyze() {
	j.diff(j.jsonConfigA, j.jsonConfigB)
}

// diff will peform a diff on keys between two maps, storing keys
// that exist in map B but are missing in map A
func (j *jsonAnalyzer) diff(a configJSON, b configJSON) {
	keysA := j.keys(a, b)
	keysB := j.keys(b, a)

	for _, str := range keysB {
		if !j.contains(keysA, str) && !j.contains(j.missingKeys, str) {
			j.missingKeys = append(j.missingKeys, str)
		}
	}
}

// keys stores known missing keys between map a and map b
func (j *jsonAnalyzer) keys(a configJSON, b configJSON) []string {
	keys := []string{}

	for k, _ := range a {
		keys = append(keys, k)

		if j.isMap(a[k]) {
			if !j.isMap(b[k]) {
				j.missingKeys = append(j.missingKeys, k)
				continue
			}

			j.diff(a[k].(map[string]interface{}), b[k].(map[string]interface{}))
		}
	}

	return keys
}

// equalKeys will scan both maps determining whether map B
// has identical keys compared with map A
func (j jsonAnalyzer) equalKeys() (bool, error) {
	j.clearValues(j.jsonConfigA)
	j.clearValues(j.jsonConfigB)

	bytesA, err := json.Marshal(j.jsonConfigA)
	if err != nil {
		return false, err
	}

	bytesB, err := json.Marshal(j.jsonConfigB)
	if err != nil {
		return false, err
	}

	return string(bytesA) == string(bytesB), nil
}

// clearValues will set empty values for each key in a given map
func (j *jsonAnalyzer) clearValues(m map[string]interface{}) {
	for k, _ := range m {
		if !j.isMap(m[k]) {
			m[k] = ""
		} else {
			j.clearValues(m[k].(map[string]interface{}))
		}
	}
}

func (j jsonAnalyzer) isMap(m interface{}) bool {
	_, ok := m.(map[string]interface{})
	return ok
}

func (j jsonAnalyzer) contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
