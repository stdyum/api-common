package grpc

import (
	"google.golang.org/grpc"
)

type Routes interface {
	ConfigureRoutes() *grpc.Server
}
