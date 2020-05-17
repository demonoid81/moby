// +build !linux

package archive // import "github.com/demonoid81/moby/pkg/archive"

func getWhiteoutConverter(format WhiteoutFormat, inUserNS bool) tarWhiteoutConverter {
	return nil
}
