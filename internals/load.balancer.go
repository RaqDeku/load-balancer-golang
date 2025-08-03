package internals

import (
	"io"
	"net/http"
	"sync"
)

type LoadBalancer struct {
	Servers          []NodeServer
	RoundRobinNumber int
	mu               sync.Mutex
}

func NewLoadBalancer(servers []NodeServer) *LoadBalancer {
	return &LoadBalancer{
		Servers:          servers,
		RoundRobinNumber: 0,
		mu:               sync.Mutex{},
	}
}

func (lb *LoadBalancer) GetNextAvailableServer() NodeServer {
	numOfServers := len(lb.Servers)

	lb.mu.Lock()
	defer lb.mu.Unlock()

	for i := range numOfServers {
		serverIdx := (lb.RoundRobinNumber + i) % len(lb.Servers)

		if ok := lb.Servers[serverIdx].IsServerHealthy(); ok {
			lb.RoundRobinNumber = (serverIdx + 1) % numOfServers

			return lb.Servers[serverIdx]
		}

		lb.RoundRobinNumber = (serverIdx + 1) % numOfServers
	}

	return NodeServer{}
}

func (lb *LoadBalancer) ForwardRequest(server NodeServer, rw http.ResponseWriter, r *http.Request) {
	addr := server.GetServerAddress()

	// Create a new request to forward
	req, err := http.NewRequest(r.Method, addr, r.Body)
	if err != nil {
		// do something
	}

	// Copy/Manipulate/Cross original request headerd here
	for name, value := range r.Header {
		for _, v := range value {
			req.Header.Add(name, v)
		}
	}

	// Forward the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		// do something
	}

	// Copy/Manipulate/Cross response headers here for the original request
	for name, value := range res.Header {
		for _, v := range value {
			rw.Header().Add(name, v)
		}
	}

	//Send the response back
	rw.WriteHeader(res.StatusCode)

	if _, err = io.Copy(rw, req.Body); err != nil {
		// do something
	}

}
