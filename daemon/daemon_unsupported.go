// +build !linux,!freebsd,!windows

package daemon // import "github.com/demonoid81/moby/daemon"

import (
	"github.com/demonoid81/moby/daemon/config"
	"github.com/demonoid81/moby/pkg/sysinfo"
)

const platformSupported = false

func setupResolvConf(config *config.Config) {
}

// RawSysInfo returns *sysinfo.SysInfo .
func (daemon *Daemon) RawSysInfo(quiet bool) *sysinfo.SysInfo {
	return sysinfo.New(quiet)
}
