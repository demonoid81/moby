// +build !linux,!windows

package service // import "github.com/demonoid81/moby/volume/service"

import (
	"github.com/demonoid81/moby/pkg/idtools"
	"github.com/demonoid81/moby/volume/drivers"
)

func setupDefaultDriver(_ *drivers.Store, _ string, _ idtools.Identity) error { return nil }
