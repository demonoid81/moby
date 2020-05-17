//+build !windows

package daemon // import "github.com/demonoid81/moby/daemon"

import (
	"github.com/demonoid81/moby/container"
	"github.com/demonoid81/moby/errdefs"
)

func (daemon *Daemon) saveAppArmorConfig(container *container.Container) error {
	container.AppArmorProfile = "" // we don't care about the previous value.

	if !daemon.apparmorEnabled {
		return nil // if apparmor is disabled there is nothing to do here.
	}

	if err := parseSecurityOpt(container, container.HostConfig); err != nil {
		return errdefs.InvalidParameter(err)
	}

	if !container.HostConfig.Privileged {
		if container.AppArmorProfile == "" {
			container.AppArmorProfile = defaultAppArmorProfile
		}

	} else {
		container.AppArmorProfile = unconfinedAppArmorProfile
	}
	return nil
}
