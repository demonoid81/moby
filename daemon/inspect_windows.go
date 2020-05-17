package daemon // import "github.com/demonoid81/moby/daemon"

import (
	"github.com/demonoid81/moby/api/types"
	"github.com/demonoid81/moby/api/types/backend"
	"github.com/demonoid81/moby/container"
	"github.com/demonoid81/moby/daemon/exec"
)

// This sets platform-specific fields
func setPlatformSpecificContainerFields(container *container.Container, contJSONBase *types.ContainerJSONBase) *types.ContainerJSONBase {
	return contJSONBase
}

// containerInspectPre120 get containers for pre 1.20 APIs.
func (daemon *Daemon) containerInspectPre120(name string) (*types.ContainerJSON, error) {
	return daemon.ContainerInspectCurrent(name, false)
}

func inspectExecProcessConfig(e *exec.Config) *backend.ExecProcessConfig {
	return &backend.ExecProcessConfig{
		Tty:        e.Tty,
		Entrypoint: e.Entrypoint,
		Arguments:  e.Args,
	}
}
