package cfg

import (
	"fmt"
	"os/exec"
)

// bash holds data for connecting to an external host via bash
type bash struct {
	hostAlias string
}

// newBash returns a new bash
func newBash(host string) *bash {
	return &bash{
		hostAlias: host,
	}
}

// ssh runs a ssh command
func (b bash) ssh() error {
	_, err := b.command(fmt.Sprintf("ssh %s", b.hostAlias))
	return err
}

// scp runs a scp (secure copy) command
func (b bash) scp(path string) ([]byte, error) {
	return b.command(
		fmt.Sprintf("scp %s:%s /dev/stdout", b.hostAlias, path))
}

// command is the executable command
func (b bash) command(cmd string) ([]byte, error) {
	return exec.Command("/bin/bash", "-c", cmd).Output()
}
