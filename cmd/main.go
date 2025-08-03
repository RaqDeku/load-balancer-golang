package cmd

import (
	"load-balancer/example/internals"
	"net/http"
)

func main() {
	server1 := internals.NewNodeServer("127.0.0.1", "3000")
	server2 := internals.NewNodeServer("127.0.0.1", "3001")
	server3 := internals.NewNodeServer("127.0.0.1", "3002")

	lb := internals.NewLoadBalancer([]internals.NodeServer{
		server1,
		server2,
		server3,
	})

	server := lb.GetNextAvailableServer()

	http.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
		lb.ForwardRequest(server, w, r)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
