// +build !windows

package ioutils // import "github.com/demonoid81/moby/pkg/ioutils"

import "io/ioutil"

// TempDir on Unix systems is equivalent to ioutil.TempDir.
func TempDir(dir, prefix string) (string, error) {
	return ioutil.TempDir(dir, prefix)
}
