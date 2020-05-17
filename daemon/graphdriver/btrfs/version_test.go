// +build linux

package btrfs // import "github.com/demonoid81/moby/daemon/graphdriver/btrfs"

import (
	"testing"
)

func TestLibVersion(t *testing.T) {
	if btrfsLibVersion() <= 0 {
		t.Error("expected output from btrfs lib version > 0")
	}
}
