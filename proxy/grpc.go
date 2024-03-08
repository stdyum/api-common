package proxy

import (
	"net"

	"github.com/mwitkow/grpc-proxy/proxy"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func StartGRPC(config Config) error {
	opt := grpc.WithTransportCredentials(insecure.NewCredentials())
	dst, err := grpc.Dial(config.ProxyUrl+":"+config.ProxyGRpcPort, opt)
	if err != nil {
		return err
	}

	listener, err := net.Listen("tcp", ":"+config.ProxyGRpcPort)
	if err != nil {
		return err
	}

	return proxy.NewProxy(dst).Serve(listener)
}
