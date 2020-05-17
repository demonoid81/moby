// +build !exclude_graphdriver_zfs,linux !exclude_graphdriver_zfs,freebsd

package register // import "github.com/demonoid81/moby/daemon/graphdriver/register"

import (
	// register the zfs driver
	_ "github.com/demonoid81/moby/daemon/graphdriver/zfs"
)
