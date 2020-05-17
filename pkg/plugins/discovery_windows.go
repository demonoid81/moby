package plugins // import "github.com/demonoid81/moby/pkg/plugins"

import (
	"os"
	"path/filepath"
)

var specsPaths = []string{filepath.Join(os.Getenv("programdata"), "docker", "plugins")}
