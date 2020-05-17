// +build !exclude_graphdriver_overlay,linux

package register // import "github.com/demonoid81/moby/daemon/graphdriver/register"

import (
	// register the overlay graphdriver
	_ "github.com/demonoid81/moby/daemon/graphdriver/overlay"
)
