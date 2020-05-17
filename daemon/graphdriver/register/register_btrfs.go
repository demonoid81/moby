// +build !exclude_graphdriver_btrfs,linux

package register // import "github.com/demonoid81/moby/daemon/graphdriver/register"

import (
	// register the btrfs graphdriver
	_ "github.com/demonoid81/moby/daemon/graphdriver/btrfs"
)
