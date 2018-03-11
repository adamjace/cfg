package cfg

import (
	"errors"
	"fmt"
	"strings"
)

// configEnv is a key value struct representing each env var
type configEnv struct {
	Key   string
	Value string
}

// envAnalyzer holds data for both working and master .env config files
type envAnalyzer struct {
	analyzer
	envWorking []configEnv
	envMaster  []configEnv
}

// newEnvAnalyzer returns a new envAnalyzer
func newEnvAnalyzer(c Config) (*envAnalyzer, error) {

	base, err := newAnalyzer(c)
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

// scan will analyze two sets of env key value pairs identifying:
// 1) keys that exist in the master file and are missing in the working file
// 2) keys that exists but have different values
func (e *envAnalyzer) scan() {
	for _, master := range e.envMaster {
		exists := false
		for _, working := range e.envWorking {
			if master.Key == working.Key {
				if master.Value != working.Value {
					e.different = append(e.different,
						fmt.Sprintf("%s=%s", working.Key, working.Value))
				}

				exists = true
			}
		}

		if !exists {
			e.missing = append(e.missing, master.Key)
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
			return nil, errors.New("Invalid key value pair found in config")
		}

		c := configEnv{
			Key:   parts[0],
			Value: parts[1],
		}

		config = append(config, c)
	}

	return config, nil
}
