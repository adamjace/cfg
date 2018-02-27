package cfganalyze

import (
	"fmt"
	"io/ioutil"
)

// analyzer contains data for analyzing config files
type analyzer struct {
	configA []byte
	configB []byte
	bash    *bash
}

// NewAnalyzer returns a new Analyzer
func NewAnalyzer() *analyzer {
	return &analyzer{}
}

// Connect will return a new connected Analyzer to an external host via SSH
func Connect(hostAlias string) (*analyzer, error) {

	bash := newBash(hostAlias)

	if err := bash.ssh(); err != nil {
		return nil, fmt.Errorf(
			"could not connect to host %s. %s", hostAlias, err)
	}

	analyzer := &analyzer{
		bash: bash,
	}

	return analyzer, nil
}

// read will read a config file to []byte
func (c *analyzer) read(pathA, pathB string) error {
	var err error

	c.configA, err = ioutil.ReadFile(pathA)
	if err != nil {
		return fmt.Errorf("could not open %s. %s", pathA, err)
	}

	// we have a remote file. read in the contents via scp
	if c.bash != nil {
		c.configB, err = c.bash.scp(pathB)
		if err != nil {
			return fmt.Errorf("could not open %s. %s", pathB, err)
		}

		return nil
	}

	c.configB, err = ioutil.ReadFile(pathB)
	if err != nil {
		return fmt.Errorf("could not open %s. %s", pathB, err)
	}

	return nil
}

// AnalyzeJson will compare two .json configuration files
// highlighting keys that are missing
func (c analyzer) AnalyzeJson(a, b string) ([]string, error) {
	if err := c.read(a, b); err != nil {
		return nil, err
	}

	analyzer, err := newJsonAnalyzer(c.configA, c.configB)
	if err != nil {
		return nil, err
	}

	analyzer.analyze()

	return analyzer.missingKeys, nil
}

// AnalyzeEnv will compare two .env configuration files
// highlighting keys that are missing
func (c analyzer) AnalyzeEnv(a, b string) ([]string, error) {
	if err := c.read(a, b); err != nil {
		return nil, err
	}

	analyzer, err := newEnvAnalyzer(c.configA, c.configB)
	if err != nil {
		return nil, err
	}

	analyzer.analyze()

	return analyzer.missingKeys, nil
}

// EqualKeys will compare two configurations identifying whether they are the same
func (c analyzer) EqualKeys(a, b string) (bool, error) {
	if err := c.read(a, b); err != nil {
		return false, err
	}

	analyzer, err := newJsonAnalyzer(c.configA, c.configB)
	if err != nil {
		return false, err
	}

	return analyzer.equalKeys()
}
