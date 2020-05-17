package daemon // import "github.com/demonoid81/moby/daemon"

import (
	"testing"

	"github.com/demonoid81/moby/api/types"
	"github.com/demonoid81/moby/dockerversion"
	"gotest.tools/v3/assert"
)

func TestFillLicense(t *testing.T) {
	v := &types.Info{}
	d := &Daemon{
		root: "/var/lib/docker/",
	}
	d.fillLicense(v)
	assert.Assert(t, v.ProductLicense == dockerversion.DefaultProductLicense)
}
