// +build linux freebsd

package daemon // import "github.com/demonoid81/moby/daemon"

import "github.com/demonoid81/moby/container"

// excludeByIsolation is a platform specific helper function to support PS
// filtering by Isolation. This is a Windows-only concept, so is a no-op on Unix.
func excludeByIsolation(container *container.Snapshot, ctx *listContext) iterationAction {
	return includeContainer
}
