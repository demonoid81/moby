package runconfig // import "github.com/demonoid81/moby/runconfig"

import (
	"github.com/demonoid81/moby/api/types/container"
	networktypes "github.com/demonoid81/moby/api/types/network"
)

// ContainerConfigWrapper is a Config wrapper that holds the container Config (portable)
// and the corresponding HostConfig (non-portable).
type ContainerConfigWrapper struct {
	*container.Config
	HostConfig       *container.HostConfig          `json:"HostConfig,omitempty"`
	NetworkingConfig *networktypes.NetworkingConfig `json:"NetworkingConfig,omitempty"`
}

// getHostConfig gets the HostConfig of the Config.
func (w *ContainerConfigWrapper) getHostConfig() *container.HostConfig {
	return w.HostConfig
}
