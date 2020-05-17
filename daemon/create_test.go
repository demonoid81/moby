package daemon // import "github.com/demonoid81/moby/daemon"

import (
	"testing"

	"github.com/demonoid81/moby/api/types/network"
	"github.com/demonoid81/moby/errdefs"
	"gotest.tools/v3/assert"
)

// Test case for 35752
func TestVerifyNetworkingConfig(t *testing.T) {
	name := "mynet"
	endpoints := make(map[string]*network.EndpointSettings, 1)
	endpoints[name] = nil
	nwConfig := &network.NetworkingConfig{
		EndpointsConfig: endpoints,
	}
	err := verifyNetworkingConfig(nwConfig)
	assert.Check(t, errdefs.IsInvalidParameter(err))
}
