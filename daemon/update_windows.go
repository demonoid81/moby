package daemon // import "github.com/demonoid81/moby/daemon"

import (
	"github.com/demonoid81/moby/api/types/container"
	libcontainerdtypes "github.com/demonoid81/moby/libcontainerd/types"
)

func toContainerdResources(resources container.Resources) *libcontainerdtypes.Resources {
	// We don't support update, so do nothing
	return nil
}
