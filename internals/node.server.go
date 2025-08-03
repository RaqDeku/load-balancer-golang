package internals

import "fmt"

type NodeServer struct {
	Address string
	Port    string
}

func NewNodeServer(addr, port string) NodeServer {
	return NodeServer{
		Address: addr,
		Port:    port,
	}
}

func (s *NodeServer) IsServerHealthy() bool {
	// make request to health check endpoint
	return true
}

func (s *NodeServer) GetServerAddress() string {
	return fmt.Sprintf("%s:%s", s.Address, s.Port)
}
