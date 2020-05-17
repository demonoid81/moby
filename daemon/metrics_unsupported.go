// +build windows

package daemon // import "github.com/demonoid81/moby/daemon"

import "github.com/demonoid81/moby/pkg/plugingetter"

func registerMetricsPluginCallback(getter plugingetter.PluginGetter, sockPath string) {
}

func (daemon *Daemon) listenMetricsSock() (string, error) {
	return "", nil
}
