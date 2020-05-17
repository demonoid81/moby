// +build !linux

package vfs // import "github.com/demonoid81/moby/daemon/graphdriver/vfs"

import "github.com/demonoid81/moby/pkg/chrootarchive"

func dirCopy(srcDir, dstDir string) error {
	return chrootarchive.NewArchiver(nil).CopyWithTar(srcDir, dstDir)
}
