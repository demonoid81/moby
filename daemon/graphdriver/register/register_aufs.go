// +build !exclude_graphdriver_aufs,linux

package register // import "github.com/demonoid81/moby/daemon/graphdriver/register"

import (
	// register the aufs graphdriver
	_ "github.com/demonoid81/moby/daemon/graphdriver/aufs"
)
