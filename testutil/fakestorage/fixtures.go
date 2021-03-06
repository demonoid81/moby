package fakestorage // import "github.com/demonoid81/moby/testutil/fakestorage"

import (
	"context"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"testing"

	"github.com/demonoid81/moby/api/types"
	"github.com/demonoid81/moby/pkg/archive"
	"gotest.tools/v3/assert"
)

var ensureHTTPServerOnce sync.Once

func ensureHTTPServerImage(t testing.TB) {
	t.Helper()
	var doIt bool
	ensureHTTPServerOnce.Do(func() {
		doIt = true
	})

	if !doIt {
		return
	}

	defer testEnv.ProtectImage(t, "httpserver:latest")

	tmp, err := ioutil.TempDir("", "docker-http-server-test")
	if err != nil {
		t.Fatalf("could not build http server: %v", err)
	}
	defer os.RemoveAll(tmp)

	goos := testEnv.OSType
	if goos == "" {
		goos = "linux"
	}
	goarch := os.Getenv("DOCKER_ENGINE_GOARCH")
	if goarch == "" {
		goarch = "amd64"
	}

	cpCmd, lookErr := exec.LookPath("cp")
	if lookErr != nil {
		t.Fatalf("could not build http server: %v", lookErr)
	}

	if _, err = os.Stat("../contrib/httpserver/httpserver"); os.IsNotExist(err) {
		goCmd, lookErr := exec.LookPath("go")
		if lookErr != nil {
			t.Fatalf("could not build http server: %v", lookErr)
		}

		cmd := exec.Command(goCmd, "build", "-o", filepath.Join(tmp, "httpserver"), "github.com/demonoid81/moby/contrib/httpserver")
		cmd.Env = append(os.Environ(), []string{
			"CGO_ENABLED=0",
			"GOOS=" + goos,
			"GOARCH=" + goarch,
		}...)
		var out []byte
		if out, err = cmd.CombinedOutput(); err != nil {
			t.Fatalf("could not build http server: %s", string(out))
		}
	} else {
		if out, err := exec.Command(cpCmd, "../contrib/httpserver/httpserver", filepath.Join(tmp, "httpserver")).CombinedOutput(); err != nil {
			t.Fatalf("could not copy http server: %v", string(out))
		}
	}

	if out, err := exec.Command(cpCmd, "../contrib/httpserver/Dockerfile", filepath.Join(tmp, "Dockerfile")).CombinedOutput(); err != nil {
		t.Fatalf("could not build http server: %v", string(out))
	}

	c := testEnv.APIClient()
	reader, err := archive.TarWithOptions(tmp, &archive.TarOptions{})
	assert.NilError(t, err)
	resp, err := c.ImageBuild(context.Background(), reader, types.ImageBuildOptions{
		Remove:      true,
		ForceRemove: true,
		Tags:        []string{"httpserver"},
	})
	assert.NilError(t, err)
	_, err = io.Copy(ioutil.Discard, resp.Body)
	assert.NilError(t, err)
}
