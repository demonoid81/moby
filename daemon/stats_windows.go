package daemon // import "github.com/demonoid81/moby/daemon"

import (
	"github.com/demonoid81/moby/api/types"
	"github.com/demonoid81/moby/container"
)

// Windows network stats are obtained directly through HCS, hence this is a no-op.
func (daemon *Daemon) getNetworkStats(c *container.Container) (map[string]types.NetworkStats, error) {
	return make(map[string]types.NetworkStats), nil
}
