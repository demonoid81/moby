// +build !linux,!darwin,!freebsd,!windows

package daemon // import "github.com/demonoid81/moby/daemon"

func (daemon *Daemon) setupDumpStackTrap(_ string) {
	return
}
