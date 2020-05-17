package system // import "github.com/demonoid81/moby/api/server/router/system"

import (
	"context"
	"time"

	"github.com/demonoid81/moby/api/types"
	"github.com/demonoid81/moby/api/types/events"
	"github.com/demonoid81/moby/api/types/filters"
	"github.com/demonoid81/moby/api/types/swarm"
)

// Backend is the methods that need to be implemented to provide
// system specific functionality.
type Backend interface {
	SystemInfo() *types.Info
	SystemVersion() types.Version
	SystemDiskUsage(ctx context.Context) (*types.DiskUsage, error)
	SubscribeToEvents(since, until time.Time, ef filters.Args) ([]events.Message, chan interface{})
	UnsubscribeFromEvents(chan interface{})
	AuthenticateToRegistry(ctx context.Context, authConfig *types.AuthConfig) (string, string, error)
}

// ClusterBackend is all the methods that need to be implemented
// to provide cluster system specific functionality.
type ClusterBackend interface {
	Info() swarm.Info
}
