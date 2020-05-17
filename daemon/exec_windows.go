package daemon // import "github.com/demonoid81/moby/daemon"

import (
	"github.com/demonoid81/moby/container"
	"github.com/demonoid81/moby/daemon/exec"
	specs "github.com/opencontainers/runtime-spec/specs-go"
)

func (daemon *Daemon) execSetPlatformOpt(c *container.Container, ec *exec.Config, p *specs.Process) error {
	if c.OS == "windows" {
		p.User.Username = ec.User
	}
	return nil
}
