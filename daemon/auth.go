package daemon // import "github.com/demonoid81/moby/daemon"

import (
	"context"

	"github.com/demonoid81/moby/api/types"
	"github.com/demonoid81/moby/dockerversion"
)

// AuthenticateToRegistry checks the validity of credentials in authConfig
func (daemon *Daemon) AuthenticateToRegistry(ctx context.Context, authConfig *types.AuthConfig) (string, string, error) {
	return daemon.RegistryService.Auth(ctx, authConfig, dockerversion.DockerUserAgent(ctx))
}
