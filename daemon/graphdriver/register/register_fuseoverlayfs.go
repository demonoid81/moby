// +build !exclude_graphdriver_fuseoverlayfs,linux

package register // import "github.com/demonoid81/moby/daemon/graphdriver/register"

import (
	// register the fuse-overlayfs graphdriver
	_ "github.com/demonoid81/moby/daemon/graphdriver/fuse-overlayfs"
)
