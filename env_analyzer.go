package cfganalyze

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
	envConfigA  []configEnv
	envConfigB  []configEnv
	missingKeys []string
}

// newEnvAnalyzer returns a new envAnalyzer
func newEnvAnalyzer(a, b []byte) (*envAnalyzer, error) {
	envAnalyzer := envAnalyzer{}

	envA := strings.Split(string(a), "\n")
	envB := strings.Split(string(b), "\n")

	var err error

	envAnalyzer.envConfigA, err = envAnalyzer.unmarshal(envA)
	if err != nil {
		return nil, err
	}

	envAnalyzer.envConfigB, err = envAnalyzer.unmarshal(envB)
	if err != nil {
		return nil, err
	}

	return &envAnalyzer, nil
}

// analyze will analyze two sets of env key value paids identifying keys that
// exist in config B but are missing in config A
func (e *envAnalyzer) analyze() {
	for _, itemB := range e.envConfigB {
		exists := false
		for _, itemA := range e.envConfigA {
			if itemB.Key == itemA.Key {
				exists = true
			}
		}

		if !exists {
			e.missingKeys = append(e.missingKeys, itemB.Key)
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
