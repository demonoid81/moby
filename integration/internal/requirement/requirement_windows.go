// +build windows

package requirement // import "github.com/demonoid81/moby/integration/internal/requirement"

func overlayFSSupported() bool {
	return false
}

// Overlay2Supported returns true if the current system supports overlay2 as graphdriver
func Overlay2Supported(kernelVersion string) bool {
	return false
}
