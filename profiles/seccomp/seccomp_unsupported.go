// +build linux,!seccomp

package seccomp // import "github.com/demonoid81/moby/profiles/seccomp"

import (
	"github.com/demonoid81/moby/api/types"
)

// DefaultProfile returns a nil pointer on unsupported systems.
func DefaultProfile() *types.Seccomp {
	return nil
}
