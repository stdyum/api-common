package proxy

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/stdyum/api-common/env"
)

type Config struct {
	ProxyUrl string `env:"PROXY_URL"`
	ProxyLog bool   `env:"PROXY_LOG"`
}

func StartDefault() {
	config, err := LoadConfig()
	if err != nil {
		log.Fatal("error loading config", err)
	}

	err = Start(config)
	if err != nil {
		log.Fatal("error launching proxy server", err)
	}
}

func LoadConfig() (config Config, err error) {
	err = env.Fill(&config)
	return
}

func Start(config Config) error {
	http.HandleFunc("/", handleRequest(config))
	return http.ListenAndServe(":8080", nil)
}

func handleRequest(config Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		proxyReq, err := http.NewRequest(r.Method, config.ProxyUrl+r.URL.String(), r.Body)
		if err != nil {
			http.Error(w, "Error creating proxy request: "+err.Error(), http.StatusInternalServerError)
			return
		}

		for name, values := range r.Header {
			for _, value := range values {
				proxyReq.Header.Add(name, value)
			}
		}

		resp, err := http.DefaultTransport.RoundTrip(proxyReq)
		if err != nil {
			http.Error(w, "Error sending proxy request: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(resp.Body)

		for name, values := range resp.Header {
			for _, value := range values {
				w.Header().Add(name, value)
			}
		}

		if config.ProxyLog {
			logProxy(r, resp, startTime)
		}

		w.WriteHeader(resp.StatusCode)
		_, _ = io.Copy(w, resp.Body)
	}
}

func logProxy(req *http.Request, resp *http.Response, startTime time.Time) {
	log.Printf("[%s] |%s| %s | \"%s\"", resp.Status, req.Method, time.Now().Sub(startTime).String(), req.URL)
}
