package infra

import (
	"fmt"
	"os"
)

type Server struct {
	port string
}

func NewServer() *Server {
	return &Server{}
}

func WithPort(port string) func(*Server) {
	return func(server *Server) {
		server.port = port
	}
}

func (s *Server) Start() {
	port := os.Getenv("PORT")
	var options []func(*Server)
	if port == "" {
		options = append(options, WithPort("8080"))
	} else {
		options = append(options, WithPort(port))
	}
	for _, option := range options {
		option(s)
	}
	fmt.Println("Starting server on port:", s.port)
}
