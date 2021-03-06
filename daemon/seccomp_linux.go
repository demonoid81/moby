// +build linux,seccomp

package daemon // import "github.com/demonoid81/moby/daemon"

import (
	"context"
	"fmt"

	"github.com/containerd/containerd/containers"
	coci "github.com/containerd/containerd/oci"
	"github.com/demonoid81/moby/container"
	"github.com/demonoid81/moby/profiles/seccomp"
	specs "github.com/opencontainers/runtime-spec/specs-go"
	"github.com/sirupsen/logrus"
)

const supportsSeccomp = true

// WithSeccomp sets the seccomp profile
func WithSeccomp(daemon *Daemon, c *container.Container) coci.SpecOpts {
	return func(ctx context.Context, _ coci.Client, _ *containers.Container, s *coci.Spec) error {
		var profile *specs.LinuxSeccomp
		var err error

		if c.HostConfig.Privileged {
			return nil
		}

		if !daemon.seccompEnabled {
			if c.SeccompProfile != "" && c.SeccompProfile != "unconfined" {
				return fmt.Errorf("seccomp is not enabled in your kernel, cannot run a custom seccomp profile")
			}
			logrus.Warn("seccomp is not enabled in your kernel, running container without default profile")
			c.SeccompProfile = "unconfined"
		}
		if c.SeccompProfile == "unconfined" {
			return nil
		}
		if c.SeccompProfile != "" {
			profile, err = seccomp.LoadProfile(c.SeccompProfile, s)
			if err != nil {
				return err
			}
		} else {
			if daemon.seccompProfile != nil {
				profile, err = seccomp.LoadProfile(string(daemon.seccompProfile), s)
				if err != nil {
					return err
				}
			} else {
				profile, err = seccomp.GetDefaultProfile(s)
				if err != nil {
					return err
				}
			}
		}

		s.Linux.Seccomp = profile
		return nil
	}
}
