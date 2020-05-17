// +build !linux,!windows

package daemon // import "github.com/demonoid81/moby/daemon"

func secretsSupported() bool {
	return false
}
