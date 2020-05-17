// +build linux freebsd

package system // import "github.com/demonoid81/moby/pkg/system"

import "golang.org/x/sys/unix"

// Unmount is a platform-specific helper function to call
// the unmount syscall.
func Unmount(dest string) error {
	return unix.Unmount(dest, 0)
}
