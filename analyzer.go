package cfg

import (
	"fmt"
	"io/ioutil"
)

type analyzer interface {
	analyze()
}

// baseAnalyzer contains data for analyzing config files
type baseAnalyzer struct {
	working     []byte
	master      []byte
	bash        *bash
	missingKeys []string
}

// newBaseAnalyzer returns a new baseAnalyzer
func newBaseAnalyzer(c Config) (*baseAnalyzer, error) {

	analyzer := baseAnalyzer{}

	if len(c.HostAlias) > 0 {
		if err := analyzer.connect(c.HostAlias); err != nil {
			return nil, err
		}
	}

	if err := analyzer.read(c.WorkingPath, c.MasterPath); err != nil {
		return nil, err
	}

	return &analyzer, nil
}

// connect will return a new connected Analyzer to an external host via SSH
func (b *baseAnalyzer) connect(hostAlias string) error {

	b.bash = newBash(hostAlias)

	if err := b.bash.ssh(); err != nil {
		fmt.Errorf("could not connect to host %s. %s", hostAlias, err)
	}

	return nil
}

// read will read a config file to []byte
func (b *baseAnalyzer) read(workingPath, masterPath string) error {
	var err error

	b.working, err = ioutil.ReadFile(workingPath)
	if err != nil {
		return fmt.Errorf("could not open %s. %s", workingPath, err)
	}

	// we have a remote file. read in the contents via scp
	if b.bash != nil {
		b.master, err = b.bash.scp(masterPath)
		if err != nil {
			return fmt.Errorf("could not open %s. %s", masterPath, err)
		}

		return nil
	}

	b.master, err = ioutil.ReadFile(masterPath)
	if err != nil {
		return fmt.Errorf("could not open %s. %s", masterPath, err)
	}

	return nil
}

// ScanJson will scan two JSON configuration files returning a slice
// of keys that are missing in the working file
func ScanJson(c Config) ([]string, error) {
	analyzer, err := newJsonAnalyzer(c)
	if err != nil {
		return nil, err
	}

	analyzer.analyze()

	return analyzer.missingKeys, nil
}

// ScanEnv will scan two ENV configuration files returning a slice
// of keys that are missing in the working file
func ScanEnv(c Config) ([]string, error) {
	analyzer, err := newEnvAnalyzer(c)
	if err != nil {
		return nil, err
	}

	analyzer.analyze()

	return analyzer.missingKeys, nil
}

// AnalyzeJson will compare two .json configuration files
// highlighting keys that are missing
func AnalyzeJson(c Config) error {
	keys, err := ScanJson(c)
	if err != nil {
		return err
	}

	if len(keys) == 0 {
		return nil
	}

	fmt.Printf("warning! missing keys from json file (%s): %+v\n", c.MasterPath, keys)

	return nil
}

// AnalyzeEnv will compare two .env configuration files
// highlighting keys that are missing
func AnalyzeEnv(c Config) error {
	keys, err := ScanEnv(c)
	if err != nil {
		return err
	}

	if len(keys) == 0 {
		return nil
	}

	fmt.Printf("warning! missing keys from env file (%s): %+v\n", c.MasterPath, keys)

	return nil
}
