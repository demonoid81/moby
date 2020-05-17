package daemon // import "github.com/demonoid81/moby/daemon"

import (
	"github.com/demonoid81/moby/api/types"
	"github.com/demonoid81/moby/dockerversion"
)

func (daemon *Daemon) fillLicense(v *types.Info) {
	v.ProductLicense = dockerversion.DefaultProductLicense
}
