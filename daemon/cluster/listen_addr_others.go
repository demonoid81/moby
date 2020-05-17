// +build !linux

package cluster // import "github.com/demonoid81/moby/daemon/cluster"

import "net"

func (c *Cluster) resolveSystemAddr() (net.IP, error) {
	return c.resolveSystemAddrViaSubnetCheck()
}
