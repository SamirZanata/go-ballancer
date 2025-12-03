package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync/atomic"
)

type Server struct {
	URL *url.URL
}

type LoadBalancer struct {
	Servers []*Server
	Current int64
}

func NewLoadBalancer(urls []string) *LoadBalancer {
	var servers []*Server
	for _, u := range urls {
		parsedURL, err := url.Parse(u)
		if err != nil {
			log.Fatalf("URL Inválida: %v", err)
		}
		servers = append(servers, &Server{URL: parsedURL})
	}
	return &LoadBalancer{
		Servers: servers,
		Current: 0,
	}
}

func (lb *LoadBalancer) GetNextServer() *Server {
	index := atomic.AddInt64(&lb.Current, 1) % int64(len(lb.Servers))
	return lb.Servers[index]
}

func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	targetServer := lb.GetNextServer()
	proxy := httputil.NewSingleHostReverseProxy(targetServer.URL)
	log.Printf("Encaminhando requisição para: %s\n", targetServer.URL.Host)
	proxy.ServeHTTP(w, r)
}

func main() {
	serverURLs := []string{
		"http://localhost:8001",
		"http://localhost:8002",
		"http://localhost:8003",
	}

	lb := NewLoadBalancer(serverURLs)

	server := http.Server{
		Addr:    ":8080",
		Handler: lb,
	}

	log.Println("Load Balancer iniciado na porta 8080. Servidores:", serverURLs)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
