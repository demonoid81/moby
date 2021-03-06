package vfs // import "github.com/demonoid81/moby/daemon/graphdriver/vfs"

import "github.com/demonoid81/moby/daemon/graphdriver/copy"

func dirCopy(srcDir, dstDir string) error {
	return copy.DirCopy(srcDir, dstDir, copy.Content, false)
}
