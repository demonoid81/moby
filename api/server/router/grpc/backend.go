package grpc // import "github.com/demonoid81/moby/api/server/router/grpc"

import "google.golang.org/grpc"

// Backend abstracts a registerable GRPC service.
type Backend interface {
	RegisterGRPC(*grpc.Server)
}
