package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var serverCount = 0

const (
	SERVER1 = "https://www.google.com"
	SERVER2 = "https://www.facebook.com"
	SERVER3 = "https://www.yahoo.com"
	PORT    = "1338"
)

func main() {
	http.HandleFunc("/", loadBalancer)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}

func loadBalancer(res http.ResponseWriter, req *http.Request) {
	url := getProxyUrl()

	logRequestPayload(url)

	serveReverseProxy(url, res, req)
}

func getProxyUrl() string {
	var servers = []string{SERVER1, SERVER2, SERVER3}
	server := servers[serverCount]
	serverCount++

	if serverCount >= len(servers) {
		serverCount = 0
	}

	return server
}

func logRequestPayload(proxyURL string) {
	log.Printf("proxy_url: %s\n", proxyURL)
}

func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	url, _ := url.Parse(target)

	proxy := httputil.NewSingleHostReverseProxy(url)

	proxy.ServeHTTP(res, req)
}