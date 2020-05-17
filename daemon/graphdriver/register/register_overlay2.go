// +build !exclude_graphdriver_overlay2,linux

package register // import "github.com/demonoid81/moby/daemon/graphdriver/register"

import (
	// register the overlay2 graphdriver
	_ "github.com/demonoid81/moby/daemon/graphdriver/overlay2"
)
