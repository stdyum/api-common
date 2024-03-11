package server

import (
	"net"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/stdyum/api-common/grpc"
	"github.com/stdyum/api-common/http"
	"github.com/stdyum/api-common/http/metrics"
)

type PortConfig struct {
	HTTP string `env:"HTTP"`
	GRPC string `env:"GRPC"`
}

type Routes struct {
	Error error
	Ports PortConfig

	GRPC grpc.Routes
	HTTP http.Routes
}

func (r Routes) Run() error {
	if r.Error != nil {
		return r.Error
	}

	var wg sync.WaitGroup

	if r.GRPC != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			listener, err := net.Listen("tcp", ":"+r.Ports.GRPC)
			if err != nil {
				logrus.Errorf("http server error: %v\n", err)
				return
			}

			routes := r.GRPC.ConfigureRoutes()
			if routes == nil {
				return
			}

			if err = routes.Serve(listener); err != nil {
				logrus.Errorf("http server error: %v\n", err)
				return
			}
		}()
	}

	if r.HTTP != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()

			e := r.HTTP.ConfigureRoutes()
			if e == nil {
				return
			}

			metrics.ApplyMetrics(e)
			if err := e.Run(":" + r.Ports.HTTP); err != nil {
				logrus.Errorf("http server error: %v\n", err)
				return
			}
		}()
	}

	wg.Wait()
	return nil
}
