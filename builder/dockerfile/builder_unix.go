// +build !windows

package dockerfile // import "github.com/demonoid81/moby/builder/dockerfile"

func defaultShellForOS(os string) []string {
	return []string{"/bin/sh", "-c"}
}
