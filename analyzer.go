package cfg

import (
	"fmt"
	"io/ioutil"
)

// baseAnalyzer contains base data for analyzing all supported types of config
// files.
//
// The working file is considered is considered to be the current local or
// active config file driving the local application.
//
// The master file is considered to be the 'compare to' file which could either
// be an example file, locally or an active remote config file on a server.
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

// ScanJson will scan two .json configuration files returning a slice
// of keys that exist in the master file and are missing in the working file
func ScanJson(c Config) ([]string, error) {
	analyzer, err := newJsonAnalyzer(c)
	if err != nil {
		return nil, err
	}

	analyzer.analyze()

	return analyzer.missingKeys, nil
}

// ScanEnv will scan two .env configuration files returning a slice
// of keys that exist in the master file and are missing in the working file
func ScanEnv(c Config) ([]string, error) {
	analyzer, err := newEnvAnalyzer(c)
	if err != nil {
		return nil, err
	}

	analyzer.analyze()

	return analyzer.missingKeys, nil
}

// PrintJson uses ScanJson to retrieve a slice of missing keys and will then
// print out the difference / discrepencies between the master and working files
func PrintJson(c Config) error {
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

// PrintEnv uses ScanEnv to retrieve a slice of missing keys and will then
// print out the difference / discrepencies between the master and working files
func PrintEnv(c Config) error {
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

// connect will return a new connected Analyzer to an external host via SSH
// currently this only supports connection via bash
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
