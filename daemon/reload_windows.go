package daemon // import "github.com/demonoid81/moby/daemon"

import "github.com/demonoid81/moby/daemon/config"

// reloadPlatform updates configuration with platform specific options
// and updates the passed attributes
func (daemon *Daemon) reloadPlatform(config *config.Config, attributes map[string]string) error {
	return nil
}
