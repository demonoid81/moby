// +build linux,cgo,!static_build

package devicemapper // import "github.com/demonoid81/moby/pkg/devicemapper"

// #cgo pkg-config: devmapper
import "C"
