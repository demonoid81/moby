// +build !windows

package listeners // import "github.com/demonoid81/moby/daemon/listeners"

import (
	"fmt"
	"strconv"

	"github.com/demonoid81/moby/pkg/idtools"
)

const defaultSocketGroup = "docker"

func lookupGID(name string) (int, error) {
	group, err := idtools.LookupGroup(name)
	if err == nil {
		return group.Gid, nil
	}
	gid, err := strconv.Atoi(name)
	if err == nil {
		return gid, nil
	}
	return -1, fmt.Errorf("group %s not found", name)
}
