package daemon // import "github.com/demonoid81/moby/daemon"

import (
	"github.com/demonoid81/moby/container"
	"github.com/demonoid81/moby/pkg/archive"
)

func (daemon *Daemon) tarCopyOptions(container *container.Container, noOverwriteDirNonDir bool) (*archive.TarOptions, error) {
	return daemon.defaultTarCopyOptions(noOverwriteDirNonDir), nil
}
