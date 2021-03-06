package libcontainerd // import "github.com/demonoid81/moby/libcontainerd"

import (
	"context"

	"github.com/containerd/containerd"
	"github.com/demonoid81/moby/libcontainerd/remote"
	libcontainerdtypes "github.com/demonoid81/moby/libcontainerd/types"
)

// NewClient creates a new libcontainerd client from a containerd client
func NewClient(ctx context.Context, cli *containerd.Client, stateDir, ns string, b libcontainerdtypes.Backend, useShimV2 bool) (libcontainerdtypes.Client, error) {
	return remote.NewClient(ctx, cli, stateDir, ns, b, useShimV2)
}
