// +build !linux

package daemon // import "github.com/demonoid81/moby/daemon"

func selinuxSetDisabled() {
}

func selinuxFreeLxcContexts(label string) {
}

func selinuxEnabled() bool {
	return false
}
