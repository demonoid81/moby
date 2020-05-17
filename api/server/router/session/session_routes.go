package session // import "github.com/demonoid81/moby/api/server/router/session"

import (
	"context"
	"net/http"

	"github.com/demonoid81/moby/errdefs"
)

func (sr *sessionRouter) startSession(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	err := sr.backend.HandleHTTPRequest(ctx, w, r)
	if err != nil {
		return errdefs.InvalidParameter(err)
	}
	return nil
}
