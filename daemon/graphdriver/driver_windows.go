package graphdriver // import "github.com/demonoid81/moby/daemon/graphdriver"

var (
	// List of drivers that should be used in order
	priority = "windowsfilter"
)

// GetFSMagic returns the filesystem id given the path.
func GetFSMagic(rootpath string) (FsMagic, error) {
	// Note it is OK to return FsMagicUnsupported on Windows.
	return FsMagicUnsupported, nil
}
