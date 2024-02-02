package server

import (
	"github.com/sirupsen/logrus"
	"github.com/stdyum/api-common/grpc"
	"github.com/stdyum/api-common/http"
	"net"
	"sync"
)

type PortConfig struct {
	HTTP string `env:"HTTP"`
	GRPC string `env:"GRPC"`
}

type Routes struct {
	GRPC grpc.Routes
	HTTP http.Routes
}

func (r *Routes) Run(ports PortConfig) error {
	var wg sync.WaitGroup

	if r.GRPC != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			listener, err := net.Listen("tcp", ":"+ports.GRPC)
			if err != nil {
				logrus.Errorf("http server error: %v\n", err)
				return
			}

			if err = r.GRPC.ConfigureRoutes().Serve(listener); err != nil {
				logrus.Errorf("http server error: %v\n", err)
				return
			}
		}()
	}

	if r.HTTP != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := r.HTTP.ConfigureRoutes().Run(":" + ports.HTTP); err != nil {
				logrus.Errorf("http server error: %v\n", err)
				return
			}
		}()
	}

	wg.Wait()
	return nil
}
