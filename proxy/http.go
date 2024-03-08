package proxy

import (
	"io"
	"net/http"
	"time"
)

func StartHttp(config Config) error {
	http.HandleFunc("/", handleHttpRequest(config))
	return http.ListenAndServe(":"+config.ProxyHttpPort, nil)
}

func handleHttpRequest(config Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		proxyReq, err := http.NewRequest(r.Method, "http://"+config.ProxyUrl+":"+config.ProxyHttpPort+r.URL.String(), r.Body)
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

		logProxy(r, resp, startTime)

		w.WriteHeader(resp.StatusCode)
		_, _ = io.Copy(w, resp.Body)
	}
}
