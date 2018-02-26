package cfganalyze

import (
	"fmt"
	"os/exec"
)

type bash struct {
	hostAlias string
}

func newBash(host string) *bash {
	return &bash{
		hostAlias: host,
	}
}

func (b bash) ssh() error {
	_, err := b.command(fmt.Sprintf("ssh %s", b.hostAlias))
	return err
}

func (b bash) scp(path string) ([]byte, error) {
	return b.command(
		fmt.Sprintf("scp %s:%s /dev/stdout", b.hostAlias, path))
}

func (b bash) command(cmd string) ([]byte, error) {
	return exec.Command("/bin/bash", "-c", cmd).Output()
}
