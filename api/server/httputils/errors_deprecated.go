package httputils // import "github.com/demonoid81/moby/api/server/httputils"
import "github.com/demonoid81/moby/errdefs"

// GetHTTPErrorStatusCode retrieves status code from error message.
//
// Deprecated: use errdefs.GetHTTPErrorStatusCode
func GetHTTPErrorStatusCode(err error) int {
	return errdefs.GetHTTPErrorStatusCode(err)
}
