package cfg

import (
	"errors"
	"strings"
)

// configEnv is a key value struct representing each env var
type configEnv struct {
	Key   string
	Value string
}

// envAnalyzer holds data for both ENV configs
type envAnalyzer struct {
	baseAnalyzer
	envWorking []configEnv
	envMaster  []configEnv
}

// newEnvAnalyzer returns a new envAnalyzer
func newEnvAnalyzer(c Config) (*envAnalyzer, error) {

	base, err := newBaseAnalyzer(c)
	if err != nil {
		return nil, err
	}

	analyzer := envAnalyzer{}

	working := strings.Split(string(base.working), "\n")
	master := strings.Split(string(base.master), "\n")

	analyzer.envWorking, err = analyzer.unmarshal(working)
	if err != nil {
		return nil, err
	}

	analyzer.envMaster, err = analyzer.unmarshal(master)
	if err != nil {
		return nil, err
	}

	return &analyzer, nil
}

// analyze will analyze two sets of env key value paids identifying keys that
// exist in config B but are missing in config A
func (e *envAnalyzer) analyze() {
	for _, master := range e.envMaster {
		exists := false
		for _, working := range e.envWorking {
			if master.Key == working.Key {
				exists = true
			}
		}

		if !exists {
			e.missingKeys = append(e.missingKeys, master.Key)
		}
	}
}

// unmarshal will unmarshal a slice of env vars into key value pairs (configEnv)
func (e envAnalyzer) unmarshal(env []string) ([]configEnv, error) {
	config := []configEnv{}

	for _, line := range env {
		parts := strings.Split(line, "=")

		if line == "" || strings.Index(line, "#") == 0 {
			continue
		}

		if len(parts) != 2 {
			return nil, errors.New("Invalid key value pair in .env config")
		}

		ce := configEnv{
			Key:   parts[0],
			Value: parts[1],
		}

		config = append(config, ce)
	}

	return config, nil
}
