package cfg

import (
	"encoding/json"
)

// jsoncfg is a basic map struct for json configs
type jsoncfg map[string]interface{}

// jsonAnalyzer holds data for both JSON maps
type jsonAnalyzer struct {
	analyzer
	jsonWorking jsoncfg
	jsonMaster  jsoncfg
}

// newJsonAnalyzer returns a new jsonAnalyzer loaded with json maps
func newJsonAnalyzer(c Config) (*jsonAnalyzer, error) {

	analyzer, err := newAnalyzer(c)
	if err != nil {
		return nil, err
	}

	working := jsoncfg{}
	if err := json.Unmarshal(analyzer.working, &working); err != nil {
		return nil, err
	}

	master := jsoncfg{}
	if err := json.Unmarshal(analyzer.master, &master); err != nil {
		return nil, err
	}

	jsonAnalyzer := jsonAnalyzer{
		analyzer:    *analyzer,
		jsonWorking: working,
		jsonMaster:  master,
	}

	return &jsonAnalyzer, nil
}

// scan will analyze two sets of json config files identifying keys that
// exist in the master file and are missing in the working file
func (j *jsonAnalyzer) scan() {
	j.diff(j.jsonWorking, j.jsonMaster)
}

// diff will peform a diff on keys between two maps, storing ones
// that exist in the master and are missing in the working file
func (j *jsonAnalyzer) diff(working jsoncfg, master jsoncfg) {
	a := j.keys(working, master)
	b := j.keys(master, working)

	for _, str := range b {
		if !j.contains(a, str) && !j.contains(j.missing, str) {
			j.missing = append(j.missing, str)
		}
	}
}

// keys stores known missing keys between the two maps
func (j *jsonAnalyzer) keys(working jsoncfg, master jsoncfg) []string {
	keys := []string{}

	for k, _ := range working {
		keys = append(keys, k)

		if j.isMap(working[k]) {
			if !j.isMap(master[k]) {
				j.missing = append(j.missing, k)
				continue
			}

			j.diff(working[k].(map[string]interface{}),
				master[k].(map[string]interface{}))
		}
	}

	return keys
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

// isMap determins if the interface passed in is a go map or not
func (j jsonAnalyzer) isMap(m interface{}) bool {
	_, ok := m.(map[string]interface{})
	return ok
}

// contains is a simple util func for determining the existance of a
// string value within a slice
func (j jsonAnalyzer) contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
