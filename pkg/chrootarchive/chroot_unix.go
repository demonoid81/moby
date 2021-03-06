// +build !windows,!linux

package chrootarchive // import "github.com/demonoid81/moby/pkg/chrootarchive"

import "golang.org/x/sys/unix"

func chroot(path string) error {
	if err := unix.Chroot(path); err != nil {
		return err
	}
	return unix.Chdir("/")
}

func realChroot(path string) error {
	return chroot(path)
}
