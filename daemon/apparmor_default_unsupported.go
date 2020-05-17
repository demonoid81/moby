// +build !linux

package daemon // import "github.com/demonoid81/moby/daemon"

func ensureDefaultAppArmorProfile() error {
	return nil
}
