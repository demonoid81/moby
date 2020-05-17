package register // import "github.com/demonoid81/moby/daemon/graphdriver/register"

import (
	// register the windows graph drivers
	_ "github.com/demonoid81/moby/daemon/graphdriver/lcow"
	_ "github.com/demonoid81/moby/daemon/graphdriver/windows"
)
