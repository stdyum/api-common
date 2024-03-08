package proxy

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/stdyum/api-common/env"
)

type Config struct {
	ProxyUrl      string `env:"PROXY_URL"`
	ProxyHttpPort string `env:"PROXY_HTTP_PORT"`
	ProxyGRpcPort string `env:"PROXY_GRPC_PORT"`
}

func StartDefault() {
	config, err := LoadConfig()
	if err != nil {
		log.Fatal("error loading config", err)
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		if err = StartHttp(config); err != nil {
			log.Fatal("error launching http proxy server", err)
		}
	}()

	go func() {
		defer wg.Done()
		if err = StartGRPC(config); err != nil {
			log.Fatal("error launching grpc proxy server", err)
		}
	}()

	wg.Wait()
}

func LoadConfig() (config Config, err error) {
	err = env.Fill(&config)
	return
}

func logProxy(req *http.Request, resp *http.Response, startTime time.Time) {
	log.Printf("[%s] |%s| %s | \"%s\"", resp.Status, req.Method, time.Now().Sub(startTime).String(), req.URL)
}
