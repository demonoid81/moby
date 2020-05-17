// +build linux freebsd darwin openbsd

package layer // import "github.com/demonoid81/moby/layer"

import "github.com/demonoid81/moby/pkg/stringid"

func (ls *layerStore) mountID(name string) string {
	return stringid.GenerateRandomID()
}
